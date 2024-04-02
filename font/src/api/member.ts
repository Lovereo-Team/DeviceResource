import request from '@/utils/request'

// 员工列列表
export function memberLists(params?: Record<string, any>) {
    return request.get({ url: '/member/list', params })
}

// 员工列详情
export function memberDetail(params: Record<string, any>) {
    return request.get({ url: '/member/detail', params })
}

// 员工列新增
export function memberAdd(params: Record<string, any>) {
    return request.post({ url: '/member/add', params })
}

// 员工列编辑
export function memberEdit(params: Record<string, any>) {
    return request.post({ url: '/member/edit', params })
}

// 员工列删除
export function memberDelete(params: Record<string, any>) {
    return request.post({ url: '/member/del', params })
}
