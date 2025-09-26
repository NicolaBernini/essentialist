package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	essentiali18n "github.com/essentialist-app/essentialist/cmd/essentialist/i18n"
	"github.com/essentialist-app/essentialist/internal"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type SettingsScreen struct{}

func NewSettingsScreen() Screen {
	return &SettingsScreen{}
}

func (s *SettingsScreen) importDirectoryButton(app Application) *widget.Button {
	importCallback := func(d fyne.ListableURI, err error) {
		if err != nil {
			app.Display(NewErrorScreen(err))
			return
		}
		if d == nil {
			return
		}
		if err = importDirectory(d); err != nil {
			app.Display(NewErrorScreen(err))
			return
		}
		app.Display(NewSplashScreen())
	}
	return widget.NewButton(essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "import_directory"}), func() {
		dialog.ShowFolderOpen(importCallback, app.Window())
	})
}

func (s *SettingsScreen) changeDirectoryButton(app Application) *widget.Button {
	changeDirectoryCallback := func(d fyne.ListableURI, err error) {
		if err != nil {
			app.Display(NewErrorScreen(err))
			return
		}
		if d == nil {
			return
		}
		setDirectory(d)
		app.Display(NewSplashScreen())
	}
	return widget.NewButton(essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "change_directory"}), func() {
		dialog.NewFolderOpen(changeDirectoryCallback, app.Window()).Show()
	})
}

func (s *SettingsScreen) selectRepetition(app Application) *widget.Select {
	selections := []string{
		essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "10_cards"}),
		essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "20_cards"}),
		essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "30_cards"}),
		essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "40_cards"}),
		essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "50_cards"}),
		essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "remaining_cards_to_learn"}),
		essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "all_cards"}),
	}
	values := []int{
		10,
		20,
		30,
		40,
		50,
		internal.CARDS_TO_REVIEW,
		internal.ALL_CARDS,
	}
	onChange := func(selected string) {
		for i, s := range selections {
			if s == selected {
				setRepetitionLenght(values[i])
				return
			}
		}
	}
	repetitions := widget.NewSelect(selections, onChange)
	repetitions.Alignment = fyne.TextAlignCenter
	cardsNb := getRepetitionLenght()
	for i, v := range values {
		if v == cardsNb {
			repetitions.SetSelected(selections[i])
			break
		}
	}
	return repetitions
}

func (s *SettingsScreen) changeThemeSelector(app Application) *widget.Select {
	selections := []string{
		essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "theme_user_default"}),
		essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "theme_light"}),
		essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "theme_dark"}),
	}
	values := []themeName{
		defaultTheme,
		lightTheme,
		darkTheme,
	}
	onChange := func(selected string) {
		for i, s := range selections {
			if s == selected {
				setThemeName(values[i])
				fyne.CurrentApp().Settings().SetTheme(getTheme())
				return
			}
		}
	}
	themeSelector := widget.NewSelect(selections, onChange)
	themeSelector.Alignment = fyne.TextAlignCenter
	currentTheme := getThemeName()
	themeSelector.SetSelected(selections[0])
	for i, t := range values {
		if t == currentTheme {
			themeSelector.SetSelected(selections[i])
			break
		}
	}
	return themeSelector
}

func (s *SettingsScreen) cleanUpStorageButton(app Application) *widget.Button {
	cb := func(yes bool) {
		if !yes {
			return
		}
		err := cleanDirectory()
		if err != nil {
			app.Display(NewErrorScreen(err))
		}
	}
	label := essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "delete_cards_in",
		TemplateData: map[string]interface{}{
			"Directory": getDirectory().Name(),
		},
	})
	return widget.NewButton(essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "erase_storage"}), func() {
		dialog.ShowConfirm(essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "erase_storage"}), label, cb, app.Window())
	})
}

func (s *SettingsScreen) newSettingsTopBar(app Application) *fyne.Container {
	back := widget.NewButton(essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "home"}), func() {
		app.Display(NewSplashScreen())
	})
	return newTopBar(essentiali18n.Localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "settings"}), back)
}

func (s *SettingsScreen) languageSelector(app Application) *widget.Select {
	selections := []string{"English", "Español", "Français", "中文", "हिन्दी", "العربية"}
	langCodes := []string{"en", "es", "fr", "zh", "hi", "ar"}

	onChange := func(selected string) {
		for i, s := range selections {
			if s == selected {
				app.Preferences().SetString("language", langCodes[i])
				essentiali18n.SetLanguage(langCodes[i])
				return
			}
		}
	}

	selector := widget.NewSelect(selections, onChange)
	selector.Alignment = fyne.TextAlignCenter
	currentLang := app.Preferences().StringWithFallback("language", "en")
	for i, code := range langCodes {
		if code == currentLang {
			selector.SetSelected(selections[i])
			break
		}
	}
	return selector
}

func (s *SettingsScreen) Show(app Application) {
	topBar := s.newSettingsTopBar(app)

	objects := make([]fyne.CanvasObject, 0)
	if fyne.CurrentDevice().IsMobile() {
		objects = append(objects, s.importDirectoryButton(app))
		objects = append(objects, s.cleanUpStorageButton(app))
	} else {
		objects = append(objects, s.changeDirectoryButton(app))
	}
	objects = append(objects, s.changeThemeSelector(app))
	objects = append(objects, s.selectRepetition(app))
	objects = append(objects, s.languageSelector(app))
	center := container.NewVScroll(container.New(layout.NewVBoxLayout(),
		objects...))
	app.SetContent(container.New(layout.NewBorderLayout(
		topBar, nil, nil, nil), topBar, center))
	app.SetOnTypedKey(EscapeKeyHandler(app))
}

func (s *SettingsScreen) Hide(app Application) {
	app.Window().Canvas().SetOnTypedKey(nil)
}
