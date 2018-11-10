package main

import (
	"fmt"

	"gitlab.com/toby3d/telegram"
)

type (
	Repo struct {
		Owner string
		Name  string
	}

	Build struct {
		Number int
		Commit string
		Branch string
		Author string
		Status string
		Link   string
	}

	Config struct {
		AccessToken string
		ChatID      int64
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
	}
)

func (p Plugin) Exec() error {
	bot, err := telegram.New(p.Config.AccessToken)
	if err != nil {
		return err
	}

	_, err = bot.SendMessage(telegram.NewMessage(p.Config.ChatID, message(p.Repo, p.Build)))
	return err
}

func message(repo Repo, build Build) string {
	return fmt.Sprintf("%s %s <%s/%s#%s> Build #%d (%s) by %s",
		emoji(build),
		build.Status,
		repo.Owner,
		repo.Name,
		build.Commit[:8],
		build.Number,
		build.Branch,
		build.Author,
	)
}

func emoji(b Build) string {
	switch b.Status {
	case "success":
		return "👍"
	case "failure", "error", "killed":
		return "😱"
	default:
		return "🤔"
	}
}
