<script setup lang="ts">
import NavBarComponent from '@/components/NavBarComponent.vue'
import { RouterView } from 'vue-router'
import { useRoute } from 'vue-router'

const route = useRoute()
</script>

<template>
  <div class="flex flex-col h-full">
    <NavBarComponent
      v-if="route.name !== 'login' && route.name !== 'register'"
    />
    <RouterView v-slot="{ Component }">
      <transition
        mode="out-in"
        enter-active-class="transition duration-200 ease-in"
        enter-from-class="-translate-y-4 opacity-0"
        enter-to-class="translate-y-0 opacity-100"
        leave-active-class="transition duration-150 ease-out"
        leave-from-class="translate-y-0 opacity-100"
        leave-to-class="-translate-y-4 opacity-0"
      >
        <component :is="Component" />
      </transition>
    </RouterView>
  </div>
</template>
