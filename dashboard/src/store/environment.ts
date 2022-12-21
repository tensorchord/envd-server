import type { Ref } from '@vue/reactivity'
import { acceptHMRUpdate, defineStore } from 'pinia'

import { useEnvdFetch } from '~/composables/request'
import type { TypesEnvironment, TypesEnvironmentListResponse } from '~/composables/types/scheme'

export const useEnvStore = defineStore('envs', () => {
  const { userInfo } = useUserStore()
  const { data, execute } = useEnvdFetch(`/users/${userInfo.username}/environments`).get().json<TypesEnvironmentListResponse>()
  const envs = ref<TypesEnvironment[]>([])

  async function fetchEnvs(): Promise<TypesEnvironment[]> {
    await execute()
    return data.value!.items!
  }

  async function refreshEnvs(): Promise<void> {
    envs.value = await fetchEnvs()
  }

  function getEnvsRef(): Ref<TypesEnvironment[]> {
    return envs
  }

  async function deleteEnv(id: string): Promise<void> {
    await useEnvdFetch(`/users/${userInfo.username}/environments/${id}`).delete().json()
    await refreshEnvs()
  }

  async function getEnvInfo(name: string): Promise<TypesEnvironment> {
    const envs = getEnvsRef()
    const env = envs.value.find(env => env.name === name)
    return env!
  }

  return { getEnvsRef, refreshEnvs, deleteEnv, getEnvInfo }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useEnvStore, import.meta.hot))
