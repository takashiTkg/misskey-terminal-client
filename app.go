package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/yitsushi/go-misskey/core"
)

func createApp(instanceName core.String) *tview.Application {
	app := tview.NewApplication()
	timelinePanel := createTimelinePanel()
	commandList := createCommandList()
	commandList.AddItem("Quit", "", 'q', func() {
		app.Stop()
	})
	commandList.AddItem("Load notes", "", 'r', func() {
		noteList = append(noteList, getLocalNoteList(50)...)
		timelinePanel = createTimelinePanel()
	})
	commandList.AddItem("Focus timeline", "", 't', func() {
		app.SetFocus(timelinePanel.GetItem(0))
	})
	commandList.AddItem("New note", "", 'n', func() {
		app.SetFocus(timelinePanel.GetItem(1))
	})

	layout := createLayout(instanceName, commandList, timelinePanel)
	// create page
	pages := tview.NewPages()
	pages.AddPage("main", layout, true, true)
	app.SetRoot(pages, true)
	// set custom key
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			app.SetFocus(commandList)
		}
		return event
	})

	return app
}
