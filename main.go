package main

import (
	"fmt"
	"net/http"

	"github.com/TalkBeCheap/hade/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	fmt.Println("Server listening on", 8888)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("err occurred", err)
	}
}
