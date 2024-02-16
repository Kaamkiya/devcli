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
  devcli read <author>/<article_slug> [<options>]
  devcli read <article_id> [<options>]`)
			os.Exit(-1)
		}
		readArticle(os.Args[2])
	case "write":
		writeArticle()
	case "latest":
	case "recent":
		recentlyPosted()
	case "following-tags":
		followingTags()
	case "followers":
		followers()
	case "user":
		if len(os.Args) < 3 {
			fmt.Println(`Usage: devcli user <username>`)
			os.Exit(-1)
		}
		displayUser(os.Args[2])
	case "readinglist":
		readingList()
	case "help":
		fmt.Println(`devcli help

Commands:
  read                print an article so that you can read it
  user                show information about a user
  following-tags      show all of the tags that you follow
  followers           write a list of all of your followers
  readinglist         print your reading list and the command to use to read the article
  latest              print the IDs of the 30 most recently written articles
  recent              same as latest`)
	default:
		fmt.Println("Not a valid command. Run \033[38;5;245mdevcli help\033[0m for help.")
	}
}
