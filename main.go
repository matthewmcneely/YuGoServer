package main

/*
	Command-line initialization support.
	See yugo_server/yugo_server.go for more information
*/

import (
	"os"
	"yugo_server"
)

var server *yugo_server.YugoServer

func main() {
	workingDir, _ := os.Getwd()
	server = yugo_server.NewServer(8090, workingDir)
	if err := server.Run(); err != nil {
		panic(err.Error())
	}
}
