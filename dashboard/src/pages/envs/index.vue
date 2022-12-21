<script setup lang="ts">
import { FlexRender, createColumnHelper, getCoreRowModel, useVueTable } from '@tanstack/vue-table'
import dayjs from 'dayjs'
import type { Component } from 'vue'
import InfoModal from '~/components/InfoModal.vue'
import StatusTag from '~/components/StatusTag.vue'
import type { TypesEnvironment } from '~/composables/types/scheme'
import IMdiBin from '~icons/mdi/bin'

const { getEnvsRef, deleteEnv, getEnvInfo, refreshEnvs } = useEnvStore()
const modalInfo = ref<TypesEnvironment>()
const modal = ref<typeof InfoModal>()
const envs = getEnvsRef()
const showEnvDetail = async (name: string) => {
  modalInfo.value = await getEnvInfo(name)
  modal.value!.openModal()
}

const columnHelper = createColumnHelper<TypesEnvironment>()
const columns = [
  columnHelper.accessor('name', {
    id: 'name',
    cell: info => h('button', {
      class: 'text-blue-500 hover:text-blue-800',
      innerText: info.getValue(),
      onClick: async () => {
        await showEnvDetail(info.getValue()!)
      },
    }),
    header: () => 'Image Name',
  }),
  columnHelper.accessor('created_at', {
    id: 'created',
    cell: info => dayjs(info.getValue()! * 1000).fromNow(),
    header: () => 'Created',
  }),
  columnHelper.accessor('spec', {
    id: 'services',
    cell: info => info.getValue()!.ports?.map(e => `${e.name} : ${e.port}`).join(' | '),
    header: () => 'Services',
  }),
  columnHelper.accessor('status', {
    id: 'status',
    cell: info => h(StatusTag as Component, { status: info.getValue()?.phase }),
    header: () => 'Status',
  }),
  columnHelper.accessor('name', {
    id: 'delete',
    cell: info => h('button', {
      class: 'hover:bg-grey-200',
    }, [
      h(IMdiBin, {
        class: 'h-6 w-6',
        onClick: async () => {
          await deleteEnv(info.getValue()!)
          await refreshEnvs()
        },
      }),
    ]),
    header: () => 'Operations',
  }),
]

const data = ref<TypesEnvironment[]>([])

const table = useVueTable<TypesEnvironment>({
  columns,
  get data() {
    return data.value
  },
  getCoreRowModel: getCoreRowModel(),
})

watch(envs, (newVal) => {
  data.value = newVal
})

onMounted(async () => {
  await refreshEnvs()
})
</script>

<template>
  <InfoModal ref="modal">
    <template #header>
      Image Detail
    </template>
    <template #body>
      <div class="pt-2">
        <dl class="mt-2 border-t border-b border-gray-200 divide-y divide-gray-200">
          <div class="py-3 flex justify-between text-sm font-medium">
            <dt class="text-gray-500 mr-6">
              Name
            </dt>
            <dd class="text-gray-900 text-">
              {{ modalInfo!.name! }}
            </dd>
          </div>
          <div class="py-3 flex justify-between text-sm font-medium">
            <dt class="text-gray-500 mr-6">
              Labels
            </dt>
            <dd class="text-gray-900 text-">
              {{ JSON.stringify(modalInfo!.labels) }}
            </dd>
          </div>
          <div class="py-3 flex justify-between text-sm font-medium">
            <dt class="text-gray-500 mr-6">
              Resource
            </dt>
            <dd class="text-gray-900 text-">
              {{ modalInfo!.resource! }}
            </dd>
          </div>
          <div class="py-3 flex justify-between text-sm font-medium">
            <dt class="text-gray-500 mr-6">
              Created
            </dt>
            <dd class="text-gray-900 text-">
              {{ dayjs(modalInfo!.created_at! * 1000).fromNow() }}
            </dd>
          </div>
        </dl>
      </div>
    </template>
  </InfoModal>
  <div class="container p-5">
    <div class="container py-5">
      <span class="font-semibold text-lg ">envd Environments</span>
    </div>
    <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
      <thead class="text-gray-600 border-b border-t">
        <tr
          v-for="headerGroup in table.getHeaderGroups()"
          :key="headerGroup.id"
        >
          <th
            v-for="header in headerGroup.headers"
            :key="header.id"
            :colSpan="header.colSpan"
            class="py-3 px-3 font-semibold"
          >
            <FlexRender
              v-if="!header.isPlaceholder"
              :render="header.column.columnDef.header"
              :props="header.getContext()"
            />
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="row in table.getRowModel().rows" :key="row.id">
          <td v-for="cell in row.getVisibleCells()" :key="cell.id" class="py-4 px-4">
            <FlexRender
              :render="cell.column.columnDef.cell"
              :props="cell.getContext()"
            />
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<route lang="yaml">
meta:
  layout: dashboard
</route>
