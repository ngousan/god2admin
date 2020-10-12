import { request } from '@/api/service'

export function ListObj(params) {
  return request({
    url: '/sys/user/roles/list',
    method: 'get',
    params
  })
}

export function NewObj(obj) {
  return request({
    url: '/sys/user/roles/new',
    method: 'post',
    data: obj
  })
}

export function UpdateObj(obj) {
  return request({
    url: '/sys/user/roles/update',
    method: 'post',
    data: obj
  })
}
