import { acceptHMRUpdate, defineStore } from 'pinia'
import type { TypesAuthNRequest, TypesAuthNResponse } from '~/composables/types/scheme'

export const useUserStore = defineStore('user', () => {
  const userInfo = ref(useLocalStorage('envd/userInfo', {
    username: '',
    jwtToken: '',
  }))
  const isLogin = ref(false)

  async function getToken(): Promise<string> {
    return userInfo.value.jwtToken
  }

  async function login(username: string, password: string): Promise<boolean> {
    const request: TypesAuthNRequest = {
      login_name: username,
      password,
    }
    const { data, statusCode } = await useEnvdFetch('/login').post(request).json<TypesAuthNResponse>()
    if (statusCode.value !== 200) {
      return false
    }
    else {
      userInfo.value = {
        username,
        jwtToken: data.value!.identity_token!,
      }
      return true
    }
  }

  function setUser(username: string) {
    userInfo.value.username = username
  }
  return {
    userInfo,
    isLogin,
    login,
    setUser,
    getToken,
  }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useUserStore, import.meta.hot))
