package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	rabbitmq "github.com/millbj92/go-events/rabbitmq/pubsub/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func Run() error {
	fmt.Println("Go RabbitMQ test")

	rmq := rabbitmq.NewRabbitMQService()
	app := App{
		Rmq: rmq,
	}

	err := app.Rmq.Connect()
	if err != nil {
		return err
	}
	defer app.Rmq.Conn.Close()


	
	err = app.Rmq.Publish("I've been trying to get ahold of you about your cars extended warranty.")
	if err != nil {
		return err
	}
	

	for{
		log.Println()
		log.Println("Press CTRL+C to exit.")
		log.Print("Enter a message to publish: ")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		err = app.Rmq.Publish(text)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error setting up application!")
		fmt.Println(err)
	}
}