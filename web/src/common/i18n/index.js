// 准备翻译的语言环境信息
let i18nMessage = {
  en: {
    oemTitle: "netdisk",
    dashboard: {
      totalInvokeNum: 'Total PV',
      weekRate: 'Week',
      dayRate: 'Day',
      yesterdayInvoke: 'Yesterday PV',
      totalUV: 'Total UV',
      yesterdayUV: 'Yesterday UV',
      totalMatterNum: 'Total file num',
      yesterdayMatterNum: 'Yesterday File Num',
      totalFileSize: 'Total File Size',
      yesterdayMatterSize: 'Yesterday File Size',
      recentDayInvokeUV: 'Recent {0} days PV/UV',
      downloadMatterTop10: 'File download TOP10',
      activeIpTop10: 'Active IP TOP10',
      loading: 'loading...',
      date: 'Date',
      num: 'Num'
    },
    install: {
      configMysql: "Config MySQL",
      port: "Port",
      schema: "Schema",
      mysqlConnectionPass: "Connect MySQL Ok",
      testMysqlConnection: "Tes MySQL Connection",
      notice: "Notice",
      mysqlNotice1: "If Mysql and oemTitle installed on the same server, Host is 127.0.0.1",
      mysqlNotice2: "Your mysql account must have access to create table, or the second step will fail.",
      validateMysqlFirst: "Please test the mysql connection firstly.",
      preStep: "Pre Step",
      nextStep: "Next Step",
      createTable: "Craete Tables",
      installed: "Installed",
      installedButMissing: "Installed, but missing fields",
      toBeInstalled: "To be installed",
      allFields: "All fields",
      missingFields: "Missing fields",
      tableNotice: "'Create Tables' will trigger the following actions:",
      tableNotice1: "If a table not exist, create it.",
      tableNotice2: "If a table exist and no fields missing, nothing will do on this table.",
      tableNotice3: "If a table exist but some fields is missing, it will add the missing fields.",
      tableNotice4: "If a table exist and some fields not necessary, nothing will do on this table.",
      oneKeyCreate: "Create Tables",
      createFinish: "Finish Creating Tables",
      createTableSuccess: "Create tables successfully",
      crateTableFirst: "Please click 'Create Tables'",
      setAdministrator: "Config Administrator",
      detectAdministrator: "Detect the following administrators:",
      useOrCreateAdministrator: "You can validate one of them, or you can create a new one.",
      validateAdministrator: "Validate administrator",
      createAdministrator: "Create administrator",
      administratorUsername: "username",
      administratorPassword: "password",
      administratorRePassword: "Enter administrator password again",
      usernameRule: "oemTitle will use username as directory name, so only lowercase letter and number and _ is permitted.",
      congratulationInstall: "Congratulations, install successfully!",
      configAdminFirst: "Please config administrator first.",
      createAdminSuccess: "Create administrator successfully!",
      validateAdminSuccess: "Validate administrator successfully!",
      enterHome: "Click to enter home"
    },
    layout: {
      allFiles: "Files",
      myShare: "My Share",
      setting: "Setting",
      dashboard: "Dashboard",
      users: "Users",
      projects: "Projects",
      logout: "Logout",
      about: "About",
      install: "Install",
      dragMouseUp: "Put file here~"
    },
    matter: {
      file: "File",
      directory: "Directory",
      rename: "Rename",
      download: "Download",
      delete: "Delete",
      more: "More",
      share: "Share",
      close: "Close",
      size: "Size",
      preview: "Preview",
      move: "Move",
      upload: "Upload",
      create: "Create",
      createTime: "Create Time",
      updateTime: "Update Time",
      root: "Root",
      fillInPicLink: "Fill in Picture Link",
      rePick: "Re Choose",
      chooseImage: "Choose Image",
      uploadMode: "Upload Mode",
      fillMode: "Fill Mode",
      sizeExceedLimit: "File size exceed limit {0}>{1}",
      setPublic: "Set as public",
      setPrivate: "Set as private",
      copyLink: "Copy Link",
      enterName: "Please enter name",
      publicFileEveryoneCanVisit: "Public file, anyone can access",
      fileDetail: "File info",
      expire: "Expire",
      copyLinkAndCode: "Copy link and code",
      uploaded: "Uploaded",
      speed: "Speed",
      fileInfo: "File basic info",
      fileName: "Filename",
      path: "Path",
      copyPath: "Copy path",
      publicOrPrivate: "Public or private",
      privateInfo: "Private file, only self or auth user can download",
      publicInfo: "Public file, anyone can download",
      downloadTimes: "Download times",
      operations: "Operation",
      oneTimeLink: "One time link",
      oneTimeLinkInfo: "One time link will expire after downloading, click to copy",
      imageCache: "Image cache",
      searchFile: "Search file",
      noContentYet: "No content under this directory yet",
      allFiles: "All Files",
      newDirectory: "New directory",
      notChoose: "Not choose any file",
      exceed1000: "Exceed file limit 1000"
    },
    router: {
      allFiles: "All Files",
      fileDetail: "File Detail",
      login: "Login",
      autoLogin: "Auto Login",
      register: "Register",
      users: "Users",
      userDetail: "User Detail",
      changePassword: "Change password",
      editUser: "Edit User",
      createUser: "Create User",
      shareDetail: "Share Detail",
      myShare: "My Share",
      dashboard: "Dashboard",
      install: "Install",
      setting: "Setting"
    },
    preference: {
      websiteName: "Website Name",
      logo: "Logo",
      logoSquare: "Logo will be cropped to square size",
      onlyAllowIco: "Only .ico allowed",
      copyright: "Copyright (support html)",
      extraInfo: "Extra info (support html)",
      zipMaxNumLimit: "Zip download max num limit",
      zipMaxSizeLimit: "Zip download max size limit(B)",
      current: "Current",
      noLimit: "No limit",
      userDefaultSizeLimit: "User default size limit(B) ",
      docLink: "Document Link",
      allowRegister: "Allow register",
      systemCleanup: "System Cleanup",
      systemCleanupDescription: "This operation will cleanup everything except administrators' data",
      systemCleanupPrompt: "This operation will cleanup everything except administrators' account data, please input login password.",
    },
    share: {
      shareDetail: "Share Detail",
      shareTime: "Share Time",
      expireTime: "Expire Time",
      noExpire: "Never Expire",
      expired: "Expired",
      copyLinkAndCode: "Copy Link And Code",
      shareSuccess: "Share Successfully",
      sharer: "Share Person",
      link: "Link",
      copyLink: "Copy Link",
      code: "Code",
      copyCode: "Copy Code",
      copySuccess: "Copy Successfully",
      more: "More",
      cancelShare: "Cancel Share",
      getLink: "Get Link",
      allFiles: "All Files",
      noContent: "No content in this directory",
      enterCode: "Please enter code",
      getFiles: "Get Files",
      codeError: "Code Error",
      cancelPrompt: "This operation will cancel sharing forever, continue?",
      hour: "1 Hour",
      day: "1 Day",
      week: "1 Week",
      month: "1 Month",
      year: "1 Year",
      infinity: "Forever",
    },
    user: {
      redirecting: "Redirecting...",
      oldPassword: "Old Password",
      newPassword: "New Password",
      confirmNewPassword: "Confirm New Password",
      cannotBeNull: "Cannot be null！",
      passwordNotSame: "Old and new password not same！",
      role: "Role",
      singleFileSizeLimit: "Single File Limit",
      totalFileSizeLimit: "Total Space Limit",
      current: "Current",
      noLimit: "No Limit",
      totalFileSize: "Used Space",
      status: "Status",
      lastLoginIp: "Last Login Ip",
      lastLoginTime: "Last Login Time",
      resetPassword: "Reset Password",
      transfiguration: "Transfiguration",
      changePassword: "Edit Password",
      enterPassword: "Enter Password",
      profile: "Profile",
      avatar: "Avatar",
      username: "Username",
      password: "Password",
      confirmPassword: "Confirm Password",
      disabled: "Disabled",
      disableUser: "Disable this user",
      activeUser: "Active this user",
      welcomeLogin: "Welcome Login",
      logining: "Login...",
      login: "Login",
      toToRegister: "Go To Register",
      welcomeRegister: "Welcome Register",
      registering: "Login...",
      register: "Register",
      goToLogin: "Go To Login",
      roleGuest: "Guest",
      roleUser: "User",
      roleAdministrator: "Administrator",
      statusActive: "Ok",
      statusDisabled: "Disabled",
      webdavLink: "WebDAV Link",
      docLink: "Document Link",
    },
    model: {
      usernameRule: "only lowercase letter and number and _ is permitted.",
      passwordRule: "Password must have more than 6 chars",
      linkCodeText: "Link:{0} Code:{1}",
      copyLinkCodeSuccess: "Copy Link and Code successfully",
      transfigurationPromptText: "Transfiguration Prompt",
      transfigurationPrompt: "You will login as this user.Please visit this link in other browser, if in current browser, you will logout.{0}",
    },
    plugin: {
      cannotPreview: "Cannot Preview",
      emptyHintDefault: "No Items",
      everyPage: "Every Page",
      items: "Items",
      total: "Total",
      clickRefresh: "Click To Refresh",
    },
    selectAll: "All",
    edit: "Edit",
    createTime: "Create Time",
    download: "Download",
    close: "Close",
    required: "Required",
    cancel: "Cancel",
    delete: "Delete",
    actionCanNotRevertConfirm: "This action cannot be reverted, confirm?",
    prompt: "Prompt",
    confirm: "Confirm",
    copy: "Copy",
    showMore: "Show More",
    username: "Username",
    password: "Password",
    submit: "Submit",
    save: "Save",
    create: "Create",
    finish: "Finish",
    operationSuccess: "Operation success",
    notFound: "404 Not Found",
    login: "Login",
    logout: "Logout",
    yes: "Yes",
    no: "No",
    all: "All",
  },
  zh: {
    oemTitle: "云盘",
    dashboard: {
      totalInvokeNum: '总PV',
      weekRate: '周环比',
      dayRate: '日环比',
      yesterdayInvoke: '昨日PV',
      totalUV: '总UV',
      yesterdayUV: '昨日UV',
      totalMatterNum: '总文件数',
      yesterdayMatterNum: '昨日文件数',
      totalFileSize: '文件总大小',
      yesterdayMatterSize: '昨日文件大小',
      recentDayInvokeUV: '最近{0}日PV/UV',
      downloadMatterTop10: '文件下载量TOP10',
      activeIpTop10: '活跃IP TOP10',
      loading: '加载中…',
      date: '日期',
      num: '数量'
    },
    install: {
      configMysql: "配置MySQL",
      port: "端口",
      schema: "库名",
      mysqlConnectionPass: "MySQL连接测试通过",
      testMysqlConnection: "测试MySQL连接",
      notice: "注意",
      mysqlNotice2: "数据库账户的权限要求要能够创建表，否则第二步\"创建表\"操作会出错",
      validateMysqlFirst: "请首先验证数据库连接",
      preStep: "上一步",
      nextStep: "下一步",
      createTable: "创建表",
      installed: "已安装",
      installedButMissing: "已安装,字段缺失",
      toBeInstalled: "待安装",
      allFields: "所有字段",
      missingFields: "缺失字段",
      tableNotice: "点击\"一键建表\"后会按照以下逻辑执行操作：",
      tableNotice1: "如果某表不存在，则直接创建表。",
      tableNotice2: "如果某表存在并且字段齐全，那么不会对该表做任何操作。",
      tableNotice3: "如果某表存在但是部分字段缺失，那么会在该表中增加缺失字段。",
      oneKeyCreate: "一键建表",
      createFinish: "建表完成",
      createTableSuccess: "建表成功",
      crateTableFirst: "请首先点击'一键建表'",
      setAdministrator: "设置管理员",
      detectAdministrator: "检测到系统中已经存在有以下管理员：",
      useOrCreateAdministrator: "你可以使用其中一位管理员的用户名和密码进行验证，或者创建一位新的管理员账户",
      validateAdministrator: "验证管理员账户",
      createAdministrator: "创建管理员账户",
      administratorUsername: "创建管理员账户",
      administratorPassword: "管理员密码",
      administratorRePassword: "再次输入密码",
      usernameRule: "由于用户名将作为文件上传的目录，因此只允许字母数字以及\"_\"。",
      congratulationInstall: "恭喜，安装成功！",
      configAdminFirst: "请首先配置管理员信息！",
      createAdminSuccess: "创建管理员成功！",
      validateAdminSuccess: "验证管理员成功！",
      enterHome: "点击进入首页"
    },
    layout: {
      allFiles: "所有文件",
      myShare: "我的分享",
      setting: "网站设置",
      dashboard: "监控统计",
      users: "用户列表",
      permission: "权限管理",
      projects: "空间列表",
      logout: "退出登录",
      about: "关于",
      install: "安装网站",
      dragMouseUp: "可以松手啦~"
    },
    matter: {
      file: "文件",
      directory: "文件夹",
      rename: "重命名",
      download: "下载",
      delete: "删除",
      more: "更多",
      share: "分享",
      close: "关闭",
      size: "大小",
      preview: "预览",
      move: "移动",
      upload: "上传",
      create: "新建",
      createTime: "创建日期",
      updateTime: "修改日期",
      root: "根目录",
      fillInPicLink: "请填写图片链接",
      rePick: "重新选择",
      chooseImage: "选择图片",
      uploadMode: "上传模式",
      fillMode: "填写模式",
      sizeExceedLimit: "文件大小超过了限制{0}>{1}",
      setPublic: "设置为公有文件",
      setPrivate: "设置为私有文件",
      copyLink: "复制下载链接",
      enterName: "请输入名称",
      publicFileEveryoneCanVisit: "公有文件，任何人可以访问",
      fileDetail: "文件详情",
      expire: "有效期",
      copyLinkAndCode: "复制链接+提取码",
      uploaded: "已上传",
      uploadDir: '上传文件夹',
      speed: "速度",
      fileInfo: "文件基本信息",
      fileName: "文件名",
      path: "路径",
      copyPath: "复制路径",
      publicOrPrivate: "文件公开性",
      privateInfo: "私有文件，只有自己或者授权的用户可以下载",
      publicInfo: "公有文件，任何人可以通过链接下载",
      downloadTimes: "下载次数",
      operations: "操作",
      oneTimeLink: "一次性链接",
      oneTimeLinkInfo: "使用一次性链接下载后链接立即失效,可以分享这个链接给朋友，点击复制",
      imageCache: "图片缓存",
      searchFile: "搜索文件",
      noContentYet: "该目录下暂无任何内容",
      allFiles: "全部文件",
      newDirectory: "新建文件夹",
      notChoose: "没有选择文件",
      exceed1000: "最多只能同时选取1000个文件"
    },
    router: {
      allFiles: "全部文件",
      fileDetail: "文件详情",
      login: "登录",
      autoLogin: "自动登录",
      register: "注册",
      users: "用户列表",
      permission: "权限管理",
      projects: "空间列表",
      userDetail: "用户详情",
      changePassword: "修改密码",
      editUser: "编辑用户",
      editProject: "编辑空间",
      createUser: "创建用户",
      createProject: "创建空间",
      shareDetail: "分享详情",
      myShare: "我的分享",
      dashboard: "监控统计",
      install: "安装网站",
      setting: "网站设置"
    },
    preference: {
      websiteName: "网站名称",
      logo: "Logo",
      logoSquare: "logo请使用正方形图片，否则在显示时会裁剪成正方形",
      onlyAllowIco: "只允许上传.ico图标",
      copyright: "版权信息(支持html)",
      extraInfo: "备案信息(支持html)",
      zipMaxNumLimit: "zip下载数量限制",
      zipMaxSizeLimit: "zip下载大小限制(B)",
      current: "当前值",
      noLimit: "无限制",
      userDefaultSizeLimit: "用户默认总大小限制(B) ",
      docLink: "文档链接",
      allowRegister: "允许自主注册",
      systemCleanup: "重置系统",
      systemCleanupDescription: "重置系统将清空除管理员账号外所有数据",
      systemCleanupPrompt: "重置系统将清空除管理员账号外所有数据，事关重大，请输入登录密码",
    },
    share: {
      shareDetail: "分享详情",
      shareTime: "分享时间",
      expireTime: "失效时间",
      noExpire: "永久有效",
      expired: "已过期",
      copyLinkAndCode: "复制链接+提取码",
      shareSuccess: "分享成功",
      sharer: "分享者",
      link: "链接",
      copyLink: "复制链接",
      code: "提取码",
      copyCode: "复制提取码",
      copySuccess: "复制成功",
      more: "更多",
      cancelShare: "取消分享",
      getLink: "获取链接",
      allFiles: "全部文件",
      noContent: "该目录下暂无任何内容",
      enterCode: "请输入提取码",
      getFiles: "提取文件",
      codeError: "提取码错误",
      cancelPrompt: "此操作将永久取消该分享, 是否继续?",
      hour: "1小时",
      day: "1天",
      week: "1周",
      month: "1个月",
      year: "1年",
      infinity: "永远有效",
    },
    user: {
      redirecting: "正在转跳...",
      oldPassword: "旧密码",
      newPassword: "新密码",
      confirmNewPassword: "确认新密码",
      cannotBeNull: "不能为空！",
      passwordNotSame: "两次输入不一致！",
      role: "角色",
      singleFileSizeLimit: "单文件限制",
      totalFileSizeLimit: "总空间限制",
      current: "当前值",
      noLimit: "无限制",
      totalFileSize: "已使用空间",
      status: "状态",
      lastLoginIp: "上次登录IP",
      lastLoginTime: "上次登录时间",
      resetPassword: "重置密码",
      transfiguration: "变身",
      changePassword: "修改密码",
      enterPassword: "输入新密码",
      profile: "个人详情",
      avatar: "头像",
      username: "用户名",
      password: "密码",
      confirmPassword: "确认密码",
      disabled: "已禁用",
      disableUser: "禁用该用户",
      activeUser: "激活该用户",
      welcomeLogin: "欢迎登录",
      logining: "正在登录...",
      login: "登录",
      toToRegister: "立即注册",
      welcomeRegister: "欢迎注册",
      registering: "正在登录...",
      register: "登录",
      goToLogin: "前往登录",
      roleGuest: "游客",
      roleUser: "普通用户",
      roleAdministrator: "系统管理员",
      statusActive: "正常",
      statusDisabled: "禁用",
      webdavLink: "WebDAV 地址",
      docLink: "文档链接",
    },
    project: {
      name: "空间名称",
      description: "空间描述",
      avatar: "头像",
      profile: "详情",
      userCount: "用户数量",
      directoryCount: "文件夹数量",
      totalFileSize: "已使用空间",
      searchProject: "搜索空间",
      deleteProject: "删除空间",
    },
    permission: {
      roleGuest: "访客",
      roleUser: "普通用户",
      roleProjectAdmin: "空间管理员",
      roleAdministrator: "系统管理员",
      selectTip: "请选择空间",
    },
    model: {
      usernameRule: "用户名只能包含字母，数字和\"_\"",
      passwordRule: "密码长度至少为6位",
      linkCodeText: "链接:{0} 提取码:{1}",
      copyLinkCodeSuccess: "复制链接提取码成功",
      transfigurationPromptText: "变身提示",
      transfigurationPrompt: "您将使用该用户的身份登录。请复制以下链接到其他浏览器访问，在当前浏览器访问会导致当前用户登录信息失效。{0}",
    },
    plugin: {
      cannotPreview: "无法预览",
      emptyHintDefault: "没有符合条件的空间",
      everyPage: "每页",
      items: "条",
      total: "共",
      clickRefresh: "点击刷新",
    },
    selectAll: "全选",
    edit: "修改",
    createTime: "创建时间",
    download: "下载",
    close: "关闭",
    required: "必填",
    cancel: "取消",
    delete: "删除",
    actionCanNotRevertConfirm: "此操作不可撤回, 是否继续?",
    prompt: "提示",
    confirm: "确定",
    copy: "复制",
    showMore: "显示更多",
    username: "用户名",
    password: "密码",
    submit: "提交",
    save: "保存",
    create: "创建",
    finish: "完成",
    operationSuccess: "操作成功",
    notFound: "404 页面找不到",
    login: "登录",
    logout: "退出",
    yes: "是",
    no: "否",
    all: "所有",

  }
}
export default i18nMessage
