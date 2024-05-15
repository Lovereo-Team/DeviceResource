import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from './router'; // 导入Vue Router的实例

const app = createApp(App)

app.use(ElementPlus) // 添加 Element Plus 插件
app.use(router) // 添加 Vue Router 实例

app.mount('#app')
