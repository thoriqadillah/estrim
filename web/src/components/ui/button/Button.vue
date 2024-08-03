<script setup lang="ts">
import type { HTMLAttributes } from 'vue'
import { Primitive, type PrimitiveProps } from 'radix-vue'
import { type ButtonVariants, buttonVariants } from '.'
import { cn } from '@/components/ui/shadcn'

interface Props extends PrimitiveProps {
  variant?: ButtonVariants['variant']
  size?: ButtonVariants['size']
  class?: HTMLAttributes['class'],
  appendIcon?: string
  prependIcon?: string
}

const props = withDefaults(defineProps<Props>(), {
  as: 'button',
})
</script>

<template>
  <Primitive
    :as="as"
    :as-child="asChild"
    :class="cn('flex gap-2', buttonVariants({ variant, size }), props.class)"
  >
    <slot name="prependIcon">
      <img v-if="prependIcon" :src="prependIcon" alt="">
    </slot>
    <slot />
    <slot name="appendIcon">
      <img v-if="appendIcon" :src="appendIcon" alt="">
    </slot>
  </Primitive>
</template>