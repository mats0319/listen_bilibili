import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import path from "path"; // require @types/node

// element-plus auto import
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),

        AutoImport({ resolvers: [ElementPlusResolver()] }),
        Components({ resolvers: [ElementPlusResolver()] }),
    ],
    resolve: {
        alias: {
            "@": path.resolve(__dirname, "./src"),
            "~": path.resolve(__dirname, "./"), // for import 'package.json'
        },
    },
    clearScreen: false,
    server: {
        host: "127.0.0.1",
        port: 9694,
    },
});
