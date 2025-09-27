package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/essentialist-app/essentialist/cmd/essentialist/i18n"
)

type AboutScreen struct {
}

func NewAboutScreen() Screen {
	return &AboutScreen{}
}

func (s *AboutScreen) Show(app Application) {
	back := widget.NewButton("Settings", func() {
		app.Display(NewSettingsScreen())
	})
	topBar := newTopBar(i18n.MustLocalize("about"), back)
	app.SetContent(
		container.NewBorder(
			topBar,
			nil,
			nil,
			nil,
			topBar,
			container.NewVBox(
				widget.NewLabel("Essentialist"),
				widget.NewLabel("Version: 0.3.18"),
				NewRichTextFromMarkdown("Author: [github.com/lugu](https://github.com/lugu)"),
				NewRichTextFromMarkdown("Site: [https://essentialist.app](https://essentialist.app)"),
			),
		),
	)
}

func (s *AboutScreen) Hide(app Application) {
}
