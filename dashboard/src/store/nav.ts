import { acceptHMRUpdate, defineStore } from 'pinia'

export const useNav = defineStore('nav', () => {
  const header = ref('')
  function setNavHeader(newHeader: string) {
    header.value = newHeader
  }
  return { header, setNavHeader }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useUserStore, import.meta.hot))
