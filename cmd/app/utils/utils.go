package utils

import (
	"github.com/dinhdobathi1992/go-github-app-v2/cmd/app/config"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v41/github"
	"log"
	"net/http"
)

func InitGitHubClient() {
	tr := http.DefaultTransport
	itr, err := ghinstallation.NewKeyFromFile(tr, 337794, 643169445, "/config/github-app.pem")

	if err != nil {
		log.Fatal(err)
	}

	config.Config.GitHubClient = github.NewClient(&http.Client{Transport: itr})
}
