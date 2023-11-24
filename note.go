package main

import (
	"log"

	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/notes"
)

func createNote(text string) notes.CreateResponse {
	response, err := client.Notes().Create(notes.CreateRequest{
		Text:       core.NewString(text),
		Visibility: models.VisibilityPublic,
	})
	if err != nil {
		log.Printf("[Create note] Error happened: %s", err)
	}
	log.Println(response.CreatedNote.ID)
	return response
}
