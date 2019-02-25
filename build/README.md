# jenkins+k8s搭建简版发布系统

1. 安装jenkins

2. 配置jenkins
    必要插件
    gitlab凭证
    ssh

3. 部署git钩子

4. 发布流程

![代码提交流程](/build/shell/code_pub.png "代码提交流程")
![代码发布流程](/build/shell/k8s_pub.png "代码发布流程")

- 提交代码，触发git钩子（git_hook.sh）
- 钩子创建jenkins流水线（jenkins.php）,流水线定义镜像创建脚本（build.sh）
- wayne选择版本，发布代码

5. 目录介绍

    build
    ├── README.md
    ├── code        #代码目录
    ├── docker      #docker所需的配置项
    │   ├── nginx   #nginx镜像
    │   └── php     #php镜像
    └── shell
        ├── build.sh        #创建镜像的脚本
        ├── git_hook.sh     #git钩子
        ├── jenkins.php     #封装好的关于jenkins的操作
        ├── jenkins.sh
        ├── jenkins.xml     #jenkins发布流水线
        └── k8s_conn.go     #封装好的k8s的示例操作