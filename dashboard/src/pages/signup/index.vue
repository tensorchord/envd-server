<script setup lang="ts">
// const props = defineProps<{ name: string }>()
// const router = useRouter()
// const user = useUserStore()

// watchEffect(() => {
//   user.setNewName(props.name)
// })
</script>

<template>
  <div class="h-screen flex">
    <div class="flex-1 flex flex-col justify-center py-2 px-4 sm:px-6 lg:flex-none lg:px-20 xl:px-24">
      <div class="mx-auto w-full max-w-sm lg:w-96">
        <div>
          <h2 class="mb-20 text-3xl font-extrabold text-gray-900 mx-auto text-center">
            Sign up
          </h2>
        </div>

        <div class="mt-8">
          <div class="mt-6">
            <form action="#" method="POST" class="space-y-6">
              <div>
                <label for="email" class="block text-sm font-medium text-gray-700 py-2"> Full Name
                </label>
                <div class="mt-1">
                  <input
                    id="fullname" name="fullname" type="fullname" autocomplete="fullname" required
                    class="appearance-none block w-full px-3 py-2 bg-[#F7F7F8] rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  >
                </div>
              </div>

              <div>
                <label for="email" class="block text-sm font-medium text-gray-700 py-2"> Email address
                </label>
                <div class="mt-1">
                  <input
                    id="email" name="email" type="email" autocomplete="email" required
                    class="appearance-none block w-full px-3 py-2 bg-[#F7F7F8] rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  >
                </div>
              </div>

              <div class="space-y-1">
                <label for="password" class="block text-sm font-medium text-gray-700 py-2"> Password
                </label>
                <div class="mt-1">
                  <input
                    id="password" name="password" type="password" autocomplete="current-password" required
                    class="appearance-none block w-full px-3 py-2 bg-[#F7F7F8] rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  >
                </div>
              </div>

              <div class="pt-4">
                <router-link to="/">
                  <button
                    type="submit"
                    class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white hover:bg-[#1949C5] bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                  >
                    Create
                    Account
                  </button>
                </router-link>
              </div>

              <div class="space-y-1">
                <div class="mt-1 mx-auto text-center">
                  Already have an account? <router-link class="text-indigo-600" to="/login">
                    Log
                    in
                  </router-link>
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
    <div class="h-screen flex w-0 flex-1  bg-gray-50">
      <LoginImg class="m-auto" />
    </div>
  </div>
</template>

<route lang="yaml">
meta:
  layout: default
</route>

