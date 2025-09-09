package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

//go:generate go run generate.go

func parseArgs() {
	if len(os.Args) == 2 {
		overrideDirectory = os.Args[1]
	}
}

func main() {
	parseArgs()
	application := app.NewWithID("io.github.essentialist_app.essentialist")
	application.Settings().SetTheme(getTheme())
	window := application.NewWindow("Essentialist")
	window.SetIcon(fyne.NewStaticResource("Icon.png", IconBytes))
	// window.Resize(fyne.NewSize(640, 480))
	NewApplication(application, window).Display(NewSplashScreen())
	window.ShowAndRun()
}
