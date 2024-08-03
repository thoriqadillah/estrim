<script setup lang="ts">
import { Button } from '@/components/ui/button'
import H1 from '@/components/ui/H1.vue';
import Dropzone from './components/Dropzone.vue';
import { ref } from 'vue';
import { Collapsible, CollapsibleContent } from '@/components/ui/collapsible'
import H3 from '@/components/ui/H3.vue';
import Setting from './components/Setting.vue';
import { Input } from '@/components/ui/input';

const input = ref<HTMLInputElement>()
const collapsed = ref(false)

const ratio = ref(0.6)
</script>

<template>
  <div class="w-full my-20">
    <div class="text-center">
      <H1>Estrim</H1>
      <p>Your background file size trimmer</p>
    </div>

    <div class="mx-auto my-5 flex flex-col gap-2 max-w-2xl">
      <input type="file" hidden ref="input">
      <Dropzone 
        #default="{ active }"
        class="text-center relative border border-muted-foreground rounded-xl cursor-pointer h-[10rem] data-[active:true]:border-primary w-full"
        @click="input?.click()">

      </Dropzone>

      <div class="flex gap-2 justify-between">
        <div class="flex gap-2">
          <Button>Image</Button>
          <Button variant="ghost">PDF</Button>
          <Button variant="ghost">Video</Button>
        </div>
        
        <Button @click="collapsed = !collapsed" variant="ghost">
          <p>Settings</p>
          <template #appendIcon>
            <img 
              :class="`transition-all ${collapsed ? 'rotate-180' : ''}`" 
              src="https://api.iconify.design/fluent:arrow-curve-down-right-20-filled.svg?color=%23000"
            >
          </template>
        </Button>
      </div>

      <Collapsible :open="collapsed" class="mt-5">
        <CollapsibleContent>
          <H3>Settings</H3>

          <div class="flex flex-col">
            <Setting title="Compression ratio">
              <Input type="number" class="max-w-[6rem]" v-model="ratio" />
            </Setting>
            <Setting title="Convert to">
              
            </Setting>
          </div>
        </CollapsibleContent>
      </Collapsible>
    </div>
  </div>
</template>
