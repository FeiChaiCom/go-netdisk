<template>
  <div class="animated fadeIn backyard-user-edit">

    <div class="user-block netdisk-box">
      <div>
<!--        <div class="row mt10">-->
<!--          <label class="col-md-2 control-label mt5">{{$t('project.avatar')}}</label>-->
<!--          <div class="col-md-10">-->
<!--            <div>-->
<!--              <MatterImage v-model="currentProject.avatarUrl" uploadHint=""/>-->
<!--            </div>-->
<!--          </div>-->
<!--        </div>-->

        <div class="row mt10" v-validator="currentProject.validatorSchema.name.error">
          <label class="col-md-2 control-label mt5 compulsory">{{$t('project.name')}}</label>
          <div class="col-md-10 validate">
            <input type="text" class="form-control"
                   :disabled="!createMode"
                   v-model="currentProject.name">
          </div>
        </div>

        <div class="row mt10">
          <label class="col-md-2 control-label mt5">{{$t('project.description')}}</label>
          <div class="col-md-10">
            <el-input
              type="textarea"
              :rows="3"
              placeholder="请输入空间描述"
              v-model="currentProject.description">
            </el-input>
          </div>
        </div>

<!--        <div class="row mt10" v-if="user.role === UserRole.ADMINISTRATOR">-->
<!--          <label class="col-md-2 control-label mt5 compulsory">{{$t('project.role')}}</label>-->
<!--          <div class="col-md-10">-->
<!--            <select class="form-control" v-model="currentProject.role">-->
<!--              <option v-for="item in UserRoleList" v-if="item.value !== UserRole.GUEST" :value="item.value">-->
<!--                {{$t(item.name)}}-->
<!--              </option>-->
<!--            </select>-->
<!--          </div>-->
<!--        </div>-->

      </div>
    </div>
    <div class="mt10 text-right">
      <CreateSaveButton :entity="currentProject" :callback="save"></CreateSaveButton>
    </div>

  </div>
</template>

<script>
  import NbRadio from '../../components/NbRadio.vue'
  import MatterImage from '../matter/widget/MatterImage'
  import CreateSaveButton from '../../components/CreateSaveButton'
  import {UserRole, UserRoleList, UserRoleMap} from "../../model/user/UserRole";
  import {UserStatus, UserStatusList, UserStatusMap} from "../../model/user/UserStatus";
  import Project from '../../model/project/Project'

  export default {

    data() {
      return {
        UserRole,
        UserRoleList,
        UserRoleMap,
        UserStatus,
        UserStatusList,
        UserStatusMap,
        createMode: false,
        confirmPassword: null,
        user: this.$store.state.user,
        currentProject: new Project(),
        breadcrumbs: this.$store.state.breadcrumbs
      }
    },
    components: {
      NbRadio,
      MatterImage,
      CreateSaveButton
    },
    methods: {
      save() {
        let that = this
        this.currentProject.httpSave(function (response) {
          that.$message.success({
            message: that.$t('operationSuccess')
          })

          that.$router.go(-1)
        })
      }
    },
    created() {
      let that = this
      if (this.user.role === UserRole.ADMINISTRATOR) {
        this.breadcrumbs.splice(0, this.breadcrumbs.length)
        this.breadcrumbs.push({
          displayDirect: true,
          title: that.$t('project.profile'),
          path: '/user/detail/' + this.user.uuid
        }, {
          displayDirect: true,
          title: that.$t('edit')
        })
      }
    },
    mounted() {
      let that = this
      this.currentProject.errorMessage = null
      this.currentProject.uuid = this.$store.state.route.params.uuid
      if (this.currentProject.uuid) {
        this.createMode = false
        this.currentProject.httpDetail()
      } else {
        this.createMode = true
        this.currentProject.role = UserRole.USER
      }
    }
  }
</script>

<style lang="less" rel="stylesheet/less">
  .backyard-user-edit {

    .user-block {
      margin-top: 10px;
      margin-bottom: 10px;
    }

  }
</style>
