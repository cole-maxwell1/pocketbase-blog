<template>
  <Card>
    <template #header>
      <h1 class="mb-2 px-4 pt-4 text-2xl font-semibold">Comments</h1>
    </template>
    <template #content>
      <div
        class="overflow-auto"
        ref="scrollContainer"
        style="max-height: 225px"
      >
        <ul v-if="comments.length > 0">
          <li
            v-for="comment in comments"
            :key="comment.id"
            class="border-b border-gray-300 p-2"
          >
            <p class="font-semibold capitalize">
              {{ `${comment.firstName} ${comment.lastName}` }}
            </p>
            <p>{{ comment.content }}</p>
            <small class="text-surface-500 dark:text-surface-400">{{
              formatDateTime(comment.created)
            }}</small>
          </li>
        </ul>
        <div
          v-if="!isLoadingComments && comments.length === 0"
          class="p-10 text-center"
        >
          No comments yet. Be the first to comment!
        </div>
        <div v-if="isLoadingComments" class="p-4 text-center">
          <ProgressSpinner />
        </div>
      </div>
    </template>
    <template #footer>
      <form
        v-if="isLoggedIn"
        @submit.prevent="postComment"
        class="mt-auto flex h-full flex-col gap-2"
      >
        <Textarea v-model="newComment" rows="3" autoResize></Textarea>
        <Button type="submit" label="Post Comment" icon="pi pi-comment" />
      </form>
      <template v-else>
        <p class="p-4 text-center">
          <router-link
            to="/login"
            class="text-primary-500 dark:text-primary-400"
          >
            Log in
          </router-link>
          to post a comment
        </p>
      </template>
    </template>
  </Card>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useInfiniteScroll } from '@vueuse/core'
import Card from 'primevue/card'
import Textarea from 'primevue/textarea'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'
import { formatDateTime } from '@/utils/formatters'
import ProgressSpinner from 'primevue/progressspinner'
import type { PostComment } from '@/interfaces/post'
import {
  createNewComment,
  getPostCommentsPaginated,
} from '@/services/postService'
import pbClient from '@/pocketbase'

const toast = useToast()
const PAGE_SIZE = 5
const pagesLoaded = ref(new Set<number>())
const totalCommentPages = ref(1)
const isLoadingComments = ref<boolean>(false)
const comments = ref<PostComment[]>([])
const newComment = ref<string>('')
const isLoggedIn = ref<boolean>(false)

interface Props {
  postId: string
}

const props = defineProps<Props>()

onMounted(() => {
  isLoggedIn.value = pbClient.authStore.isValid
  loadCommentsChunk(1)
})

const scrollContainer = ref(null)
useInfiniteScroll(
  scrollContainer,
  async () => {
    const nextPage = pagesLoaded.value.size + 1
    await loadCommentsChunk(nextPage)
  },
  { distance: PAGE_SIZE },
)

async function loadCommentsChunk(page: number) {
  if (
    isLoadingComments.value ||
    pagesLoaded.value.has(page) ||
    page > totalCommentPages.value
  ) {
    return
  }

  try {
    isLoadingComments.value = true
    const fetchedCommentsPage = await getPostCommentsPaginated(
      props.postId,
      page,
      PAGE_SIZE,
    )

    comments.value.push(...fetchedCommentsPage.comments)

    pagesLoaded.value.add(page)
    totalCommentPages.value = fetchedCommentsPage.totalPages
  } catch (error) {
    console.error('Error loading comments:', error)
    toast.add({
      severity: 'error',
      summary: 'Error Loading Comments',
      detail: 'An error occurred while loading comments',
      life: 3000,
    })
  } finally {
    isLoadingComments.value = false
  }
}

async function postComment() {
  try {
    const resultComment = await createNewComment(props.postId, newComment.value)

    comments.value.unshift(resultComment)

    newComment.value = ''

    toast.add({
      severity: 'success',
      summary: 'Comment Posted',
      detail: 'Your comment has been posted successfully',
      life: 3000,
    })
  } catch (error: unknown) {
    console.error(error)
    toast.add({
      severity: 'error',
      summary: 'Error Posting Comment',
      detail: 'An error occurred while posting your comment',
      life: 3000,
    })
  }
}
</script>

<style scoped>
/* Optional styling to give the comments section a scrollable container */
.overflow-auto {
  max-height: 400px;
  overflow-y: auto;
}
</style>
