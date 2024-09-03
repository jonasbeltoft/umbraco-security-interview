<script>
import {
  Card,
  CardTitle,
  CardHeader,
  CardDescription,
} from "./components/ui/card";

import draggable from "vuedraggable";

let id = 1;
export default {
  name: "simple",
  display: "Simple",
  order: 0,
  components: {
    draggable,
    Card,
    CardTitle,
    CardHeader,
    CardDescription,
  },
  data() {
    return {
      enabled: true,
      list: [
        {
          name: "English",
          code: "en",
        },
        {
          name: "Spanish",
          code: "es",
        },
        {
          name: "French",
          code: "fr",
        },
      ],
      dragging: false,
    };
  },
};
</script>

<template>
  <draggable
    :list="list"
    item-key="code"
    class="rounded-md border p-6 flex flex-col gap-3"
    ghost-class="ghost"
    @start="dragging = true"
    @end="dragging = false"
  >
    <template #item="{ element: language }">
      <Card class="transition-all cursor-move">
        <CardHeader>
          <CardTitle>{{ language.name }}</CardTitle>
          <CardDescription>{{ language.code }}</CardDescription>
        </CardHeader>
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
