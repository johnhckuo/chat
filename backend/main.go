package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/johnhckuo/chat/backend/pkg/chat"
)

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple server")
		log.Printf("Someone logged in...")
	})
}

var connectionString = "redis://localhost:6379"

func main() {

	setupRoutes()
	go http.ListenAndServe(":8080", nil)

	client := chat.NewRedis(connectionString)

	client.Push("hehe", "haha")
	res, err := client.Pop("hehe")
	if err != nil {
		panic(err)
	}
	log.Println(res)
	/*
		err := client.Set(ctx, "key", "test", 0).Err()
		if err != nil {
			panic(err)
		}

		val, err := client.Get(ctx, "key").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key", val)

		var mq []int = []int{1, 2, 3, 4}
		//var mq2 []string = []string{"123", "dcw"}

		r, _ := json.Marshal(mq)
		result := client.LPush(ctx, "mq", r)
		fmt.Println(result)
		len, _ := client.LLen(ctx, "mq").Uint64()

		res := client.LRange(ctx, "mq", 0, int64(len))
		fmt.Println(res)

	*/
	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

}
