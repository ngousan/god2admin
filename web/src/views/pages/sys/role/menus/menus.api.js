import { request } from '@/api/service'

export function ListObj(params) {
  return request({
    url: '/sys/role/menus/list',
    method: 'get',
    params
  })
}

export function NewObj(obj) {
  return request({
    url: '/sys/role/menus/new',
    method: 'post',
    data: obj
  })
}

export function UpdateObj(obj) {
  return request({
    url: '/sys/role/menus/update',
    method: 'post',
    data: obj
  })
}
