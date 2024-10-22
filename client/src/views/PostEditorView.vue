<template>
  <div class="p-6">
    <Card>
      <template #title>
        <h1 v-if="hasPostId">Edit Post</h1>
        <h1 v-else>New Post</h1>
      </template>
      <template #content>
        <form @submit.prevent="submitPost" class="flex flex-col gap-6 h-fit">
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
      </template>
    </Card>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import Card from 'primevue/card'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import FloatLabel from 'primevue/floatlabel'
import MultiSelect from 'primevue/multiselect'
import client from '@/pocketbase'
import { useRoute, useRouter } from 'vue-router'
import { useToast } from 'primevue/usetoast'
import RichEditor from '@/components/RichEditor.vue'

const toast = useToast()

const route = useRoute()
const router = useRouter()
const hasPostId: boolean = !!route.params.postId
const postId = route.params.postId as string

const title = ref<string>('')
const content = ref<string>('')

const isLoadingTags = ref<boolean>(false)
const selectedTags = ref<Tag[]>([]) // Array to hold selected tags
const tags = ref<Tag[]>([]) // Array to hold all available tags
const filterValue = ref('')

onMounted(async () => {
  // Load existing tags
  isLoadingTags.value = true
  const result = await client.collection('tags').getFullList()
  tags.value = result.map(tag => ({ id: tag.id, name: tag.name }))
  isLoadingTags.value = false

  if (hasPostId) {
    const post = await client.collection('posts').getOne(postId, {
      expand: 'tags',
    })
    title.value = post.title
    content.value = post.content
    if (post.expand && post.expand.tags) {
      selectedTags.value = post.expand.tags
    }
  }
})

import type { MultiSelectFilterEvent } from 'primevue/multiselect'
import type { Tag } from '@/interfaces/post'

function onFilter(event: MultiSelectFilterEvent) {
  filterValue.value = event.value
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
      const result = await client.collection('tags').create({ name: tagName })

      const newTag = { id: result.id, name: result.name }
      // Add to tags list
      tags.value.push(newTag)
      // Add to selectedTags
      selectedTags.value.push(newTag)
    } catch (error) {
      console.error('Error creating tag:', error)
    }
  }
}

async function submitPost() {
  const tagIds = selectedTags.value.map(tag => tag.id)

  try {
    if (hasPostId) {
      const post = await client.collection('posts').update(postId, {
        title: title.value,
        content: content.value,
        userId: client.authStore?.model?.id,
        tags: tagIds,
      })
      router.push({ name: 'post-view', params: { postId: post.id } })
    } else {
      const newPost = await client.collection('posts').create({
        title: title.value,
        content: content.value,
        userId: client.authStore?.model?.id,
        tags: tagIds,
      })
      toast.add({
        severity: 'success',
        summary: 'Success',
        detail: 'Post saved successfully',
        life: 3000,
      })
      router.push({ name: 'post-view', params: { postId: newPost.id } })
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
