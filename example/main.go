package main

import (
	"fmt"
	"github.com/yyangl/mars"
)

func main() {
	//fmt.Printf("hello")
	server := mars.NewServer()
	if err := server.Run(); err != nil {
		fmt.Printf("server run error err %v", err)
	}
}
