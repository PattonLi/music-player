const pathMap = {
  home: '首页',
  addsong: '添加商品'
}

//浏览器获取存储方法
const localGet = (key: string) => {
  const value = window.localStorage.getItem(key)
  try {
    return JSON.parse(value as string)
  } catch (error) {
    return null
  }
}

const localSet = (key: string, value: object) => {
  window.localStorage.setItem(key, JSON.stringify(value))
}

const localRemove = (key: string) => {
  window.localStorage.removeItem(key)
}

export { pathMap, localGet, localSet, localRemove }
