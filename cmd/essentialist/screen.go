package main

import (
	"fyne.io/fyne/v2"
)

type Application interface {
	Display(screen Screen)
	Storage() fyne.Storage
	Window() fyne.Window
	SetContent(o fyne.CanvasObject)
	SetOnTypedKey(func(*fyne.KeyEvent))
	Close()
}

type Screen interface {
	Show(app Application)
	Hide(app Application)
}

type application struct {
	app    fyne.App
	win    fyne.Window
	screen Screen
}

// Storage returns the storage a screen can use.
func (a *application) Storage() fyne.Storage {
	return a.app.Storage()
}

// Display hides the previous screen if it exists and show the new screen
func (a *application) Display(screen Screen) {
	if a.screen != nil {
		a.screen.Hide(a)
	}
	a.screen = screen
	screen.Show(a)
}

// Window returns the window where a screen should paint itself.
func (a *application) Window() fyne.Window {
	return a.win
}

// Close terminate the application
func (a *application) Close() {
	a.win.Close()
}

// SetOnTypedKey changes the keyboard shortcuts
func (a *application) SetOnTypedKey(f func(*fyne.KeyEvent)) {
	a.win.Canvas().SetOnTypedKey(f)
}

// Window returns the window where a screen should paint itself.
func (a *application) SetContent(o fyne.CanvasObject) {
	fyne.Do(func() {
		a.win.SetContent(o)
	})
}

func NewApplication(app fyne.App, window fyne.Window) Application {
	return &application{
		app: app,
		win: window,
	}
}
