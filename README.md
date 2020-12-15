# goStudy01
本项目为go学习项目 <br>
最终利用gin和orm框架生成一个简单的web项目示例
# 国内代理问题导致无法正常下载依赖使用下面的命令
 go env -w GOPROXY=https://goproxy.cn
# 本项目所用主要框架或库
    1.gin
    2.gorm
    3.zap
    4.mysql
    5.yaml
# 目录简介
    api:
        .后台接口及相关业务的.go文件
    cache：
        模拟缓存，缓存token令牌
    code:
        简易且只适合本项目使用的利用模板生成.go .html .js 等文件工具
        生成的代码需在router中添加路由才能正常访问，路由.go文件会在生成的相关模块的router目录下
        serviceGo.go: 主要逻辑代码
        serviceGo_test.go:工具入口，具体查看文件内示例
        tmpl目录：模板文件所在目录
    config:
        加载读取yaml配置文件，对外调用使用c.go文件，请勿直接使用config.go文件
    db:
        顾名思义，数据库的
    log:
        zap相关配置
    router:
        gin相关包含端口、路由等相关设置，也是本项目全局路由所在
        跨域相关
        接口相关限制，doc见restrict.go内
    static:
        前端静态资源目录，css、js等
    util:
        工具类，包括文件上传、md5加密工具类等
    views:
        html页面
       
