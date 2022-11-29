<script setup lang="ts">
// const user = useUserStore()
// const name = ref(user.savedName)

// const router = useRouter()
// const go = () => {
//   if (name)
//     router.push(`/hi/${encodeURIComponent(name.value)}`)
// }
</script>

<template>
  <div class="min-h-screen flex">

    <Sidebar class="flex-none z-40 h-screen overflow-y-auto bg-white w-80 dark:bg-gray-800" />

    <!-- 
    <div class="grid grid-rows-1 justify-center">
      <img src="/pwa-512x512.png" alt="ViteTail Logo" width="100" heigh="100">
    </div>

    <div class="flex flex-row justify-between pt-5">
      <i-logos:vitejs class="h-12 w-12" />
      <i-logos:tailwindcss-icon class="h-14 w-14" />
      <i-fxemoji:blossom class="h-12 w-12" />
    </div>

    <div class="mockup-code mt-5 bg-base-200">
      <pre data-prefix="$"><code>ViteTail</code></pre>
      <pre data-prefix=">" class="text-error"><code class="text-error">Vite + TailwindCSS + DaisyUI</code></pre>
      <pre data-prefix=">" class="text-success"><code>Inspired by Vitesse</code></pre>
    </div> -->

    <div class="container mx-auto flex-1 py-5 mx-5">
      <Navbar />
      <div class="container p-5">
        <div class="container py-5">
          <span class="font-semibold	text-lg	">envd Environments</span>
        </div>
        <Table />
      </div>
    </div>
  </div>
</template>

<route lang="yaml">
meta:
  layout: home
</route>
