package main

import (
	"github.com/rivo/tview"
	"github.com/yitsushi/go-misskey/core"
)

func createCommandList() (commandList *tview.List) {
	commandList = tview.NewList()
	commandList.SetBorder(true).SetTitle("command")
	return commandList
}

func createTimlieTable() (timelineTable *tview.Table) {
	timelineTable = tview.NewTable()
	timelineTable.SetBorder(true).SetTitle("timeline")

	for i, v := range noteList {
		timelineTable.SetCellSimple(i, 0, v.User.Name)
		timelineTable.GetCell(i, 0).SetAlign(tview.AlignRight)
		note := tview.NewTableCell(v.Text)
		timelineTable.SetCell(i, 1, note)
	}
	return timelineTable
}

func createTimelinePanel() (timelinePanel *tview.Flex) {
	timelineTable := createTimlieTable().SetSelectable(true, false)
	timelineForm := createNoteForm()
	timelinePanel = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(timelineTable, 0, 1, true).
		AddItem(timelineForm, 0, 1, true)

	return timelinePanel
}

func createLayout(instanceName core.String, cList tview.Primitive, recvPanel tview.Primitive) (layout *tview.Flex) {
	bodyLayout := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(cList, 20, 1, true).
		AddItem(recvPanel, 0, 1, true)

	header := tview.NewTextView()
	header.SetBorder(true)
	header.SetText(*instanceName)
	header.SetTextAlign(tview.AlignCenter)

	layout = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 3, 1, false).
		AddItem(bodyLayout, 0, 1, true)
	return layout
}
