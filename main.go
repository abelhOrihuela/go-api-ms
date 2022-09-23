package main

import (
	"banking.com/abelh/app"
	"banking.com/abelh/logger"
)

func main() {
	logger.Info("Starting our application...")
	app.Start()
}
