<template>
  <div class="p-6">
    <Editor v-model="text" editorStyle="height: 320px">
      <template v-slot:toolbar>
        <span class="ql-formats">
          <button v-tooltip.bottom="'Bold'" class="ql-bold"></button>
          <button v-tooltip.bottom="'Italic'" class="ql-italic"></button>
          <button v-tooltip.bottom="'Underline'" class="ql-underline"></button>
            <button v-tooltip.bottom="'Strike'" class="ql-strike"></button>
            <button v-tooltip.bottom="'Link'" class="ql-link"></button>
            <button v-tooltip.bottom="'Code'" class="ql-code"></button>
        </span>
      </template>
    </Editor>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import Editor from 'primevue/editor'

const text = ref<string>('')

onMounted(() => {
  text.value = '<p>Hi <strong>how</strong> are ya</p>'
})
</script>
