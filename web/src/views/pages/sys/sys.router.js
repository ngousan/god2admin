// import layoutHeaderAside from '@/layout/header-aside'
// 由于懒加载页面太多的话会造成webpack热更新太慢，所以开发环境不使用懒加载，只有生产环境使用懒加载
const _import = require('@/libs/util.import.' + process.env.NODE_ENV)

const meta = { auth: true }

export default {
  path: '/sys',
  name: 'system',
  meta: {
    ...meta,
    title: '系统管理'
  },
  // redirect: { name: 'sys-user' },
  // component: layoutHeaderAside,
  children: (pre => [
    {
      path: 'user',
      name: `${pre}user`,
      component: _import('pages/sys/user'),
      props: true,
      meta: {
        ...meta,
        title: '用户管理'
      }
      // children: (pre => [
      //   {
      //     path: 'roles',
      //     name: `${pre}roles`,
      //     component: _import('pages/sys/user/roless'),
      //     props: true,
      //     meta: {
      //       ...meta,
      //       title: '角色分配'
      //     }
      //   }
      // ])(`${pre}user-`)
    },
    {
      path: 'menu',
      name: `${pre}menu`,
      component: _import('pages/sys/menu'),
      props: true,
      meta: {
        ...meta,
        title: '菜单管理'
      }
    },
    {
      path: 'role',
      name: `${pre}role`,
      component: _import('pages/sys/role'),
      props: true,
      meta: {
        ...meta,
        title: '角色管理'
      }
    },
    // 临时菜单路由
    {
      path: 'temp',
      name: `${pre}temp`,
      props: true,
      component: _import('pages/sys/user/userRoles'),
      meta: {
        ...meta,
        title: '临时菜单'
      }
      // children: (pre => [
      //   {
      //     path: 'roles',
      //     name: `${pre}roles`,
      //     component: _import('pages/sys/user/userRoles'),
      //     props: true,
      //     meta: {
      //       ...meta,
      //       title: '角色分配2'
      //     }
      //   }
      // ])(`${pre}temp-`)
    }
  ])('sys-')
}
