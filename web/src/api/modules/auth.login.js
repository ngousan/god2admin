export default ({ request }) => ({
  AUTH_LOGIN(data) {
    return request({
      url: '/auth/login',
      method: 'post',
      data
    })
  }
})
