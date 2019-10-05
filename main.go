package main

import (
	"github.com/yamagn/qbox/db"
	"github.com/yamagn/qbox/server"
)

func main() {
	db.Init()
	server.Init()
}

