package main

import (
	"github.com/pubgo/golug"
	"github.com/pubgo/golugtest/ossync/rsync"
)

func main() {
	golug.Init()
	golug.Run(rsync.GetEntry())
}
