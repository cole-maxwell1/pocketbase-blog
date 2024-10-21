<template>
  <div class="p-6 flex flex-col w-full h-full max-h-full min-h-0">
    <ButtonGroup class="h-fit" v-if="isUserPost">
      <Button
        as="router-link"
        :to="{ name: 'edit-post', params: { postId } }"
        text
        label="Edit"
        severity="secondary"
        icon="pi pi-fw pi-pencil"
      />
      <Button
        text
        label="Delete"
        severity="secondary"
        icon="pi pi-fw pi-trash"
        @click="deletePostConfirm"
      />
    </ButtonGroup>
    <div class="flex flex-col gap-1 mb-1">
      <h1 class="font-semibold text-2xl">{{ title }}</h1>
      <p>
        By: <span class="capitalize">{{ author }}</span>
      </p>
      <p class="text-sm">
        <span
          >Published on <span>{{ created }}</span></span
        >
        <span v-if="hasBeenUpdated">
          | Updated on <span>{{ updated }}</span></span
        >
      </p>
    </div>
    <article
      class="overflow-auto h-full"
      v-html="content"
      id="post-content"
    ></article>
    <PostComments class="mt-6" :post-id="postId"></PostComments>
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

const confirm = useConfirm()
const toast = useToast()

const route = useRoute()
const router = useRouter()
const isUserPost = ref<boolean>(false)

const postId = route.params.postId as string
const title = ref<string>('')
const content = ref<string>('')
const author = ref<string>('')
const created = ref<string>('')
const hasBeenUpdated = ref<boolean>(false)
const updated = ref<string>('')

onMounted(async () => {
  const post = await client.collection('posts_author_vw').getOne(postId)
  isUserPost.value = post.userId === client.authStore?.model?.id
  title.value = post.title
  content.value = post.content
  author.value = `${post.firstName} ${post.lastName}`

  created.value = formatDateTime(new Date(post.created))
  hasBeenUpdated.value = post.updated !== post.created
  updated.value = formatDateTime(new Date(post.updated))
})

async function deletePost() {
  await client.collection('posts').delete(postId)
  toast.add({
    severity: 'info',
    summary: 'Confirmed',
    detail: 'Record deleted',
    life: 3000,
  })
  router.push({ name: 'home' })
}

const deletePostConfirm = () => {
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
    reject: () => {
      toast.add({
        severity: 'error',
        summary: 'Rejected',
        detail: 'Failed to delete the post. Please try again.',
        life: 3000,
      })
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
