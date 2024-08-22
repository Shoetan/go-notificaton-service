package main

import (
	"log"

	"github.com/Shoetan/database"
	"github.com/Shoetan/models"
	"github.com/Shoetan/pkg/rabbitmq"
	"github.com/Shoetan/pkg/server"
	_ "github.com/lib/pq"
)


func main()  {
	

	db, err := database.Database()

	if err != nil {
		return
	}

	
	db.MustExec(models.TABLES)

	conn, err := rabbitmq.RabbitMqConn()

	if err != nil {
		log.Println("Could not connect to rabbitmq server")
	}

	rabbitmq.ReceiveFromQueue(conn)

	server := server.NewAPISERVER(":8080")

	server.Run()

	
}