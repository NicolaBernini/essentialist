package i18n

import (
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func init() {
	AppBundle = i18n.NewBundle(language.English)
	AppBundle.AddMessages(language.English, &i18n.Message{ID: "congratulations", Other: "Congratulations"})
	AppBundle.AddMessages(language.English, &i18n.Message{ID: "hello", Other: "Hello {{.Name}}"})
	Localizer = i18n.NewLocalizer(AppBundle, language.English.String())
}

func TestMustLocalize(t *testing.T) {
	// Simple case
	translation := MustLocalize("congratulations")
	assert.Equal(t, "Congratulations", translation)

	// With template data
	translation = MustLocalize("hello", "Name", "John")
	assert.Equal(t, "Hello John", translation)

	// Panic with odd number of arguments
	assert.Panics(t, func() {
		MustLocalize("hello", "Name")
	})

	// Panic with non-string key
	assert.Panics(t, func() {
		MustLocalize("hello", 1, "John")
	})
}
