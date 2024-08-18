<script setup lang="ts">
import { cn } from '@/components/ui/shadcn';
import { Tooltip } from '@/components/ui/tooltip';
import { Icon } from '@iconify/vue';

const props = defineProps<{
  id: string
  type: string
  name: string
  size: number
  progress?: number
  status?: string,
  class?: string
}>()
</script>

<template>
  <div :class="cn('w-full h-[4rem] rounded-lg border border-muted bg-muted/60 p-2 flex items-center relative', props.class)">
    <div class="flex gap-2 items-center justify-between w-full">
      <Icon v-if="type === 'image'" icon="fluent:image-32-regular"  class="text-foreground" width="40" />
      <Icon v-else-if="type === 'pdf'" icon="fluent:document-pdf-32-regular" class="text-foreground" width="40" />
      <Icon v-else-if="type === 'video'" icon="fluent:video-32-regular" class="text-foreground" width="40" />
  
      <div class="flex flex-col w-full">
        <p class="font-medium text-sm line-clamp-1">{{ name }}</p>
        <p class="text-xs line-clamp-1 text-muted-foreground">{{ size }} KB</p>
      </div>

      <Tooltip v-if="progress && progress < 100" text="Cancel" side="right">
        <Button variant="ghost" class="mr-2 group">
          <Icon icon="fluent:presence-blocked-16-regular" class="text-destructive group-hover:text-opacity-70"/>
        </Button>
      </Tooltip>

      <Tooltip v-else text="Success" side="right">
        <Button disabled variant="ghost" class="mr-2">
          <Icon icon="fluent:checkmark-16-filled" class="text-success"/>
        </Button>
      </Tooltip>
    </div>
    <div v-if="progress && progress < 100" class="h-[4rem] -z-10 w-full bg-success/10 absolute left-0 rounded-lg" :style="`width: ${progress || 0}%;`" /> 
  </div>
</template>