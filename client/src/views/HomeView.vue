<template>
  <div class="p-6 flex flex-col gap-4 h-full max-h-full overflow-auto">
    <main class="gap-6 grid grid-cols-1 lg:grid-cols-2 overflow-auto">
      <PostDisplayCard v-for="post in posts" :post="post" :key="post.id" />
    </main>
    <Paginator
      class="mt-auto"
      :totalRecords="totalPosts"
      :rows="itemsPerPage"
      :rowsPerPageOptions="[5, 10, 20, 50]"
      @page="handelPageChange"
    />
  </div>
</template>

<script setup lang="ts">
import client from '@/pocketbase'
import { onMounted, ref } from 'vue'
import Paginator, { type PageState } from 'primevue/paginator'
import PostDisplayCard from '@/components/PostDisplayCard.vue'
import { useToast } from 'primevue/usetoast'
import type { Post, Tag } from '@/interfaces/post'

const toast = useToast()

const posts = ref<Post[]>([])

const currentPage = ref(1)
const totalPosts = ref(0)
const itemsPerPage = ref(10)

onMounted(async () => {
  await getPostsPage(currentPage.value, itemsPerPage.value)
})

async function handelPageChange(event: PageState) {
  await getPostsPage(event.page + 1, event.rows)
}

async function getPostsPage(page: number, totalItems: number) {
  try {
    const resultList = await client
      .collection('posts_author_vw')
      .getList(page, totalItems, { expand: 'tags' })

    totalPosts.value = resultList.totalItems

    // Update the posts array
    posts.value = resultList.items.map(item => ({
      id: item.id,
      title: item.title,
      content: item.content,
      firstName: item.firstName,
      lastName: item.lastName,
      tags: item.expand?.tags.map((tag: Tag) => {
        return {
          name: tag?.name,
          id: tag?.id,
        }
      }),
      created: item.created,
      updated: item.updated,
    }))
  } catch (error: unknown) {
    console.error(error)
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Could not get posts. Please try again.',
      life: 3000,
    })
  }
}
</script>
