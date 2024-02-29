package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

/*
	Main function.

Program entry point.
This function recieves all of the command line arguments and runs the
appropriate function or subcommand for each one.

For example:

	In the shell:
		$ devcli user kaamkiya
	This file would then recieve the subcommand "user".
	Then it would call the displayUser function, defined in user.go
*/
func main() {
	/*if len(os.Args) < 2 {
		fmt.Println("Usage: devcli <command> [<args>] [<options>]")
		os.Exit(-1)
	}*/

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "show-comments",
				Aliases: "c",
				Usage: "show comments on an article",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "read",
				Usage: "read an article. (devcli read <author>/<article_slug>)",
				Action: func(*cli.Context) {
					readArticle()
				},
			},
			{
				Name:  "my-articles",
				Usage: "print a list of all of your articles",
				Action: func(*cli.Context) {
					myArticles()
				},
			},
			{
				Name: "write",
				Usage: "write an article - NOT YET WORKING",
				Action: func(*cli.Context) {
					writeArticle()
				},
			},
			{
				Name: "latest",
				Aliases: []string{"recent", "new"},
				Usage: "fetch the 30 most recently posted articles on dev",
				Action: func(*cli.Context) {
					recentlyPosted()
				},
			},
			{
				Name: "following-tags",
				Usage: "fetch all of the tags that you follow",
				Action: func(*cli.Context) {
					followingTags()
				},
			},
			{
				Name: "followers",
				Usage: "print a list of your followers",
				Action: func(*cli.Context) {
					followers()
				},
			},
			{
				Name: "user",
				Usage: "display information about a user",
				Action: func(*cli.Context) {
					// TODO: allow the user to pick the username
					displayUser("kaamkiya")
				},
			},
			{
				Name: "readinglist",
				Usage: "print your reading list",
				Action: func(*cli.Context) {
					readingList()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

	/*switch os.Args[1] {
	case "my-articles":
		// fetch the user's articles based on $DEV_API_KEY
		myArticles()
	case "write":
		// FIXME: this does not currently work
		writeArticle()
	case "latest":
	case "recent":
		// fetch the 30 most recent articles
		recentlyPosted()
	case "following-tags":
		// fetch the tags a user follows
		followingTags()
	case "followers":
		// fetch a user's follower list
		followers()
	case "user":
		// this command requires an argument, so...
		if len(os.Args) < 3 {
			// ... if the argument is not given, exit with a help message
			fmt.Println(`Usage: devcli user <username>`)
			os.Exit(-1)
		}
		// otherwise, display the user
		displayUser(os.Args[2])
	case "readinglist":
		// fetch a user's reading list. Requires $DEV_API_KEY
		readingList()
	case "help":
		// help message
		fmt.Println(`devcli - a CLI for dev.to

Options:
  -sc, --show-comments    with read subcommand, print the comments an article recieved

Commands:
  read                    print an article so that you can read it
  user                    show information about a user
  following-tags          show all of the tags that you follow
  followers               write a list of all of your followers
  readinglist             print your reading list and the command to use to read the article
  latest                  print the IDs of the 30 most recently written articles
  recent                  same as latest
  my-articles             print a list of all of your articles`)
	default:
		fmt.Println("Not a valid command. Run \033[38;5;245mdevcli help\033[0m for help.")
	}*/
}
