<script setup lang="ts">
import { storeToRefs } from 'pinia'

const { header } = storeToRefs(useNav())
</script>

<template>
  <nav class="bg-white border-gray-200 px-2 sm:px-4 py-2.5 rounded dark:bg-gray-900">
    <div class="container flex flex-wrap justify-between mx-auto">
      <span class="text-2xl font-bold">{{ header }}</span>
      <div id="navbar-default" class="w-auto">
        <button
          type="button"
          class="text-white bg-blue-600 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800"
        >
          +
          &nbsp; Create New Environment
        </button>
      </div>
    </div>
  </nav>
</template>
