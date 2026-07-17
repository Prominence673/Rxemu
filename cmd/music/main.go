package main

import (
	"os"
	"fmt"
	"github.com/Prominence673/rxemu/internal/commands"
)

func main(){
	if err := commands.Execute(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}