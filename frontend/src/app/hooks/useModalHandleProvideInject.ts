import { inject, provide, ref } from 'vue'

interface IHandle {
  (): void
}

export const useModalHandleProvide = (nameDialog: string) => {
  const dialogVisible = ref(false)
  const dialogOpenHandle: IHandle = () => dialogVisible.value = true
  const dialogClosedHandle: IHandle = () => dialogVisible.value = false

  provide(`${nameDialog}Visible`, dialogVisible)
  provide(`${nameDialog}OpenHandle`, dialogOpenHandle)
  provide(`${nameDialog}ClosedHandle`, dialogClosedHandle)

  return {
    dialogVisible,
    dialogOpenHandle,
    dialogClosedHandle,
  }
}

export const useModalHandleInject = (nameDialog: string) => {
  const dialogVisible = inject(`${nameDialog}Visible`)
  const dialogOpenHandle = inject(`${nameDialog}OpenHandle`) as IHandle
  const dialogClosedHandle = inject(`${nameDialog}ClosedHandle`) as IHandle

  return {
    dialogVisible,
    dialogOpenHandle,
    dialogClosedHandle,
  }
}
