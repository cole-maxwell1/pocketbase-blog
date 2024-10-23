<template>
  <Card>
    <template #title>
      <div>
        <h1>{{ post.title }}</h1>
        <p class="text-sm capitalize">
          {{ `By: ${post.firstName} ${post.lastName}` }}
        </p>
        <p class="text-sm">{{ formatDateTime(post.created) }}</p>
      </div>
    </template>
    <template #content>
      <p class="m-0">{{ snippet }}</p>
    </template>
    <template #footer>
      <div class="flex flex-col gap-1">
        <div class="flex justify-end">
          <Button
            outlined
            size="small"
            icon="pi pi-arrow-right"
            label="Read"
            v-tooltip.left="'Read More'"
            icon-pos="right"
            as="router-link"
            :to="{ name: 'post-view', params: { postId: post.id } }"
          />
        </div>
        <div class="flex gap-1 items-center">
          <Tag
            v-for="tag in post.tags"
            :key="tag.id"
            :value="tag.name"
            severity="secondary"
            icon="pi pi-tag"
            class="capitalize"
          />
        </div>
      </div>
    </template>
  </Card>
</template>

<script setup lang="ts">
import Card from 'primevue/card'
import Button from 'primevue/button'
import Tag from 'primevue/tag'
import { computed } from 'vue'
import type { Post } from '@/interfaces/post'
import { formatDateTime } from '@/utils/formatters'

interface Props {
  post: Post
}

const props = defineProps<Props>()

const snippet = computed(() => {
  const parser = new DOMParser()

  const doc = parser.parseFromString(props.post.content, 'text/html')

  const firstParagraph = doc.querySelector('p')

  let snippetContent = ''
  if (firstParagraph) {
    snippetContent = firstParagraph.innerHTML

    // Optionally, truncate the content if it's too long
    const maxLength = 200 // Customize the length as needed
    if (snippetContent.length > maxLength) {
      // Remove any trailing punctuation or whitespace from the end of the string
      snippetContent = snippetContent.replace(
        /[.,\/#!$%\^&\*;:{}=\-_`~()\s]+$/,
        '',
      )

      // Truncate the string and add an ellipsis
      snippetContent = snippetContent.substring(0, maxLength) + '...'
    }
  }
  return snippetContent
})
</script>
