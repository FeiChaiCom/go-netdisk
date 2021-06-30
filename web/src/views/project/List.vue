<template>
  <div class="backyard-user-list animated fadeIn">
    <div class="row">
      <div class="col-md-6  text-left mb10">
        <div class="input-group w300">
          <input type="text" class="form-control" v-model="searchText" @keyup.enter="searchProject"
                 :placeholder="$t('project.searchProject')">
          <span class="input-group-btn">
          <button type="button" class="btn btn-primary" @click.prevent.stop="searchProject">
            <i class="fa fa-search"></i>
          </button>
        </span>
        </div>
      </div>
      <div class="col-md-6 text-right mb10" v-if="accessControl.role === UserRole.ADMINISTRATOR">
        <router-link class="btn btn-primary btn-sm" to="/project/create">
          <i class="fa fa-plus"></i>
          创建空间
        </router-link>
      </div>

    </div>
    <div class="row">
      <div class="col-md-6" v-for="(projectItem,index) in pager.data">
        <div class="bg-white border br4 p10 mb10">
          <div class="media">
            <div class="pull-left">
<!--              <router-link :to="'/project/detail/'+projectItem.uuid">-->
                <img class="img-circle img-md" :src="projectItem.getAvatarUrl()">
<!--              </router-link>-->
            </div>
            <div class="media-body">
              <div>
							<span class="f16">
<!--								<router-link class="black" :to="'/project/detail/'+projectItem.uuid">-->
                  <span>
                    {{projectItem.name}}
                  </span>
<!--								</router-link>-->
							</span>
              </div>
              <div class="mv5 text-muted one-line">
                {{projectItem.description || '暂无描述'}}
              </div>
              <div class="mv5">
                <span class="mr10">
                  {{$t('project.userCount')}}:
                    <span>
                        {{ projectItem.userCount || '0' }}
                    </span>
                </span>
                <span class="mr10">
                  {{$t('project.directoryCount')}}:
                    <span>
                        {{ projectItem.directoryCount || '0' }}
                    </span>
                </span>
                <span class="mr10">
                  {{$t('project.totalFileSize')}}:
                    <span>
                        {{ projectItem.totalFileSize || '0' }}
                    </span>
                </span>

              </div>

              <div class="mv5">
                <span class="pull-right action-buttons">

                  <router-link :to="'/project/edit/'+projectItem.uuid" :title="$t('edit')">
										<i class="fa fa-pencil text-info f18"></i>
									</router-link>

									<a href="javascript:void(0)"
                     v-if="accessControl.role === UserRole.ADMINISTRATOR"
                     :title="$t('project.deleteProject')" @click.stop.prevent="deleteProject(projectItem)">
                    <i class="fa fa-close text-danger f18"></i>
									</a>
							</span>

              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-md-12 mt20">
        <NbPager :pager="pager" :callback="refresh"></NbPager>
      </div>
    </div>
  </div>
</template>

<script>
  import NbPlainFilter from '../../components/filter/NbPlainFilter.vue'
  import NbPager from '../../components/NbPager.vue'
  import Pager from '../../model/base/Pager'
  import {UserRole, UserRoleMap} from "../../model/user/UserRole";
  import {UserStatus} from "../../model/user/UserStatus";
  import {handleImageUrl} from "../../common/util/ImageUtil";
  import {SortDirection} from "../../model/base/SortDirection";
  import Project from '../../model/project/Project'
  import { Message, MessageBox } from 'element-ui'
  import { mapState } from 'vuex'

  export default {

    data() {
      return {
        UserRole,
        UserRoleMap,
        UserStatus,
        //搜索的文字
        searchText: null,
        pager: new Pager(Project),
        user: this.$store.state.user
      }
    },
    computed: {
      ...mapState(['accessControl']),
    },
    components: {
      NbPlainFilter,
      NbPager
    },
    methods: {
      handleImageUrl,
      searchProject() {

        let that = this
        if (that.searchText) {

          //刷新面包屑
          that.pager.resetFilter()
          that.pager.setFilterValue('name__contains', that.searchText)
          that.pager.httpFastPage()

        } else {
          that.pager.resetFilter()
          that.refresh()
        }

      },
      search() {
        this.pager.page = 1
        this.refresh()
      },
      refresh() {
        this.pager.httpFastPage()
      },
      deleteProject(project) {
        let that = this
        MessageBox.confirm(that.$t("actionCanNotRevertConfirm"), that.$t("prompt"), {
          confirmButtonText: that.$t("confirm"),
          cancelButtonText: that.$t("cancel"),
          type: 'warning',
          callback: function (action, instance) {
            if (action === 'confirm') {
              project.httpDelete((response) => {
                Message.success(that.$t('operationSuccess'))
                that.refresh()
              })
            }

          }
        })
      }
    },
    mounted() {
      this.pager.enableHistory()
      // this.pager.setFilterValue("orderLastTime", SortDirection.DESC)
      this.refresh()
    }
  }
</script>

<style lang="less" rel="stylesheet/less">
  .backyard-user-list {

  }
</style>
