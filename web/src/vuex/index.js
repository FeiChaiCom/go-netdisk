import Vue from 'vue'
import Vuex from 'vuex'
import User from "../model/user/User"
import Preference from '../model/preference/Preference'
import BrowserUtil from "../common/util/BrowserUtil"
import Permission from '../model/permission/Permission'
import Cookies from "js-cookie"

Vue.use(Vuex)

let user = new User()
let permission = new Permission()
user.renderFromLocalStorage()

let lang = BrowserUtil.browserLang()
let localLang = Cookies.get("_lang");
if (localLang === "zh" || localLang === "en") {
  lang = localLang
}

const state = {
  config: {
    mobile: false,
    showDrawer: true
  },
  //当前版本信息。
  versionName: '0.0.6',
  //当前用户，即使没有登录依然有游客的用户在。
  user,
  accessControl: {},
  project: {},
  breadcrumbs: [],
  //全局正在上传的文件
  uploadMatters: [],
  //当前接受上传的那个Matter List.vue实例
  uploadListInstance: null,

  //当前的语言
  lang: lang,

  //网站设置
  preference: new Preference(),
  //上次报没有登录错误的时间戳，用于控制登录提示框的个数不能太频繁。
  lastLoginErrorTimestamp: 0

}

const getters = {
  getConfig(state) {
    return state.config
  }
}

const mutations = {
  setAccessControl(state, data) {
    state.accessControl = data
  },
  setProject(state, data) {
    state.project = data
  }
}

const actions = {
  getAccessControl(context) {
    user.getAccessControl((res) => {
      if (res.data.result) {
        context.commit('setAccessControl', res.data.data)
      }
    });
  },
  getMyProject(context) {
    permission.getMyProject((res) => {
      if (res.data.result) {
        context.commit('setProject', res.data.data)
      }
    });
  }
}

export default new Vuex.Store({
  state,
  getters,
  mutations,
  actions
})
