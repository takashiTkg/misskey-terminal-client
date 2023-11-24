package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/models"
)

var client *misskey.Client
var noteList []models.Note
var formText string

func main() {
	loadEnv()
	// create misskey client
	cl, err := misskey.NewClientWithOptions(
		misskey.WithAPIToken(os.Getenv("MISSKEY_TOKEN")),
		misskey.WithBaseURL("https", os.Getenv("MISSKEY_HOST"), ""),
		misskey.WithLogLevel(logrus.ErrorLevel),
	)
	if err != nil {
		logrus.Error(err.Error())
	}
	client = cl

	// get instnce meta.
	meta := getInstanceMeta()
	// get note init.
	noteList = getLocalNoteList(50)
	// create tview app.
	app := createApp(meta.Name)

	if err := app.Run(); err != nil {
		panic(err)
	}

}
