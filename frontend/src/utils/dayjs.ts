import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import 'dayjs/locale/ru'

dayjs.locale('ru')
dayjs.extend(relativeTime)

export { dayjs }
