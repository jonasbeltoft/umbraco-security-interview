package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Translation struct {
	LanguageCode string `json:"language_code"`
	Description  string `json:"description"`
	ButtonText   string `json:"button_text"`
}

func GetTranslation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	languageCode := params["languageCode"]

	// Fetch the translation based on the language code
	translation := Translation{}
	for _, lang := range TRANSLATIONS {
		if lang.LanguageCode == languageCode {
			translation = lang
		}
	}

	if translation.LanguageCode == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Convert the translation to JSON
	jsonResponse, err := json.Marshal(translation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the response headers and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func GetAllLanguages(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//languageCode := params["languageCode"]

	// Fetch all languages based on the language code

	// Convert the languages to JSON
	// jsonResponse, err := json.Marshal(languages)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// Set the response headers and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ALL_LANGUAGES)
}

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define the routes and their respective handlers
	router.HandleFunc("/translations/{languageCode}", GetTranslation).Methods("GET")
	router.HandleFunc("/allLanguages/{languageCode}", GetAllLanguages).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

type Languages struct {
	Languages []Language `json:"languages"`
}

type Language struct {
	Name         string                `json:"name"`
	Code         string                `json:"code"`
	Translations []LangNameTranslation `json:"translations"`
}

type LangNameTranslation struct {
	Code        string `json:"code"`
	Translation string `json:"translation"`
}

var TRANSLATIONS = []Translation{
	{
		LanguageCode: "en",
		Description:  "Move, add and remove new languages to display",
		ButtonText:   "Add languages",
	},
	{
		LanguageCode: "en-US",
		Description:  "Move, add and remove new languages to display",
		ButtonText:   "Add languages",
	},
	{
		LanguageCode: "da",
		Description:  "Flyt, tilføj og fjern nye sprog for at vise",
		ButtonText:   "Tilføj sprog",
	},
	{
		LanguageCode: "da-DK",
		Description:  "Flyt, tilføj og fjern nye sprog for at vise",
		ButtonText:   "Tilføj sprog",
	},
	{
		LanguageCode: "sv",
		Description:  "Flytta, lägg till och ta bort nya språk att visa",
		ButtonText:   "Lägg till språk",
	},
	{
		LanguageCode: "no",
		Description:  "Flytt, legg til og fjern nye språk for visning",
		ButtonText:   "Legg til språk",
	},
	{
		LanguageCode: "de",
		Description:  "Sprachen verschieben, hinzufügen und entfernen, um sie anzuzeigen",
		ButtonText:   "Sprachen hinzufügen",
	},
	{
		LanguageCode: "nl",
		Description:  "Verplaats, voeg toe en verwijder nieuwe talen om weer te geven",
		ButtonText:   "Talen toevoegen",
	},
}

var ALL_LANGUAGES = []byte(`{
	{
	  name: "English",
	  current_locale: "en",
	  chosen: false, // You can update this based on selected languages
	  locales: ["en-US", "en"],
	  translations: [
		{ code: "en", translation: "English" },
		{ code: "en-US", translation: "English - US" },
		{ code: "da", translation: "Danish" },
		{ code: "da-DK", translation: "Danish - Denmark" },
		{ code: "sv", translation: "Swedish" },
		{ code: "no", translation: "Norwegian" },
		{ code: "de", translation: "German" },
		{ code: "nl", translation: "Dutch" },
	  ],
	},
	{
	  name: "Dansk",
	  current_locale: "da",
	  chosen: false,
	  locales: ["da-DK", "da"],
	  translations: [
		{ code: "da", translation: "Dansk" },
		{ code: "da-DK", translation: "Dansk - Danmark" },
		{ code: "en", translation: "Engelsk" },
		{ code: "en-US", translation: "Engelsk - US" },
		{ code: "sv", translation: "Svensk" },
		{ code: "no", translation: "Norsk" },
		{ code: "de", translation: "Tysk" },
		{ code: "nl", translation: "Hollandsk" },
	  ],
	},
	{
	  name: "Svenska",
	  current_locale: "sv",
	  chosen: false,
	  locales: ["sv"],
	  translations: [
		{ code: "sv", translation: "Svenska" },
		{ code: "en", translation: "Engelska" },
		{ code: "en-US", translation: "Engelska - US" },
		{ code: "da", translation: "Danska" },
		{ code: "da-DK", translation: "Danska - Danmark" },
		{ code: "no", translation: "Norska" },
		{ code: "de", translation: "Tyska" },
		{ code: "nl", translation: "Holländska" },
	  ],
	},
	{
	  name: "Norsk",
	  current_locale: "no",
	  chosen: false,
	  locales: ["no"],
	  translations: [
		{ code: "no", translation: "Norsk" },
		{ code: "en", translation: "Engelsk" },
		{ code: "en-US", translation: "Engelsk - US" },
		{ code: "da", translation: "Dansk" },
		{ code: "da-DK", translation: "Dansk - Danmark" },
		{ code: "sv", translation: "Svensk" },
		{ code: "de", translation: "Tysk" },
		{ code: "nl", translation: "Nederlandsk" },
	  ],
	},
	{
	  name: "Deutsch",
	  current_locale: "de",
	  chosen: false,
	  locales: ["de"],
	  translations: [
		{ code: "de", translation: "Deutsch" },
		{ code: "en", translation: "Englisch" },
		{ code: "en-US", translation: "Englisch - US" },
		{ code: "da", translation: "Dänisch" },
		{ code: "da-DK", translation: "Dänisch - Dänemark" },
		{ code: "sv", translation: "Schwedisch" },
		{ code: "no", translation: "Norwegisch" },
		{ code: "nl", translation: "Niederländisch" },
	  ],
	},
	{
	  name: "Nederlands",
	  current_locale: "nl",
	  chosen: false,
	  locales: ["nl"],
	  translations: [
		{ code: "nl", translation: "Nederlands" },
		{ code: "en", translation: "Engels" },
		{ code: "en-US", translation: "Engels - US" },
		{ code: "da", translation: "Deens" },
		{ code: "da-DK", translation: "Deens - Denemarken" },
		{ code: "sv", translation: "Zweeds" },
		{ code: "no", translation: "Noors" },
		{ code: "de", translation: "Duits" },
	  ],
	},
}`)
