export default {
  title: '系统管理',
  icon: 'cog',
  children: (pre => [
    { title: '用户管理 ', icon: 'user', path: `${pre}/user` },
    { title: '菜单管理 ', icon: 'user', path: `${pre}/menu` },
    { title: '角色管理 ', icon: 'user', path: `${pre}/role` },
    { title: '临时菜单 ', icon: 'user', path: `${pre}/temp` }
  ])('/sys')
}
