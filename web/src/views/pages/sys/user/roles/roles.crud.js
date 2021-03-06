export const crudOptions = () => {
  return {
    options: {
      height: '100%'
    },
    formOptions: {
      labelPosition: 'top'
    },
    searchOptions: {
      size: 'small',
      show: true, // 是否显示搜索工具条
      debounce: {
        wait: 500 // 延迟500毫秒
      }
    },
    // rowHandle: {
    //   width: 145,
    //   fixed: 'right',
    //   view: {
    //     thin: true,
    //     text: null
    //   },
    //   edit: {
    //     thin: true,
    //     text: null
    //   },
    //   remove: {
    //     thin: true,
    //     text: null,
    //     disabled: true
    //   }
    // },
    rowHandle: false,
    columns: [
      {
        title: 'ID',
        key: 'ID',
        fixed: 'left',
        form: {
          addDisabled: true,
          component: {
            disabled: true
          }
        }
      },
      {
        title: '用户',
        key: 'userId',
        type: 'select',
        fixed: 'left',
        component: {
          type: 'text'
        },
        form: {
          rules: [{ required: true, message: '请输入用户' }]
        },
        editForm: {},
        search: {
          disabled: false
        },
        dict: {
          url: '/sys/user/dict'
        }
      },
      {
        title: '角色',
        key: 'roleId',
        type: 'select',
        fixed: 'left',
        component: {
          type: 'text'
        },
        form: {
          component: {
            multiple: true
          },
          rules: [{ required: true, message: '请输入角色' }]
        },
        editForm: {
          component: {
            disabled: true
          }
        },
        search: {
          disabled: false
        },
        dict: {
          url: '/sys/role/dict'
        }
      },
      {
        title: '启用日期',
        key: 'dateEnable',
        width: 90,
        type: 'date',
        form: {
          addDisabled: true
        },
        editForm: {
          component: {
            disabled: true
          }
        }
      },
      {
        title: '停用日期',
        key: 'dateDisable',
        width: 90,
        type: 'date',
        form: {
          addDisabled: true
        },
        editForm: {
          component: {
            disabled: true
          }
        }
      },
      {
        title: '状态',
        key: 'status',
        type: 'switch',
        width: 70,
        form: {
          addDisabled: true
        }
      },
      {
        title: '主要角色',
        key: 'mainRole',
        type: 'switch',
        width: 70,
        form: {
          addDisabled: true
        }
      }
    ]
  }
}
