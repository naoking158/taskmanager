<template>
  <v-container class="fill-height" fluid>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="6" lg="4">
        <v-card class="elevation-12">
          <v-toolbar color="primary" dark flat>
            <v-toolbar-title>Login</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-form @submit.prevent="handleLogin">
              <v-text-field
                v-model="username"
                label="Username"
                name="username"
                prepend-icon="mdi-account"
                type="text"
                required
              />
              <v-text-field
                v-model="password"
                label="Password"
                name="password"
                prepend-icon="mdi-lock"
                type="password"
                required
              />
              <v-card-actions>
                <v-spacer />
                <v-btn type="submit" color="primary">Login</v-btn>
              </v-card-actions>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref } from 'vue'
import { useAuth } from '~/composables/useAuth.ts'

const username = ref('')
const password = ref('')
const { login } = useAuth()

const router = useRouter()

const handleLogin = async () => {
  const success = await login(username.value, password.value);
  if (success) {
    // ログイン成功後にダッシュボードページにリダイレクト
    router.push('/dashboard')
  } else {
    alert('Login failed. Please check your credentials.')
  }
}

/* const login = async () => {
 *   // ここに login ロジックを実装します
 *   console.log('Login attempt', { username: username.value, password: password.value })
 *   // TODO: API呼び出しとエラーハンドリングを実装
 * } */
</script>
