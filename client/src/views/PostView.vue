<template>
  <div
    v-if="post"
    class="flex h-full max-h-full min-h-0 w-full flex-col overflow-auto p-6"
  >
    <div class="w-full md:flex">
      <div class="mb-1 flex flex-col gap-1">
        <h1 class="text-3xl font-semibold">{{ post.title }}</h1>
        <router-link
          v-if="post.userId"
          :to="{ name: 'profile', params: { profileId: post.userId } }"
          class="flex items-center gap-1 transition hover:translate-x-2"
        >
          By:
          <span class="capitalize">{{
            `${post.firstName} ${post.lastName}`
          }}</span>
          <span class="pi pi-chevron-right" style="font-size: 0.75rem"></span>
        </router-link>
        <div class="text-sm text-surface-600 dark:text-surface-400">
          <p>
            Published on <span>{{ formatDateTime(post.created) }}</span>
          </p>

          <p v-if="hasBeenUpdated">
            Updated on <span>{{ formatDateTime(post.updated) }}</span>
          </p>
        </div>
        <!-- Share buttons -->
        <ShareButtonGroup :title="post.title"></ShareButtonGroup>
      </div>
      <ButtonGroup class="ml-auto self-end" v-if="isUserPost">
        <Button
          as="router-link"
          :to="{ name: 'edit-post', params: { postId } }"
          text
          label="Edit"
          v-tooltip.bottom="'Edit Post'"
          severity="secondary"
          icon="pi pi-fw pi-pencil"
        />
        <Button
          text
          label="Delete"
          severity="danger"
          icon="pi pi-fw pi-trash"
          @click="deletePostConfirm"
        />
      </ButtonGroup>
    </div>
    <article v-html="post.content" class="rich-text"></article>
    <div class="flex items-center gap-1">
      <Tag
        v-for="tag in post.tags"
        :key="tag.id"
        :value="tag.name"
        severity="secondary"
        icon="pi pi-tag"
        class="mt-4 capitalize"
      />
    </div>
    <div class="min-h-0">
      <PostComments class="mt-6 h-fit" :post-id="postId"></PostComments>
    </div>
  </div>
  <div v-else class="flex justify-center">
    <ProgressSpinner />
  </div>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import pbClient from '@/pocketbase'
import { onMounted, ref } from 'vue'
import ButtonGroup from 'primevue/buttongroup'
import Button from 'primevue/button'
import { useConfirm } from 'primevue/useconfirm'
import { useToast } from 'primevue/usetoast'
import PostComments from '@/components/PostComments.vue'
import { formatDateTime } from '@/utils/formatters'
import ShareButtonGroup from '@/components/ShareButtonGroup.vue'
import { deletePost, getPostById } from '@/services/postService'
import type { Post } from '@/interfaces/post'
import ProgressSpinner from 'primevue/progressspinner'
import Tag from 'primevue/tag'

const confirm = useConfirm()
const toast = useToast()

const route = useRoute()
const router = useRouter()
const isUserPost = ref<boolean>(false)

const postId = route.params.postId as string
const post = ref<Post | undefined>(undefined)
const hasBeenUpdated = ref<boolean>(false)

onMounted(async () => {
  try {
    post.value = await getPostById(postId)
    isUserPost.value = post.value.userId === pbClient.authStore?.model?.id
    hasBeenUpdated.value = post.value.updated !== post.value.created
  } catch (error: unknown) {
    console.log(error)
    toast.add({
      severity: 'error',
      summary: 'Rejected',
      detail: 'Failed to fetch the post. Please try again.',
      life: 5000,
    })
  }
})

async function handelDeletePost() {
  try {
    await deletePost(postId)
    toast.add({
      severity: 'info',
      summary: 'Confirmed',
      detail: 'Post has been deleted',
      life: 3000,
    })
    router.push({ name: 'home' })
  } catch (error: unknown) {
    console.log(error)
    toast.add({
      severity: 'error',
      summary: 'Rejected',
      detail: 'Failed to delete the post. Please try again.',
      life: 5000,
    })
  }
}

function deletePostConfirm() {
  confirm.require({
    message:
      'Are you sure you want to delete this post? This action cannot be undone.',
    header: 'Delete Post?',
    icon: 'pi pi-info-circle',
    rejectLabel: 'Cancel',
    rejectProps: {
      label: 'Cancel',
      severity: 'secondary',
      outlined: true,
    },
    acceptProps: {
      label: 'Delete',
      severity: 'danger',
    },
    accept: async () => {
      await handelDeletePost()
    },
  })
}
</script>

<style>
#post-content {
  margin-top: 1rem;
}

#post-content h1 {
  font-size: 2rem;
  margin-bottom: 0.5rem;
}

/* h2 in post-content */
#post-content h2 {
  font-size: 1.5rem;
  margin-top: 1rem;
  margin-bottom: 0.5rem;
}
</style>
