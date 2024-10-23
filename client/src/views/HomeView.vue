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
import { onMounted, ref } from 'vue'
import Paginator, { type PageState } from 'primevue/paginator'
import PostDisplayCard from '@/components/PostDisplayCard.vue'
import { useToast } from 'primevue/usetoast'
import type { Post } from '@/interfaces/post'
import { getAllPostsPaginated } from '@/services/postService'

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
    const postsPage = await getAllPostsPaginated(page, totalItems)

    totalPosts.value = postsPage.totalItems

    // Update the posts array
    posts.value = postsPage.posts
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
