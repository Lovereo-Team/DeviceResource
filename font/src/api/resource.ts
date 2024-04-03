import request from '@/utils/request'

// 溯源列列表
export function resourceLists(params?: Record<string, any>) {
    return request.get({ url: '/resource/list', params })
}

// 溯源列详情
export function resourceDetail(params: Record<string, any>) {
    return request.get({ url: '/resource/detail', params })
}

// 溯源列新增
export function resourceAdd(params: Record<string, any>) {
    return request.post({ url: '/resource/add', params })
}

// 溯源列编辑
export function resourceEdit(params: Record<string, any>) {
    return request.post({ url: '/resource/edit', params })
}

// 溯源列删除
export function resourceDelete(params: Record<string, any>) {
    return request.post({ url: '/resource/del', params })
}
