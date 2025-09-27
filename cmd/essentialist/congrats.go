package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/essentialist-app/essentialist/cmd/essentialist/i18n"
	"github.com/essentialist-app/essentialist/internal"
)

type CongratsScreen struct {
	game *internal.Game
}

func NewCongratsScreen(game *internal.Game) Screen {
	return &CongratsScreen{game: game}
}

func (s *CongratsScreen) Show(app Application) {
	topBar := newProgressTopBar(app, s.game)
	label := container.New(layout.NewCenterLayout(),
		widget.NewLabel(i18n.MustLocalize("congratulations")))
	button := bottomButton(i18n.MustLocalize("press_to_continue"), func() {
		app.Display(NewSplashScreen())
	})

	box := container.New(layout.NewBorderLayout(topBar, button, nil, nil),
		topBar, button, label)
	app.SetContent(box)
	app.SetOnTypedKey(EscapeKeyHandler(app, NewSplashScreen()))
}

func (s *CongratsScreen) Hide(app Application) {
	app.SetOnTypedKey(nil)
}
