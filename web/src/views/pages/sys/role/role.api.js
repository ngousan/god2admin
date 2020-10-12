import { request } from '@/api/service'

export function ListObj(params) {
  return request({
    url: '/sys/role/list',
    method: 'get',
    params
  })
}

export function NewObj(obj) {
  return request({
    url: '/sys/role/new',
    method: 'post',
    data: obj
  })
}

export function UpdateObj(obj) {
  return request({
    url: '/sys/role/update',
    method: 'post',
    data: obj
  })
}
