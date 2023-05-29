/* router根据路由名称和id跳转 */
const router = useRouter()

export const routerPush = (name: string, id: number) => {
  router.push({ name: name, query: { id: id } }).then(() => {})
}
