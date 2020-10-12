<template>
  <d2-container :class="{ 'page-compact': crud.pageOptions.compact }">
    <template slot="header">
      用户管理
    </template>
    <d2-crud-x
      ref="d2Crud"
      v-bind="_crudProps"
      v-on="_crudListeners"
      @emitAssignRoles="assignRoles"
    >
      <div slot="header">
        <crud-search ref="search" :options="crud.searchOptions" @submit="handleSearch" />
        <el-button slot="header" size="small" type="primary" @click="addRow"
          ><i class="el-icon-plus" />新增</el-button
        >
        <!-- v-permission="'permission:role:add'" -->
        <crud-toolbar
          :search.sync="crud.searchOptions.show"
          :compact.sync="crud.pageOptions.compact"
          @refresh="doRefresh()"
        />
        <!-- <crud-toolbar
          :search.sync="crud.searchOptions.show"
          :compact.sync="crud.pageOptions.compact"
          :columns="crud.columns"
          @refresh="doRefresh()"
          @columns-filter-changed="handleColumnsFilterChanged"
        /> -->
      </div>
      <div></div>
    </d2-crud-x>
    <el-drawer
      :visible.sync="drawer"
      :with-header="false"
      size="40%"
      title="角色配置"
      v-if="drawer"
    >
      <el-tabs class="role-box" type="border-card">
        <el-tab-pane label="用户角色分配"> </el-tab-pane>
      </el-tabs>
    </el-drawer>
  </d2-container>
</template>

<script>
import { d2CrudPlus } from 'd2-crud-plus'
import { crudOptions } from './user.crud'
import { ListObj, NewObj, UpdateObj, DelObj } from './user.api'
export default {
  name: 'User',
  mixins: [d2CrudPlus.crud],
  data() {
    return {
      activeNames: [],
      roleList: [],
      checked: [],
      dialogPermissionVisible: false,
      currentUserId: undefined,
      drawer: false
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
    delRequest(row) {
      return DelObj(row.id)
    },
    assignRoles({ index, row }) {
      this.drawer = true
      console.log(row)
      this.$message('自定义操作按钮：' + row)
    }
  }
}
</script>
