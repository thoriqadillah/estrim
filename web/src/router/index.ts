import { createRouter, createWebHistory } from 'vue-router'

export enum Routes {
  Home = 'home',
  Pricing = 'pricing',
  Login = 'sign-in',
  Register = 'sign-up'
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

export default router
