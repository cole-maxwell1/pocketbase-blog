<template>
  <main class="gap-6 grid grid-cols-1 lg:grid-cols-2 p-6 overflow-auto">
    <pre>
      {{ posts }}
    </pre>
    <PostDisplayCard v-for="post in posts" :post="post" :key="post.id" />
  </main>
</template>

<script setup lang="ts">
import client from '@/pocketbase'
import { onMounted, ref } from 'vue'
import PostDisplayCard from '@/components/PostDisplayCard.vue'

const posts = ref<
  {
    id: string
    title: string
    content: string
    firstName: string
    lastName: string
    tags: any[]
    created: string
    updated: string
  }[]
>([])

onMounted(async () => {
  const resultList = await client.collection('posts_author_vw').getList(1, 10,
    {expand: 'tags'}
  )
  console.log(resultList)
  posts.value = resultList.items.map(item => ({
    id: item.id,
    title: item.title,
    content: item.content,
    firstName: item.firstName,
    lastName: item.lastName,
    tags: item.expand?.tags.map((tag: { name: any; id: any; }) => {
      return {
        name: tag?.name,
        id: tag?.id
      }
    }),
    created: item.created,
    updated: item.updated,
  }))
})
</script>
