package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
	"time"
)

const keys = "myJobQueue"

func main() {

	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	fmt.Println("Waiting for jobs on jobQueue: ", keys)

	go func() {
		for {
			result, err := c.BLPop(0*time.Second, keys).Result()

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Executing job: ", result[1])
		}
	}()

	// block for ever, used for testing only
	select {}
}
