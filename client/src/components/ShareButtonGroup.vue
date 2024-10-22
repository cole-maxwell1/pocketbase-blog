<template>
  <div class="flex gap-2">
    <Button
      v-if="isSupported"
      icon="pi pi-share-alt"
      aria-label="Share"
      text
      v-tooltip.bottom="'Share this!'"
      @click="startShare"
    />
    <Button
      v-for="button in shareButtons"
      :key="button.label"
      :icon="button.icon"
      :arial-label="button.label"
      :class="button.color"
      v-tooltip.bottom="{
        value: button.label,
      }"
      text
      @click="button.command"
    />
  </div>
</template>

<script setup lang="ts">
import Button from 'primevue/button'
import { ref } from 'vue'
import { useShare } from '@vueuse/core'
import { isClient } from '@vueuse/shared'

const props = defineProps<{
  title: string
}>()

const options = ref({
  title: props.title,
  text: 'Checkout this post!',
  url: isClient ? location.href : '',
})

const { share, isSupported } = useShare(options)

function startShare() {
  share({
    title: 'Hello',
    text: 'Hello my friend!',
    url: location.href,
  })
}

const shareButtons = ref([
  {
    label: 'Facebook',
    icon: 'pi pi-facebook',
    color: 'text-[#3b5998]',
    command: () => myShare('facebook'),
  },
  {
    label: 'Twitter (x.com)',
    icon: 'pi pi-twitter',
    color: 'text-black dark:text-white',
    command: () => myShare('twitter'),
  },
  {
    label: 'LinkedIn',
    icon: 'pi pi-linkedin',
    color: '#0077b5',
    command: () => myShare('linkedin'),
  },
  {
    label: 'Reddit',
    icon: 'pi pi-reddit',
    color: 'text-[#ff4500]',
    command: () => myShare('reddit'),
  },
  {
    label: 'Email',
    icon: 'pi pi-envelope',
    color: 'text-black dark:text-white',
    command: () => myShare('email'),
  },
])

function myShare(platform: string) {
  const url = encodeURIComponent(window.location.href)
  const title = encodeURIComponent(props.title)

  switch (platform) {
    case 'facebook':
      window.open(
        `https://www.facebook.com/sharer/sharer.php?u=${url}`,
        '_blank',
      )
      break
    case 'twitter':
      window.open(
        `https://twitter.com/intent/tweet?text=${title}&url=${url}`,
        '_blank',
      )
      break
    case 'linkedin':
      window.open(
        `https://www.linkedin.com/shareArticle?mini=true&url=${url}&title=${title}`,
        '_blank',
      )
      break
    case 'reddit':
      window.open(
        `https://www.reddit.com/submit?url=${url}&title=${title}`,
        '_blank',
      )
      break
    case 'email':
      window.open(`mailto:?subject=${title}&body=${url}`, '_self')
      break
    default:
      console.error('Unsupported platform:', platform)
  }
}
</script>

<style scoped></style>
