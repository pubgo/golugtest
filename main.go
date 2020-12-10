package main

import (
	"github.com/pubgo/golug"
	"github.com/pubgo/golugtest/ctl_entry"
	"github.com/pubgo/golugtest/grpc_entry"
	"github.com/pubgo/golugtest/http_entry"
)

func main() {
	golug.Init()
	golug.Run(
		http_entry.GetEntry(),
		ctl_entry.GetEntry(),
		grpc_entry.GetEntry(),
		grpc_entry.GetHttpEntry(),
	)
}
