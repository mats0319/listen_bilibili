# listen_bilibili

使用B站作为音源的听歌工具。  
A tool for listening music, use Bilibili source.

## 使用

> 参考`build.ps1`

推荐环境：（项目开发期间使用以下环境）

1. node：18
2. go：1.21
3. 浏览器：chrome

编译：

1. 前端：下载依赖，构建代码
    - `cd ui`
    - `npm install && npm run build`
2. 生成后端可执行程序(.exe)
    - `./build/ps1`
    - 下载工具：`github.com/josephspurrier/goversioninfo`，为可执行程序绑定icon；注释脚本的`go generate`行可以跳过这一步

编译完成得到`listen`文件夹，包含以下内容：

- ui：文件夹，是编译好的前端文件
- listen_bilibili.exe：可执行文件，实际操作歌单文件
- list.yaml：歌单
- manual.md：本文档，用作使用手册
- log/backup：文件夹，使用过程中产生，记录程序日志/历史歌单

使用：

1. 启动exe程序，在windows系统中会自动打开网页，程序运行日志将同步记录在命令行和`log.log`文件
2. 如果无法自动播放，请在浏览器设置中授权
    - chrome浏览器允许使用js控制播放声音：设置-隐私和安全-网站设置-（更多内容设置）声音-添加网址`127.0.0.1`
3. 项目附带歌单仅供参考

## todo

- 前端引入eslint、prettier
- 前端实现功能：（统一使用modify list接口，考虑在每次执行修改后自动保存，思考新的备份规则）
    - 播放列表内歌曲排序（拟拖动排序），现在还只能通过修改list文件的方式改变歌曲排序
    - 修改歌曲信息，例如描述、bv地址等
    - 操作播放列表（新增、删除）
- 前端优化：pinia持久化，目前遇到的问题是使用组件会导致编译出错，错误代码ts2345，类型不匹配；考虑自己实现持久化功能，通过storage
