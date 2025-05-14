package main

import (
	"fmt"
	"yim-go/api/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		panic(err)		
	}
	fmt.Println("Server stopped.")
}