package main

import (
	"github.com/uvmain/uvsonic/db"
)

func main() {
	db.Init("uvsonic.db")
	StartServer()
}
