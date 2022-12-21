<script setup lang="ts">
// import { CalendarIcon, ChartBarIcon, FolderIcon, HomeIcon, InboxIcon, UsersIcon } from '@heroicons/vue/outline'
import IMdiDriveDocument from '~icons/mdi/drive-document'
import IMaterialSymbolsInsertChartRounded from '~icons/material-symbols/insert-chart-rounded'
import IMaterialSymbolsCalendarMonth from '~icons/material-symbols/calendar-month'
import IMdiGear from '~icons/mdi/gear'

const navigation = [
  { name: 'Environments', icon: IMdiDriveDocument, href: '/envs', current: true },
  { name: 'Images', icon: IMaterialSymbolsInsertChartRounded, href: '/images', current: true },
  { name: 'Data', icon: IMaterialSymbolsCalendarMonth, href: '#', current: false },
  { name: 'Settings', icon: IMdiGear, href: '#', current: false },
]

const index = ref(0)

const { setNavHeader } = useNav()

watch(index, (val) => {
  setNavHeader(navigation[val].name)
})
</script>

<template>
  <aside class="w-64 flex" aria-label="Sidebar">
    <div class="overflow-y-auto flex-1 py-4 px-3  bg-gray-100 bg-opacity-50 rounded dark:bg-gray-800">
      <a href="https://envd.tensorchord.ai" class="flex items-center pl-2.5 mb-5 h-20">
        <img
          src="https://user-images.githubusercontent.com/12974685/200007223-cd94fe9a-266d-4bbd-ac23-e71043d0c3dc.svg"
          class="mr-3 h-10" alt="envd Logo"
        >
        <span class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">envd</span>
      </a>
      <ul class="space-y-2">
        <li
          v-for="(item, i) in navigation" :key="item.name"
        >
          <router-link
            :to="item.href"
            :class="[index === i ? 'bg-gray-200 text-gray-900' : 'text-gray-600 hover:bg-gray-100 hover:text-gray-900']"
            class="flex items-center px-2 py-3 text-base font-normalrounded-lg"
            @click="index = i"
          >
            <component :is="item.icon" class="mr-3 flex-shrink-0 h-6 w-6" :class="[index === i ? 'text-gray-500' : 'text-gray-400 group-hover:text-gray-500']" aria-hidden="true" />
            <span class="ml-3">{{ item.name }}</span>
          </router-link>
        </li>
      </ul>
    </div>
  </aside>
</template>
