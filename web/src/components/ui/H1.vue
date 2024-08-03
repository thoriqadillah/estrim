<script setup lang="ts">
import { cn } from './shadcn';
import { type PrimitiveProps } from 'radix-vue';

interface Props extends PrimitiveProps {
  class?: string
}

const props = withDefaults(defineProps<Props>(), {
  as: 'h1'
})

</script>

<template>
  <h1 :class="cn('text-5xl font-bold tracking-tighter', props.class)">
    <slot/>
  </h1>
</template>