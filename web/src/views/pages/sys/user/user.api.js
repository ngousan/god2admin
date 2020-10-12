import { request } from '@/api/service'

export function ListObj(params) {
  return request({
    url: '/sys/user/list',
    method: 'get',
    params
  })
}

export function NewObj(obj) {
  return request({
    url: '/sys/user/new',
    method: 'post',
    data: obj
  })
}

export function UpdateObj(obj) {
  return request({
    url: '/sys/user/update',
    method: 'post',
    data: obj
  })
}

export function DelObj(id) {
  return request({
    url: '/sys/user/disable',
    method: 'post',
    data: { id }
  })
}
