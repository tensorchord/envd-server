import { createFetch } from '@vueuse/core'

export const useEnvdFetch = createFetch({
  baseUrl: `${import.meta.env.VITE_BASE_HOST ? import.meta.env.VITE_BASE_HOST : ''}/api/v1`,
  options: {
    async beforeFetch({ options }) {
      const { getToken } = useUserStore()
      const myToken = await getToken()
      options.headers = {
        ...options.headers,
        Authorization: `Bearer ${myToken}`,
      }
      return { options }
    },
  },
})
