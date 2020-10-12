import { request } from '@/api/service'

export function ListObj(params) {
  return request({
    url: '/sys/menu/list',
    method: 'get',
    params
  })
}

export function NewObj(obj) {
  return request({
    url: '/sys/menu/new',
    method: 'post',
    data: obj
  })
}

export function UpdateObj(obj) {
  return request({
    url: '/sys/menu/update',
    method: 'post',
    data: obj
  })
}
