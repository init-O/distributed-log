package main

import (
	"log"

	"github.con/init-O/distributed-log/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
