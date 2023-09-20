# 根据go代码，生成http请求ts代码

命名：
 - 生成http请求ts代码：generate http ts code
 - http数据结构统一管理工具：http structure manage

根据go代码，生成本文件夹中的4种ts文件：
1. utils：暂时仅包括object转formdata一个函数，代码完全固定
2. config：axios config，代码基本固定，参数取值根据`config.go`文件的参数，没找到该文件则使用默认参数
3. xxx.pb.ts：根据对应的`xxx.go`生成，主要包含结构体（class）定义，是对应的关系，相对清晰
4. xxx.http.ts：根据对应的`xxx.go`生成，主要包含http请求方法，从形如`const URI_getList = "/list/get"`的uri开始，最为复杂
