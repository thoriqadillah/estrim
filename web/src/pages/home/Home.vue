<script setup lang="ts">
import { Button } from '@/components/ui/button'
import H1 from '@/components/ui/H1.vue';
import Dropzone from './components/Dropzone.vue';
import { ref } from 'vue';

const input = ref<HTMLInputElement>()
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
      @click="input?.click()"
      >
    </Dropzone>

    <div class="flex gap-2 justify-between">
      <div>
        <Button>Image</Button>
        <Button variant="ghost">PDF</Button>
        <Button variant="ghost">Video</Button>
      </div>
      
      <Button variant="ghost" append-icon="https://api.iconify.design/fluent:arrow-curve-down-right-20-filled.svg?color=%23000">
        Setting
      </Button>
    </div>
    </div>
  </div>
</template>
