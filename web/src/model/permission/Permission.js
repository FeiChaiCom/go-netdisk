import BaseEntity from '../base/BaseEntity'
import Filter from '../base/Filter'

import { FilterType } from '../base/FilterType'

export default class Permission extends BaseEntity {

  static URL_API_BATCH_DELETE = '/api/permission/batch_delete/'
  static URL_API_FETCH_USERS = '/api/permission/users/'
  static LOCAL_STORAGE_KEY = 'permission'

  constructor(args) {
    super(args)
    this.uuid = null
    this.username = null
    this.projectUuid = null
    this.role = null
    this.updatedBy = null
    this.updatedAt = null

    this.validatorSchema = {
      username: {
        rules: [
          { required: true, message: 'username required' },
          {
            type: 'string',
            pattern: /^[\u4e00-\u9fa5.0-9a-zA-Z_]+$/,
            message: '仅支持中英文数字及下划线，且长度小于64个字符.'
          }],
        error: null
      },
    }
  }


  getUrlPrefix() {
    return '/api/permission'
  }

  render(obj) {
    super.render(obj)
  }

  getFilters() {
    return [
      ...super.getFilters(),
      new Filter(FilterType.INPUT, '用户名', 'username__contains', null, Permission, false),
      new Filter(FilterType.INPUT, '角色', 'role', null, Permission, false),
      new Filter(FilterType.INPUT, '空间', 'project_uuid', null, Permission, false),
    ]
  }

  getForm() {
    return {
      username: this.username,
      role: this.role,
      projectUuid: this.projectUuid,
      uuid: this.uuid ? this.uuid : null
    }
  }

  validate() {
    return super.validate()
  }

  httpBatchDelete(params, successCallback, errorCallback) {
    let that = this
    this.httpPost(Permission.URL_API_BATCH_DELETE, params, function (response) {
      that.render(response.data.data)
      that.safeCallback(successCallback)(response)
    }, errorCallback)
  }

  httpFetchNoAuthUsers(successCallback, errorCallback) {
    let that = this
    this.httpGet(Permission.URL_API_FETCH_USERS, {}, function (response) {
      that.safeCallback(successCallback)(response)
    }, errorCallback)
  }

  getMyProject(successCallback, errorCallback) {
    let that = this
    this.httpGet(this.getUrlPrefix() + '/get_my_project/', {}, function (response) {
      that.safeCallback(successCallback)(response)
    }, errorCallback)
  }

}
