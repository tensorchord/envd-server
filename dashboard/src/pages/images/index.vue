<script setup lang="ts">
import { FlexRender, createColumnHelper, getCoreRowModel, useVueTable } from '@tanstack/vue-table'
import dayjs from 'dayjs'
import { storeToRefs } from 'pinia'
import InfoModal from '~/components/InfoModal.vue'
import type { TypesImageMeta } from '~/composables/types/scheme'

const { getImageInfo, refreshImages } = useImageStore()
const { imgs } = storeToRefs(useImageStore())
// const imgs = getImages()
const modalInfo = ref<TypesImageMeta>()
const modal = ref<typeof InfoModal>()

const showImageDetail = async (name: string) => {
  modalInfo.value = await getImageInfo(name)
  modal.value!.openModal()
}

const columnHelper = createColumnHelper<TypesImageMeta>()
const columns = [
  columnHelper.accessor('name', {
    id: 'name',
    cell: info => h('button', {
      class: 'text-blue-500 hover:text-blue-800',
      innerText: info.getValue(),
      onClick: async () => {
        await showImageDetail(info.getValue()!)
      },
    }),
    header: () => 'Image Name',
  }),
  columnHelper.accessor('created', {
    id: 'created',
    cell: info => dayjs(info.getValue()! * 1000).fromNow(),
    header: () => 'Created',
  }),
  columnHelper.accessor('size', {
    id: 'size',
    cell: info => formatBytes(info.getValue()!),
    header: () => 'Size',
  }),
]
const data = ref<TypesImageMeta[]>([])

const table = useVueTable<TypesImageMeta>({
  columns,
  get data() {
    return data.value
  },
  getCoreRowModel: getCoreRowModel(),
})

watch(imgs, (newVal) => {
  data.value = newVal
})

onMounted(async () => {
  await refreshImages()
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
              Digest
            </dt>
            <dd class="text-gray-900 text-">
              {{ modalInfo!.digest! }}
            </dd>
          </div>
          <div class="py-3 flex justify-between text-sm font-medium">
            <dt class="text-gray-500 mr-6">
              Created
            </dt>
            <dd class="text-gray-900 text-">
              {{ dayjs(modalInfo!.created! * 1000).fromNow() }}
            </dd>
          </div>
        </dl>
      </div>
    </template>
  </InfoModal>
  <div class="container p-5">
    <div class="container py-5">
      <span class="font-semibold text-lg ">envd Images</span>
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
