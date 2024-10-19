<template>
  <main
    v-if="posts"
    class="gap-6 grid grid-cols-1 lg:grid-cols-2 p-6 overflow-auto"
  >
    <Card class="" v-for="post in posts" :key="post.id">
      <template #title
        ><div>
          <h1>{{ post.title }}</h1>
          <p class="text-sm capitalize">
            {{ `By: ${post.firstName} ${post.lastName}` }}
          </p>
          <p class="text-sm">{{ new Date(post.created_at).toDateString() }}</p>
        </div></template
      >
      <template #content>
        <p class="m-0">
          {{ post.content }}
        </p>
      </template>
    </Card>
  </main>
</template>

<script setup lang="ts">
import client from '@/pocketbase'
import { onMounted, ref } from 'vue'
import type { RecordModel } from 'pocketbase'
import Card from 'primevue/card'

const posts = ref<RecordModel[]>([])

onMounted(async () => {
  const resultList = await client.collection('posts_author_vw').getList(1, 10)
  posts.value = resultList.items
})
</script>
