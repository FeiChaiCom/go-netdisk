# go-netdisk（基于Gin的简易网盘）


## 一、产品介绍

&nbsp;&nbsp;&nbsp;&nbsp;本项目基于Gin实现了一个简易的`云盘系统`，前端采用`Vue`实现具体业务逻辑，包括：用户登录、用户管理、文件夹管理、
文件管理（上传、下载、查看、删除、批量操作）等功能，可用于扩展开发`网盘`、`企业云盘`等面向个人和企业等不同用户的云存储平台.

### 1.体验环境

地址：http://netdisk.bksaas.com

普通用户：`user0/miya.12345`       
系统管理员：`admin/miya.12345`

### 2.产品截图

> 登录
![image](https://user-images.githubusercontent.com/84143485/119234765-177a1d00-bb62-11eb-889b-4f87988d1023.png)

> 登陆成功
![image](https://user-images.githubusercontent.com/84143485/119234954-ccacd500-bb62-11eb-868c-0dce25124678.png)

> 上传文件
![image](https://user-images.githubusercontent.com/84143485/119235171-cc610980-bb63-11eb-9937-e43794bd356d.png)

> 新建文件夹
![image](https://user-images.githubusercontent.com/84143485/119235208-eef32280-bb63-11eb-8b4f-96d3d92d18b8.png)

> 用户管理
![image](https://user-images.githubusercontent.com/84143485/119235261-2bbf1980-bb64-11eb-80dd-8313c99217b7.png)

> 手机端操作   
![image](https://user-images.githubusercontent.com/84143485/119238695-8e211580-bb76-11eb-9c29-6c6829e1a11f.png)

##  二、代码结构
```
├── Dockerfile
├── build.sh                # 镜像打包脚本
├── config                  # 全局配置
├── gin-contrib             # session管理
│   └── sessions
├── go.mod                  # 依赖管理
├── go.sum
├── init.go                 # 站点数据初始化
├── k8s                     # k8s部署配置
│   ├── deploy
│   └── go
├── main.go                 # 总入口
├── middleware              # gin中间件
├── models                  # gorm模型定义
│   ├── db
│   └── form
├── render
├── services                # 服务层接口
│   ├── demo
│   ├── login         # 登录管理  
│   ├── main.go       # 路由注册入口
│   ├── matter        # 文件管理
│   ├── monitor       # 服务健康监控  
│   ├── permission    # 用户权限管理
│   ├── preference    # 站点偏好
│   └── user          # 用户管理
├── static                  # 静态资源目录，前端打包输出css/js文件到该目录        
└── templates               # 静态资源目录，前端打包输出html文件到该目录
├── utils
└── web                     # 前端工程
```

##  三、本地开发

### 1.设置环境变量

无

### 2.启动后台


首先，需要安装go依赖环境（`go>=1.16.x`）：

```
go mod tidy
go mod download
```


> 站点初始化信息

普通用户：`user0/miya.12345`  
系统管理员：`admin/miya.12345`


最后启动后台：

```
air -c .air.toml
```

### 3.启动前端

```
cd web
npm install
npm run serve
```

启动成功后，访问地址：http://127.0.0.1:8080

端口可有根据需要修改`vue.config.js`中的设置，默认前端绑定`8080`，后端绑定`8000`

```
  devServer: {
    https: false,
    port: 8080,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8000',
        changeOrigin: true,
        pathRewrite: {
          '^/api': '/api'
        }
      }
    }
  },

```

##  四、线上部署

### 1.环境变量

无

### 2.前端打包

```
npm run build
```

打包完成后，会生成一份静态资源文件到 `static` 和 `templates` 目录下


### 3.启动命令


```
./server
```



> 站点初始化信息

普通用户：`user0/miya.12345`  
系统管理员：`admin/miya.12345`

最后启动后台服务：

```
./start.sh
```

说明：生产环境建议结合`nginx`来部署本云盘项目

## License

[MIT](LICENSE)
