<template>
  <v-form @submit.prevent="register">
    <v-container>
      <v-text-field
        v-model="username"
        label="Username"
        required
        :rules="[v => !!v || 'ユーザー名は必須です']"
      ></v-text-field>
      <v-text-field
        v-model="password"
        label="Password"
        type="password"
        required
        :rules="[
          v => !!v || 'パスワードは必須です',
          v => (v && v.length >= 8) || 'パスワードは8文字以上である必要があります'
        ]"
      ></v-text-field>
      <v-text-field
        v-model="confirmPassword"
        label="Confirm Password"
        type="password"
        required
        :rules="[
          v => !!v || 'パスワード（確認）は必須です',
          v => v === password || 'パスワードが一致しません'
        ]"
      ></v-text-field>
      <v-text-field
        v-model="displayName"
        label="Display Name"
      ></v-text-field>
      <v-btn type="submit" color="primary" block :disabled="!isFormValid">Register</v-btn>
      
      <!-- エラーメッセージ表示 -->
      <v-alert
        v-if="errorMessage"
        type="error"
        class="mt-4"
      >
        {{ errorMessage }}
      </v-alert>
    </v-container>
  </v-form>
</template>

<script setup lang="ts">
import { useAPI } from '@/composables/useApi'
import type { errorResponse } from '@/types';

const username = ref('')
const password = ref('')
const confirmPassword = ref('')
const displayName = ref('')
const errorMessage = ref('')
const router = useRouter()

const isFormValid = computed(() => {
  return username.value && password.value && confirmPassword.value && (password.value === confirmPassword.value)
})

const handleError = (e: errorResponse) => {
  if (e.message.includes('username already exists')) {
    errorMessage.value = 'このユーザー名は既に使用されています'
  } else {
    errorMessage.value = '登録に失敗しました。もう一度お試しください。'
  }
}

const register = async () => {
  if (!isFormValid.value) {
    errorMessage.value = 'フォームの入力内容を確認してください。'
    return
  }

  try {
    errorMessage.value = '' // エラーメッセージをリセット
    const { data, error } = await useAPI('/auth/register', {
      method: 'POST',
      body: {
        username: username.value,
        password: password.value,
        display_name: displayName.value
      },
    })

    if (error?.value) {
      handleError(error.value.data)
      return
    }

    // 登録成功
    console.log('Registration successful', data.value)
    navigateTo('/login') // ログインページにリダイレクト
  } catch (err) {
    console.error('Registration failed:', err)
    errorMessage.value = '予期せぬエラーが発生しました。'
  }
}
</script>
