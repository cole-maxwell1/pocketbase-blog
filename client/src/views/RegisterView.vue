<template>
  <main class="flex flex-col justify-center items-center h-screen">
    <Card>
      <template v-if="errorOccurred" #header>
        <div class="px-4 pt-4">
          <Message closable severity="error">
            {{ errorMessage }}
          </Message>
        </div>
      </template>
      <template #title>
        <h1>Register an Account</h1>
      </template>
      <template #content>
        <form class="gap-6 grid grid-cols-1" @submit.prevent="createUser">
          <div class="gap-6 grid grid-cols-1 lg:grid-cols-2">
            <FloatLabel variant="on">
              <InputText
                type="text"
                id="firstName"
                v-model="form.firstName"
                aria-describedby="firstName-help"
                class="w-full"
              />
              <label for="firstName">First Name</label>
            </FloatLabel>
            <FloatLabel variant="on">
              <InputText
                type="text"
                id="lastName"
                v-model="form.lastName"
                aria-describedby="lastName-help"
                class="w-full"
              />
              <label for="lastName">Last Name</label>
            </FloatLabel>
          </div>
          <FloatLabel variant="on">
            <InputText
              type="text"
              id="username"
              v-model="form.username"
              aria-describedby="username-help"
              class="w-full"
            />
            <label for="username">Username</label>
          </FloatLabel>
          <FloatLabel variant="on">
            <InputText
              type="email"
              id="email"
              v-model="form.email"
              aria-describedby="email-help"
              class="w-full"
            />
            <label for="email">Email</label>
          </FloatLabel>
          <FloatLabel variant="on">
            <Password
              v-model="form.password"
              toggleMask
              class="w-full"
              pt:pcInputText:root:class="w-full"
            >
              <template #header>
                <div class="mb-4 font-semibold text-xm">Pick a password</div>
              </template>
              <template #footer>
                <Divider />
                <ul class="my-0 ml-2 pl-2 leading-normal">
                  <li>At least one lowercase</li>
                  <li>At least one uppercase</li>
                  <li>At least one numeric</li>
                  <li>Minimum 8 characters</li>
                </ul>
              </template>
            </Password>
            <label for="password"> Password</label>
          </FloatLabel>
          <FloatLabel variant="on" pt:root:class="w-full">
            <Password
              v-model="form.passwordConfirm"
              toggleMask
              :feedback="false"
              class="w-full"
              pt:pcInputText:root:class="w-full"
            />
            <label for="password"> Repeat Password</label>
          </FloatLabel>
          <Button type="submit"> Sign Up </Button>
        </form>
      </template>
      <template #footer>
        <div class="flex flex-col items-center">
          <p>Already have an account?</p>
          <router-link
            class="text-sky-500 hover:text-sky-700 dark:hover:text-sky-300 underline underline-offset-2 transition-colors duration-200"
            to="/login"
          >
            Login here
          </router-link>
        </div>
      </template>
    </Card>
  </main>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'
import client from '@/pocketbase'
import FloatLabel from 'primevue/floatlabel'
import Password from 'primevue/password'
import Divider from 'primevue/divider'
import InputText from 'primevue/inputtext'
import Card from 'primevue/card'
import Button from 'primevue/button'
import Message from 'primevue/message'

// Init the store for user
const userStore = useUserStore()

const router = useRouter()

const form = ref({
  firstName: '',
  lastName: '',
  email: '',
  username: '',
  password: '',
  passwordConfirm: '',
})

const errorOccurred = ref(false)
const errorMessage = ref('')

// Function to create a new user
async function createUser() {
  errorOccurred.value = false
  try {
    if (validateInput()) {
      // Create new user
      const user = await client.collection('users').create({
        firstName: form.value.firstName.toLowerCase(),
        lastName: form.value.lastName.toLowerCase(),
        username: form.value.username,
        email: form.value.email.toLowerCase(),
        password: form.value.password,
        passwordConfirm: form.value.passwordConfirm,
        emailVisibility: false,
      })
      if (user) {
        // Authenticate the user in order to set the username
        await authUser()

        // After successful user registration, redirect to dashboard
        router.push({ path: '/about' })
      } else {
        console.log('Error')
      }
    } else {
      alert("Password doesn't match")
    }
  } catch (error: unknown) {
    errorOccurred.value = true
    errorMessage.value = 'Could not create user. Please try again.'
    console.log(error)
  }
}

// Function to authenticate the user based on email and password
async function authUser() {
  try {
    // Authenticate the user via email and password
    const userData = await client
      ?.collection('users')
      .authWithPassword(form.value.email, form.value.password)
    if (userData) {
      userStore.userID = userData.record.id
      userStore.username = userData.record.profile?.username
      userStore.userProfileID = userData.record.profile?.id
      router.push({ path: '/' })
    }
  } catch (error) {
    console.log(error)
  }
}

// Simple utility function to validate input. Easily extendable with additional checks if needed
const validateInput = () => {
  if (form.value.password !== form.value.passwordConfirm) {
    return false
  } else {
    return true
  }
}
</script>
