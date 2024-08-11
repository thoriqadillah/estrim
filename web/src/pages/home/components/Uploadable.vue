<script setup lang="ts">
import { Button } from '@/components/ui/button';
import { Icon } from '@iconify/vue';
import type { UploadableFile } from '../types';

defineProps<{
  type: string
  file: UploadableFile
}>()

const emits = defineEmits<{
  (e: 'cancel'): void
}>()

</script>

<template>
  <div class="h-28 flex flex-col relative aspect-square border rounded-lg bg-background p-2" >
    <Icon v-if="type === 'image'" icon="fluent:image-32-regular"  class="text-muted-foreground mx-auto" width="70" />
    <Icon v-else-if="type === 'pdf'" icon="fluent:document-pdf-32-regular" class="text-muted-foreground mx-auto" width="70" />
    <Icon v-else-if="type === 'video'" icon="fluent:video-32-regular" class="text-muted-foreground mx-auto" width="70" />
    <p class="text-xs max-w-24 truncate text-left mt-2 font-medium">{{ file.name }}</p>
    <p class="text-muted-foreground text-xs max-w-20 truncate text-left">{{ file.size }} KB</p>
    <Button 
      variant="ghost" 
      size="icon" 
      class="absolute w-6 h-6 top-1 right-1" 
      @click="(e) => {
        e.stopPropagation()
        emits('cancel')
      }"
    >
      <Icon icon="mdi:close" />
    </Button>
  </div>
</template>