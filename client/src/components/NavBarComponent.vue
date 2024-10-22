<template>
  <nav
    class="flex flex-wrap justify-between items-center dark:border-zinc-600 shadow-sm p-2 border-b"
  >
    <!-- Logo -->
    <div class="flex flex-shrink-0 items-center mr-6">
      <img :src="pocketBaseLogo" alt="PocketBase Logo" class="w-8 h-8" />
    </div>

    <!-- Mobile menu button -->
    <div class="block md:hidden">
      <Button
        @click="isMobileMenuOpen = !isMobileMenuOpen"
        icon="pi pi-bars"
        class="p-button-plain p-button-text"
        aria-label="Menu"
      />
    </div>

    <!-- Navigation links -->
    <div
      :class="[
        'w-full',
        'md:flex',
        'md:items-center',
        'md:w-auto',
        'md:grow',
        isMobileMenuOpen ? 'block' : 'hidden',
      ]"
    >
      <div
        class="flex md:flex-row flex-col md:items-center gap-2 mt-2 md:mt-0 md:grow"
      >
        <!-- Left-side navigation -->
        <div
          class="flex md:flex-row flex-col md:items-center gap-2 md:place-self-start"
        >
          <Button
            v-if="isUserLoggedIn"
            as="router-link"
            to="/"
            text
            label="Home"
            severity="secondary"
            icon="pi pi-fw pi-home"
          />
          <Button
            v-if="isUserLoggedIn"
            as="router-link"
            to="/feed"
            text
            label="Feed"
            severity="secondary"
            icon="pi pi-fw pi-book"
          />
          <Button
            v-if="isUserLoggedIn"
            as="router-link"
            to="/posts/new"
            text
            label="New Post"
            severity="secondary"
            icon="pi pi-fw pi-plus"
          />
        </div>

        <!-- Right-side navigation -->
        <div class="flex md:flex-row flex-col md:items-center gap-2 md:ml-auto">
          <Button
            v-if="isUserLoggedIn"
            as="router-link"
            :to="{
              name: 'profile',
              params: { profileId: userId },
            }"
            text
            v-tooltip.left="'Profile'"
            :label="`Hi, ${firstName}`"
            severity="secondary"
            icon="pi pi-fw pi-user"
          />
          <Button
            v-if="!isUserLoggedIn"
            as="router-link"
            to="login"
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
            aria-label="Logout"
            severity="secondary"
            icon="pi pi-fw pi-sign-out"
          />
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import client from '@/pocketbase'
import { ref } from 'vue'
import Button from 'primevue/button'
import pocketBaseLogo from '@/assets/svg/logos-pocket-base.svg'

// Router composable
const router = useRouter()

const isUserLoggedIn = ref(!!client.authStore.token)
const userId = ref(client.authStore.model?.id || '')
const isMobileMenuOpen = ref(false)
const firstName = ref(
  client.authStore.model?.firstName[0].toUpperCase() +
    client.authStore.model?.firstName.slice(1),
)

const logoutUser = () => {
  // Remove the PocketBase token
  client.authStore.clear()
  // Redirect to the login page

  isUserLoggedIn.value = false
  router.push({ path: '/' })
}
</script>
