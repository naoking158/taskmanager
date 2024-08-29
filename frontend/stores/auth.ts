import { ref } from 'vue'
import { defineStore } from 'pinia'

interface User {
  id: string;
  username: string;
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(null)
  const user = ref<User | null>(null)

  function setToken(newToken: string) {
    token.value = newToken
  }

  function setUser(newUser: User) {
    user.value = newUser
  }

  function logout() {
    token.value = null
    user.value = null
  }

  return { token, user, setToken, setUser, logout }
})
