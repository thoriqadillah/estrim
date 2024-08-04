<script setup lang="ts">
import { cn } from '@/components/ui/shadcn';
import type { PrimitiveProps } from 'radix-vue';
import { ref } from 'vue';

const emits = defineEmits<{
  (e: 'change', file: File[]): void,
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
    emits('change', [...e.dataTransfer.files])
    files.value.push(...e.dataTransfer.files)
  }
}

const input = ref<HTMLInputElement>()
const files = defineModel<File[]>({
  default: []
})

function fileChange(e: Event) {
  const fileinput = e.target as HTMLInputElement
  if (!fileinput.files) return
  
  files.value.push(...fileinput.files)
  emits('change', [...fileinput.files])
}

</script>

<template>
  <div 
  :class="cn(props.class)" 
  :data-active="active" 
  @dragenter.prevent
  @dragover.prevent="() => active = true"
  @dragleave="() => active = false"
  @drop.prevent="onDrop"
  @click="input?.click()"
  >
    <input type="file" @change="fileChange" multiple hidden ref="input">
    <slot :active="active" />
  </div>
</template>