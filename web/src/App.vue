<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router'
import Nav from './components/Nav.vue';
import Main from './components/Main.vue';
import { eventBus } from './lib';
import { Routes } from './router';
import { Toaster, useToast } from './components/ui/toast';
import { onMounted, onUnmounted } from 'vue';

const router = useRouter()
const { toast } = useToast()

const logout = () => router.push({ name: Routes.Login })

onMounted(() => {
  eventBus.on('logout', logout)
  eventBus.on<any>('toast', toast)
})

onUnmounted(() => {
  eventBus.off('logout', logout)
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