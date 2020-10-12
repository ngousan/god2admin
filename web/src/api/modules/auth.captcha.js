export default ({ request }) => ({
  AUTH_CAPTCHA(data) {
    return request({
      url: '/auth/captcha',
      method: 'post',
      data
    })
  }
})
