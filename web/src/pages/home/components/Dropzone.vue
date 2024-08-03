<script setup lang="ts">
import { cn } from '@/components/ui/shadcn';
import type { PrimitiveProps } from 'radix-vue';
import { ref } from 'vue';

const emits = defineEmits<{
  (e: 'droped', file: File[]): void
}>()

interface Props extends PrimitiveProps {
  class?: string
}

const props = withDefaults(defineProps<Props>(), {
  as: 'div'
})

const active = ref(false)

function activate() {
  active.value = true
}

function deactivate() {
  active.value = true
}

function onDrop(event: DragEvent) {
  deactivate()
  if (event.dataTransfer) {
    emits('droped', [...event.dataTransfer.files])
  }
}

</script>

<template>
  <div 
    :class="cn(props.class)" 
    :data-active="active" 
    @dragenter.prevent="activate"
    @dragleave.prevent="deactivate"
    @drop.prevent="onDrop"
  >
    <slot :active="active" />
  </div>
</template>