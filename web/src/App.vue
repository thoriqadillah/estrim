<script setup lang="ts">
import { RouterView } from 'vue-router'
import Nav from './components/Nav.vue';
import Main from './components/Main.vue';
</script>

<template>
  <Nav />
  <Main>
    <RouterView />
  </Main>
</template>