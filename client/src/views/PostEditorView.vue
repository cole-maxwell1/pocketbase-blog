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
          <Editor v-model="content" editorStyle="height: 320px">
            <template v-slot:toolbar>
              <select class="ql-header" v-tooltip.bottom="'Header Levels'">
                <option selected></option>
                <option value="1"></option>
                <option value="2"></option>
                <option value="3"></option>
                <option value="4"></option>
                <option value="5"></option>
                <option value="6"></option>
              </select>
              <span class="ql-formats">
                <button v-tooltip.bottom="'Bold'" class="ql-bold"></button>
                <button v-tooltip.bottom="'Italic'" class="ql-italic"></button>
                <button
                  v-tooltip.bottom="'Underline'"
                  class="ql-underline"
                ></button>
                <button
                  v-tooltip.bottom="'Ordered List'"
                  class="ql-list"
                  value="ordered"
                ></button>
                <button
                  v-tooltip.bottom="'Bullet List'"
                  class="ql-list"
                  value="bullet"
                ></button>
                <button v-tooltip.bottom="'Strike'" class="ql-strike"></button>
                <button v-tooltip.bottom="'Link'" class="ql-link"></button>
                <select v-tooltip.bottom="'Text Color'" class="ql-color">
                  <option value="#e60000"></option>
                  <option value="#ff9900"></option>
                  <option value="#ffff00"></option>
                  <option value="#008a00"></option>
                  <option value="#0066cc"></option>
                  <option value="#9933ff"></option>
                  <option value="#ffffff"></option>
                  <option value="#facccc"></option>
                  <option value="#ffebcc"></option>
                  <option value="#ffffcc"></option>
                  <option value="#cce8cc"></option>
                  <option value="#cce0f5"></option>
                  <option value="#ebd6ff"></option>
                  <option value="#bbbbbb"></option>
                  <option value="#f06666"></option>
                  <option value="#ffc266"></option>
                  <option value="#ffff66"></option>
                  <option value="#66b966"></option>
                  <option value="#66a3e0"></option>
                  <option value="#c285ff"></option>
                  <option value="#888888"></option>
                  <option value="#a10000"></option>
                  <option value="#b26b00"></option>
                  <option value="#b2b200"></option>
                  <option value="#006100"></option>
                  <option value="#0047b2"></option>
                  <option value="#6b24b2"></option>
                  <option value="#444444"></option>
                  <option value="#5c0000"></option>
                  <option value="#663d00"></option>
                  <option value="#666600"></option>
                  <option value="#003700"></option>
                  <option value="#002966"></option>
                  <option value="#3d1466"></option>
                </select>
                <select class="ql-background">
                  <option value="#e60000"></option>
                  <option value="#ff9900"></option>
                  <option value="#ffff00"></option>
                  <option value="#008a00"></option>
                  <option value="#0066cc"></option>
                  <option value="#9933ff"></option>
                  <option value="#ffffff"></option>
                  <option value="#facccc"></option>
                  <option value="#ffebcc"></option>
                  <option value="#ffffcc"></option>
                  <option value="#cce8cc"></option>
                  <option value="#cce0f5"></option>
                  <option value="#ebd6ff"></option>
                  <option value="#bbbbbb"></option>
                  <option value="#f06666"></option>
                  <option value="#ffc266"></option>
                  <option value="#ffff66"></option>
                  <option value="#66b966"></option>
                  <option value="#66a3e0"></option>
                  <option value="#c285ff"></option>
                  <option value="#888888"></option>
                  <option value="#a10000"></option>
                  <option value="#b26b00"></option>
                  <option value="#b2b200"></option>
                  <option value="#006100"></option>
                  <option value="#0047b2"></option>
                  <option value="#6b24b2"></option>
                  <option value="#444444"></option>
                  <option value="#5c0000"></option>
                  <option value="#663d00"></option>
                  <option value="#666600"></option>
                  <option value="#003700"></option>
                  <option value="#002966"></option>
                  <option value="#3d1466"></option>
                </select>
                <button
                  v-tooltip.bottom="'Subscript'"
                  class="ql-script"
                  value="sub"
                ></button>
                <button
                  v-tooltip.bottom="'Superscript'"
                  class="ql-script"
                  value="super"
                ></button>
              </span>
            </template>
          </Editor>
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
                  @click="createTag(filterValue )"
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
import Editor from 'primevue/editor'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import FloatLabel from 'primevue/floatlabel'
import MultiSelect from 'primevue/multiselect'
import client from '@/pocketbase'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const hasPostId: boolean = !!route.params.postId
const postId = route.params.postId as string

const title = ref<string>('')
const content = ref<string>('')

const isLoadingTags = ref<boolean>(false)
const selectedTags = ref<any[]>([]) // Array to hold selected tags
const tags = ref<any[]>([]) // Array to hold all available tags
  const filterValue = ref('');


onMounted(async () => {
  // Load existing tags
  isLoadingTags.value = true
  const result = await client.collection('tags').getFullList()
  tags.value = result
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

import type { MultiSelectFilterEvent } from 'primevue/multiselect';

function onFilter(event: MultiSelectFilterEvent) {
  filterValue.value = event.value;
}

async function createTag(tagName: string) {
  // Check if tag already exists
  const existingTag = tags.value.find(
    (tag) => tag.name.toLowerCase() === tagName.toLowerCase()
  )
  if (existingTag) {
    // Add to selectedTags if not already selected
    if (!selectedTags.value.some((tag) => tag.id === existingTag.id)) {
      selectedTags.value.push(existingTag)
    }
  } else {
    try {
      // Create new tag in 'tags' collection
      const newTag = await client.collection('tags').create({ name: tagName })
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
  const tagIds = selectedTags.value.map((tag) => tag.id)

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
      router.push({ name: 'post-view', params: { postId: newPost.id } })
    }
  } catch (error) {
    console.error(error)
  }
}
</script>
