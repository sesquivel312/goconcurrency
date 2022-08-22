package client

import be "concurr/async-simple/backend"

func CallBackend(id string, sleepSecs int64) {
	go be.DoStuff(id, sleepSecs) // call the backend, in this case, "don't care" about the result
}
