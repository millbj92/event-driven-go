package main

import (
	"fmt"

	"github.com/millbj92/go-rabbitmq/internal/rabbitmq"
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

	err = app.Rmq.Publish("Hi")
	if err != nil {
		return err
	}

	app.Rmq.Consume()
	
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error setting up application!")
		fmt.Println(err)
	}
}