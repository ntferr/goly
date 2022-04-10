package main

import (
	"github.com/ntferr/goly/app/model"
	"github.com/ntferr/goly/app/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
