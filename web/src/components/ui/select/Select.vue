<script setup lang="ts">
import {
  SelectContainer,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
  type Option,
} from '.'
import { cn } from '../shadcn';

const props = defineProps<{
  label: string
  options: Option[] | string[]
  class?: string,
  placeholder?: string
}>()

const model = defineModel<Option | string>()
const emits = defineEmits<{
  (e: 'select', option: Option | string): void,
}>()

function select(option: Option | string) {
  if (typeof option !== 'string' && option.disabled) {
    return
  }

  model.value = option
  emits('select', option)
}

function toValue(option: Option | string) {
  return typeof option === 'string' ? option : option.value
}

function toLabel(option: Option | string) {
  return typeof option === 'string' ? option : option.label
}

function isDisabled(option: Option | string): boolean {
  return typeof option === 'string' ? false : (option.disabled || false)
}

</script>

<template>
  <SelectContainer :defaultValue="model ? toValue(model) : undefined">
    <SelectTrigger :class="cn('border border-input', props.class)">
      <slot>
        <SelectValue :placeholder="model ? toLabel(model) : placeholder" class="mr-2" />
      </slot>
    </SelectTrigger>
    <SelectContent>
      <SelectGroup>
        <SelectLabel>{{ label }}</SelectLabel>
        <SelectItem
          v-for="option in options" 
          ref="r"
          :value="toValue(option)"
          :key="toValue(option)"
          :disabled="isDisabled(option)"
          @click="() => select(option)"
          class="cursor-pointer">
          {{ toLabel(option) }}
        </SelectItem>
      </SelectGroup>
    </SelectContent>
  </SelectContainer>
</template>