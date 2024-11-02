<template>
  <main class="flex h-full flex-col items-center justify-center">
    <Card>
      <template #title>Login</template>
      <template #content>
        <form class="grid grid-cols-1 gap-6 lg:w-96" @submit.prevent="authUser">
          <FloatLabel variant="on">
            <InputText
              type="email"
              id="email"
              v-model="email"
              aria-describedby="email-help"
              class="w-full"
            />
            <label for="email">Email</label>
          </FloatLabel>

          <FloatLabel variant="on">
            <Password
              v-model="password"
              toggleMask
              :feedback="false"
              class="w-full"
              pt:pcInputText:root:class="w-full"
            />
            <label for="password"> Password</label>
          </FloatLabel>

          <Button type="submit"> Login </Button>
        </form>
      </template>
      <template #footer>
        <div class="flex flex-col items-center">
          <p>Don't have an account?</p>
          <router-link
            class="text-sky-500 underline underline-offset-2 transition-colors duration-200 hover:text-sky-700 dark:hover:text-sky-300"
            to="/register"
          >
            Register here
          </router-link>
        </div>
      </template>
    </Card>
  </main>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import pbClient from '@/pocketbase'
import Card from 'primevue/card'
import FloatLabel from 'primevue/floatlabel'
import InputText from 'primevue/inputtext'
import Password from 'primevue/password'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'

const router = useRouter()

const toast = useToast()

// Local reactive variables
const email = ref('')
const password = ref('')

// Function to authenticate the user based on email and password
const authUser = async () => {
  try {
    // Authenticate the user via email and password
    const userData = await pbClient
      .collection('users')
      .authWithPassword(email.value, password.value)

    if (userData) {
      router.push({ path: '/' })
    }
  } catch (error) {
    console.log(error)
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Invalid email or password',
      life: 5000,
    })
  }
}
</script>
