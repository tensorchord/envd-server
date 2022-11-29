<script setup lang="ts">
const props = defineProps<{ name: string }>()
const router = useRouter()
const user = useUserStore()

watchEffect(() => {
  user.setNewName(props.name)
})
</script>

<template>
  <div class="container mx-auto grid grid-rows-2 justify-center py-10">
    <span class="text-lg text-center">Hello <span class="animate-pulse">{{ props.name }}!</span></span>

    <button class="btn btn-wide btn-error" @click="router.back()">
      Back
    </button>
  </div>
</template>
