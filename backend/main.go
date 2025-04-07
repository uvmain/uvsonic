package main

import (
	"github.com/uvmain/uvsonic/db"
	"github.com/uvmain/uvsonic/logic"
)

func main() {
	logic.LoadEnv()
	db.Init()
	logic.GetDirContents(logic.AudioFilesDirectory, []string{})
	StartServer()
}
