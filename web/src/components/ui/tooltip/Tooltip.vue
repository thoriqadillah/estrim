<script setup lang="ts">
import {
  TooltipContainer,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger
} from '@/components/ui/tooltip'

defineProps<{
  text?: string
  side?: 'top' | 'right' | 'bottom' | 'left'
}>()
</script>

<template>
  <TooltipProvider>
    <TooltipContainer>
      <TooltipTrigger>
        <slot/>
      </TooltipTrigger>
      <TooltipContent v-if="text" :side="side">
        <p>{{ text }}</p>
      </TooltipContent>
    </TooltipContainer>
  </TooltipProvider>
</template>