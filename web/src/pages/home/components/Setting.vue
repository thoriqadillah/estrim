<script setup lang="ts">
import { Tooltip } from '@/components/ui/tooltip';

defineProps<{
  title: string
  tooltip?: string
}>()
</script>

<template>
  <div class="flex justify-between items-center py-3 px-1 border-b">
    <span class="flex gap-2 items-center">
      <p>{{ title }}</p>
      <Tooltip :text="tooltip">
        <img v-if="tooltip" src="https://api.iconify.design/fluent:info-24-regular.svg?color=%23888888">
      </Tooltip>
    </span>
    <slot />
  </div>
</template>