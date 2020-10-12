export default ({ request }) => ({
  AUTH_USER_ROUTERS(param) {
    return request({
      url: '/auth/user/routers',
      method: 'get',
      param
    })
  }
})
