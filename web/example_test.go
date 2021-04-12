package web

import (
	"log"
	"time"
)

// ExampleServer is an example of how to launch a server
// If you use docker, [docker stop](https://docs.docker.com/compose/reference/stop/) has a default timeout of 10 seconds, the graceful timeout should be set to expire before then.
func ExampleServer() {
	s := Server{
		Addr:    ":8080",
		Timeout: 5 * time.Second,
	}
	log.Println("Starting...")
	log.Println(s.ListenAndServe())
}
