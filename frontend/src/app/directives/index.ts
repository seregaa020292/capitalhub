import { createApp } from 'vue'

export function initDirectives(app: ReturnType<typeof createApp>): void {
  const files = import.meta.globEager('./*.ts')

  for (const path in files) {
    if (!Object.prototype.hasOwnProperty.call(files, path)) {
      continue
    }
    if (path !== './index.ts') {
      const file = files[path].default
      app.directive(file.name, file)
    }
  }
}
