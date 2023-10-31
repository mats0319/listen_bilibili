import { createApp } from "vue"
import app from "./app.vue"

import { router } from "./router.ts";

// element message
import { ElMessage } from "element-plus"
import "element-plus/es/components/message/style/css"

// pinia, with persist
import { createPinia } from "pinia"
// import piniaPluginPersistedState from "pinia-plugin-persistedstate"
// todo：使用持久化组件库，无法build，错误代码为ts2345，即参数类型错误

createApp(app)
    .use(router)
    .use(ElMessage)
    .use(createPinia()/*.use(piniaPluginPersistedState)*/)
    .mount("#app")
