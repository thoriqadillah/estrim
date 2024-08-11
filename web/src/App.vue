<script setup lang="ts">
import { RouterView } from 'vue-router'
import Nav from './components/Nav.vue';
import Main from './components/Main.vue';
import { eventBus } from './lib';
import { Toaster, useToast } from './components/ui/toast';
import { onMounted, onUnmounted } from 'vue';

const { toast } = useToast()

onMounted(() => {
  eventBus.on<any>('toast', toast)
})

onUnmounted(() => {
  eventBus.off<any>('toast', toast)
})

</script>

<template>
  <Toaster />
  <Nav />
  <Main>
    <RouterView />
  </Main>
</template>