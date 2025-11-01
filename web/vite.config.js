// 导入必要的模块
import { defineConfig, loadEnv } from 'vite'
import { fileURLToPath } from 'url' // 用于处理路径别名
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools' // 确保已安装该插件

// defineConfig 接收一个函数，函数参数包含 mode（环境模式）
export default defineConfig(({ mode }) => {
    // 在函数内部加载环境变量（必须放在函数里，因为需要 mode 参数）
    const env = loadEnv(mode, process.cwd())

    // 返回配置对象
    return {
        plugins: [
            vue(),
            vueDevTools(), // 确保插件正确调用
        ],
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url)) // 路径别名配置
            },
        },
        server: {
            // 前端端口：使用 .env 中的 VITE_CLI_PORT（7788）
            port: Number(env.VITE_CLI_PORT),
            strictPort: true,

            // 代理配置：转发 /api 请求到后端 7789 端口
            proxy: {
                [env.VITE_BASE_API]: { // 匹配前端请求的 /api 前缀
                    target: `${env.VITE_BASE_PATH}:${env.VITE_SERVER_PORT}`, // 后端地址（http://127.0.0.1:7789）
                    changeOrigin: true, // 解决跨域
                    // 若后端接口不带 /api 前缀，取消下面注释（去掉请求中的 /api）
                    // rewrite: (path) => path.replace(env.VITE_BASE_API, '')
                }
            }
        }
    }
})