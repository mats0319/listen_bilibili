import { createApp } from "vue";
import app from "./app.vue";

import { router } from "./router.ts";

// element message
import { ElMessage } from "element-plus";
import "element-plus/es/components/message/style/css";

// pinia
import { createPinia } from "pinia";

createApp(app).use(router).use(ElMessage).use(createPinia()).mount("#app");
