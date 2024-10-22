<template>
  <Card>
    <template #header>
      <h1 class="text-2xl font-semibold pt-4 px-4 mb-2">Comments</h1>
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
            class="p-2 border-b border-gray-300"
          >
            <p class="font-semibold">{{ comment.author }}</p>
            <p>{{ comment.content }}</p>
            <small class="text-gray-500">{{ comment.created }}</small>
          </li>
        </ul>
        <div
          v-if="!isLoadingComments && comments.length === 0"
          class="text-center p-10"
        >
          No comments yet. Be the first to comment!
        </div>
        <div v-if="isLoadingComments" class="text-center p-4">
          <ProgressSpinner />
        </div>
      </div>
    </template>
    <template #footer>
      <form
        @submit.prevent="postComment"
        class="mt-auto flex flex-col gap-2 h-full"
      >
        <Textarea v-model="newComment" rows="3" autoResize></Textarea>
        <Button type="submit" label="Post Comment" icon="pi pi-comment" />
      </form>
    </template>
  </Card>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useInfiniteScroll } from '@vueuse/core'
import Card from 'primevue/card'
import Textarea from 'primevue/textarea'
import Button from 'primevue/button'
import client from '@/pocketbase'
import { useToast } from 'primevue/usetoast'
import { formatDateTime } from '@/utils/formatters'
import ProgressSpinner from 'primevue/progressspinner'

const toast = useToast()
const PAGE_SIZE = 5
const pagesLoaded = ref(new Set<number>())
const totalCommentPages = ref(1)
const isLoadingComments = ref<boolean>(false)
const comments = ref<
  {
    id: string
    author: string
    content: string
    created: string
    updatedAt: string
  }[]
>([])
const newComment = ref<string>('')

interface Props {
  postId: string
}

const props = defineProps<Props>()

onMounted(() => {
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
    const result = await client
      .collection('comments_author_vw')
      .getList(page, PAGE_SIZE, {
        filter: `postId='${props.postId}'`,
        sort: 'created',
      })

    comments.value.push(
      ...result.items.map(comment => ({
        id: comment.id,
        author: `${comment.firstName} ${comment.lastName}`,
        content: comment.content,
        created: formatDateTime(new Date(comment.created)),
        updatedAt: formatDateTime(new Date(comment.updated)),
      })),
    )

    pagesLoaded.value.add(page)
    totalCommentPages.value = result.totalPages
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
    const resultComment = await client.collection('comments').create({
      userId: client.authStore?.model?.id,
      postId: props.postId,
      content: newComment.value,
    })

    const commentDetails = await client
      .collection('comments_author_vw')
      .getOne(resultComment.id)

    comments.value.push({
      id: commentDetails.id,
      author: `${commentDetails.firstName} ${commentDetails.lastName}`,
      content: commentDetails.content,
      created: commentDetails.created,
      updatedAt: commentDetails.updated,
    })

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
