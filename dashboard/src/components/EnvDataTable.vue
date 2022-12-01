<script setup lang="ts">
import type { TypesEnvironmentListResponse } from '~/composables/types/scheme'
const props = defineProps<{ datas: TypesEnvironmentListResponse }>()
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
            SSH | Jupyter | Tensorboard
          </td>
          <td class="py-4 px-6">
            "Unknown for now"
          </td>
          <td class="py-4 px-6">
            <span
              class="bg-green-100 text-green-800 text-xs font-semibold mr-2 px-2.5 py-0.5 rounded dark:bg-green-200 dark:text-green-900"
            >
              {{ e.status!.phase }}
            </span>
          </td>
          <td class="py-4 px-6">
            <i-mdi-bin class="h-6 w-6" />
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
