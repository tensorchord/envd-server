<script setup lang="ts">
import dayjs from 'dayjs'
import type { TypesEnvironmentListResponse } from '~/composables/types/scheme'
const props = defineProps<{ datas: TypesEnvironmentListResponse }>()
defineEmits(['deleteEnv'])
</script>

<template>
  <div class="overflow-x-auto relative">
    <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
      <thead class="text-gray-600 border-b border-t">
        <tr>
          <th scope="col" class="py-3 px-3 font-semibold">
            Name
          </th>
          <th scope="col" class="py-3 px-6 font-semibold">
            Docker Image
          </th>
          <th scope="col" class="py-3 px-6 font-semibold">
            Services
          </th>
          <th scope="col" class="py-3 px-6 font-semibold">
            Created
          </th>
          <th scope="col" class="py-3 px-6 font-semibold">
            Status
          </th>
          <th scope="col" class="py-3 px-6 font-semibold">
            Operation
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="e in props.datas.items" :key="e.name" class="bg-white dark:bg-gray-800 dark:border-gray-700 text-gray-500">
          <th scope="row" class="py-4 px-3 font-medium text-gray-500 whitespace-nowrap dark:text-white">
            {{ e.name }}
          </th>
          <td class="py-4 px-6">
            <span class="text-blue-500">{{ e.spec!.image }}</span>
          </td>
          <td class="py-4 px-6">
            {{ e.spec?.ports?.map((e) => `${e.name} : ${e.port}`).join(' | ') }}
          </td>
          <td class="py-4 px-6">
            {{ dayjs(e.created_at! * 1000).fromNow() }}
          </td>
          <td class="py-4 px-6">
            <StatusTag :status="e.status!.phase" />
          </td>
          <td class="py-4 px-6">
            <button class="hover:bg-grey-200">
              <i-mdi-bin class="h-6 w-6" @click="$emit('deleteEnv', e.name)" />
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
