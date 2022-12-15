<script setup lang="ts">
const { userInfo, login } = useUserStore()
// const props = defineProps<{ name: string }>()
// const router = useRouter()
// const user = useUserStore()

// watchEffect(() => {
//   user.setNewName(props.name)
// })
const username = ref(userInfo.username)
const password = ref('')
const router = useRouter()
const failWarning = ref(false)

async function loginAction() {
  const isLogin = await login(username.value, password.value)
  if (isLogin)
    router.push('/envs')
  else
    failWarning.value = true
}
</script>

<template>
  <div class="h-screen flex">
    <div class="flex-1 flex flex-col justify-center py-2 px-4 sm:px-6 lg:flex-none lg:px-20 xl:px-24">
      <div class="mx-auto w-full max-w-sm lg:w-96">
        <div>
          <h2 class="mb-20 text-3xl font-extrabold text-gray-900 mx-auto text-center">
            Log in
          </h2>
        </div>

        <div class="mt-8">
          <div class="mt-6">
            <form class="space-y-6">
              <div>
                <label for="email" class="block text-sm font-medium text-gray-700 py-2"> Username
                </label>
                <div class="mt-1">
                  <input
                    id="username" v-model="username" name="username" autocomplete="username"
                    required
                    class="appearance-none block w-full px-3 py-2 bg-[#F7F7F8] rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  >
                </div>
              </div>

              <div class="space-y-1">
                <label for="password" class="block text-sm font-medium text-gray-700 py-2"> Password
                </label>
                <div class="mt-1">
                  <input
                    id="password" v-model="password" name="password" type="password" autocomplete="current-password"
                    class="appearance-none block w-full px-3 py-2 bg-[#F7F7F8] rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  >
                </div>
              </div>

              <div class="flex items-center justify-between">
                <div class="flex items-center">
                  <input
                    id="remember-me" name="remember-me" type="checkbox"
                    class="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded"
                  >
                  <label for="remember-me" class="ml-2 block text-sm text-gray-900"> Remember me
                  </label>
                </div>

                <div class="text-sm">
                  <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500"> Forgot your
                    password? </a>
                </div>
              </div>

              <div class="flex item-center">
                <label for="" />
              </div>

              <div>
                <button
                  type="button" class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white hover:bg-[#1949C5] bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                  @click="loginAction"
                >
                  Log
                  in
                </button>
              </div>

              <div class="space-y-1">
                <div class="mt-1 mx-auto text-center">
                  Don't have an account yet? <router-link class="text-indigo-600" to="/signup">
                    New
                    Account
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
