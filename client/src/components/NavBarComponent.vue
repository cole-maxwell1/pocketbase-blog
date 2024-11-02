<template>
  <nav
    class="flex flex-wrap items-center justify-between border-b p-2 shadow-sm dark:border-zinc-600"
  >
    <!-- Logo -->
    <div class="mr-6 flex flex-shrink-0 items-center">
      <img :src="pocketBaseLogo" alt="PocketBase Logo" class="h-8 w-8" />
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
        class="mt-2 flex flex-col gap-2 md:mt-0 md:grow md:flex-row md:items-center"
      >
        <!-- Left-side navigation -->
        <div
          class="flex flex-col gap-2 md:flex-row md:items-center md:place-self-start"
        >
          <Button
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
            to="/posts/new"
            text
            label="New Post"
            severity="secondary"
            icon="pi pi-fw pi-plus"
          />
        </div>

        <!-- Right-side navigation -->
        <div class="flex flex-col gap-2 md:ml-auto md:flex-row md:items-center">
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
import pbClient from '@/pocketbase'
import { ref } from 'vue'
import Button from 'primevue/button'
import pocketBaseLogo from '@/assets/svg/logos-pocket-base.svg'

// Router composable
const router = useRouter()

const isUserLoggedIn = ref(!!pbClient.authStore.token)
const userId = ref(pbClient.authStore.model?.id || '')
const isMobileMenuOpen = ref(false)
const firstName = ref(
  pbClient.authStore.model?.firstName[0].toUpperCase() +
    pbClient.authStore.model?.firstName.slice(1),
)

const logoutUser = () => {
  // Remove the PocketBase token
  pbClient.authStore.clear()
  // Redirect to the login page

  isUserLoggedIn.value = false
  router.push({ path: '/' })
}
</script>
