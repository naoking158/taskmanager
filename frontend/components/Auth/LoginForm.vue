<template>
  <v-form @submit.prevent="login">
    <v-container>
      <v-text-field
        v-model="username"
        label="Username"
        required
      ></v-text-field>
      <v-text-field
        v-model="password"
        label="Password"
        type="password"
        required
      ></v-text-field>
      <v-btn type="submit" color="primary" block>Login</v-btn>

      <!-- エラーメッセージ表示 -->
      <v-alert
        v-if="errorMessage"
        type="error"
        class="mt-4">
        {{ errorMessage }}
      </v-alert>
    </v-container>
  </v-form>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useAPI } from '@/composables/useApi'
import type { errorResponse } from '@/types'

const username = ref('')
const password = ref('')
const errorMessage = ref('')
const authStore = useAuthStore()
const router = useRouter()
const headers = useRequestHeaders(['cookie'])

const handleError = (e: errorResponse) => {
  if (e.message.includes('User not found')) {
    errorMessage.value = 'ユーザー名が間違っています'
  } else if (e.message.includes('Invalid password')) {
    errorMessage.value = 'パスワードが間違っています'
  } else {
    errorMessage.value = 'ログインに失敗しました。もう一度お試しください。'
  }
}

const login = async () => {
  try {
    const { data, error } = await useAPI('/auth/login', {
      method: 'POST',
      body: { username: username.value, password: password.value },
    })
    // const { data, error } = await useAPI('/auth/login', {
    //   method: 'POST',
    //   body: { username: username.value, password: password.value },
    // })

    if (error?.value) {
      handleError(error.value.data)
      return
    }
    
    authStore.setToken(data.value)
    // トークンを使用してユーザー情報を取得する処理をここに追加
    
    // ログイン後のリダイレクト
    router.push('/')
  } catch (err) {
    console.error('Login failed:', err)
    errorMessage.value = '予期せぬエラーが発生しました。'
    // エラーメッセージの表示処理をここに追加
  }
}

</script>
