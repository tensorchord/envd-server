import { acceptHMRUpdate, defineStore } from 'pinia'
import { useFetch } from '@vueuse/core'

export const useUserStore = defineStore('user', () => {

  const userInfo = ref(useLocalStorage("userInfo", {
    username: '',
    jwtToken: '',
  }))

  function logIn() {
    // TODO
    let { } = useFetch('/api/login', {})
  }

  function setUser(username: string) {
    userInfo.value.username = username
  }
  return {
    userInfo,
    logIn,
    setUser,
  }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useUserStore, import.meta.hot))
