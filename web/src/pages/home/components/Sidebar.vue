<script setup lang="ts">
import { Button } from '@/components/ui/button';
import { eventBus } from '@/lib';
import { Icon } from '@iconify/vue';
import { onMounted, onUnmounted, ref } from 'vue';
import FileItem from './FileItem.vue';
import { Skeleton } from '@/components/ui/skeleton';

const open = defineModel('open', { default: false });
const files = ref([])
const loading = ref(true)

async function getAllFiles() {
  loading.value = false
}

function appendFile(fileId?: string) {
  if (!open.value) open.value = true
  // TODO
}

onMounted(async () => {
  await getAllFiles()
  eventBus.on('uploaded', appendFile)
})

onUnmounted(() => {
  eventBus.off('uploaded', appendFile)
})

</script>

<template>
  <div class="wrapper flex">
    <div ref="sidebar" :class="['sidebar', { open }, 'flex flex-col gap-2 relative']">
      <Skeleton v-if="loading" v-for="_ in Array(3)" class="w-full h-[4rem] rounded-lg" />
      <div class="absolute w-[15rem] left-1/2 -translate-x-1/2 top-1/2 -translate-y-1/2 text-center">
        <Icon icon="fluent:document-question-mark-16-regular" width="50" class="text-muted-foreground opacity-50 mx-auto"/>
        <p class="text-xs text-muted-foreground opacity-50 mt-1">Your compressed file will appear here</p>
      </div>
      <FileItem 
        v-if="files.length > 0" 
        v-for="file in files" 
        class="w-[15rem]"
        id="" 
        type="image" 
        name="Singo" 
        :size="9287348" 
      />
    </div>
    <Button  @click="() => open = !open" :class="open && 'ml-2'" size="icon" variant="ghost">
      <Icon class="text-xl" :icon="!open ? 'fluent:panel-left-expand-24-regular' : 'fluent:panel-right-expand-24-regular'"/>
    </Button>
  </div>
</template>

<style>
.wrapper {
  transition: width 0.3s ease;
}

.sidebar {
  width: 0;
  overflow: hidden;
  transition: width 0.3s ease;
}

.sidebar.open {
  width: 15rem;
}
</style>