package i18n

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Xuanwo/go-locale"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales/*.json
var localeFS embed.FS
var AppBundle *i18n.Bundle
var Localizer *i18n.Localizer
var Languages []string

func init() {
	AppBundle = i18n.NewBundle(language.English)
	AppBundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	localeFiles, err := localeFS.ReadDir("locales")
	if err != nil {
		panic(fmt.Sprintf("failed to read locales directory: %v", err))
	}

	for _, localeFile := range localeFiles {
		fileName := localeFile.Name()
		file, err := localeFS.ReadFile("locales/" + fileName)
		if err != nil {
			panic(fmt.Sprintf("failed to read locale file %s: %v", fileName, err))
		}
		AppBundle.ParseMessageFileBytes(file, fileName)
		lang, _ := strings.CutSuffix(fileName, path.Ext(fileName))
		Languages = append(Languages, lang)
	}

	userLocale, err := locale.Detect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to detect locale: %v", err)
		userLocale = language.English
	}

	lang, _ := userLocale.Base()
	SetLanguage(lang.String())
}

func SetLanguage(lang string) {
	for _, l := range Languages {
		if lang == l {
			Localizer = i18n.NewLocalizer(AppBundle, lang)
		}
	}
}

func MustLocalize(messageID string, args ...interface{}) string {
	if len(args)%2 != 0 {
		panic("odd number of arguments passed to MustLocalize")
	}

	templateData := make(map[string]interface{}, len(args)/2)
	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string)
		if !ok {
			panic(fmt.Sprintf("template data key is not a string: %v", args[i]))
		}
		templateData[key] = args[i+1]
	}

	return Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})
}
