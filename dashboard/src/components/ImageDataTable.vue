<script setup lang="ts">
import dayjs from 'dayjs'
import type { TypesImageListResponse } from '~/composables/types/scheme'
const props = defineProps<{ data: TypesImageListResponse }>()
</script>

<template>
  <div class="overflow-x-auto relative">
    <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
      <thead class="text-gray-600 border-b border-t">
        <tr>
          <th scope="col" class="py-3 px-3 font-semibold">
            Image Name
          </th>
          <th scope="col" class="py-3 px-6 font-semibold">
            Digest
          </th>
          <th scope="col" class="py-3 px-6 font-semibold">
            Labels
          </th>
          <th scope="col" class="py-3 px-6 font-semibold">
            Created
          </th>
          <th scope="col" class="py-3 px-6 font-semibold">
            Size
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="e in props.data.items" :key="e.name" class="bg-white dark:bg-gray-800 dark:border-gray-700 text-gray-500">
          <th scope="row" class="py-4 px-3 font-medium text-gray-500 whitespace-nowrap dark:text-white">
            {{ e.name }}
          </th>
          <td class="py-4 px-6">
            <span class="text-blue-500">{{ e.digest }}</span>
          </td>
          <td class="py-4 px-6">
            {{ Object.entries(e.labels!).map((value) => {
              return `${value[0]} : ${value[1]}`
            }).join(' | ') }}
          </td>
          <td class="py-4 px-6">
            {{ dayjs(e.created! * 1000).fromNow() }}
          </td>
          <td class="py-4 px-6">
            {{ e.size }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
