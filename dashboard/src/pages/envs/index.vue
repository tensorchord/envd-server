<script setup lang="ts">
import { useEnvdFetch } from '~/composables/request'
import type { TypesEnvironmentListResponse } from '~/composables/types/scheme'
const { userInfo } = useUserStore()
const { data, isFinished, execute } = useEnvdFetch(`/users/${userInfo.username}/environments`).get().json<TypesEnvironmentListResponse>()

const deleteEnv = async (envName: string) => {
  const { isFinished } = await useEnvdFetch(`/users/${userInfo.username}/environments/${envName}`).delete()
  if (isFinished)
    execute()
}
</script>

<template>
  <div class="container flex-1 py-5 mx-5">
    <Navbar />
    <div class="container p-5">
      <div class="container py-5">
        <span class="font-semibold text-lg ">envd Environments</span>
      </div>
      <EnvDataTable v-if="isFinished" :datas="data!" @delete-env="deleteEnv" />
    </div>
  </div>
</template>

<route lang="yaml">
meta:
  layout: dashboard
</route>
