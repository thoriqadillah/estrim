<script setup lang="ts">
import { cn } from '@/components/ui/shadcn';
import type { PrimitiveProps } from 'radix-vue';
import { ref } from 'vue';
import type { UploadableFile } from '../types';
import { eventBus, api } from '@/lib';

const emits = defineEmits<{
  (e: 'change', file: UploadableFile[]): void,
}>()

interface Props extends PrimitiveProps {
  class?: string
  type?: string,
}

const props = withDefaults(defineProps<Props>(), {
  as: 'div'
})

const active = ref(false)

function onDrop(e: DragEvent) {
  active.value = false

  if (e.dataTransfer) {
    const uploadables: UploadableFile[] = []
    for (const file of e.dataTransfer.files) {
      uploadables.push({
        name: file.name,
        size: file.size,
      })
    }

    emits('change', [...uploadables])
    files.value.push(...e.dataTransfer.files)
  }
}

const input = ref<HTMLInputElement>()
const model = defineModel<UploadableFile[]>({
  default: []
})

const files = ref<File[]>([])

function fileChange(e: Event) {
  const fileinput = e.target as HTMLInputElement
  if (!fileinput.files) return
  
  const uploadables: UploadableFile[] = []
  for (const file of fileinput.files) {
    uploadables.push({
      name: file.name,
      size: file.size,
    })
  }
  
  files.value.push(...fileinput.files)
  model.value.push(...uploadables)
  emits('change', [...uploadables])
}

async function upload() {
  files.value.forEach(async (file, i) => {
    const form = new FormData()
    form.append('file', file)
    form.append('type', props.type || 'image')
  
    const uploaded = await api.compress(form, progress => {
      model.value[i].progress = progress
    })
  
    if (uploaded) {
      eventBus.emit('uploaded', uploaded.id)
      model.value = model.value.filter(el => el.progress !== 100)
    }
  })
}

defineExpose({ upload })

</script>

<template>
  <div 
  :class="cn(props.class)" 
  :data-active="active" 
  @dragenter.prevent
  @dragover.prevent="() => active = true"
  @dragleave="() => active = false"
  @drop.prevent="onDrop"
  @click="input?.click"
  >
    <input type="file" @change="fileChange" multiple hidden ref="input">
    <slot :active="active" />
  </div>
</template>