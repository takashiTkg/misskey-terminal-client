package main

import (
	"log"

	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/notes/timeline"
)

func getLocalNoteList(limit uint) []models.Note {
	noteList, err := client.Notes().Timeline().Local(timeline.LocalRequest{
		Limit: limit,
	})
	if err != nil {
		log.Printf("[Get Notes] Error happened: %s", err)
	}
	return noteList
}
