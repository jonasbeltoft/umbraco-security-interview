<script setup lang="ts">
import { ref } from "vue";
import DragList from "./DragList.vue";
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

const page_text = useStorage("page_text", {
  button_text: "Tilføg sprog",
  description: "This is a description",
});

let current_language = "";
let all_languages = ref([
  {
    name: "Danish",
    chosen: true,
    current_locale: "da",
    locales: ["da", "da-DK"],
  },
  {
    name: "English",
    chosen: true,
    current_locale: "en",
    locales: ["en", "en-GB"],
  },
  {
    name: "German",
    chosen: false,
    current_locale: "de",
    locales: ["de", "de-DE"],
  },
  {
    name: "Spanish",
    chosen: false,
    current_locale: "es",
    locales: ["es", "es-ES"],
  },
  {
    name: "French",
    chosen: false,
    current_locale: "fr",
    locales: ["fr", "fr-FR"],
  },
  {
    name: "Italian",
    chosen: false,
    current_locale: "it",
    locales: ["it", "it-IT"],
  },
  {
    name: "Dutch",
    chosen: false,
    current_locale: "nl",
    locales: ["nl", "nl-NL"],
  },
  {
    name: "Norwegian",
    chosen: false,
    current_locale: "no",
    locales: ["no", "no-NO"],
  },
  {
    name: "Swedish",
    chosen: false,
    current_locale: "sv",
    locales: ["sv", "sv-SE"],
  },
  {
    name: "Finnish",
    chosen: false,
    current_locale: "fi",
    locales: ["fi", "fi-FI"],
  },
  {
    name: "Portuguese",
    chosen: false,
    current_locale: "pt",
    locales: ["pt", "pt-PT"],
  },
  {
    name: "Polish",
    chosen: false,
    current_locale: "pl",
    locales: ["pl", "pl-PL"],
  },
  {
    name: "Chinese",
    chosen: false,
    current_locale: "zh",
    locales: ["zh", "zh-CN"],
  },
  {
    name: "Japanese",
    chosen: false,
    current_locale: "ja",
    locales: ["ja", "ja-JP"],
  },
  {
    name: "Korean",
    chosen: false,
    current_locale: "ko",
    locales: ["ko", "ko-KR"],
  },
  {
    name: "Arabic",
    chosen: false,
    current_locale: "ar",
    locales: ["ar", "ar-AR"],
  },
  {
    name: "Turkish",
    chosen: false,
    current_locale: "tr",
    locales: ["tr", "tr-TR"],
  },
  {
    name: "Greek",
    chosen: false,
    current_locale: "el",
    locales: ["el", "el-EL"],
  },
  {
    name: "Czech",
    chosen: false,
    current_locale: "cs",
    locales: ["cs", "cs-CZ"],
  },
]);

// Load chosen languages from local storage
const chosen_languages = useStorage(
  "chosen_languages",
  all_languages.value
    .filter((lang) => lang.chosen)
    .map((lang) => lang.current_locale)
);

let search = ref("");

function onlanguagechange(lang: string) {
  if (lang !== current_language) {
    current_language = lang;
    page_text.value.description = "This is a description in " + lang;
  }
}

function addLanguages() {
  console.log("Adding languages");
}
</script>

<template>
  <main
    class="w-full max-w-5xl min-w-[648px] mx-auto flex flex-col px-8 pt-4 gap-4 pb-4"
  >
    <h4 class="scroll-m-20 text-xl font-semibold tracking-tight">
      {{ page_text.description }}
    </h4>
    <DragList
      v-bind:current_languages="chosen_languages"
      v-bind:all_languages="all_languages"
      @lang_change="onlanguagechange"
    />
    <!-- MARK: Add Languages Dialog / Modal -->
    <Dialog>
      <!-- v-on:update:open="(open) => (open ? '' : (search = ''))" -->
      <DialogTrigger as-child>
        <Button>
          {{ page_text.button_text }}
        </Button>
      </DialogTrigger>
      <DialogContent
        class="sm:max-w-3xl grid-rows-[auto_minmax(0,1fr)_auto] p-0 max-h-[90dvh]"
      >
        <DialogHeader class="p-6 pb-0 flex flex-row justify-between">
          <DialogTitle class="text-2xl">{{
            page_text.button_text + " " + search
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
              .filter((lang) => !lang.chosen)
              .filter(
                (lang) =>
                  lang.name.toLowerCase().includes(search.toLowerCase()) ||
                  lang.current_locale
                    .toLowerCase()
                    .includes(search.toLowerCase()) ||
                  lang.locales.some((locale) =>
                    locale.toLowerCase().includes(search.toLowerCase())
                  )
              )"
            :key="i"
            class="transition-all flex flex-row"
          >
            <CardHeader>
              <CardTitle class="text-xl font-medium">{{
                language.name
              }}</CardTitle>
            </CardHeader>
            <Checkbox :id="language.name" class="my-auto" />
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

<style></style>
