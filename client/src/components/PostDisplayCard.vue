<template>
  <Card>
    <template #title>
      <div>
        <h1>{{ post.title }}</h1>
        <p class="text-sm capitalize">
          {{ `By: ${post.firstName} ${post.lastName}` }}
        </p>
        <p class="text-sm">{{ new Date(post.created).toDateString() }}</p>
      </div>
    </template>
    <template #content>
      <p class="m-0">{{ post.content.slice(0, 100) }}...</p>
    </template>
    <template #footer>
      <div class="flex justify-end">
        <Button
          outlined
          size="small"
          icon="pi pi-arrow-right"
          aria-label="Read More"
          v-tooltip.left="'Read More'"
          as="router-link"
          :to="{ name: 'post-view', params: { postId: post.id } }"
        />
      </div>
    </template>
  </Card>
</template>

<script setup lang="ts">
import Card from 'primevue/card'
import Button from 'primevue/button'

interface Post {
  id: string
  title: string
  content: string
  firstName: string
  lastName: string
  created: string
  updated: string
}

interface Props {
  post: Post
}

defineProps<Props>()
</script>
