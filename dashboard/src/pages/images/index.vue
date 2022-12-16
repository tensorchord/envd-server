<script setup lang="ts">
import { useEnvdFetch } from '~/composables/request'
import type { TypesImageListResponse } from '~/composables/types/scheme'

const { userInfo } = useUserStore()
const { data, isFinished } = useEnvdFetch(`/users/${userInfo.username}/images`).get().json<TypesImageListResponse>()
</script>

<template>
  <div class="container flex-1 py-5 mx-5">
    <Navbar />
    <div class="container p-5">
      <div class="container py-5">
        <span class="font-semibold text-lg ">envd Environments</span>
      </div>
      <ImageDataTable v-if="isFinished" :data="data!" />
    </div>
  </div>
</template>

<route lang="yaml">
meta:
  layout: dashboard
</route>
