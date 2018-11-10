package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	Version   = "unknown"
	BuildDate = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "telegram plugin"
	app.Usage = "telegram plugin"
	app.Action = run
	app.Version = fmt.Sprintf("%s+%s", Version, BuildDate)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "access-token",
			Usage:  "telegram access token",
			EnvVar: "TELEGRAM_ACCESS_TOKEN,PLUGIN_ACCESS_TOKEN",
		},
		cli.StringFlag{
			Name:   "chat-id",
			Usage:  "telegram chat id, you can obtain it from @chatid_echo_bot",
			EnvVar: "TELEGRAM_CHAT_ID,PLUGIN_CHAT_ID",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
			Value:  "00000000",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Commit: c.String("commit.sha"),
			Branch: c.String("commit.branch"),
			Author: c.String("commit.author"),
			Number: c.Int("build.number"),
			Status: c.String("build.status"),
			Link:   c.String("build.link"),
		},
		Config: Config{
			AccessToken: c.String("access-token"),
			ChatID:      c.Int64("chat-id"),
		},
	}
	return plugin.Exec()
}
