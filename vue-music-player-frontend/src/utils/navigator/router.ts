/* router根据路由名称和id跳转 */
const router = useRouter()

//带参数跳转
export const routerPush = (name: string, id: number) => {
  router.push({ name: name, query: { id: id } }).then(() => {})
}
