import {createApp} from "vue"
import app from "./app.vue"

import {router} from "./router.ts";

// element message
import {ElMessage} from "element-plus"
import "element-plus/es/components/message/style/css"

createApp(app)
    .use(router)
    .use(ElMessage)
    .mount("#app")
