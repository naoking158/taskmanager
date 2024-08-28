import { ref } from 'vue'

export const useAuth = () => {
  const token = useCookie('auth_token')
  const user = ref(null)

  const login = async (username: string, password: string) => {
    return true
    // try {
    //   // ここに実際のAPI呼び出しを実装します
    //   const response = await $fetch('/api/auth/login', {
    //     method: 'POST',
    //     body: { username, password },
    //   })
    //   token.value = response.token
    //   user.value = response.user
    //   return true
    // } catch (error) {
    //   console.error('Login failed', error)
    //   return false
    // }
  }

  const logout = () => {
    token.value = null
    user.value = null
  }

  return {
    login,
    logout,
    user,
  }
}
