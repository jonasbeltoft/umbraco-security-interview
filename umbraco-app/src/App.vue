<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { onMounted } from "vue";

function handleClick() {
  document.title = "hello there";
}

onMounted(() => {
  document.title = "Beginning";
});
</script>
<template>
  <div>
    <h1>Bubba</h1>
    <Button @click="handleClick()">Hello there</Button>
  </div>
</template>
