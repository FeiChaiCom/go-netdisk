<template>
  <div class="backyard-user-list animated fadeIn">
    <div class="row query-form">
      <div class="col-md-4 left-panel">
        <el-form :inline="true">
          <el-form-item>
            <el-button icon="el-icon-plus" type="primary" size="small" @click="handleEdit">
              添加
            </el-button>
          </el-form-item>
          <el-form-item>
            <el-button icon="el-icon-delete" type="default" size="small" @click="handleDelete">
              批量删除
            </el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="col-md-8 right-panel">
        <el-form :inline="true" :model="queryForm" @submit.native.prevent>
          <el-form-item>
            <el-select
              v-model="queryForm.projectUuid"
              size="small"
              clearable
              filterable
              :placeholder="$t('permission.selectTip')">
              <el-option
                v-for="item in projectList"
                :key="item.uuid"
                :label="item.name"
                :value="item.uuid">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-select
              v-model="queryForm.role"
              size="small"
              clearable
              placeholder="请选择角色">
              <el-option
                v-for="item in projectRole"
                v-if="item.key !== 'GUEST'"
                :key="item.key"
                :label="$t(item.name)"
                :value="item.key">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-input
              v-model.trim="queryForm.username"
              placeholder="请输入用户名"
              clearable
              size="small"
            />
          </el-form-item>
          <el-form-item>
            <el-button icon="el-icon-search" type="primary" size="small" @click="searchPermission">
              查询
            </el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <div class="row user-table-wrapper">
      <el-table
        ref="permissionTable"
        v-loading="listLoading"
        :element-loading-text="elementLoadingText"
        @selection-change="setSelectRows"
        :data="pager.data"
        max-height="705"
        tooltip-effect="dark"
        style="width: 100%">

        <el-table-column show-overflow-tooltip type="selection"></el-table-column>
        <el-table-column
          show-overflow-tooltip
          prop="username"
          label="用户名"
        ></el-table-column>
        <el-table-column
          show-overflow-tooltip
          prop="projectName"
          :label="$t('project.name')"
        ></el-table-column>

        <el-table-column show-overflow-tooltip
                         prop="role"
                         label="角色">
          <template slot-scope="scope">
            <i :class="projectRole[scope.row.role].icon"></i>
            <span style="margin-left: 10px">{{ $t(projectRole[scope.row.role].name) }}</span>
          </template>
        </el-table-column>

        <el-table-column
          show-overflow-tooltip
          prop="updatedAt"
          label="更新时间"
          width="200"
        ></el-table-column>

        <el-table-column
          show-overflow-tooltip
          prop="updatedBy"
          label="更新人"
        ></el-table-column>

        <el-table-column show-overflow-tooltip label="操作"
                         width="250"
                         align="center">
          <template slot-scope="scope">
            <el-button type="text"
              size="mini"
              @click="handleEdit(scope.row)">编辑</el-button>
            <el-button type="text"
              size="mini"
              @click="handleDelete(scope.row)">删除</el-button>
          </template>
        </el-table-column>

      </el-table>

      <el-pagination
        background
        :current-page="queryForm.page"
        :page-size="queryForm.pageSize"
        :layout="layout"
        :total="queryForm.total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      ></el-pagination>
      <edit ref="edit" @fetch-data="refresh"></edit>

<!--      <div class="col-md-12 mt20">-->
<!--        <NbPager :pager="pager" :callback="refresh"></NbPager>-->
<!--      </div>-->
    </div>
  </div>
</template>

