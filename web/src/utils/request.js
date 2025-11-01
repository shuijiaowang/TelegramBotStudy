import axios from 'axios'
import { useUserStore } from '@/stores/user'
import { ElLoading, ElMessage } from 'element-plus'
import router from '@/router'

// 创建axios实例
const service = axios.create({
    baseURL: import.meta.env.VITE_BASE_API,
    timeout: 10000,
    headers: {
        'Content-Type': 'application/json'
    }
})

// 加载状态管理（解决多请求冲突问题）
let loadingCount = 0
let loadingInstance = null

// 显示加载
const showLoading = (config) => {
    if (config.donNotShowLoading) return

    loadingCount++
    if (loadingCount === 1) {
        loadingInstance = ElLoading.service({
            fullscreen: true,
            text: '加载中...',
            background: 'rgba(0, 0, 0, 0.1)'
        })
    }
}

// 隐藏加载
const hideLoading = (config) => {
    if (config.donNotShowLoading) return

    loadingCount--
    if (loadingCount <= 0) {
        loadingInstance?.close()
        loadingInstance = null
        loadingCount = 0
    }
}

// 请求拦截器
service.interceptors.request.use(
    (config) => {
        showLoading(config)
        const userStore = useUserStore()

        // 添加认证信息
        if (userStore.token) {
            config.headers['Authorization'] = `Bearer ${userStore.token}`; // 拼接 Bearer 前缀
        }
        return config
    },
    (error) => {
        hideLoading(error.config)
        ElMessage.error(`请求准备失败: ${error.message || '未知错误'}`)
        return Promise.reject(error)
    }
)

// 响应拦截器
service.interceptors.response.use(
    (response) => {
        hideLoading(response.config)
        const userStore = useUserStore()
        const res = response.data

        // 更新token（如果有新token）
        if (response.headers['new-token']) {
            userStore.setToken(response.headers['new-token'])
        }

        // 处理业务状态码
        if (typeof res.code !== 'undefined') {
            // 成功状态码（根据实际业务调整）
            if (res.code === 0) {
                return res
            } else {
                ElMessage.error(res.msg || `操作失败（${res.code}）`)
                return Promise.reject(res)
            }
        }

        return res
    },
    (error) => {
        hideLoading(error.config)

        // 网络错误处理
        if (!error.response) {
            ElMessage.error('网络连接异常，请检查网络')
            return Promise.reject(error)
        }

        // 401 未授权处理
        if (error.response.status === 401) {
            ElMessage.error('登录已过期，请重新登录')
            const userStore = useUserStore()
            userStore.clearUserState() // 清理用户状态
            router.replace({ name: 'login' })
        } else {
            // 其他HTTP错误
            const status = error.response.status
            const msg = error.response.data?.msg || `请求错误（${status}）`
            ElMessage.error(msg)
        }

        return Promise.reject(error)
    }
)

export default service