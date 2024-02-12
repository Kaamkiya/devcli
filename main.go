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
			fmt.Println(`Usage: 
  devcli read <author>/<article> [<options>]
  devcli read <article_id> [<options>]`)
			os.Exit(-1)
		}
		readArticle(os.Args[2])
	case "write":
		writeArticle()
	case "following-tags":
		followingTags()
	case "followers":
		followers()
	case "user":
		if len(os.Args) < 3 {
			fmt.Println(`Usage: devcli user <username>`)
		}
		displayUser(os.Args[2])
	case "readinglist":
		readingList()
	}
}