<script>
  import NbPlainFilter from '../../components/filter/NbPlainFilter.vue'
  import NbPager from '../../components/NbPager.vue'
  import Pager from '../../model/base/Pager'
  import Permission from '../../model/permission/Permission'
  import {UserRoleMap} from "../../model/permission/UserAclRole";
  import Edit from './Edit'
  import { Message, MessageBox } from 'element-ui'
  import Project from '../../model/project/Project'

  export default {

    data() {
      return {
        listLoading: false,
        layout: 'total, sizes, prev, pager, next, jumper',
        selectRows: '',
        elementLoadingText: '正在加载...',
        queryForm: {
          total: 0,
          page: 1,
          pageSize: 10,
          username: '',
          role: '',
          projectUuid: '',
        },
        project: new Project(),
        loadingProjects: true,
        projectList: [],
        projectRole: UserRoleMap,

        pager: new Pager(Permission),
        permission: new Permission(),
        user: this.$store.state.user
      }
    },
    components: {
      Edit,
      NbPlainFilter,
      NbPager
    },
    created() {
    },
    mounted() {
      this.refresh()
      this.getProjects()
    },
    methods: {
      refresh() {
        this.listLoading = true;
        this.pager.httpFastPage((resp) => {
          this.listLoading = false
          const { data } = resp.data
          this.queryForm.total = data.totalItems
          this.queryForm.page = data.page
          this.queryForm.pageSize = data.pageSize
        })
      },
      getProjects() {
        this.loadingProjects = true
        this.project.httpFetchAll((res) => {
          this.loadingProjects = false
          this.projectList = res.data.data
        })
      },
      searchPermission() {
        let that = this

        that.pager.resetFilter()

        if (that.queryForm.username) {
          that.pager.setFilterValue('username__contains', that.queryForm.username)
        }

        if (that.queryForm.projectUuid) {
          that.pager.setFilterValue('project_uuid', that.queryForm.projectUuid)
        }

        if (that.queryForm.role) {
          that.pager.setFilterValue('role', that.queryForm.role)
        }

        that.refresh()
      },
      setSelectRows(val) {
        this.selectRows = val
      },
      handleEdit(row) {
        if (row.uuid) {
          this.$refs['edit'].showEdit(row)
        } else {
          this.$refs['edit'].showEdit()
        }
      },
      handleDelete(row) {
        let that = this
        if (row.uuid) {
          MessageBox.confirm('你确定要删除当前项吗', that.$t("prompt"), {
            confirmButtonText: that.$t("confirm"),
            cancelButtonText: that.$t("cancel"),
            type: 'warning',
            callback: function (action, instance) {
              if (action === 'confirm') {
                row.httpDelete((response) => {
                  if(response.data.result) {
                    that.$message.success(that.$t('operationSuccess'))
                    that.refresh()
                  } else {
                    that.$message.error(response.data.msg)
                  }
                })
              }

            }
          })
        } else {
          if (this.selectRows.length > 0) {
            const uuids = this.selectRows.map((item) => item.uuid).join(',')
            let that = this
            MessageBox.confirm('你确定要删除选中项吗', that.$t("prompt"), {
              confirmButtonText: that.$t("confirm"),
              cancelButtonText: that.$t("cancel"),
              type: 'warning',
              callback: function (action, instance) {
                if (action === 'confirm') {
                  const params = { uuids: uuids }
                  that.permission.httpBatchDelete(params, response => {
                    if(response.data.result) {
                      that.$message.success(that.$t('operationSuccess'))
                      that.refresh()
                    } else {
                      that.$message.error(response.data.msg)
                    }
                  })
                }

              }
            })

          } else {
            that.$message.error({
              message: '未选中任何行'
            })
            return false
          }
        }
      },
      handleSizeChange(val) {
        this.queryForm.pageSize = val
        this.pager.pageSize = val
        this.refresh()
      },
      handleCurrentChange(val) {
        this.queryForm.page = val
        this.pager.page = val
        this.refresh()
      },
    },
  }
</script>

<style lang="less" rel="stylesheet/less">
  .backyard-user-list {
    .user-table-wrapper {
      padding: 0 20px;
    }
    .query-form {
      padding: 0 5px;
      margin-bottom: 10px;
      .left-panel {
        display: flex;
        flex-wrap: wrap;
        align-items: center;
        justify-content: flex-start;
      }
      .right-panel {
        display: flex;
        flex-wrap: wrap;
        align-items: center;
        justify-content: flex-end;
      }
      .el-form-item {
        margin-bottom: 0;
      }
    }
    .el-pagination {
      padding: 2px 5px;
      margin: 15px 0 0 0;
      font-weight: 400;
      color: #000;
      text-align: center;
    }
  }
</style>
