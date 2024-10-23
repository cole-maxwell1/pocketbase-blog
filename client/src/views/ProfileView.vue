<template>
  <div
    class="p-6 flex flex-col gap-4 w-full h-full max-h-full min-h-0 overflow-auto"
  >
    <div v-if="author">
      <h1 class="mb-2 font-semibold text-3xl capitalize">
        {{ `${author.firstName} ${author.lastName}` }}
      </h1>
      <Inplace
        :disabled="!isUserProfile"
        v-model:active="isEditingBio"
        :pt="{
          display: {
            role: isUserProfile ? 'button' : 'article',
          },
        }"
      >
        <template #display>
          <article
            v-if="author.bio"
            class="rich-text"
            v-tooltip.top="{ disabled: !isUserProfile, value: 'Click to edit' }"
            v-html="author.bio"
            id="bio-content"
          ></article>
          <p
            v-else
            class="text-surface-600 dark:text-surface-400"
            v-tooltip.top="{ disabled: !isUserProfile, value: 'Click to edit' }"
          >
            {{ isUserProfile ? 'Click to add a bio' : 'No bio available' }}
          </p>
        </template>
        <template #content="{ closeCallback }">
          <form @submit.prevent="saveBio">
            <RichEditor v-model="editedBio" />
            <div class="mt-2 flex gap-4 justify-end">
              <Button
                label="Save"
                icon="pi pi-check"
                type="submit"
                severity="primary"
              />
              <Button
                label="Cancel"
                icon="pi pi-times"
                severity="secondary"
                @click="closeCallback"
              />
            </div>
          </form>
        </template>
      </Inplace>
    </div>
    <div v-else class="flex justify-center">
      <ProgressSpinner />
    </div>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <PostDisplayCard v-for="post in posts" :post="post" :key="post.id" />
      <Paginator
        class="col-span-full"
        :totalRecords="totalPosts"
        :rows="itemsPerPage"
        :rowsPerPageOptions="[5, 10, 20]"
        @page="handelPageChange"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import pbClient from '@/pocketbase'
import { onMounted, ref } from 'vue'
import Inplace from 'primevue/inplace'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'
import RichEditor from '@/components/RichEditor.vue'
import PostDisplayCard from '@/components/PostDisplayCard.vue'
import type { Post } from '@/interfaces/post'
import Paginator from 'primevue/paginator'
import { getPostByUserIdPaginated } from '@/services/postService'
import type { Author } from '@/interfaces/author'
import { getAuthorById, updateAuthorBio } from '@/services/authorService'
import ProgressSpinner from 'primevue/progressspinner'

const toast = useToast()

const route = useRoute()
//const router = useRouter()
const isUserProfile = ref<boolean>(false)
const profileId = route.params.profileId as string

const author = ref<Author | undefined>(undefined)

const isEditingBio = ref<boolean>(false)
const editedBio = ref<string>('')

const posts = ref<Post[]>([])
const currentPage = ref(1)
const totalPosts = ref(0)
const itemsPerPage = ref(5)

onMounted(async () => {
  try {
    author.value = await getAuthorById(profileId)
    isUserProfile.value = author.value.id === pbClient.authStore?.model?.id

    editedBio.value = author.value.bio
  } catch (error: unknown) {
    console.log(error)
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Failed to load profile. Please try again.',
      life: 5000,
    })
  }

  try {
    await getProfilePostsPage(currentPage.value, itemsPerPage.value)
  } catch (error: unknown) {
    console.error(error)
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Could not get posts. Please try again.',
      life: 3000,
    })
  }
})

async function saveBio() {
  try {
    if (author.value && editedBio.value !== author.value.bio) {
      await updateAuthorBio(editedBio.value)
      author.value.bio = editedBio.value
      isEditingBio.value = false
      toast.add({
        severity: 'success',
        summary: 'Success',
        detail: 'Bio updated successfully.',
        life: 3000,
      })
    } else {
      isEditingBio.value = false
    }
  } catch (error: unknown) {
    console.log(error)
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Failed to update bio. Please try again.',
      life: 5000,
    })
  }
}

async function getProfilePostsPage(page: number, totalItems: number) {
  try {
    const postsPage = await getPostByUserIdPaginated(
      profileId,
      page,
      totalItems,
    )

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

function handelPageChange(event: { page: number; rows: number }) {
  getProfilePostsPage(event.page + 1, event.rows)
}
</script>
