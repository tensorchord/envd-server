import type { Ref } from '@vue/reactivity'
import { acceptHMRUpdate, defineStore } from 'pinia'

import { useEnvdFetch } from '~/composables/request'
import type { TypesEnvironment, TypesEnvironmentListResponse } from '~/composables/types/scheme'

export const useEnvStore = defineStore('envs', () => {
  const { userInfo } = useUserStore()
  const { data, execute } = useEnvdFetch(`/users/${userInfo.username}/environments`, { immediate: false }).get().json<TypesEnvironmentListResponse>()
  const envs = ref<TypesEnvironment[]>([])

  async function fetchEnvs(): Promise<TypesEnvironment[]> {
    await execute()
    return data.value!.items!
  }

  async function refreshEnvs(): Promise<void> {
    envs.value = await fetchEnvs()
  }

  async function getEnvs(): Promise<Ref<TypesEnvironment[]>> {
    if (envs.value.length === 0)
      await refreshEnvs()
    return envs
  }

  async function deleteEnvs(id: string): Promise<void> {
    await useEnvdFetch(`/environments/${id}`).delete().json()
    await refreshEnvs()
  }

  return { getEnvs, refreshEnvs, deleteEnvs }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useEnvStore, import.meta.hot))
