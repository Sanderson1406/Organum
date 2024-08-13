import { createApp } from 'vue'
import App from '/src/App.vue'
import router from '/src/router/router.js'
import './assets/main.css'
import axios from 'axios'

const app = createApp(App)

app.use(router)
app.mount('#app')

const instance = axios.create({
  baseURL: 'http://localhost:8081'
})

export default instance
