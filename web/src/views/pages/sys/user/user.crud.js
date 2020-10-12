// import { mobileValidator } from 'el-phone-number-input'
// import { request } from '@/api/service'
export const crudOptions = () => {
  return {
    options: {
      height: '100%' // 表格高度100%, 使用toolbar必须设置
    },
    formOptions: {
      // 编辑对话框及el-form的配置
      labelPosition: 'top'
    },
    searchOptions: {
      // 查询配置参数, 支持el-form的配置参数
      form: {
        // 默认搜索参数
        username: '%' // 请求列表默认会带上此处配置参数,重置后会恢复成此处配置的值
      },
      size: 'small',
      show: false, // 是否显示搜索工具条
      // disabled: false, //是否禁用搜索工具条
      debounce: {
        // 自动查询防抖,debounce:false关闭自动查询
        wait: 500 // 延迟500毫秒
        // ... //options : https://www.lodashjs.com/docs/lodash.debounce
      }
    },
    rowHandle: {
      // columnHeader: '操作',
      width: 145,
      fixed: 'right',
      view: {
        thin: true,
        text: null
        // disabled() {
        //   return !this.hasPermissions('permission:role:view')
        // }
      },
      edit: {
        thin: true,
        text: null
      },
      remove: {
        thin: true,
        text: null,
        disabled: true
      },
      custom: [
        {
          text: ' 设置角色',
          // icon: 'el-icon-s-flag'
          emit: 'emitAssignRoles',
          // type: 'warning',
          size: 'small'
        }
      ]
    },
    // 定义列
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
        title: '用户名',
        key: 'username',
        fixed: 'left',
        form: {
          rules: [{ required: true, message: '请输入用户名' }]
        },
        editForm: {
          component: {
            disabled: true
          }
        },
        search: {}
      },
      {
        title: '密码',
        key: 'password',
        show: false,
        form: {
          component: {
            type: 'password'
          },
          rules: [
            { required: true, message: '请输入密码' }
            // { min: 6, max: 25, message: '长度在 6 到 25个字符', trigger: 'blur' },
            // {
            //   pattern: /^(\w){6,25}$/,
            //   message: '只能输入6-25个字母、数字、下划线',
            //   trigger: 'blur'
            // }
          ]
        }
      },
      {
        title: '电子邮箱',
        key: 'email',
        form: {
          rules: [
            { required: true, message: '请输入邮箱' },
            { type: 'email', message: '请正确输入邮箱', trigger: 'blur' }
          ]
        }
      },
      {
        title: '移动电话',
        key: 'telephone',
        width: 120,
        form: {
          rules: [
            // { type: 'number', message: '必须为数字', trigger: 'blur' }
            //   {
            //     validator: function(value, callback) {
            //       if (/^1[34578]\d{9}$/.test(value) == false) {
            //         callback(new Error('请输入正确的手机号'))
            //       } else {
            //         callback()
            //       }
            //     },
            //     trigger: 'blur'
            //   }
          ]
        }
      },
      {
        title: '启用日期',
        key: 'dateEnable',
        width: 100,
        type: 'date',
        form: {
          addDisabled: true,
          component: {
            disabled: true
          }
        }
      },
      {
        title: '停用日期',
        key: 'dateDisable',
        width: 100,
        type: 'date',
        form: {
          addDisabled: true,
          component: {
            disabled: true
          }
        }
      },
      {
        title: '启用状态',
        key: 'status',
        width: 80,
        type: 'switch',
        component: {
          disabled: true
        },
        form: {
          addDisabled: true
        }
      }
    ]
  }
}
