<template>
<aside class="w-64 flex" aria-label="Sidebar">
    <div class="overflow-y-auto flex-1 py-4 px-3  bg-gray-100 rounded dark:bg-gray-800">
       <a href="https://envd.tensorchord.ai" class="flex items-center pl-2.5 mb-5 h-20">
          <img src="https://user-images.githubusercontent.com/12974685/200007223-cd94fe9a-266d-4bbd-ac23-e71043d0c3dc.svg" class="mr-3 h-10" alt="envd Logo" />
          <span class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">envd</span>
       </a>
       <ul class="space-y-2">
          <li>
             <a href="#" class="flex items-center px-2 py-3 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-200 dark:hover:bg-gray-700">
               <i-mdi-drive-document class="h-6 w-6 text-gray-500"/>
                <span class="ml-3">Environments</span>
             </a>
          </li>
          <li>
             <a href="#" class="flex items-center px-2 py-3 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-200 dark:hover:bg-gray-700">
               <i-material-symbols-insert-chart-rounded class="h-6 w-6 text-gray-500"/>
                <span class="flex-1 ml-3 whitespace-nowrap">Images</span>
             </a>
          </li>
          <li>
             <a href="#" class="flex items-center px-2 py-3 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-200 dark:hover:bg-gray-700">
                <i-material-symbols-calendar-month class="h-6 w-6 text-gray-500"/>
                <span class="flex-1 ml-3 whitespace-nowrap">Data</span>               
             </a>
          </li>
          <li>
             <a href="#" class="flex items-center px-2 py-3 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-200 dark:hover:bg-gray-700">
               <i-mdi-gear class="h-6 w-6 text-gray-500"/>
                <span class="flex-1 ml-3 whitespace-nowrap">Settings</span>
             </a>
          </li>
       </ul>
    </div>
 </aside>
</template>