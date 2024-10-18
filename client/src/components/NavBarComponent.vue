<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'
import client from '@/pocketbase'
import { computed, ref } from 'vue'
import Menubar from 'primevue/menubar'
import Button from 'primevue/button'
import pocketBaseLogo from '@/assets/svg/logos-pocket-base.svg'
import type { MenuItem } from 'primevue/menuitem'

// Init the store
const userStore = useUserStore()

// Router composable
const router = useRouter()

const isUserLoggedIn = ref(!!client.authStore.token)

const items = computed<MenuItem[]>(() => [
  {
    label: 'Home',
    icon: 'pi pi-fw pi-home',
    url: '/',
  },
  {
    label: 'Feed',
    icon: 'pi pi-fw pi-book',
    url: '/feed',
    visible: isUserLoggedIn.value,
  },
  {
    label: 'New Post',
    icon: 'pi pi-fw pi-plus',
    url: '/posts/new',
    visible: isUserLoggedIn.value,
  },
])

const logoutUser = () => {
  // Manual reset because Pinia using the composition API does not support the $reset function
  userStore.clear()
  // Remove the PocketBase token
  client.authStore.clear()
  // Redirect to the login page

  isUserLoggedIn.value = false
  router.push({ path: '/' })
}
</script>

<template>
  <Menubar :model="items">
    <template #start>
      <img :src="pocketBaseLogo" alt="PocketBase Logo" class="w-8 h-8" />
    </template>
    <template #end>
      <div class="flex items-center gap-2">
        <Button
          v-if="isUserLoggedIn"
          as="router-link"
          to="/profile"
          text
          size="large"
          v-tooltip.left="'Profile'"
          aria-label="Profile"
          severity="secondary"
          icon="pi pi-fw pi-user"
        />
        <Button
          v-if="!isUserLoggedIn"
          as="router-link"
          to="/login"
          outlined
          v-tooltip.left="'Login'"
          label="Login"
          severity="secondary"
          icon="pi pi-fw pi-sign-in"
        />
        <Button
          v-if="isUserLoggedIn"
          @click="logoutUser"
          v-tooltip.left="'Logout'"
          text
          size="large"
          aria-label="Logout"
          severity="secondary"
          icon="pi pi-fw pi-sign-out"
        />
      </div>
    </template>
  </Menubar>
</template>
