<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { Button } from '@/components/ui/button'
</script>

<template>
  <main>
    <p>Home</p>
    <RouterLink to="/about">
      <Button>About</Button>
    </RouterLink>    
  </main>
</template>
