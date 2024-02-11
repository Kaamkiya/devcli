package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: devcli <command> [<args>] [<options>]")
		os.Exit(-1)
	}

	switch os.Args[1] {
	case "read":
		if len(os.Args) < 3 {
			fmt.Println("Error")
			os.Exit(-1)
		}
		readArticle(os.Args[2])
	}
}
