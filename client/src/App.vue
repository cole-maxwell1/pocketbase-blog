<script setup lang="ts">
import NavBarComponent from '@/components/NavBarComponent.vue'
import { RouterView } from 'vue-router'
</script>

<template>
  <NavBarComponent
    v-if="!$route.path.includes('login') && !$route.path.includes('register')"
  />
  <RouterView />
</template>
