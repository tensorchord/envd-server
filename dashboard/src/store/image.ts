import type { Ref } from '@vue/reactivity'
import { acceptHMRUpdate, defineStore } from 'pinia'
import type { TypesImageListResponse, TypesImageMeta } from '~/composables/types/scheme'

export const useImageStore = defineStore('image', () => {
  const imgs = ref<TypesImageMeta[]>([])
  const user = useUserStore()
  const { data, execute } = useEnvdFetch(`/users/${encodeURIComponent(user.userInfo.username)}/images`).get().json<TypesImageListResponse>()

  async function fetchImages(): Promise<TypesImageMeta[]> {
    await execute()
    return data.value!.items!
  }
  async function refreshImages(): Promise<void> {
    imgs.value = await fetchImages()
  }
  function getImages(): Ref<TypesImageMeta[]> {
    return imgs
  }
  async function getImageInfo(name: string): Promise<TypesImageMeta> {
    const imgs = getImages()
    const img = imgs.value.find(img => img.name === name)
    return img!
  }

  return { imgs, getImages, refreshImages, getImageInfo }
},
)

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useImageStore, import.meta.hot))
