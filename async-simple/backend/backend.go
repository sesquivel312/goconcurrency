package be

import (
	"fmt"
	"time"
)

func DoStuff(id string, sleepSecs int64) {
	if id == "" { // if we didn't get a name for this task, call it undefined
		id = "undefined"
	}

	time.Sleep(time.Duration(sleepSecs) * time.Second)                  // this is meant to represent a synchronous call to some backend thingy
	fmt.Printf("   backend completed (%v), slept: %v\n", id, sleepSecs) // and it finally completes
}
