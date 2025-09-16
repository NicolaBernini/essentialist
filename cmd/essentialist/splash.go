package main

import (
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	welcomeMessage = "Please set the directory containing your flashcards."
)

type SplashScreen struct{}

func (s *SplashScreen) load(app Application) {
	decks, err := loadDecks()
	if err != nil {
		app.Display(NewErrorScreen(err))
		return
	}
	sort.SliceStable(decks, func(i, j int) bool {
		return decks[i].DeckName() < decks[j].DeckName()
	})
	app.Display(NewHomeScreen(decks))
}

func newWelcomeTopBar(app Application) *fyne.Container {
	home := widget.NewButton("Settings", func() {
		app.Display(NewSettingsScreen())
	})
	return newTopBar("Welcome", home)
}

func (s *SplashScreen) keyHandler(app Application) func(*fyne.KeyEvent) {
	return func(key *fyne.KeyEvent) {
		if key.Name != "" {
			switch key.Name {
			case fyne.KeyQ, fyne.KeyEscape:
				app.Close()
			case fyne.KeyReturn, fyne.KeySpace:
				app.Display(NewSettingsScreen())
			}
		} else {
			switch key.Physical {
			case fyne.HardwareKey{ScanCode: 9}, fyne.HardwareKey{ScanCode: 24}: // Escape
				app.Close()
			case fyne.HardwareKey{ScanCode: 36}, fyne.HardwareKey{ScanCode: 55}: // Enter or Space
				app.Display(NewSettingsScreen())
			}
		}
	}
}

// Show an empty screen until the games are loaded, then shows HomeScreen.
func (s *SplashScreen) Show(app Application) {
	// Welcome message when the application is launched for the first time.
	prefs := fyne.CurrentApp().Preferences()
	dir := prefs.StringWithFallback("directory", "")
	if dir == "" {
		topBar := newWelcomeTopBar(app)
		welcomeButton := widget.NewButton(welcomeMessage, func() {
			app.Display(NewSettingsScreen())
		})
		center := container.NewVScroll(welcomeButton)
		app.SetContent(container.New(layout.NewBorderLayout(
			topBar, nil, nil, nil), topBar, center))
		app.SetOnTypedKey(s.keyHandler(app))
		return
	}
	emptyContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer())
	app.SetContent(emptyContainer)
	go s.load(app) // load the games in the background
}

func (s *SplashScreen) Hide(app Application) {}

func NewSplashScreen() Screen {
	return &SplashScreen{}
}
