package main

import (
	"github.com/uvmain/uvsonic/config"
	"github.com/uvmain/uvsonic/db"
	"github.com/uvmain/uvsonic/files"
	"github.com/uvmain/uvsonic/logic"
)

func main() {
	logic.LoadEnv()
	config.LoadConfig()
	db.Init()
	files.Init()
	StartServer()
}
