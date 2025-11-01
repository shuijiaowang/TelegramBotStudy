<template>
  <div class="login-container">
    <div class="login-card">
      <h2 class="login-title">系统登录</h2>
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label>用户名</label>
          <input
              type="text"
              v-model="loginForm.username"
              placeholder="请输入用户名"
              required
          >
        </div>
        <div class="form-group">
          <label>密码</label>
          <input
              type="password"
              v-model="loginForm.password"
              placeholder="请输入密码"
              required
          >
        </div>
        <button type="submit" class="login-btn">登录</button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useUserStore } from '@/stores/user.js'
import { ElMessage } from 'element-plus'
import {login} from "@/api/user.js";

// 登录表单数据
const loginForm = ref({
  username: '',
  password: ''
})

// 获取用户存储
const userStore = useUserStore()

// 处理登录逻辑
const handleLogin = async () => {
  if (!loginForm.value.username || !loginForm.value.password) {
    ElMessage.warning('请输入用户名和密码')
    return
  }

  const success = await userStore.loginIn(loginForm.value)
  if (success) {
    ElMessage.success('登录成功')
  } else {
    ElMessage.error('登录失败，请检查用户名和密码')
  }
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--color-background-soft);
}

.login-card {
  width: 350px;
  padding: 2rem;
  background-color: var(--color-background);
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

.login-title {
  text-align: center;
  margin-bottom: 1.5rem;
  color: var(--color-heading);
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: var(--color-text);
}

.form-group input {
  width: 100%;
  padding: 0.8rem;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  background-color: var(--color-background);
  color: var(--color-text);
}

.form-group input:focus {
  outline: none;
  border-color: hsla(160, 100%, 37%, 1);
}

.login-btn {
  width: 100%;
  padding: 0.8rem;
  background-color: hsla(160, 100%, 37%, 1);
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.login-btn:hover {
  background-color: hsla(160, 100%, 30%, 1);
}
</style>