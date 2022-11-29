<script setup lang="ts">
const router = useRouter()
</script>

<template>
  <main>
    <RouterView />
    <div class="container mx-auto flex justify-center">
      <button class="btn btn-wide btn-error" @click="router.back()">
        Go Back
      </button>
    </div>
    <div class="mt-5 mx-auto text-center">
      [404 Layout]
    </div>
  </main>
</template>
