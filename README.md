# Umbraco Security Interview Writeup
Demo app and write up for Umbraco Security Engineering position

## Security issues in free trial

Pass, can't find anything in large public commercial software

## Languages app

The application is a simple website that allows a user to select one or more languages, order them by preference, and have the website be displayed in the top most language.

### Frontend UI

### Frontend functionality

* The current settings are stored in local storage to persist in sessions / tabs. This is represented with the key ```deviceSettings``` that has the JSON formatted language codes. This object is represented like so:
  ```
  {
    "languages" : [
      <language code string>
    ]
  }
  ```
  When adding several chosen languages, this is the array that is appended to, to represent state accross tabs. This is also the array that dictates the order of the languages.
  
* The top most (first) element in the array is the chosen one. The language specific translations of the sites text, such as the initial description and the add button is recieved and stored in local storage as well. This is set as ```translations``` with the JSON value of:
  ```
  {
    "SortLanguagesPrompt" : <description text string>,
    "ChoosePreferredLanguagesPrompt" : <add button text string>,
    "NextButtonText" : <next button text string> (unused?)
  }
  ```

* To keep state the same accross tabs, the site needs to listen to the local storage changes, thus it must implement a listener on ```StorageEvent``` to check for changes, and reload with the new settings.

* Upon changing the specific local language (using the dropdown on a specific language), no matter the language, a request is made to ```/translations/<code>``` for the currently selected (top most) language to update descriptions.
* If a change is made to the local language (using the dropdown) that is at the top, then both a request to ```/translations/<code>``` and ```/allLanguages/<code>``` is made to update the language used to show the other languages.
* When a language is set as the first (top), a request is again sent to both previous endpoints with the new language code, and the text of the site and languages are updated, just as previously described. See pictures below, where the top 2 languages are switched.


  

### Backend
