import { createFetch } from "@vueuse/core"

export const useEnvdFetch = createFetch({
  baseUrl: `${import.meta.env.VITE_BASE_HOST ? import.meta.env.VITE_BASE_HOST : ""}/api/v1`,
})
