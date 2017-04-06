package main

import (
	"os"
	"os/signal"
)

// Wait for process to be terminated.
func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	os.Exit(0)
}
