<script setup lang="ts">
// https://github.com/vueuse/head
// you can use this to manipulate the document head in any components,
// they will be rendered correctly in the html results with vite-ssg
useHead({
  title: 'envd dashboard',
  meta: [
    {
      name: 'description',
      content: 'TensorChord envd',
    },
    {
      name: 'theme-color',
      content: '#141824',
    },
  ],
  link: [
    {
      rel: 'icon',
      type: 'image/svg+xml',
      href: '/favicon.ico',
    },
  ],
})
</script>

<template>
  <RouterView />
</template>
