import { createApp } from 'vue'

export function initPlugins(app: ReturnType<typeof createApp>): void {
  const files = import.meta.globEager('./**/*.ts')

  for (const path in files) {
    if (!Object.prototype.hasOwnProperty.call(files, path)) continue
    if (typeof files[path].default === 'function') {
      if (path !== './index.ts') files[path].default(app)
    }
  }
}
