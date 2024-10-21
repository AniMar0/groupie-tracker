package main

import (
	"fmt"

	TRC "TRC/lib"
)

func main() {
	var Server TRC.Server
	fmt.Println("http://localhost:8080/")
	Server.Run()
}
