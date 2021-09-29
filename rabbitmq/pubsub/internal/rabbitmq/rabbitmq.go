package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Service interface{
	Connect() error
	Publish(message string) error
}

type RabbitMQ struct {
	Conn *amqp.Connection
	Channel *amqp.Channel
	Queue amqp.Queue
}

func (r *RabbitMQ) Connect() error {
	fmt.Println("Connecting to RabbitMQ")
	var err error
	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to RabbitMQ")

	r.Channel, err = r.Conn.Channel()
	if err != nil {
		return err
	}

	err = r.Channel.ExchangeDeclare(
		"warranty",  
		"fanout",
		true,     
		false,   
		false,   
		false,    
		nil,     
    ); 
	if err != nil {
		return err
	}
        

	r.Queue, err = r.Channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = r.Channel.QueueBind(
		r.Queue.Name, 
		"",    
		"warranty",
		false,
		nil,
    )
	if err != nil {
		return err
	}

	return nil
}

//Publish - takes in a string 'message' and pusblished to a rmq queue.
func (r *RabbitMQ) Publish(message string) error {
	err := r.Channel.Publish(
		"warranty",
		r.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		},
	)

	if err != nil {
		return err
	}

	fmt.Println("Successfully published message to queue.")
	return nil
}

func (r *RabbitMQ) Consume() {

	msgs, err := r.Channel.Consume(
		r.Queue.Name,
		"",     
		true,  
		false, 
		false, 
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Application Error on Consume: %s", err)
	}

	infinity := make(chan bool)

	go func() {
		for d := range msgs {
		   log.Printf("Received New Message: %s", d.Body)
		}
     }()

	
	log.Println("Waiting for new messages. Press CTRL+C to exit.")
	<- infinity
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}