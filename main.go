package main

import (
	"fmt"
	"os"
	"time"

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
	var showComments bool

	app := &cli.App{
		Name:    "devcli",
		Version: "0.2.0",
		UseShortOptionHandling: true,
		Usage: "A CLI for dev.to",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Kaamkiya",
				Email: "codeberg.org/kaamkiya",
			},
		},
		Copyright: fmt.Sprintf("(c) 2024-%d under the GNU AGPLv3 License", time.Time.Year(time.Now())),
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "show-comments",
				Aliases: []string{"c"},
				Usage:   "show comments on an article",
				Destination: &showComments,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "read",
				Usage: "read an article. (devcli read <author>/<article_slug>)",
				Action: func(ctx *cli.Context) error {
					readArticle(ctx.Args().Get(0), showComments)
					return nil
				},
			},
			{
				Name:  "my-articles",
				Usage: "print a list of all of your articles",
				Action: func(*cli.Context) error {
					myArticles()
					return nil
				},
			},
			{
				Name:  "write",
				Usage: "write an article - NOT YET WORKING",
				Action: func(*cli.Context) error {
					writeArticle()
					return nil
				},
			},
			{
				Name:    "latest",
				Aliases: []string{"recent", "new"},
				Usage:   "fetch the 30 most recently posted articles on dev",
				Action: func(*cli.Context) error {
					recentlyPosted()
					return nil
				},
			},
			{
				Name:  "following-tags",
				Usage: "fetch all of the tags that you follow",
				Action: func(*cli.Context) error {
					followingTags()
					return nil
				},
			},
			{
				Name:  "followers",
				Usage: "print a list of your followers",
				Action: func(*cli.Context) error {
					followers()
					return nil
				},
			},
			{
				Name:  "user",
				Usage: "display information about a user",
				Action: func(*cli.Context) error {
					// TODO: allow the user to pick the username
					displayUser("kaamkiya")
					return nil
				},
			},
			{
				Name:  "readinglist",
				Usage: "print your reading list",
				Action: func(*cli.Context) error {
					readingList()
					return nil
				},
			},
		},
	}

	// add a version flag (-v, -V, --version) to print the program version
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v", "V"},
		Usage:   "print version info and exit",
	}

	// suggest commands if the user entered a non-existant one
	app.Suggest = true

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
