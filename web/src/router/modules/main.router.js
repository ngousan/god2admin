import layoutHeaderAside from '@/layout/header-aside'

// 由于懒加载页面太多的话会造成webpack热更新太慢，所以开发环境不使用懒加载，只有生产环境使用懒加载
const _import = require('@/libs/util.import.' + process.env.NODE_ENV)

const meta = { auth: true }

export default {
  path: '/main',
  name: 'main',
  meta,
  redirect: { name: 'main-index' },
  component: layoutHeaderAside,
  children: (pre => [
    {
      path: 'index',
      name: `${pre}index`,
      component: _import('pages/index'),
      meta: { ...meta, title: '首页' }
    },
    // {
    //   path: 'sys',
    //   name: `${pre}sys`,
    //   // component: _import('pages/sys'),
    //   meta: { ...meta, title: '系统管理' },
    //   children: (pre => [
    //     {
    //       path: 'user',
    //       name: `${pre}user`,
    //       component: _import('pages/sys/user'),
    //       meta: { ...meta, title: '用户管理' }
    //     }
    //   ])(`${pre}sys-`)
    // }
    {
      path: 'sys/user',
      name: `${pre}user`,
      component: _import('pages/sys/user'),
      meta: { ...meta, title: '用户管理' }
    },
    {
      path: 'sys/menu',
      name: `${pre}menu`,
      component: _import('pages/sys/menu'),
      meta: { ...meta, title: '菜单管理' }
    },
    {
      path: 'sys/role',
      name: `${pre}role`,
      component: _import('pages/sys/role'),
      meta: { ...meta, title: '角色管理' }
    },
    {
      path: 'sys/user/userRoles',
      name: `${pre}user-userRoles`,
      component: _import('pages/sys/user/userRoles'),
      meta: { ...meta, title: '用户角色分配' }
    }
  ])('main-')
}
