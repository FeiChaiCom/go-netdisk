import BaseEntity from '../base/BaseEntity'
import Filter from '../base/Filter'

import { FilterType } from '../base/FilterType'
import { handleImageUrl } from '../../common/util/ImageUtil'
import { MessageBox, Message } from 'element-ui'

let defaultAvatarPath = require('../../assets/img/avatar-project.png')

export default class Project extends BaseEntity {

  static LOCAL_STORAGE_KEY = 'project'

  constructor(args) {
    super(args)
    this.name = null
    this.description = null
    this.avatarUrl = null

    // 只读属性
    this.userCount = 0
    this.directoryCount = 0
    this.totalFileSize = 0
    this.createdBy = null
    this.createdAt = null

    this.validatorSchema = {
      name: {
        rules: [
          { required: true, message: 'name required' },
          {
            type: 'string',
            pattern: /^[\u4e00-\u9fa5.0-9a-zA-Z_]+$/,
            message: '仅支持中英文数字及下划线，且长度小于64个字符.'
          }],
        error: null
      }
    }
  }

  getAvatarUrl() {
    if (this.avatarUrl) {
      return handleImageUrl(this.avatarUrl)
    } else {
      return defaultAvatarPath
    }
  }

  getUrlPrefix() {
    return '/api/project'
  }

  render(obj) {
    super.render(obj)
    this.renderEntity('lastTime', Date)
  }

  getFilters() {
    return [
      ...super.getFilters(),
      new Filter(FilterType.INPUT, '空间名称', 'name__contains', null, Project, false),
    ]
  }

  getForm() {
    return {
      name: this.name,
      description: this.description,
      avatarUrl: this.avatarUrl,
      uuid: this.uuid ? this.uuid : null
    }
  }

  httpFetchAll(successCallback, errorCallback) {
    let that = this
    this.httpGet(this.getUrlPrefix() + '/', {}, function (response) {
      that.render(response.data.data)
      that.safeCallback(successCallback)(response)
    }, errorCallback)
  }

  validate() {
    return super.validate()
  }

}
