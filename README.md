# Umbraco Security System Requirements

Demo app and write up for Umbraco Security Engineering position

## Security issues in free trial

I looked at the github and the issues. There were several, however these are known.
I looked at some code in the repo, but didn't spot anything.
I downloaded and decompiled the Nuget package using dotpeak, but also didn't find anything.

## Application writeup

This web application allows users to select and prioritize languages to display the website in their preferred language. It supports multilingual text and settings synchronized across browser tabs but isolated across different sessions.
For very detailed description see the other readme.

### Features

1. **Language Selection and Ordering**:

   - Users can select languages and reorder them by dragging and dropping.
   - Users can change locale specific versions of the language by using a dropdown on each language.
   - The top-selected language determines the display language of the website.
   - Changes are saved in the browser’s local storage to persist across sessions.
   - Users can delete languages by clicking the trash icon.

2. **Add Language Modal**:

   - A modal window allows users to search for and add new languages.
   - The button to open the modal is located at the bottom of the language list.
   - It includes real-time, case-insensitive search and multiple selection (checkboxes).

3. **Dynamic Text and UI Updates**:

   - Text elements such as descriptions and buttons are translated into the top-selected language.
   - Translations are fetched from the `/translations/<language code>` endpoint and stored in local storage.
   - The application fetches available languages from the `/allLanguages/<language code>` endpoint.
   - The site main and specific languages text is updated in real-time when the user changes the selected language or it's locale.

4. **Local Storage Usage**:

   - Language preferences and translations are stored in local storage.
   - The storage contains an array of selected language codes, dictating the order of languages and localized text for various UI elements.

5. **Cross-Tab Synchronization**:

   - The application uses a `StorageEvent` listener or similar, to detect changes in local storage and update all open tabs accordingly.

6. **Backend Endpoints**:

   - **GET `/allLanguages/<language code>`**: Returns a list of all languages translated into the specified language.
   - **GET `/translations/<language code>`**: Returns the site text translations for the specified language.
   - The `<language code>` is an ISO language code such as **da**.

### Backend Storage

- **Data Storage**:

  - Data could be stored in a NoSQL database like MongoDB for efficiency and flexibility.
  - Language data and translations can be organized by language code to optimize read performance and support easy additions of new languages.

- **Schemas**:
  - **Languages Schema**: Each document represents a language with its corresponding translations and locales.
  - **Translations Schema**: Stores UI text in different languages to support localization.

### Potential Improvements

- Consider using a relational database for structured data and complex queries.
- Implement caching strategies to minimize database reads and enhance performance.
