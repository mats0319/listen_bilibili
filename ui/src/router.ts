import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router"

const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        name: "home",
        component: () => import("@/views/home.vue"),
        children: [
            {
                path: "/listen-bilibili",
                name: "ListenBilibili",
                component: () => import("@/views/listen_bilibili.vue"),
            }
        ]
    },
]

export const router = createRouter({
    history: createWebHashHistory(),
    routes: routes,
})
