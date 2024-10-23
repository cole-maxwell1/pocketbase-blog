<template>
  <main class="flex flex-col justify-center items-center h-screen">
    <Card>
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
import { useRouter } from 'vue-router'
import pbClient from '@/pocketbase'
import FloatLabel from 'primevue/floatlabel'
import Password from 'primevue/password'
import Divider from 'primevue/divider'
import InputText from 'primevue/inputtext'
import Card from 'primevue/card'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'

const router = useRouter()
const toast = useToast()

const form = ref({
  firstName: '',
  lastName: '',
  email: '',
  username: '',
  password: '',
  passwordConfirm: '',
})

// Function to create a new user
async function createUser() {
  try {
    if (validateInput()) {
      // Create new user
      const user = await pbClient.collection('users').create({
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

        // After successful user registration, redirect to home
        router.push({ name: 'home' })
      } else {
        console.log('Error')
      }
    }
  } catch (error: unknown) {
    console.log(error)
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Could not register user. Please try again.',
      life: 3000,
    })
  }
}

// Function to authenticate the user based on email and password
async function authUser() {
  try {
    // Authenticate the user via email and password
    await pbClient
      ?.collection('users')
      .authWithPassword(form.value.email, form.value.password)

    router.push({ path: '/' })
  } catch (error) {
    console.error(error)
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Could not authenticate user. Please try again.',
      life: 3000,
    })
    router.push({ path: '/login' })
  }
}

const validateInput = () => {
  if (form.value.password !== form.value.passwordConfirm) {
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Passwords do not match. Please try again.',
      life: 5000,
    })
    return false
  } else {
    return true
  }
}
</script>
