import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router"

const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        name: "home",
        component: () => import("@/views/home.vue"),
        children: [
            {
                path: "/listen",
                name: "listen",
                component: () => import("@/views/listen.vue"),
            },
            {
                path: "/modify",
                name: "modify",
                component: () => import("@/views/modify.vue"),
            }
        ]
    },
]

export const router = createRouter({
    history: createWebHashHistory(),
    routes: routes,
})
