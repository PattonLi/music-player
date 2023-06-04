import type { TopList } from "@/models/toplist"
import myAxios from "./myAxios"

export async function apiTopListDetail() {
    const {list} = await myAxios.get<{ list: TopList[] }>('/toplist/detail')
    return list
}