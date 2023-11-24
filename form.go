package main

import (
	"github.com/rivo/tview"
)

func createNoteForm() (form *tview.Form) {
	form = tview.NewForm().
		AddTextArea("text", "", 200, 0, 0, func(text string) {
			formText = text
		}).
		AddButton("Note", func() {
			createNote(formText)
		})
	form.SetBorder(true)
	return form
}
