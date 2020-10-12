export const crudOptions = () => {
  return {
    options: {
      height: '100%',
      rowKey: 'ID'
      // lazy: false
      // load: (tree, treeNode, resolve) => {
      //   request({
      //     url: '/sys/menu/list/children',
      //     method: 'get',
      //     data: { id: tree.id }
      //   }).then(ret => {
      //     console.log('懒加载数据', ret.data)
      //     resolve(ret.data)
      //   })
      // }
    },
    formOptions: {
      labelPosition: 'top'
    },
    searchOptions: {
      // form: {
      //   name: '%'
      // },
      size: 'small',
      show: false,
      debounce: {
        wait: 500
      }
    },
    rowHandle: {
      width: 145,
      fixed: 'right',
      view: {
        thin: true,
        text: null
      },
      edit: {
        thin: true,
        text: null
      },
      remove: {
        thin: true,
        text: null,
        disabled: true
      }
    },
    columns: [
      // { fixed: 'left' },
      {
        title: 'ID',
        key: 'ID',
        width: 100,
        fixed: 'left',
        form: {
          addDisabled: true,
          component: {
            disabled: true
          }
        }
      },
      {
        title: 'UUID',
        key: 'uuid',
        show: false,
        form: {
          addDisabled: true,
          component: {
            disabled: true
          }
        }
      },
      {
        title: '上级',
        key: 'parentId',
        type: 'number'
      },
      {
        title: '菜单',
        key: 'name',
        fixed: 'left',
        form: {
          rules: [{ required: true, message: '请输入菜单名称' }]
        },
        editForm: {
          component: {
            disabled: true
          }
        },
        search: {}
      },
      {
        title: '标题',
        key: 'title',
        form: {
          rules: [{ required: true, message: '请输入菜单标题' }]
        },
        search: {}
      },
      {
        title: '描述',
        key: 'description',
        show: false,
        form: {
          rules: [{ required: true, message: '请输入菜单描述' }]
        },
        search: {}
      },
      {
        title: '路径',
        key: 'path',
        form: {
          rules: [{ required: true, message: '请输入路由Path' }]
        }
      },
      {
        title: '图标',
        key: 'icon',
        type: 'icon-selector',
        form: {
          component: {
            props: {
              'user-input': true
            }
          }
        }
      },
      {
        title: '组件',
        key: 'component'
      },
      {
        title: '重定向',
        key: 'redirectName'
      },
      {
        title: '权限',
        key: 'auth',
        type: 'switch'
      },
      {
        title: '缓存',
        key: 'cache',
        type: 'switch'
      },
      {
        title: '状态',
        key: 'status',
        type: 'switch'
      },
      {
        title: '排序',
        key: 'sort',
        type: 'number',
        show: false
      }
    ]
  }
}
