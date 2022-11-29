<script setup lang="ts">
import { themeChange } from 'theme-change'
import { tryOnMounted } from '@vueuse/core'


// Theme Change
tryOnMounted(() => {
  themeChange(false)
})
</script>

<template>
  <div class="flex flex-row justify-center">
    <RouterLink class="mx-2" to="/">
      <i-ri:home-line class="h-6 w-6 hover:text-green-500" />
    </RouterLink>

    <RouterLink class="mx-2" to="/about">
      <i-carbon-align-box-top-center class="h-6 w-6 hover:text-blue-500" />
    </RouterLink>

    <a class="mx-2" rel="noreferrer" href="https://github.com/compilekaiten/vitetail" target="_blank" title="GitHub">
      <i-carbon-logo-github class="h-6 w-6 hover:text-purple-500" />
    </a>
  </div>

  <div class="flex flex-row justify-center pt-5">
    <select class="select select-bordered w-full max-w-xs scrollbar-thin scrollbar-thumb-slate-700 scrollbar-track-inherit scroll-smooth" data-choose-theme>
      <option disabled selected>
        Select a DaisyUI Theme
      </option>
      <option value="dark">
        Dark
      </option>
      <option value="light">
        Light
      </option>
      <option value="cupcake">
        Cupcake
      </option>
      <option value="bumblebee">
        Bumblebee
      </option>
      <option value="emerald">
        Emerald
      </option>
      <option value="corporate">
        Corporate
      </option>
      <option value="synthwave">
        Synthwave
      </option>
      <option value="retro">
        Retro
      </option>
      <option value="cyberpunk">
        Cyberpunk
      </option>
      <option value="valentine">
        Valentine
      </option>
      <option value="halloween">
        Halloween
      </option>
      <option value="garden">
        Garden
      </option>
      <option value="forest">
        Forest
      </option>
      <option value="aqua">
        Aqua
      </option>
      <option value="lofi">
        Lofi
      </option>
      <option value="pastel">
        Pastel
      </option>
      <option value="fantasy">
        Fantasy
      </option>
      <option value="wireframe">
        Wireframe
      </option>
      <option value="black">
        Black
      </option>
      <option value="luxury">
        Luxury
      </option>
      <option value="dracula">
        Dracula
      </option>
      <option value="cmyk">
        CMYK
      </option>
      <option value="autumn">
        Autumn
      </option>
      <option value="business">
        Business
      </option>
      <option value="acid">
        Acid
      </option>
      <option value="lemonade">
        Lemonade
      </option>
      <option value="night">
        Night
      </option>
      <option value="coffee">
        Coffee
      </option>
      <option value="winter">
        Winter
      </option>
    </select>
  </div>
</template>
