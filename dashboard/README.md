- - - -

<br>

<div align="center">
<img src="images/../.github/images/vt.png" height="100" />
</div>

<br>

<p align='center'>
<b>ViteTail</b> is a Vue + Vite + TailwindCSS + DaisyUI template, heavily based on Vitesse.
</p>

<p align='center'>
Almost all the great features of Vitesse, but with your favorite CSS framework!
</p>

<p align='center'>
ViteTail comes with DaisyUI out of the box. Feel free to switch to any component library you want.
</p>

<br>

<p align='center'>
<a href="https://vitetail.netlify.app/">Live Demo</a>
</p>

<br>

- - - -
## Features

- ⚡️ [Vue 3](https://github.com/vuejs/core), [Vite 3](https://github.com/vitejs/vite), [pnpm](https://pnpm.io/), [ESBuild](https://github.com/evanw/esbuild) - born with fastness

- 🗂 [File based routing](./src/pages)

- 📦 [Components auto importing](./src/components)

- 🍍 [State Management via Pinia](https://pinia.vuejs.org/)

- 📑 [Layout system](./src/layouts)

- 📲 [PWA](https://github.com/antfu/vite-plugin-pwa)

- 🎨 [TailwindCSS](https://tailwindcss.com/)

- 😃 [Use icons from any icon sets with auto importing](https://github.com/antfu/unplugin-icons)

- 🔎 [Component Preview](https://github.com/johnsoncodehk/vite-plugin-vue-component-preview)

- 🔥 Use the [new `<script setup>` syntax](https://github.com/vuejs/rfcs/pull/227)

- 🤙🏻 [Reactivity Transform](https://vuejs.org/guide/extras/reactivity-transform.html) enabled

- 📥 [APIs auto importing](https://github.com/antfu/unplugin-auto-import) - use Composition API and others directly

- 🖨 Static-site generation (SSG) via [vite-ssg](https://github.com/antfu/vite-ssg)

- 🦔 Critical CSS via [critters](https://github.com/GoogleChromeLabs/critters)

- 🦾 TypeScript, of course

- ⚙️ Unit Testing with [Vitest](https://github.com/vitest-dev/vitest), E2E Testing with [Cypress](https://cypress.io/) on [GitHub Actions](https://github.com/features/actions)

- - - -

## Features Removed
-  I18n
-  Markdown Support
-  UnoCSS

- - - -
## Pre-packed

### UI Frameworks

- [TailwindCSS](https://tailwindcss.com/)
  - [@TailwindCSS/Aspect-Ratio](https://github.com/tailwindlabs/tailwindcss-aspect-ratio)
  - [@TailwindCSS/Typography](https://github.com/tailwindlabs/tailwindcss-typography)
- [TailwindCSS-Scrollbar](https://github.com/adoxography/tailwind-scrollbar)
- [DaisyUI](https://daisyui.com/) - Easily change to whatever component library you want
- [Theme-Change](https://github.com/saadeghi/theme-change) - Fast CSS Theme changing

### Icons

- [Iconify](https://iconify.design) - use icons from any icon sets
  - [🔍Icônes](https://icones.netlify.app/) - Find Icons here
- [Auto Importing Icons](https://github.com/antfu/unplugin-icons)

### Plugins

- [`Vue Router`](https://github.com/vuejs/router)
  - [`vite-plugin-pages`](https://github.com/hannoeru/vite-plugin-pages) - file system based routing
  - [`vite-plugin-vue-layouts`](https://github.com/JohnCampionJr/vite-plugin-vue-layouts) - layouts for pages
- [`Pinia`](https://pinia.vuejs.org) - Intuitive, type safe, light and flexible Store for Vue using the composition api
- [`unplugin-vue-components`](https://github.com/antfu/unplugin-vue-components) - components auto import
- [`unplugin-auto-import`](https://github.com/antfu/unplugin-auto-import) - Directly use Vue Composition API and others without importing
- [`unplugin-icons`](https://github.com/antfu/unplugin-icons) - Auto import Icons from any set
- [`vite-plugin-pwa`](https://github.com/antfu/vite-plugin-pwa) - PWA
- [`vite-plugin-vue-component-preview`](https://github.com/johnsoncodehk/vite-plugin-vue-component-preview) - Preview single component in VSCode
- [`VueUse`](https://github.com/antfu/vueuse) - collection of useful composition APIs
- [`vite-ssg-sitemap`](https://github.com/jbaubree/vite-ssg-sitemap) - Sitemap generator
- [`@vueuse/head`](https://github.com/vueuse/head) - manipulate document head reactively

### Coding Style

- Use Composition API with [`<script setup>` SFC syntax](https://github.com/vuejs/rfcs/pull/227)
- [ESLint](https://eslint.org/) with [@antfu/eslint-config](https://github.com/antfu/eslint-config), single quotes, no semi.

### Dev tools

- [`TypeScript`](https://www.typescriptlang.org/)
- [`Vitest`](https://github.com/vitest-dev/vitest) - Unit testing powered by Vite
- [`Cypress`](https://cypress.io/) - E2E testing
- [`pnpm`](https://pnpm.js.org/) - fast, disk space efficient package manager
- [`vite-ssg`](https://github.com/antfu/vite-ssg) - Static-site generation
  - [`critters`](https://github.com/GoogleChromeLabs/critters) - Critical CSS
- [VS Code Extensions](./.vscode/extensions.json)
  - [Vite](https://marketplace.visualstudio.com/items?itemName=antfu.vite) - Fire up Vite server automatically
  - [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) - Vue 3 `<script setup>` IDE support
  - [Iconify IntelliSense](https://marketplace.visualstudio.com/items?itemName=antfu.iconify) - Icon inline display and autocomplete
  - [ESLint](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint)
  - [TailwindCSS IntelliSense](https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss)

## Try it now!
### GitHub Template

[Create a repo from this template on GitHub](https://github.com/compilekaiten/ViteTail/generate).

### Clone to local

If you prefer to do it manually with the cleaner git history

```bash
npx degit compilekaiten/vitetail my-vitetail-app
cd my-vitetail-app
pnpm i # If you don't have pnpm installed, run: npm install -g pnpm
```

## Checklist

When you use this template, try follow the checklist to update your info properly

- [ ] Change the author name in `LICENSE`
- [ ] Change the title in `App.vue`
- [ ] Change the hostname in `vite.config.ts`
- [ ] Change the favicon in `public`
- [ ] Remove the `.github` folder which contains the funding info
- [ ] Clean up the READMEs and remove routes

And, enjoy :)

## Usage

### Development

Visit http://localhost:3333 to view the project

```bash
pnpm open # To open automatically
pnpm dev # To run the dev server
```

Update Dependencies

```bash
pnpm dep # Runs Taze to update Major and Minor dependencies
```

### Build

To build your App, run

```bash
pnpm build # Build for Production
pnpm tbuild # Lint and Typecheck before building
```

And you will see the generated file in `dist` that are ready to be served.

### Deploy on Netlify

Go to [Netlify](https://app.netlify.com/start) and select your clone, `OK` along the way, and your App will be live in a minute.

#### PR's are very welcome
