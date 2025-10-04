package main

import (
	"fyne.io/fyne/v2"
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

func (s *CongratsScreen) keyHandler(app Application) func(*fyne.KeyEvent) {
	return func(key *fyne.KeyEvent) {
		if key.Name != "" {
			switch key.Name {
			case fyne.KeyQ, fyne.KeyEscape, fyne.KeyReturn, fyne.KeySpace:
				app.Display(NewSplashScreen())
			}
		} else {
			switch key.Physical {
			case fyne.HardwareKey{ScanCode: 9}, fyne.HardwareKey{ScanCode: 24},
				fyne.HardwareKey{ScanCode: 36}, fyne.HardwareKey{ScanCode: 55}: // Enter or Space
				app.Display(NewSplashScreen())
			}
		}
	}
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
	app.SetOnTypedKey(s.keyHandler(app))
}

func (s *CongratsScreen) Hide(app Application) {
	app.SetOnTypedKey(nil)
}
