<template>
  <d2-container :class="{ 'page-compact': crud.pageOptions.compact }">
    <template slot="header">
      用户管理 - 角色分配
    </template>
    <d2-crud-x ref="d2Crud" v-bind="_crudProps" v-on="_crudListeners">
      <div slot="header">
        <crud-search ref="search" :options="crud.searchOptions" @submit="handleSearch" />
        <el-button slot="header" size="small" type="primary" @click="addRow"
          ><i class="el-icon-plus" />新增</el-button
        >
        <crud-toolbar
          :search.sync="crud.searchOptions.show"
          :compact.sync="crud.pageOptions.compact"
          @refresh="doRefresh()"
        />
      </div>
      <div></div>
    </d2-crud-x>
  </d2-container>
</template>

<script>
import { d2CrudPlus } from 'd2-crud-plus'
import { crudOptions } from './roles.crud'
import { ListObj, NewObj, UpdateObj } from './roles.api'
export default {
  name: 'UserRoles',
  mixins: [d2CrudPlus.crud],
  data() {
    return {
      activeNames: [],
      roleList: [],
      checked: [],
      dialogPermissionVisible: false,
      currentUserId: undefined
    }
  },
  methods: {
    getCrudOptions() {
      return crudOptions(this)
    },
    pageRequest(query) {
      return ListObj(query)
    },
    addRequest(row) {
      return NewObj(row)
    },
    updateRequest(row) {
      return UpdateObj(row)
    },
    addBefore(row) {
      // 格式化数据
      row.userId = parseInt(row.userId)
      row.roleId.forEach((item, index, array) => {
        row.roleId[index] = parseInt(item)
      })
      // 如果选择Ultraman，则删除。
      if (row.roleId.indexOf(1) >= 0) {
        row.roleId.splice(row.roleId.indexOf(1), 1)
      }
      // 如果没有选择默认用户，则添加。
      if (row.roleId.indexOf(2) < 0) {
        row.roleId.push(2)
      }
    },
    doCellDataChange(context) {
      return UpdateObj(context.row).then(res => {
        this.doRefresh()
      })
    },
    doAfterRowChange(row) {}
  }
}
</script>
