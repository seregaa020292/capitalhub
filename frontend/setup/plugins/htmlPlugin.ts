import project from '../data/project.json'
import meta from '../data/meta.json'

const htmlPlugin = () => ({
  name: 'html-transform',
  transformIndexHtml(html: string) {
    return html
      .replace(/<% title %>/, `${project.name} - ${meta.title}`)
      .replace(/<% description %>/, meta.description)
      .replace(/<% keywords %>/, meta.keywords)
  },
})

export default htmlPlugin
