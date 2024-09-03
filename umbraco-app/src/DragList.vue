<script setup lang="ts">
import { Card, CardTitle, CardHeader } from "./components/ui/card";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

import draggable from "vuedraggable";
import { onBeforeMount, ref, watch, computed, toRaw } from "vue";
import type { Language } from "./lib/types";

const chosen_languages = ref<Language[]>([
  { name: "Danish", current_locale: "da", locales: ["da", "da-DK"] },
  { name: "English", current_locale: "en", locales: ["en", "en-GB"] },
  { name: "German", current_locale: "de", locales: ["de", "de-DE"] },
  { name: "Spanish", current_locale: "es", locales: ["es", "es-ES"] },
  { name: "French", current_locale: "fr", locales: ["fr", "fr-FR"] },
  { name: "Italian", current_locale: "it", locales: ["it", "it-IT"] },
  { name: "Dutch", current_locale: "nl", locales: ["nl", "nl-NL"] },
  { name: "Norwegian", current_locale: "no", locales: ["no", "no-NO"] },
  { name: "Swedish", current_locale: "sv", locales: ["sv", "sv-SE"] },
  { name: "Finnish", current_locale: "fi", locales: ["fi", "fi-FI"] },
  { name: "Portuguese", current_locale: "pt", locales: ["pt", "pt-PT"] },
  { name: "Russian", current_locale: "ru", locales: ["ru", "ru-RU"] },
  { name: "Polish", current_locale: "pl", locales: ["pl", "pl-PL"] },
  { name: "Chinese", current_locale: "zh", locales: ["zh", "zh-CN"] },
  { name: "Japanese", current_locale: "ja", locales: ["ja", "ja-JP"] },
  { name: "Korean", current_locale: "ko", locales: ["ko", "ko-KR"] },
  { name: "Arabic", current_locale: "ar", locales: ["ar", "ar-AR"] },
  { name: "Turkish", current_locale: "tr", locales: ["tr", "tr-TR"] },
  { name: "Greek", current_locale: "el", locales: ["el", "el-EL"] },
  { name: "Czech", current_locale: "cs", locales: ["cs", "cs-CZ"] },
  { name: "Hungarian", current_locale: "hu", locales: ["hu", "hu-HU"] },
]);

const emit = defineEmits(["lang_change"]);

function setStorage() {
  localStorage.setItem(
    "chosen_languages",
    JSON.stringify(chosen_languages.value)
  );
}

// Load chosen languages from local storage
onBeforeMount(() => {
  const langs = localStorage.getItem("chosen_languages");
  if (langs) {
    chosen_languages.value = JSON.parse(langs);
  } else {
    // If no languages are chosen, set Danish as default
    chosen_languages.value = [
      { name: "Danish", current_locale: "da", locales: ["da", "da-DK"] },
    ];
    setStorage();
  }
});

function updateLocale(locale: string) {
  chosen_languages.value.forEach((lang) => {
    if (lang.locales.includes(locale)) {
      lang.current_locale = locale;
    }
  });
  setStorage();
}

// Emit event when language is changed
watch(
  () => chosen_languages,
  (lang, prevLang) => {
    // Check if the current locale has changed
    emit("lang_change", lang.value[0].current_locale);
  },
  { deep: true }
);
</script>

<template>
  <draggable
    :list="chosen_languages"
    item-key="current_locale"
    class="rounded-md border p-6 flex flex-col gap-3"
    ghost-class="ghost"
    @end="setStorage()"
  >
    <template #item="{ element: language }">
      <Card class="transition-all cursor-move flex flex-row">
        <CardHeader>
          <CardTitle>{{ language.name }}</CardTitle>
        </CardHeader>
        <Select
          :default-value="language.current_locale"
          v-on:update:model-value="updateLocale"
        >
          <SelectTrigger class="max-w-xs ml-auto my-auto mr-3">
            <SelectValue placeholder="Select a language" />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem v-for="locale in language.locales" :value="locale">
                {{ locale }}
              </SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </Card>
    </template>
  </draggable>
</template>

<style scoped>
.ghost {
  opacity: 1;
  @apply bg-accent;
}
</style>
