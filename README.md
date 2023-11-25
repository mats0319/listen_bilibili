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
2. 生成后端可执行程序(.exe程序)
    - `./build/ps1`

编译完成得到`listen`文件夹，包含以下内容：

- ui：文件夹，是编译好的前端文件
- listen_bilibili.exe：可执行文件，实际操作歌单文件
- list.yaml：歌单
- readme.md：本文档，用作使用手册

使用：

1. 启动exe程序，在windows系统中会自动打开网页，程序运行日志将同步记录在命令行和`log.log`文件
2. 如果无法自动播放，请在浏览器设置中授权
    - chrome浏览器允许使用js控制播放声音：设置-隐私和安全-网站设置-（更多内容设置）声音-添加网址`127.0.0.1`

## todo

- 前端引入eslint、prettier
- 前端实现功能：播放列表内歌曲排序（拟拖动排序）
