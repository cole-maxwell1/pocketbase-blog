<template>
  <div
    class="p-6 flex flex-col gap-4 w-full h-full max-h-full min-h-0 overflow-auto"
  >
    <div>
      <h1 class="mb-2 font-semibold text-3xl capitalize">{{ authorName }}</h1>
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
            v-if="bio"
            class="rich-text"
            v-tooltip.top="{ disabled: !isUserProfile, value: 'Click to edit' }"
            v-html="bio"
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
            <div class="flex gap-4 justify-end">
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
import client from '@/pocketbase'
import { onMounted, ref } from 'vue'
import Inplace from 'primevue/inplace'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'
import RichEditor from '@/components/RichEditor.vue'
import PostDisplayCard from '@/components/PostDisplayCard.vue'
import type { Post, Tag } from '@/interfaces/post'
import Paginator from 'primevue/paginator'

const toast = useToast()

const route = useRoute()
//const router = useRouter()
const isUserProfile = ref<boolean>(false)
const profileId = route.params.profileId as string

const bio = ref<string>('')
const authorName = ref<string>('')

const isEditingBio = ref<boolean>(false)
const editedBio = ref<string>('')

const posts = ref<Post[]>([])
const currentPage = ref(1)
const totalPosts = ref(0)
const itemsPerPage = ref(5)

onMounted(async () => {
  try {
    const author = await client.collection('authors_vw').getOne(profileId)
    isUserProfile.value = author.id === client.authStore?.model?.id
    bio.value = author.bio
    editedBio.value = bio.value
    authorName.value = `${author.firstName} ${author.lastName}`
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
    if (editedBio.value !== bio.value) {
      const updatedUser = await client
        .collection('users')
        .update(profileId, { bio: editedBio.value })

      bio.value = updatedUser.bio
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
    const resultList = await client
      .collection('posts_author_vw')
      .getList(page, totalItems, {
        expand: 'tags',
        sort: 'created',
        filter: `userId='${profileId}'`,
      })

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

function handelPageChange(event: { page: number; rows: number }) {
  getProfilePostsPage(event.page + 1, event.rows)
}
</script>
