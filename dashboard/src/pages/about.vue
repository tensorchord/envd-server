<script setup lang="ts">
const router = useRouter()
</script>

<template>
  <div class="flex justify-center">
    <div class="hero min-h-screen">
      <div class="hero-content text-center">
        <div class="max-w-md space-y-3">
          <h1 class="text-5xl font-bold pb-5">
            ViteTail
          </h1>
          <div class="divider">
            About
          </div>
          <p class="leading-relaxed text-left">
            ViteTail is heavily inspired by <a class="link link-primary" href="https://github.com/antfu/vitesse">Vitesse</a>, including Vue3 and Vite3, file-based routing, component auto importing, icon auto importing, Pinia, Vite-SSG, and PWA support.
          </p>
          <p class="leading-relaxed text-left">
            Styling out of the box with <a class="link" href="https://tailwindcss.com/">TailwindCSS</a>, <a class="link" href="https://postcss.org/">PostCSS</a>, and <a class="link" href="https://daisyui.com/">DaisyUI Components</a>.
          </p>
          <p class="leading-relaxed text-left">
            Including <a class="link" href="https://github.com/adoxography/tailwind-scrollbar">TailwindCSS Scrollbars</a> and <a class="link" href="https://github.com/saadeghi/theme-change">Theme-Change</a>.
          </p>
          <p class="leading-relaxed text-left">
            Check out the repo <a class="link link-primary" href="https://github.com/compilekaiten/ViteTail">here</a>, and give it a Star if you like it.
          </p>
          <p class="">
            PR's welcome!
          </p>
          <button id="back" class="btn btn-wide btn-outline btn-error" @click="router.back()">
            Back
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<route lang="yaml">
meta:
  layout: default
</route>

