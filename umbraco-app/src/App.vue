<script setup lang="ts">
import { ref } from "vue";
import DragList from "./DragList.vue";

const description = ref("This is a description");
let current_language = "";

function onlanguagechange(lang: string) {
  if (lang !== current_language) {
    current_language = lang;
    description.value = "This is a description in " + lang;
  }
}
</script>

<template>
  <main
    class="w-full max-w-5xl min-w-[648px] mx-auto flex flex-col pt-4 gap-2 pb-4"
  >
    <h4 class="scroll-m-20 text-xl font-semibold tracking-tight">
      {{ description }}
    </h4>
    <DragList @lang_change="onlanguagechange" />
  </main>
</template>

<style></style>
