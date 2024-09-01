# Umbraco Security Interview Writeup

Demo app and write up for Umbraco Security Engineering position

## Security issues in free trial

Pass, can't find anything in large public commercial software

## Languages app

The application is a simple website that allows a user to select one or more languages, order them by preference, and have the website be displayed in the top most language.

### Frontend UI

- The site has text at the top to describe the purpose of the site. This text is translated to the currently selected language and is gotten from the `translations` object in local storage. The key for the text value is `SortLanguagesPrompt`.

- Below the description, there is a list of languages that the user has chosen. The languages are displayed in the language of the currently selected language (the language at the top). The languages are fetched from the `/allLanguages/<code>` endpoint, where `<code>` is the language code of the currently selected language.

- The list consists of cards, where each card has the language written to the left, and a delete button to the right. Each has a dropdown arrow in the middle (right aligned) which allows the user to change the localized version of the language. The dropdown is populated with the locales of the language, and the language is changed when a locale is selected.

- The list a drag and drop functionality, where the user can drag and drop the languages to change the order. The order is saved in local storage, and is used to determine the order of the languages (read ahead for implementation).

- There is a button at the bottom of the list that allows the user to add more languages to the list. It's text is specified by the currently selected language (read ahead for language implementation), the text is also gotten from the `translations` object in local storage, with the key for the text value being `ChoosePreferredLanguagesPrompt`. This button spawns a modal (dialog) that allows the user to search for languages and add them to the list.

#### Add language modal

- The modal has a search bar at the top that allows the user to search for languages. The search is done in real time, and filters the languages based on the input. The search is case insensitive.

- Next to the search bar, there is a button with an X that closes the modal.

- Below the search bar, there is a list of languages that can be selected. Each language has a checkbox to the left of it, and the language name to the right. The language name is in the language of the currently selected (top most) language (read ahead for translations implementation).

- Below, the list to the left, there is a duplicate search bar, which i synced with the one above.

- To the bottom right, the modal has a button to add the selected languages to the list of chosen languages. This button is not disabled if no languages are selected.

- The button has the text **"Add \<language\>"** and then if several are selected, it will comma sepperate them like this **"Add \<language\>, \<language\>, \<language\>"** and so on. The **"Add"** is always in english, and the language names are in the language of the currently selected (top most) language.

### Frontend functionality

- All the languages available are fetched from the `/allLanguages/<code>` endpoint, and are stored in an array in memory. See the backend section for more information on the endpoint.

- The current settings are stored in local storage to persist in sessions / tabs. This is represented with the key `deviceSettings` that has the JSON formatted language codes. This object is represented like so:

  ```
  {
    "languages" : [
      <language code string>
    ]
  }
  ```

  When adding several chosen languages, this is the array that is appended to, to represent state accross tabs. This is also the array that dictates the order of the languages. So when the page loads, it checks if this array is in local storage, and if it is, it uses this array to populate the list of chosen languages. If it is not, it uses danish as the default language, and sets the array to `["da"]`.

- The top most (first) element in the array is the chosen one.

- The language specific translations of the sites text, such as the initial description and the add button is recieved from the `/translations/<code>` endpoint, and stored in local storage. This is set with the key `translations` with the JSON value of:

  ```
  {
    "SortLanguagesPrompt" : <description text string>,
    "ChoosePreferredLanguagesPrompt" : <add button text string>,
    "NextButtonText" : <next button text string> (unused?)
  }
  ```

- To keep state the same accross tabs, the site needs to listen to the local storage changes, thus it must implement a listener on `StorageEvent` to check for changes, and reload with the new settings.

- Upon changing the specific local language (using the dropdown on a specific language), no matter the language, a request is made to `/translations/<code>` for the currently selected (top most) language to update descriptions.

- If a change is made to the local language (using the dropdown) that is at the top, then both a request to `/translations/<code>` and `/allLanguages/<code>` is made to update the language used to show the other languages.

- When a language is set as the first (top), a request is again sent to both previous endpoints with the new language code, and the text of the site and languages are updated, just as previously described. See pictures below, where the top 2 languages are switched. Notice how both the site text changes, and the names of the languages are changed to represent the new top language.

![danish language](images/lang_code_0.png)
![Azerbaijani language](images/lang_code_1.png)

### Backend

- The `/allLanguages/<code>` endpoint returnins a list of all languages, translated into the specified language. It has the `isSelected` field in each language be false, except for the specified language, which is true. A shortened list of the returned languages would look like this:

  ```
  [
    {
      "isSelected": false,
      "languageIsoCodesWithLocales": {
        "cy": "Cymraeg (walisisk)",
        "cy-GB": "Cymraeg (walisisk) - Y Deyrnas Unedig (United Kingdom)"
      }
    },
    {
      "isSelected": true,
      "languageIsoCodesWithLocales": {
        "da": "dansk",
        "da-DK": "dansk - Danmark (Denmark)",
        "da-GL": "dansk - Grønland (Greenland)"
      }
    },
    {
      "isSelected": false,
      "languageIsoCodesWithLocales": {
        "dav": "Kitaita (taita)",
        "dav-KE": "Kitaita (taita) - Kenya"
      }
    }
  ]
  ```

  _Notice how the `languageIsoCodesWithLocales` object has the language code as the key, and the language name as the value. This is used to display the language locales in the dropdown on each language._

- The `/translations/<code>` endpoint returns the translations of the site text in the specified language. The response is a JSON object with the keys `SortLanguagesPrompt`, `ChoosePreferredLanguagesPrompt`, and `NextButtonText` with the corresponding text in the specified language. An example response would look like this:

  ```
  {
    "SortLanguagesPrompt": "Vælg dine foretrukne sprog",
    "ChoosePreferredLanguagesPrompt": "Tilføj sprog",
    "NextButtonText": "Næste"
  }
  ```

  If no translations are found, the response is filled with the key as the value, like this:

  ```
  {
    "SortLanguagesPrompt": "SortLanguagesPrompt",
    "ChoosePreferredLanguagesPrompt": "ChoosePreferredLanguagesPrompt",
    "NextButtonText": "NextButtonText"
  }
  ```
