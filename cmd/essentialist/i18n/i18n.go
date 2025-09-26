package i18n

import (
	"encoding/json"
	"embed"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales/*.json
var fs embed.FS

var AppBundle *i18n.Bundle
var Localizer *i18n.Localizer

func init() {
	AppBundle = i18n.NewBundle(language.English)
	AppBundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	AppBundle.LoadMessageFileFS(fs, "locales/en.json")
	AppBundle.LoadMessageFileFS(fs, "locales/ar.json")
	AppBundle.LoadMessageFileFS(fs, "locales/es.json")
	AppBundle.LoadMessageFileFS(fs, "locales/fr.json")
	AppBundle.LoadMessageFileFS(fs, "locales/hi.json")
	AppBundle.LoadMessageFileFS(fs, "locales/zh.json")
}

func SetLanguage(lang string) {
	Localizer = i18n.NewLocalizer(AppBundle, lang)

}