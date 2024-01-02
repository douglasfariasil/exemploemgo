package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/douglasfariasil/exemploemgo/exemplo/infra/database"
	"github.com/douglasfariasil/exemploemgo/exemplo/usecase"
	"github.com/douglasfariasil/exemploemgo/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close() // ele espera tudo rodar e depois executa o close
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgRabbitmgChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmgChannel) // escutando a fila // trava // T2
	rabbitmgWorker(msgRabbitmgChannel, uc)      // T1
}

func rabbitmgWorker(msgChan chan amqp.Delivery, uc *usecase.CalcuculateFinalPrice) {
	fmt.Println("Starting rabbitmq")
	for msg := range msgChan {
		var imput usecase.OrderInput
		err := json.Unmarshal(msg.Body, &imput)
		if err != nil {
			panic(err)
		}
		output, err := uc.Execute(imput)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println("Mensagem processada e salva no banco:", output)
	}
}
