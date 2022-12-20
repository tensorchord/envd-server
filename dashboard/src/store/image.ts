import type { Ref } from '@vue/reactivity'
import { acceptHMRUpdate, defineStore } from 'pinia'
import type { TypesImageListResponse, TypesImageMeta } from '~/composables/types/scheme'

export const useImageStore = defineStore('image', () => {
  const imgs = ref<TypesImageMeta[]>([])
  const user = useUserStore()
  const { data, execute } = useEnvdFetch(`/users/${user.userInfo.username}/images`, { immediate: false }).get().json<TypesImageListResponse>()

  async function fetchImages(): Promise<TypesImageMeta[]> {
    await execute()
    return data.value!.items!
  }
  async function refreshImages(): Promise<void> {
    imgs.value = await fetchImages()
  }
  function getImages(): Ref<TypesImageMeta[]> {
    if (imgs.value.length === 0)
      refreshImages()
    return imgs
  }

  return { getImages, refreshImages }
},
)

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useImageStore, import.meta.hot))
