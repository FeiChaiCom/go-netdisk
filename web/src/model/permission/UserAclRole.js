let UserRole = {
  GUEST: 'GUEST',
  USER: 'USER',
  ADMINISTRATOR: 'ADMINISTRATOR',
  PROJECT_ADMIN: 'PROJECT_ADMIN',
  PROJECT_PROVIDER: 'PROJECT_PROVIDER',
  GOD: 'GOD'
}

let UserRoleMap = {
  GUEST: {
    name: 'permission.roleGuest',
    value: 'GUEST',
    key: 'GUEST',
    style: "warning",
    icon: 'el-icon-user'
  },
  USER: {
    name: 'permission.roleUser',
    value: 'USER',
    key: 'USER',
    style: "primary",
    icon: 'el-icon-user'
  },
  ADMINISTRATOR: {
    name: 'permission.roleAdministrator',
    value: 'ADMINISTRATOR',
    key: 'ADMINISTRATOR',
    style: "success",
    icon: 'el-icon-s-custom'
  },
  PROJECT_ADMIN: {
    name: 'permission.roleProjectAdmin',
    value: 'PROJECT_ADMIN',
    key: 'PROJECT_ADMIN',
    style: "success",
    icon: 'el-icon-s-custom'
  }
}


let UserRoleList = [];
for (let key in UserRoleMap) {
  if (UserRoleMap.hasOwnProperty(key)) {
    UserRoleList.push(UserRoleMap[key]);
  }
}

export {UserRole, UserRoleMap, UserRoleList}
