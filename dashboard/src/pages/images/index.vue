<script setup lang="ts">
import { FlexRender, createColumnHelper, getCoreRowModel, useVueTable } from '@tanstack/vue-table'
import dayjs from 'dayjs'
import ImageDetailModal from '~/components/imageDetailModal.vue'
import { useEnvdFetch } from '~/composables/request'
import type { TypesImageListResponse } from '~/composables/types/scheme'

const { getImages } = useImageStore()
const imgs = getImages()

interface ImageDisplay {
  name: string
  created: number
  size: number
}

const imgDisplay: ImageDisplay[] = imgs.value.map((e) => {
  return {
    name: e.name!,
    created: e.created!,
    size: e.size!,
  }
})

// const imgDisplay: ImageDisplay[] = [{
//   name: 'test1',
//   created: 2000000000000,
//   size: 20000,
// }]

const columnHelper = createColumnHelper<ImageDisplay>()
const columns = [
  columnHelper.accessor('name', {
    id: 'name',
    cell: info => h('span', { innerText: info.getValue() }),
    header: () => 'Image Name',
  }),
  columnHelper.accessor('created', {
    id: 'created',
    cell: info => dayjs(info.getValue() * 1000).fromNow(),
    header: () => 'Created',
  }),
  columnHelper.accessor('size', {
    id: 'size',
    cell: info => info.getValue(),
    header: () => 'Size',
  }),
]

const table = useVueTable<ImageDisplay>({
  columns,
  data: imgDisplay,
  getCoreRowModel: getCoreRowModel(),
})
</script>

<template>
  <div class="container flex-1 py-5 mx-5">
    <Navbar />
    <ImageDetailModal :open-modal="true" />
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
  </div>
</template>

<route lang="yaml">
meta:
  layout: dashboard
</route>
