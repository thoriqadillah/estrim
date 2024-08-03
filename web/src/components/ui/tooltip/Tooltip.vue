<script setup lang="ts">
import {
  TooltipContainer,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger
} from '@/components/ui/tooltip'

defineProps<{
  text?: string
}>()
</script>

<template>
  <TooltipProvider>
    <TooltipContainer>
      <TooltipTrigger>
        <slot/>
      </TooltipTrigger>
      <TooltipContent v-if="text">
        <p>{{ text }}</p>
      </TooltipContent>
    </TooltipContainer>
  </TooltipProvider>
</template>