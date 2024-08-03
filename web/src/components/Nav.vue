<script setup lang="ts">
import { RouterLink, useRoute } from 'vue-router';
import { Button } from './ui/button';
import { Routes } from '@/router';
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from '@/components/ui/sheet'

const route = useRoute()
</script>

<template>
  <nav class="p-3 flex items-center justify-between max-w-screen-2xl mx-auto">
    <span class="flex gap-3">
      <img src="https://api.iconify.design/fluent:document-folder-24-regular.svg?color=%23000" width="23" alt="">
      <p class="font-semibold text-lg">Estrim</p>
    </span>

    <Sheet>
      <SheetTrigger>
        <Button variant="ghost" size="icon" class="md:hidden">
          <img src="https://api.iconify.design/fluent:list-rtl-16-filled.svg?color=%23000" width="20" alt="">
        </Button>
      </SheetTrigger>
      <SheetContent>
        <SheetHeader>
          <SheetTitle class="-mt-3.5">Navigation</SheetTitle>
        </SheetHeader>
        
        <div class="flex flex-col gap-3 mt-3">
          <RouterLink :to="{ name: Routes.Home }">
            <Button size="lg" variant="link" :class="`pl-0.5 hover:no-underline ${route.path === '/' ? '' : 'text-muted-foreground hover:text-foreground'}`">
              Home
            </Button>
          </RouterLink>
          <RouterLink :to="{ name: Routes.Home }">
            <Button size="lg" variant="link" :class="`pl-0.5 hover:no-underline ${route.path === '/pricing' ? '' : 'text-muted-foreground hover:text-foreground'}`">
              Pricing
            </Button>
          </RouterLink :to="{ name: Routes.Home }">
          <RouterLink :to="{ name: Routes.Home }">
            <Button size="lg" variant="link" :class="`pl-0.5 hover:no-underline ${route.path === '/sign-in' ? '' : 'text-muted-foreground hover:text-foreground'}`">
              Sign in
            </Button>
          </RouterLink>
          <Button size="lg" class="w-fit">Sign up</Button>
        </div>
      </SheetContent>
    </Sheet>


    <div class="hidden md:flex gap-3">
      <RouterLink :to="{ name: Routes.Home }">
        <Button variant="link" :class="`hover:no-underline ${route.path === '/' ? '' : 'text-muted-foreground hover:text-foreground'}`">Home</Button>
      </RouterLink>
      <RouterLink :to="{ name: Routes.Home }">
        <Button variant="link" :class="`hover:no-underline ${route.path === '/pricing' ? '' : 'text-muted-foreground hover:text-foreground'}`">Pricing</Button>
      </RouterLink>
      <RouterLink :to="{ name: Routes.Home }">
        <Button variant="link" :class="`hover:no-underline ${route.path === '/sign-in' ? '' : 'text-muted-foreground hover:text-foreground'}`">Sign in</Button>
      </RouterLink>
      <Button>Sign up</Button>
    </div>
  </nav>
</template>