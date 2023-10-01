# listenBilibili

使用B站作为音源的听歌工具。  
A tool for listening music, use Bilibili source.

## 使用

> 参考`build.ps1`

推荐环境：（项目开发期间使用以下环境运行）

1. node：16
2. go：1.21
3. 浏览器：chrome

使用步骤：

1. 编译前端代码
2. 编译后端代码
3. 启动exe程序，在windows系统中会自动打开网页
4. 如果无法自动连播，请在浏览器设置中授权
    1. chrome浏览器允许使用js控制播放声音：设置-隐私和安全-网站设置-（更多内容设置）声音-添加网址

## todo

1. 编写http请求go、ts数据结构统一工具：实现根据go数据结构生成ts代码的工具。
    1. 使用grpc protocol buffer文件类型生成的go代码，难以应对**字段是数组**的情况，  
       因为工具生成的go数组使用指针形式（形如`List *[]Playlist`)，不能简单的使用json包反序列化
2. 实现操作歌单功能
3. 前端引入eslint、prettier
4. 代码优化
5. 记录日志到文件，包括启动阶段的错误日志和执行日志
6. 歌单文件做成配置，当前值调整为默认值
7. 导入时检查歌单文件：
    1. 全部music id不能相同
    2. 一个music list内，bv不能相同
