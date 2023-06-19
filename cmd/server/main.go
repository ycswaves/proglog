package main

import (
	"fmt"
	"github.com/ycswaves/proglog/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	fmt.Println("serving at 8080")
	log.Fatal(srv.ListenAndServe())
}
