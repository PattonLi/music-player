import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

//引入tailwind
import '@/assets/style/base.scss'
import '@/assets/style/theme.scss'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
