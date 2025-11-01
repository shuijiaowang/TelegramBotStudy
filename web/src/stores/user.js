import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useStorage } from '@vueuse/core'
import { login, getUserInfo } from '@/api/user'
import { ElLoading, ElMessage } from 'element-plus'
import router from '@/router'

export const useUserStore = defineStore('user', () => {
    // 状态管理（仅保留username和uuid）
    const token = useStorage('token', '')
    const userInfo = ref({
        username: '',  // 仅保留用户名
        uuid: ''       // 仅保留uuid
    })

    // 设置用户信息（只处理username和uuid）
    const setUserInfo = (info) => {
        userInfo.value = {
            username: info.username || userInfo.value.username,
            uuid: info.uuid || userInfo.value.uuid
        }
    }

    // 设置token
    const setToken = (newToken) => {
        token.value = newToken
    }

    // 清理用户状态（仅重置保留的字段）
    const clearUserState = () => {
        token.value = ''
        userInfo.value = {
            username: '',
            uuid: ''
        }
    }

    // 登录逻辑（适配精简后的用户信息）
    const loginIn = async (loginForm) => {
        const loading = ElLoading.service({
            fullscreen: true,
            text: '登录中...'
        })

        try {
            // 验证表单
            if (!loginForm.username || !loginForm.password) {
                ElMessage.warning('请输入用户名和密码')
                return false
            }

            // 调用登录接口
            const res = await login(loginForm)
            if (res.code !== 0) {
                return false
            }

            // 保存用户信息（从响应data中直接获取username和uuid）
            setToken(res.data.token)
            setUserInfo({
                username: res.data.username,
                uuid: res.data.uuid
            })

            // 直接跳转首页（不再依赖authority字段）
            const targetRoute = 'home'
            if (router.hasRoute(targetRoute)) {
                await router.replace({ name: targetRoute })
                console.log("为什么不跳转？")
            } else {
                ElMessage.error('首页路由不存在，请联系管理员')
                return false
            }

            ElMessage.success('登录成功')
            return true
        } catch (error) {
            console.error('登录失败:', error)
            return false
        } finally {
            loading.close()
        }
    }

    // 退出登录
    const logout = async () => {
        clearUserState()
        await router.replace({ name: 'login' })
        ElMessage.success('已退出登录')
    }

    return {
        token,
        userInfo,
        setUserInfo,
        setToken,
        clearUserState,
        loginIn,
        logout
    }
})