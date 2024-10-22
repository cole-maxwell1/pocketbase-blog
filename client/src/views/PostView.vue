<template>
  <div class="p-6 flex flex-col w-full h-full max-h-full min-h-0 overflow-auto">
    <div class="md:flex w-full">
      <div class="flex flex-col gap-1 mb-1">
        <h1 class="font-semibold text-3xl">{{ title }}</h1>
        <router-link
          v-if="authorId"
          :to="{ name: 'profile', params: { profileId: authorId } }"
          class="flex gap-1 items-center transition hover:translate-x-2"
        >
          By: <span class="capitalize">{{ author }}</span>
          <span class="pi pi-chevron-right" style="font-size: 0.75rem"></span>
        </router-link>
        <p class="text-sm text-surface-600 dark:text-surface-400">
          <span
            >Published on <span>{{ created }}</span></span
          >
          
          <span v-if="hasBeenUpdated">
            Updated on <span>{{ updated }}</span>
          </span>
        </p>
        <!-- Share buttons -->
        <ShareButtonGroup :title="title"></ShareButtonGroup>
      </div>
      <ButtonGroup class="self-end ml-auto" v-if="isUserPost">
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
    <article v-html="content" class="rich-text"></article>
    <div class="min-h-0">
      <PostComments class="mt-6 h-fit" :post-id="postId"></PostComments>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import client from '@/pocketbase'
import { onMounted, ref } from 'vue'
import ButtonGroup from 'primevue/buttongroup'
import Button from 'primevue/button'
import { useConfirm } from 'primevue/useconfirm'
import { useToast } from 'primevue/usetoast'
import PostComments from '@/components/PostComments.vue'
import { formatDateTime } from '@/utils/formatters'
import ShareButtonGroup from '@/components/ShareButtonGroup.vue'

const confirm = useConfirm()
const toast = useToast()

const route = useRoute()
const router = useRouter()
const isUserPost = ref<boolean>(false)

const postId = route.params.postId as string
const title = ref<string>('')
const content = ref<string>('')
const author = ref<string>('')
const authorId = ref<string>('')
const created = ref<string>('')
const hasBeenUpdated = ref<boolean>(false)
const updated = ref<string>('')

onMounted(async () => {
  const post = await client.collection('posts_author_vw').getOne(postId)
  isUserPost.value = post.userId === client.authStore?.model?.id
  title.value = post.title
  content.value = post.content
  author.value = `${post.firstName} ${post.lastName}`
  authorId.value = post.userId

  created.value = formatDateTime(new Date(post.created))
  hasBeenUpdated.value = post.updated !== post.created
  updated.value = formatDateTime(new Date(post.updated))
})

async function deletePost() {
  try {
    await client.collection('posts').delete(postId)
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
      await deletePost()
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
