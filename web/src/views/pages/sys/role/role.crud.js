export const crudOptions = () => {
  return {
    options: {
      height: '100%'
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
        disabled: true,
        show: false
      },
      custom: [
        {
          text: ' 分配菜单',
          // icon: 'el-icon-s-flag'
          emit: 'emitAssignMenu',
          // type: 'warning',
          size: 'small'
        }
      ]
    },
    columns: [
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
        title: '角色',
        key: 'name',
        fixed: 'left',
        form: {
          rules: [{ required: true, message: '请输入角色名称' }]
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
          rules: [{ required: true, message: '请输入角色标题' }]
        },
        search: {}
      },
      {
        title: '描述',
        key: 'description',
        show: false,
        form: {
          rules: [{ required: true, message: '请输入角色描述' }]
        },
        search: {}
      },
      {
        title: '状态',
        key: 'status',
        type: 'switch'
      }
    ]
  }
}
