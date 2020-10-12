// import sysMenu from '../../views/pages/sys/sys.menu'

export default {
  title: '主菜单',
  icon: 'home',
  path: '',
  children: (pre => [
    {
      title: '首页',
      icon: 'home',
      path: `${pre}/index`
    },
    {
      title: '系统管理',
      icon: 'cog',
      children: (pre => [
        {
          title: '用户管理 ',
          icon: 'user',
          path: '',
          children: (pre => [
            { title: '用户管理 ', icon: 'user', path: `${pre}` },
            { title: '用户角色分配 ', icon: 'user', path: `${pre}/userRoles` }
          ])(`${pre}/user`)
        },
        { title: '菜单管理 ', icon: 'user', path: `${pre}/menu` },
        { title: '角色管理 ', icon: 'user', path: `${pre}/role` },
        { title: '临时菜单 ', icon: 'user', path: `${pre}/temp` }
      ])(`${pre}/sys`)
    }
  ])('/main')
}
