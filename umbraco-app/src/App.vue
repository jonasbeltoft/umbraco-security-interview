<script setup lang="ts">
import DragList from "./DragList.vue";

let description = "This is a description";
</script>

<template>
  <main class="w-full max-w-5xl min-w-[648px] mx-auto flex flex-col pt-4 gap-2">
    <h4 class="scroll-m-20 text-xl font-semibold tracking-tight">
      {{ description }}
    </h4>
    <DragList />
  </main>
</template>

<style></style>
