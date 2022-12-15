<script setup lang="ts">
const props = defineProps<{ status: string }>()

const statusTagClass = computed(() => {
  switch (props.status) {
    case 'Running':
      return 'bg-green-100 text-green-800'
    case 'Pending':
      return 'bg-yellow-100 text-yellow-800'
    case 'Terminating':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
})
</script>

<template>
  <span
    class="text-xs font-semibold mr-2 px-2.5 py-0.5 rounded"
    :class="statusTagClass"
  >
    {{ $props.status }}
  </span>
</template>
