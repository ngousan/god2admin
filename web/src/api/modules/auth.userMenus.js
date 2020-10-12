export default ({ request }) => ({
  AUTH_USER_MENUS(param) {
    return request({
      url: '/auth/user/menus',
      method: 'get',
      param
    })
  }
})
