<template>
  <div class="p-6">
    <Card>
      <template #title>
        <h1 v-if="hasPostId">Edit Post</h1>
        <h1 v-else>New Post</h1>
      </template>
      <template #content>
        <form
          v-if="!isLoadingPost"
          @submit.prevent="submitPost"
          class="flex flex-col gap-6 h-fit"
        >
          <FloatLabel variant="on">
            <InputText
              type="title"
              id="title"
              v-model="title"
              aria-describedby="title-help"
              class="w-full"
            />
            <label for="email">Title</label>
          </FloatLabel>
          <RichEditor v-model="content" />
          <!-- Tag Selector -->
          <MultiSelect
            v-model="selectedTags"
            :options="tags"
            optionLabel="name"
            dataKey="id"
            placeholder="Select Tags"
            class="w-full"
            :filter="true"
            filter-icon="pi pi-tag"
            display="chip"
            :loading="isLoadingTags"
            @filter="onFilter"
          >
            <template #emptyfilter="">
              <div class="p-3">
                <Button
                  :label="`Create Tag: ${filterValue}`"
                  @click="createTag(filterValue)"
                  class="p-button-text"
                />
              </div>
            </template>
          </MultiSelect>
          <div class="inline-flex justify-end items-center">
            <Button
              :label="hasPostId ? 'Update Post' : 'Create Post'"
              type="submit"
            />
          </div>
        </form>
        <div v-else class="flex justify-center">
          <ProgressSpinner />
        </div>
      </template>
    </Card>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import Card from 'primevue/card'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import FloatLabel from 'primevue/floatlabel'
import MultiSelect from 'primevue/multiselect'
import { useRoute, useRouter } from 'vue-router'
import { useToast } from 'primevue/usetoast'
import RichEditor from '@/components/RichEditor.vue'
import type { MultiSelectFilterEvent } from 'primevue/multiselect'
import type { Tag } from '@/interfaces/post'
import {
  createNewPost,
  createNewTag,
  getPostById,
  getPostTagsPage,
  searchPostTags,
  updatePost,
} from '@/services/postService'
import ProgressSpinner from 'primevue/progressspinner'

const toast = useToast()

const route = useRoute()
const router = useRouter()

const hasPostId: boolean = !!route.params.postId
const postId = route.params.postId as string

const isLoadingPost = ref<boolean>(false)
const title = ref<string>('')
const content = ref<string>('')

const isLoadingTags = ref<boolean>(false)
const selectedTags = ref<Tag[]>([]) // Array to hold selected tags
const tags = ref<Tag[]>([]) // Array to hold all available tags
const filterValue = ref('')

onMounted(async () => {
  try {
    isLoadingTags.value = true
    isLoadingPost.value = hasPostId
    // Get first 10 tags for now
    const result = await getPostTagsPage(1, 10)
    tags.value = result.tags
    isLoadingTags.value = false

    if (hasPostId) {
      const post = await getPostById(postId)
      title.value = post.title
      content.value = post.content
      selectedTags.value = post.tags
      isLoadingPost.value = false
    }
  } catch (error) {
    console.error('Error fetching post:', error)
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'An error occurred while fetching the post',
      life: 5000,
    })
  }
})

// Watch for route changes to update post data
watch(
  () => route.params.postId,
  async newPostId => {
    if (newPostId) {
      isLoadingPost.value = true
      const post = await getPostById(newPostId as string)
      title.value = post.title
      content.value = post.content
      selectedTags.value = post.tags
      isLoadingPost.value = false
    } else {
      title.value = ''
      content.value = ''
      selectedTags.value = []
    }
  },
)

async function onFilter(event: MultiSelectFilterEvent) {
  const trimmedValue = event.value.trim()
  filterValue.value = trimmedValue
  // Only search if value is not empty
  if (trimmedValue !== '') {
    isLoadingTags.value = true
    const result = await searchPostTags(trimmedValue)
    tags.value = result
    isLoadingTags.value = false
  }
}

async function createTag(tagName: string) {
  // Check if tag already exists
  const existingTag = tags.value.find(
    tag => tag.name.toLowerCase() === tagName.toLowerCase(),
  )
  if (existingTag) {
    // Add to selectedTags if not already selected
    if (!selectedTags.value.some(tag => tag.id === existingTag.id)) {
      selectedTags.value.push(existingTag)
    }
  } else {
    try {
      // Create new tag in 'tags' collection
      const newTag = await createNewTag(tagName)
      // Add to tags list
      tags.value.push(newTag)
      // Add to selectedTags
      selectedTags.value.push(newTag)
    } catch (error) {
      console.error('Error creating tag:', error)
      toast.add({
        severity: 'error',
        summary: 'Error',
        detail: 'An error occurred while creating the tag.',
        life: 5000,
      })
    }
  }
}

async function submitPost() {
  const tagIds = selectedTags.value.map(tag => tag.id)

  try {
    if (hasPostId) {
      const post = await updatePost(postId, {
        title: title.value,
        content: content.value,
        tagIds: tagIds,
      })
      router.push({ name: 'post-view', params: { postId: post.id } })
    } else {
      const newPostId = await createNewPost(title.value, content.value, tagIds)
      toast.add({
        severity: 'success',
        summary: 'Success',
        detail: 'Post saved successfully',
        life: 3000,
      })
      router.push({ name: 'post-view', params: { postId: newPostId } })
    }
  } catch (error) {
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'An error occurred while saving the post',
      life: 3000,
    })
    console.error(error)
  }
}
</script>
