package main

import (
	"github.com/pubgo/golug"
	"github.com/pubgo/golugtest/tickrun/server"
	"github.com/pubgo/golugtest/tickrun/worker"
)

func main() {
	golug.Init()
	golug.Run(
		server.GetEntry(),
		worker.GetEntry(),
	)
}
