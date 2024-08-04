<script setup lang="ts">
import { Button } from '@/components/ui/button'
import H1 from '@/components/ui/H1.vue';
import Dropzone from './components/Dropzone.vue';
import { ref, watch } from 'vue';
import { Collapsible, CollapsibleContent } from '@/components/ui/collapsible'
import H3 from '@/components/ui/H3.vue';
import Setting from './components/Setting.vue';
import { Input } from '@/components/ui/input';
import { Select, type Option } from '@/components/ui/select';
import { Icon } from '@iconify/vue';
import Uploadable from './components/Uploadable.vue';

const collapsed = ref(false)
const ratio = ref(0.6)

const ext = ref<Option>({ label: 'WEBP', value: 'webp' })
const options = [
  { label: 'WEBP', value: 'webp' },
  { label: 'JPEG', value: 'jpeg' },
  { label: 'PNG', value: 'png' },
]

const type = ref<Option>({ label: 'Image', value: 'image' })
const types = [
  { label: 'Image', value: 'image' },
  { label: 'PDF', value: 'pdf', disabled: true },
  { label: 'Video', value: 'video', disabled: true },
]

watch(type, () => targets.value = [])

const targets = ref<File[]>([])

</script>

<template>
  <div class="w-full my-20">
    <div class="text-center">
      <H1>Estrim</H1>
      <p>Your background file size trimmer</p>
    </div>

    <div class="mx-auto mt-5 flex flex-col gap-2 max-w-2xl">
      <Dropzone 
        v-model="targets"
        #default="{ active }"
        :type="type.value"
        class="text-center relative border bg-muted bg-opacity-25 border-muted-foreground rounded-xl cursor-pointer min-h-[10rem] data-[active:true]:border-primary data-[active:true]:bg-opacity-80 w-full">
        <p v-if="!active && targets.length === 0" class="text-muted-foreground text-sm absolute top-1/2 -translate-y-1/2 left-1/2 -translate-x-1/2">
          Drag and drop your file here
        </p>
        <p v-else-if="active && targets.length === 0" class="text-muted-foreground text-sm absolute top-1/2 -translate-y-1/2 left-1/2 -translate-x-1/2">
          Release your file
        </p>
        <div v-else class="flex flex-wrap gap-3 p-3">
          <Uploadable v-for="(target, i) in targets" :file="target" :key="i" :type="type.value" />
        </div>
      </Dropzone>

      <div class="flex gap-2 justify-between">
        <div class="flex gap-2">
          <Button>Trim</Button>
          <Select :options="types" label="File type" class="max-w-[20rem]">
            <p class="w-[5rem] text-left">{{ type.label || 'Select type' }}</p>
          </Select>
        </div>
        
        <Button @click="collapsed = !collapsed" variant="ghost">
          <p>Settings</p>
          <template #appendIcon>
            <Icon icon="fluent:arrow-curve-down-right-20-filled" :class="`text-foreground transition-all ${collapsed ? 'rotate-180' : ''}`" />
          </template>
        </Button>
      </div>

      <Collapsible :open="collapsed">
        <CollapsibleContent class="mt-5">
          <H3>Settings</H3>

          <div class="flex flex-col">
            <Setting title="Compression ratio">
              <Input type="number" class="max-w-[6rem]" v-model="ratio" />
            </Setting>
            <Setting title="Convert to">
              <Select :options="options" label="Extension" class="max-w-[10rem]">
                {{ ext.label || 'Select extension' }}
              </Select>
            </Setting>
          </div>
        </CollapsibleContent>
      </Collapsible>
    </div>
  </div>
</template>
