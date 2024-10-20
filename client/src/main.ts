import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import PrimeVue from 'primevue/config'
import Aura from '@primevue/themes/aura'
import { definePreset } from '@primevue/themes'
import { createPinia } from 'pinia'
import Tooltip from 'primevue/tooltip'
import 'primeicons/primeicons.css'
import Editor from 'primevue/editor'
import type Quill from 'quill'
import ConfirmationService from 'primevue/confirmationservice'
import ToastService from 'primevue/toastservice'

const app = createApp(App)

app.use(router)

const MyPreset = definePreset(Aura, {
  semantic: {
    primary: {
      50: '{sky.50}',
      100: '{sky.100}',
      200: '{sky.200}',
      300: '{sky.300}',
      400: '{sky.400}',
      500: '{sky.500}',
      600: '{sky.600}',
      700: '{sky.700}',
      800: '{sky.800}',
      900: '{sky.900}',
      950: '{sky.950}',
    },
  },
})

app.use(PrimeVue, {
  theme: {
    preset: MyPreset,
    options: {
      cssLayer: {
        name: 'primevue',
        order: 'tailwind-base, primevue, tailwind-utilities',
      },
    },
  },
})

const pinia = createPinia()
app.use(pinia)

app.directive('tooltip', Tooltip)

// Fix needed for Quill v2: https://github.com/primefaces/primevue/issues/5606#issuecomment-2203975395
// eslint-disable-next-line @typescript-eslint/no-explicit-any
;(Editor as any).methods.renderValue = function renderValue(
  this: { quill?: Quill },
  value: string,
) {
  if (this.quill) {
    if (value) {
      const delta = this.quill.clipboard.convert({ html: value })
      this.quill.setContents(delta, 'silent')
    } else {
      this.quill.setText('')
    }
  }
}

app.use(ConfirmationService)
app.use(ToastService)

app.mount('#app')
