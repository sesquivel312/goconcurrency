package main

import (
	"concurr/async-simple/client"
	"fmt"
	"math/rand"
	"time"
)

var seed = int64(12141012)

func main() {
	rand.Seed(seed)
	ids := []string{"one", "two"} // these are the names given to the backend jobs; could represent a thing created, updated, deleted, etc.

	sleepTotalSec := int64(0)
	for _, id := range ids {
		thisSleepSec := int64(rand.Intn(10)) + 1 // how long this call to backend should sleep; +1 to avoid 0 sleep
		sleepTotalSec += thisSleepSec            // will need main to wait this long else we won't see the backend output
		fmt.Printf("calling backend with id: %v, sleep: %v\n", id, thisSleepSec)
		client.CallBackend(id, thisSleepSec)
	}

	time.Sleep(time.Duration(sleepTotalSec) * time.Second) // go ends when main ends; this is required to give the backend calls time to complete
}
