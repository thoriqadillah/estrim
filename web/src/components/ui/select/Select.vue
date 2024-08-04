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

defineProps<{
  label: string
  options: Option[]
  class?: string,
  placeholder?: string
}>()

const model = defineModel<string>()
const emits = defineEmits<{
  (e: 'select', option: Option): void
}>()


function select(option: Option) {
  if (option.disabled) return

  model.value = option.value
  emits('select', option)
}

</script>

<template>
  <SelectContainer plac>
    <SelectTrigger :class="class">
      <slot>
        <SelectValue :placeholder="placeholder" class="mr-2" />
      </slot>
    </SelectTrigger>
    <SelectContent>
      <SelectGroup>
        <SelectLabel>{{ label }}</SelectLabel>
        <SelectItem
          v-for="option in options" 
          :value="option.value"
          :key="option.value"
          @click="() => select(option)"
          :disabled="option.disabled">
          {{ option.label }}
        </SelectItem>
      </SelectGroup>
    </SelectContent>
  </SelectContainer>
</template>