package main

import (
	rclient "concurr/async-redis/redis"
	"fmt"
	"github.com/go-redis/redis/v9"
)

func main() {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	d := rclient.Device{"d1", false}

	_ = rclient.SaveDevice(d, c)

	dd := rclient.RetrieveDevice(d.Id, c)

	fmt.Printf("set d, then retrieved it and got: %+v\n", dd)
}
