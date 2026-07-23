package main

import (
	"fmt"
	"github.com/Prominence673/rxemu/internal/commands"
	"os"
)

func main() {
	if err := commands.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
