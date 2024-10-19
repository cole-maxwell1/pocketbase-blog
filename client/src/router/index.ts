import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import client from '@/pocketbase'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('../views/RegisterView.vue'),
    },
    {
      path: '/feed',
      name: 'feed',
      meta: { requiresAuth: true },
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/FeedView.vue'),
    },
    {
      path: '/posts/new',
      name: 'new-post',
      meta: { requiresAuth: true },
      component: () => import('../views/NewPostView.vue'),
    },
  ],
})

router.beforeEach(to => {
  // Init the store within the beforeEach function as per the documentation:
  // https://pinia.vuejs.org/core-concepts/outside-component-usage.html#single-page-applications
  if (to.meta.requiresAuth && !client?.authStore.token) {
    return {
      path: '/',
    }
  }
})

// If user is logged in, redirect to feed
router.beforeEach(to => {
  if (
    (to.name === 'login' || to.name === 'register') &&
    client?.authStore.token
  ) {
    return {
      path: '/feed',
    }
  }
})

export default router
