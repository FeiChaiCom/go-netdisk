<template>
  <el-dialog
    :title="title"
    :visible.sync="dialogFormVisible"
    width="650px"
    :modal="false"
    @close="close"
    @open="open"
  >
    <el-form ref="form" :model="form" :rules="rules" label-width="80px">
      <el-form-item label="用户名" prop="username">
        <el-select
          v-model="form.username"
          size="small"
          filterable
          allow-create
          default-first-option
          :disabled="isUserReadOnly"
          :loading="loadingUsers"
          placeholder="请选择用户">
          <el-option
            v-for="item in userList"
            :key="item.username"
            :label="item.username"
            :value="item.username">
          </el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="角色" prop="role">
        <el-radio-group v-model="form.role">
          <template v-for="role in UserRoleMap">
            <el-radio :label="role.value"
                      v-if="role.value !== 'ADMINISTRATOR' && role.value !== 'GUEST'">
              {{ $t(role.name) }}
            </el-radio>
          </template>
          <el-radio label="ADMINISTRATOR" v-if="accessControl.role === 'ADMINISTRATOR'">系统管理员</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item :label="$t('project.name')" prop="projectUuid" v-if="form.role !== 'ADMINISTRATOR'">
        <el-select
          v-model="form.projectUuid"
          size="small"
          filterable
          clearable
          default-first-option
          :loading="loadingProjects"
          placeholder="请选择">
          <el-option
            v-for="item in projectList"
            :key="item.uuid"
            :label="item.name"
            :value="item.uuid">
          </el-option>
        </el-select>
      </el-form-item>

    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="close">取 消</el-button>
      <el-button type="primary" @click="save">确 定</el-button>
    </div>
  </el-dialog>
</template>

<script>
import Project from '../../model/project/Project'
import User from '../../model/user/User'
import Permission from '../../model/permission/Permission'
import { mapState } from 'vuex'
import { UserRoleMap } from '../../model/permission/UserAclRole'

export default {
  name: 'Edit',
  data() {
    return {
      UserRoleMap,
      isUserReadOnly: false,
      project: new Project(),
      user: new User(),
      permission: new Permission(),
      loadingUsers: true,
      loadingProjects: true,
      userList: [],
      projectList: [],
      form: {
        username: '',
        projectUuid: '',
        role: 'USER',
      },
      rules: {
        username: [
          { required: true, trigger: 'blur', message: '请输入用户名' },
        ],
        projectUuid: [
          { required: true, trigger: 'blur', message: '请选择空间' },
        ],
        role: [{ required: true, trigger: 'blur', message: '请选择角色' }],
      },
      title: '',
      dialogFormVisible: false,
    }
  },
  computed: {
    ...mapState(['accessControl']),
  },
  mounted() {
    this.getUsers()
    this.getProjects()
  },
  methods: {
    getUsers() {
      this.loadingUsers = true
      this.permission.httpFetchNoAuthUsers((res) => {
        this.loadingUsers = false
        this.userList = res.data.data
      })
    },
    getProjects() {
      this.loadingProjects = true
      this.project.httpFetchAll((res) => {
        this.loadingProjects = false
        this.projectList = res.data.data
      })
    },
    showEdit(row) {
      if (!row) {
        this.title = '添加用户权限'
        this.isUserReadOnly = false
        if (this.projectList.length) {
          this.form.projectUuid = this.projectList[0].uuid
        }
      } else {
        this.title = '编辑用户权限'
        this.isUserReadOnly = true
        this.form = Object.assign({}, row)
      }
      this.dialogFormVisible = true
    },
    open() {
      this.getUsers()
    },
    close() {
      this.$refs['form'].resetFields()
      this.form = this.$options.data().form
      this.dialogFormVisible = false
    },
    save() {
      let that = this
      const permission = new Permission()
      permission.render(this.form)
      permission.httpSave(function (response) {
        if (response.data.result) {
          that.$message.success({
            message: that.$t('operationSuccess')
          })
          that.$emit('fetch-data')
          that.close();
        } else {
          that.$message.error(response.data.msg);
        }
      })
    }
  },
}
</script>

<style lang="less" rel="stylesheet/less">

.el-select {
  width: 100%;
}

.el-radio-group {
  vertical-align: initial !important;
}

.el-dialog__header {
  padding: 15px;
  border-bottom: 1px solid #dcdfe6;
}

.el-dialog__body {
  padding: 20px 25px 0 20px !important;
}

.el-dialog__footer {
  padding: 15px;
  text-align: right;
  border-top: 1px solid #dcdfe6;

  .dialog-footer {
    .el-button {
      padding: 9px 15px;
    }

    .el-button {
      font-size: 12px;
      border-radius: 3px;
    }
  }
}

</style>
