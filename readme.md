### oxhorse(Wooden Ox and Gliding Horse)

木牛流马,拒绝在游戏里做牛马。

这是一款脚本开发工具，适用于windows x64系统，旨在使各种技术小白简单上手游戏脚本开发，解放双手，提高生产力。


其功能基于依赖如下：

### 1、驱动级键鼠操作。

​    基于https://github.com/oblitum/Interception封装开发。内置cgo。
   使用方法：

#### 1.1、项目引入，确保下面4个集成到你的项目中。
​        import (
​            "github.com/syzhang42/oxhorse/api"
​            _ "github.com/syzhang42/oxhorse/interception/include"
​            _ "github.com/syzhang42/oxhorse/interception/lib/x64"
​            _ "github.com/syzhang42/oxhorse/interception/installer"
​        )
    #### 1.2、驱动加载
​      加载：  /你的主目录路径/vendor/github.com/syzhang42/oxhorse/interception/installer/install-interception.exe install
​      卸载：  /你的主目录路径/vendor/github.com/syzhang42/oxhorse/interception/installer/install-interception.exe uninstall
    #### 1.3、环境变量配置
​        /你的主目录路径/vendor/github.com/syzhang42/oxhorse/interception/lib/x64  添加到系统环境变量中
​        或者将其子文件拷贝到你的配置了系统环境变量的目录中

#### 1.4、代码编写

#### 1.5、run 起来吧


​    
