<script setup lang="ts">
import { useEnvdFetch } from '~/composables/request'
import type { TypesEnvironmentListResponse } from '~/composables/types/scheme'
const { userInfo } = useUserStore()
const { data, isFinished } = useEnvdFetch(`/users/${userInfo.username}/environments`).get().json<TypesEnvironmentListResponse>()
</script>

<template>
  <div class="min-h-screen flex">
    <Sidebar class="flex-none z-40 h-screen overflow-y-auto bg-white w-80 dark:bg-gray-800" />
    <div class="container flex-1 py-5 mx-5">
      <Navbar />
      <div class="container p-5">
        <div class="container py-5">
          <span class="font-semibold text-lg ">envd Environments</span>
        </div>
        <EnvDataTable v-if="isFinished" :datas="data!" />
      </div>
    </div>
  </div>
</template>

<route lang="yaml">
meta:
  layout: default
</route>
