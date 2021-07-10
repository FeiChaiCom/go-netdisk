import Vue from 'vue'
import Router from 'vue-router'
import ByFrameView from '../views/Frame.vue'
import ShareList from '../views/share/List'
import ShareDetail from '../views/share/Detail'
import MatterList from '../views/matter/List'
import MatterDetail from '../views/matter/Detail'
import UserLogin from '../views/user/Login'
import UserAuthentication from '../views/user/Authentication'
import UserRegister from '../views/user/Register'
import UserList from '../views/user/List'
import ProjectList from '../views/project/List'
import PermissionList from '../views/permission/List'
import UserDetail from '../views/user/Detail'
import UserChangePassword from '../views/user/ChangePassword'
import UserEdit from '../views/user/Edit'
import ProjectEdit from '../views/project/Edit'
import DashboardIndex from '../views/dashboard/Index'
import PreferenceIndex from '../views/preference/Index'
import PreferenceEdit from '../views/preference/Edit'
import NotFound from '../views/layout/NotFound'
import store from '../vuex'
import Cookies from "js-cookie";

Vue.use(Router)

const router = new Router({
  mode: 'hash',
  linkActiveClass: 'is-link-active',
  scrollBehavior: () => ({y: 0}),
  routes: [
    {
      path: '/',
      component: ByFrameView,
      children: [
        {
          path: '',
          name: 'MatterList',
          component: MatterList,
          meta: {
            //here is i18n key
            title: 'router.allFiles',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'MatterList',
                title: 'router.allFiles'
              }
            ]
          }
        },
        {
          path: 'matter/detail/:uuid',
          name: 'MatterDetail',
          component: MatterDetail,
          meta: {
            title: 'router.fileDetail',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'MatterList',
                title: 'router.allFiles'
              },
              {
                name: 'MatterDetail',
                title: 'router.fileDetail'
              }
            ]
          }
        },
        {
          path: 'user/login',
          name: 'UserLogin',
          component: UserLogin,
          meta: {
            title: 'router.login',
            requiresAuth: false,
            breadcrumbs: []
          }
        },
        {
          path: 'user/authentication/:authentication',
          name: 'UserAuthentication',
          component: UserAuthentication,
          meta: {
            title: 'router.autoLogin',
            requiresAuth: false,
            breadcrumbs: []
          }
        },
        {
          path: 'user/register',
          name: 'UserRegister',
          component: UserRegister,
          meta: {
            title: 'router.register',
            requiresAuth: false,
            breadcrumbs: []
          }
        },
        {
          path: 'user/list',
          name: 'UserList',
          component: UserList,
          meta: {
            title: 'router.users',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'UserList',
                title: 'router.users'
              }
            ]
          }
        },
        {
          path: 'project/list',
          name: 'ProjectList',
          component: ProjectList,
          meta: {
            title: 'router.projects',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'ProjectList',
                title: 'router.projects'
              }
            ]
          }
        },
        {
          path: 'permission/list',
          name: 'PermissionList',
          component: PermissionList,
          meta: {
            title: 'router.permission',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'PermissionList',
                title: 'router.permission'
              }
            ]
          }
        },
        {
          path: 'user/detail/:uuid',
          name: 'UserDetail',
          component: UserDetail,
          meta: {
            title: 'router.userDetail',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'UserList',
                title: 'router.users'
              },
              {
                name: 'UserDetail',
                title: 'router.userDetail'
              }
            ]
          }
        },
        {
          path: 'user/change/password',
          name: 'UserChangePassword',
          component: UserChangePassword,
          meta: {
            title: 'router.changePassword',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'UserChangePassword',
                title: 'router.changePassword'
              }
            ]
          }
        },
        {
          path: 'user/create',
          name: 'UserCreate',
          component: UserEdit,
          meta: {
            title: 'router.createUser',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'UserList',
                title: 'router.users'
              },
              {
                name: 'UserCreate',
                title: 'router.createUser'
              }
            ]
          }
        },

        {
          path: 'project/create',
          name: 'ProjectCreate',
          component: ProjectEdit,
          meta: {
            title: 'router.createProject',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'ProjectList',
                title: 'router.projects'
              },
              {
                name: 'ProjectCreate',
                title: 'router.createProject'
              }
            ]
          }
        },

        {
          path: 'user/edit/:uuid',
          name: 'UserEdit',
          component: UserEdit,
          meta: {
            title: 'router.editUser',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'UserList',
                title: 'router.users'
              },
              {
                name: 'UserEdit',
                title: 'router.editUser'
              }
            ]
          }
        },

        {
          path: 'project/edit/:uuid',
          name: 'ProjectEdit',
          component: ProjectEdit,
          meta: {
            title: 'router.editUser',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'ProjectList',
                title: 'router.projects'
              },
              {
                name: 'ProjectEdit',
                title: 'router.editProject'
              }
            ]
          }
        },

        {
          path: 'share/detail/:uuid',
          name: 'ShareDetail',
          component: ShareDetail,
          meta: {
            title: 'router.shareDetail',
            requiresAuth: false,
            breadcrumbs: []
          }
        },
        {
          path: 'share/list',
          name: 'ShareList',
          component: ShareList,
          meta: {
            title: 'router.myShare',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'ShareList',
                title: 'router.myShare'
              }
            ]
          }
        },
        {
          path: 'dashboard/index',
          name: 'DashboardIndex',
          component: DashboardIndex,
          meta: {
            title: 'router.dashboard',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'DashboardIndex',
                title: 'router.dashboard'
              }
            ]
          }
        },
        {
          path: 'preference',
          name: 'PreferenceIndex',
          component: PreferenceIndex,
          meta: {
            title: 'router.setting',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'PreferenceIndex',
                title: 'router.setting'
              }
            ]
          }
        },

        {
          path: 'preference/edit',
          name: 'PreferenceEdit',
          component: PreferenceEdit,
          meta: {
            title: 'router.setting',
            requiresAuth: true,
            breadcrumbs: [
              {
                name: 'PreferenceIndex',
                title: 'router.setting'
              },
              {
                name: 'PreferenceEdit',
                title: 'router.setting'
              }
            ]
          }
        },
        //未被上面处理的route被视为404
        {
          path: '*',
          component: NotFound,
          meta: {requiresAuth: false}
        }
      ]
    }
  ]
})

//装填面包屑
function fillBreadcrumbs(to) {
  //清空数组
  store.state.breadcrumbs.splice(0, store.state.breadcrumbs.length);
  if (to.meta.breadcrumbs) {
    //追加一个数组
    store.state.breadcrumbs.push.apply(store.state.breadcrumbs, to.meta.breadcrumbs)
  }
}

//add global interceptor.
router.beforeEach((to, from, next) => {

  //handle auth feature.
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // this route requires auth, check if logged in
    // if not, redirect to login page.
    let uid = Cookies.get("bk_uid");
    if (!store.state.user.uuid && !uid) {
      next({
        path: '/user/login',
        query: {redirect: to.fullPath}
      })
    } else {

      fillBreadcrumbs(to);
      next()
    }
  } else {

    fillBreadcrumbs(to);
    next()
  }
})

export default router
