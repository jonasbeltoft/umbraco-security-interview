<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  DialogClose,
} from "@/components/ui/dialog";
import { Card, CardTitle, CardHeader } from "./components/ui/card";
import { Checkbox } from "@/components/ui/checkbox";
import { Search } from "lucide-vue-next";
import { Input } from "@/components/ui/input";
import { useStorage } from "@vueuse/core";
import { Trash } from "lucide-vue-next";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

import draggable from "vuedraggable";

const page_text = useStorage("page_text", {
  button_text: "Add languages",
  description: "This is a description",
});
const DEFAULT_LANGUAGES = ["da"];
// Load chosen languages from local storage
const chosen_languages = useStorage("chosen_languages", DEFAULT_LANGUAGES);

const current_language = computed(() => chosen_languages.value[0]);

let search = ref("");

function addLanguages() {
  console.log(checked_languages.value);
  chosen_languages.value.push(...checked_languages.value);
  checked_languages.value = [];
}

function updateLocale(new_locale: string, locale: string) {
  const index = all_languages.value.findIndex((lang) =>
    lang.locales.includes(new_locale)
  );
  all_languages.value[index].current_locale = new_locale;
  const index2 = chosen_languages.value.findIndex((lang) => lang === locale);
  chosen_languages.value[index2] = new_locale;
}

function deleteLanguage(language: any) {
  if (chosen_languages.value.length === 1) {
    chosen_languages.value = DEFAULT_LANGUAGES.flat();
    for (let i = 0; i < all_languages.value.length; i++) {
      if (
        !all_languages.value[i].locales.some((locale: string) =>
          DEFAULT_LANGUAGES.includes(locale)
        )
      ) {
        all_languages.value[i].chosen = false;
      }
    }
    return;
  }

  for (let i = 0; i < all_languages.value.length; i++) {
    if (all_languages.value[i].name === language) {
      all_languages.value[i].chosen = false;
    }
  }
  const index = chosen_languages.value.indexOf(language);
  chosen_languages.value.splice(index, 1);
}

const getItemKey = (element: string, index: number) => `${element}-${index}`;

const checked_languages = ref<string[]>([]);

function updateChecked(checked: boolean, lang: any) {
  if (checked) {
    checked_languages.value.push(lang.current_locale);
  } else {
    const index = checked_languages.value.indexOf(lang.current_locale);
    checked_languages.value.splice(index, 1);
  }
}

function handleClose(open: boolean) {
  if (!open) {
    search.value = "";
    checked_languages.value = [];
  }
}

const all_languages = ref([
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
]);

const text_translations = ref([
  {
    LanguageCode: "en",
    Description: "Move, add and remove new languages to display",
    ButtonText: "Add languages",
  },
  {
    LanguageCode: "en-US",
    Description: "Move, add and remove new languages to display",
    ButtonText: "Add languages",
  },
  {
    LanguageCode: "da",
    Description: "Flyt, tilføj og fjern nye sprog at vise",
    ButtonText: "Tilføj sprog",
  },
  {
    LanguageCode: "da-DK",
    Description: "Flyt, tilføj og fjern nye sprog at vise",
    ButtonText: "Tilføj sprog",
  },
  {
    LanguageCode: "sv",
    Description: "Flytta, lägg till och ta bort nya språk att visa",
    ButtonText: "Lägg till språk",
  },
  {
    LanguageCode: "no",
    Description: "Flytt, legg til og fjern nye språk for visning",
    ButtonText: "Legg til språk",
  },
  {
    LanguageCode: "de",
    Description:
      "Sprachen verschieben, hinzufügen und entfernen, um sie anzuzeigen",
    ButtonText: "Sprachen hinzufügen",
  },
  {
    LanguageCode: "nl",
    Description:
      "Verplaats, voeg toe en verwijder nieuwe talen om weer te geven",
    ButtonText: "Talen toevoegen",
  },
]);

watch(
  current_language,
  (lang) => {
    page_text.value.description = text_translations.value.find(
      (text) => text.LanguageCode === lang
    )!.Description;
    page_text.value.button_text = text_translations.value.find(
      (text) => text.LanguageCode === lang
    )!.ButtonText;
  },
  {
    immediate: true,
  }
);

