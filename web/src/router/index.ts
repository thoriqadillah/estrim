import { eventBus } from '@/lib'
import { useAccount } from '@/stores/account'
import type { Plugin } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'

export enum Routes {
  Home = 'home',
  Pricing = 'pricing',
  Login = 'auth:sign-in',
  Register = 'auth:sign-up'
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: Routes.Home,
      component: () => import('../pages/home/Home.vue'),
    },
  ]
})

const AuthGuard: Plugin = {
  install: () => {
    const store = useAccount()
    eventBus.on('logout', store.logout)

    router.beforeEach(async (to, from, next) => {
      await store.getAccount()
      
      if (to.meta.useAuth && !store.authenticated) {
        next({ 
          name: Routes.Login,
          query: {
            redirect: encodeURIComponent(to.fullPath)      
          }
        })
      }

      next()
    })
  }
}

export {
  router,
  AuthGuard
}
