<script setup lang="ts">
import Navbar from '~/components/Navbar.vue'
</script>

<template>
  <main>
    <div class="h-screen flex flex-nowrap w-screen overflow-hidden box-border">
      <Sidebar class="flex-none z-40 h-screen w-80" />
      <div class="container flex-1 py-5 mx-5">
        <Navbar />
        <RouterView />
      </div>
    </div>
  </main>
</template>

