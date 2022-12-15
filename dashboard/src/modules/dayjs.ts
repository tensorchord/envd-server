import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import { type UserModule } from '~/types'

export const install: UserModule = () => {
  dayjs.extend(relativeTime)
}