const translations = computed(() => {
  return all_languages.value.find(
    (lang) =>
      lang.current_locale === current_language.value ||
      lang.locales.includes(current_language.value)
  );
});
</script>

<template>
  <main
    class="w-full max-w-5xl min-w-[648px] mx-auto flex flex-col px-8 pt-4 gap-4 pb-4"
  >
    <h4 class="scroll-m-20 text-xl font-semibold tracking-tight">
      {{ page_text.description }}
    </h4>
    <draggable
      :list="chosen_languages"
      :itemKey="getItemKey"
      class="rounded-md border p-6 flex flex-col gap-3"
      ghost-class="ghost"
    >
      <template #item="{ element: lang_code }">
        <Card class="transition-all cursor-move flex flex-row">
          <CardHeader>
            <CardTitle>{{
              translations?.translations.find((lang) => lang.code === lang_code)
                ?.translation
            }}</CardTitle>
          </CardHeader>
          <Select
            :default-value="lang_code"
            v-on:update:model-value="updateLocale($event, lang_code)"
          >
            <SelectTrigger class="max-w-xs ml-auto my-auto mr-3">
              <SelectValue placeholder="Select a language" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectItem
                  v-for="locale in all_languages.find((lang) =>
                    lang.locales.includes(lang_code)
                  )?.locales"
                  :value="locale"
                  :key="getItemKey"
                >
                  {{ locale }}
                </SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
          <Button
            @click="deleteLanguage(lang_code)"
            class="my-auto mr-4"
            variant="outline"
            size="icon"
          >
            <Trash class="w-4 h-4" />
          </Button>
        </Card>
      </template>
    </draggable>
    <!-- MARK: Add Languages Dialog / Modal -->
    <Dialog @update:open="handleClose">
      <!-- v-on:update:open="(open) => (open ? '' : (search = ''))" -->
      <DialogTrigger as-child>
        <Button>
          {{ page_text.button_text }}
        </Button>
      </DialogTrigger>
      <DialogContent
        class="sm:max-w-2xl grid-rows-[auto_minmax(0,1fr)_auto] p-0 max-h-[90dvh]"
      >
        <DialogHeader class="p-6 pb-0 flex flex-row justify-between">
          <DialogTitle class="text-2xl">{{
            page_text.button_text
          }}</DialogTitle>
          <div class="relative w-full max-w-xs items-center mr-8">
            <Input
              id="search"
              type="text"
              placeholder="Search..."
              class="pl-10"
              v-model:model-value="search"
            />
            <span
              class="absolute start-0 inset-y-0 flex items-center justify-center px-2"
            >
              <Search class="size-6 text-muted-foreground" />
            </span>
          </div>
        </DialogHeader>
        <div class="flex flex-col gap-4 py-4 overflow-y-auto px-6">
          <Card
            v-for="(language, i) in all_languages
              .filter((lang) => {
                return !lang.locales.some((item: string) =>
                  chosen_languages.includes(item)
                );
              })
              .filter(
                (lang) =>
                  lang.name.toLowerCase().includes(search.toLowerCase()) ||
                  lang.current_locale
                    .toLowerCase()
                    .includes(search.toLowerCase()) ||
                  lang.locales.some((locale: string) =>
                    locale.toLowerCase().includes(search.toLowerCase())
                  )
              )
              .sort((a, b) => a.name.localeCompare(b.name))"
            :key="i"
            class="transition-all flex flex-row"
          >
            <Checkbox
              @update:checked="updateChecked($event, language)"
              :id="language.name"
              class="my-auto ml-6"
            />
            <label :for="language.name" class="w-full">
              <CardHeader>
                <CardTitle class="text-xl font-medium">{{
                  translations?.translations.find(
                    (lang) => lang.code === language.current_locale
                  )?.translation
                }}</CardTitle>
              </CardHeader>
            </label>
          </Card>
        </div>
        <DialogFooter class="p-6 pt-0">
          <DialogClose>
            <Button variant="outline" @click="addLanguages"
              >Add languages</Button
            >
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  </main>
</template>

<style scoped>
.ghost {
  opacity: 1;
  @apply bg-accent;
}
</style>
