package translator

import (
	"context"
	"fmt"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	"google.golang.org/api/translate/v2"
)

func CreateTranslateClient() (*translate.Service, error) {
	jsonKeyPath := "./home-lab-232620-e815cc794f9b.json" // Replace with the path to your JSON key file
	ctx := context.Background()

	client, err := translate.NewService(ctx, option.WithCredentialsFile(jsonKeyPath))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func TranslateText(client *translate.Service, sourceLang, targetLang, text string) (string, error) {
	source := language.MustParse(sourceLang)
	target := language.MustParse(targetLang)

	resp, err := client.Translations.List([]string{text}, target.String()).Source(source.String()).Do()
	if err != nil {
		return "", err
	}

	if len(resp.Translations) > 0 {
		return resp.Translations[0].TranslatedText, nil
	}

	return "", fmt.Errorf("Translation not found")
}
