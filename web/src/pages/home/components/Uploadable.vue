<script setup lang="ts">
import { Icon } from '@iconify/vue';

defineProps<{
  type: string
  file: File
}>()
</script>

<template>
  <div class="h-28 flex flex-col relative aspect-square border rounded-lg bg-background p-2" >
    <Icon v-if="type === 'image'" icon="fluent:image-32-regular"  class="text-muted-foreground mx-auto" width="70" />
    <Icon v-else-if="type === 'pdf'" icon="fluent:document-pdf-32-regular" class="text-muted-foreground mx-auto" width="70" />
    <Icon v-else-if="type === 'video'" icon="fluent:video-32-regular" class="text-muted-foreground mx-auto" width="70" />
    <p class="text-xs max-w-24 truncate text-left mt-2 font-medium">{{ file.name }}</p>
    <p class="text-muted-foreground text-xs max-w-20 truncate text-left">{{ file.size }} KB</p>
  </div>
</template>