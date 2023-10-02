# 根据go代码，生成http请求ts代码

命名：`httpc_ts`，参考protoc

前端将请求参数写成form data，后端通过`req.PostFormValue("music_id")`获取

根据go代码，生成本文件夹中的4种ts文件：

1. utils：暂时仅包括object转formdata一个函数，代码完全固定
2. config：axios config，代码基本固定，参数取值根据`config.go`文件的参数，没找到该文件则使用默认参数
3. xxx.go.ts：根据对应的`xxx.go`生成，主要包含结构体（class）定义，是对应的关系，相对清晰
4. xxx.http.ts：根据对应的`xxx.go`生成，主要包含http请求方法，从形如`const URI_getList = "/list/get"`的uri开始，相对最复杂

设计：

1. 在`config.go`里找axios配置参数，找不到的话使用默认值
2. utils根据是否使用到指定函数，动态改变其内的函数组合
    1. objectToFormData：当存在http请求，其req参数不为空时，需要该函数
3. 根据go文件生成http请求数据结构描述：
    1. 自己使用有限状态机实现一遍，按byte读取，不用正则了；画完状态机，数据结构就出来了
4. 根据生成的http请求数据结构描述，生成ts message数据
5. 根据生成的http请求数据结构描述，生成ts http请求函数数据：
    1. 在上一步完成之后，根据找到的`URI_xxx`，把名称对应的`xxxReq`/`xxxRes`编成一组
    2. 对于每一个请求（函数），其内容可能因为**是否需要req参数**而有所不同
    3. 至此，分析go代码步骤完成
6. 开始写入ts文件，包括config.ts/utils.ts(optional)/xxx.go.ts/xxx.http.ts 4大部分

数据结构：

```go
type API struct {
config apiConfig
utils apiUtils
message []message // 包含所有go文件中的message，message概念参考自protocol buffer里的message
service []service // 包含所有go文件中的`URI_xxx`,service概念参考自pb service
}

type apiConfig struct { // 初始化时，设置为默认值
baseURL string
timeout int64 // unit: micro-second
}

type apiUtils struct {
needObjectToFormData bool // 如果添加工具函数，这个结构中可能有很多布尔类型变量，当这些布尔变量中，至少有一个为true时，写utils.ts文件
objectToFormData []byte   // function code, const string
}

type message struct {
filename string // go file name, use to generate ts file(s)，生成ts文件时，根据这个name，然后分别写入不同的ts文件
name string     // message name，与go文件中的type xxx struct一一对应
fields []field
}

// 针对可能存在的自定义类型，field type正常记录，生成ts代码时，会通过一个默认类型转换表来检查一个类型是不是自定义类型
type field struct {
name string // field name, get from tag 'json'
fieldType string 
}

type service struct {
filename string // 同message.filename
// service name，会通过名字严格匹配，例如URI_GetList意味着一个http请求，它的message是GetListReq和GetListRes
// 一个http请求可以没有req，所以req message可以没有
// 见apiUtils.needObjectToFormData，如果一个service有req message，则需要该工具函数（设置其值为true）
name string
url string // http req url
}

```

解析go文件，有限状态机：

基础规则：

1. 换行符重置状态，如果在type状态内，换行符重置field状态，即仅重置一层状态、不会重置全部
2. 检测到行注释(`//`)时，忽略后续内容，直到换行符——需要一个注释状态
3. 大部分状态读到空格时结算

状态分类：

1. 解析service状态：（const）
    1. 正常的service是一行代码，形如`const URI_GetList = "/list/get"`
    2. is service：空白状态下，读取到连续的`const`，视为进入**解析service状态**，读到空格时结算
    3. parse name：解析service状态下，读取到常量名，提取出service name，此时标记为**解析service name状态**，读到空格时结算
    4. parse url：解析service name状态下，读取到双引号括起来的部分，记为service url，读到空格时结算
    5. 读到换行符时，退出到空白状态，记录service
2. 解析message状态：（type）
    1. 正常的message形如：
       ```go
       type Music struct {
         ID          string `json:"id" yaml:"id"`
         Name        string `json:"name" yaml:"name"`
         Bv          string `json:"bv" yaml:"bv"`
         Description string `json:"description" yaml:"description"`
         Volume      int32  `json:"volume" yaml:"volume"`
       }

       ```
    2. is message：空白状态下，读取到连续的`type`，视为进入**解析message状态**
    3. parse name：解析message状态下，读取到结构体名，此时标记为**解析message name状态**，读到空格时结算
    4. 读到换行符，退出到解析message状态
    5. is field：解析message状态下，读到字段名，此时标记为**解析field name状态**，读到空格时结算
    6. is filed type：解析field name状态下，读到字段类型，此时标记为**解析field type状态**，读到空格时结算
    7. is filed tag：解析field type状态下，读到反引号(\`)，提取json tag，读到下一个反引号时结算
    8. 读到换行符，退出到解析message状态，记录field
    9. 只有读取到`}`字符、且处于**解析message状态**时，退出到空白状态，如果一个文件读完了也没有退出到空白状态，终止程序，报错
3. 行注释规则：忽略直到换行符

其实我这个状态机和按行读文件差不多了，重点状态的转换都放在换行符上
