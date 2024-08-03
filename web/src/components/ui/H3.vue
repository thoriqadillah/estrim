<script setup lang="ts">
import { cn } from './shadcn';
import { Primitive, type PrimitiveProps } from 'radix-vue';

interface Props extends PrimitiveProps {
  class?: string
}

const props = withDefaults(defineProps<Props>(), {
  as: 'h1'
})

</script>

<template>
  <Primitive :class="cn('text-xl font-medium tracking-tight', props.class)">
    <slot/>
  </Primitive>
</template>