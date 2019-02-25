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

![目录说明](/build/shell/path.png "目录说明")
