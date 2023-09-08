# listenBilibili

使用B站作为音源的听歌工具。
A tool for listening music, use Bilibili source.

## 文档整理(draft)

计划将文档拆分成*使用手册*和*开发过程*两部分，因为手册需要程序基本完成才能开始写，所以可以先写在一起，到时候再整理和拆分。

## 为什么想做这样一个工具

随着一些标志性事件的发生，国内版权意识逐渐“觉醒”，但这个觉醒的版权可能和我的理解不太一样。

在我朴素的观念里，版权是保障原创作者权益的约定，比如代码库里的license。
但国内的版权，已经变成了保护获益人/资本的借口，举几个例子：

1. 《神探狄仁杰》核心制作人钱雁秋，因为所谓版权问题，在第4部作品中无法使用《神探狄仁杰4》作为名字，只能改用《神断狄仁杰》
2. 知名网络小说网站起点，签约合同中将小说作者创作内容的版权据为己有，使创作者变成了自带创意的枪手
3. 知名网络小说网站起点，在网站订阅作品，实际上只获得了15天的阅读权，15天之后，网站甚至可以随意删除你订阅的内容
4. 某知名主播因直播期间演唱某歌曲，被版权方索赔10万元，从此事件中还诞生了一批经典语句，如“索要的赔偿金额其实不重要，我们主要是希望唤起大家保护版权的意识”等

综上所述，“版权”原本神圣的光环，在我眼中逐渐褪色，我也不再重视任何我无法从中获益的版权问题。  
当然了，以上内容也可以解释成我为自己心安理得地看盗版小说、听无版权音乐找的借口，我也认可这种思路。

又因为我在一些场景中确实有听音乐的需求——比如做家务的时候，而现有音乐播放器或多或少都不符合我的情况：

1. 网易云、酷狗等播放器因为版权问题，部分歌曲直接没有，即使今天有的歌曲，可能明天就没了
2. [listen1](https://github.com/listen1/listen1_desktop)，我在使用期间经常遇到歌曲消失的情况：一首4分钟的歌，直接变成15s的“……请到酷我音乐手机端播放……”
3. B站，在一个100首歌的合集里找到一首自己不喜欢的歌可太简单了，B站听歌的问题在于up主上传什么、我就得听什么，没有办法跳过一首歌、修改播放顺序等

所以我产生了自己写一个听歌工具的想法，它应该有这些功能：

1. 能根据B站链接播放音频
2. 能显示歌词（可能会自己写歌词文件类型，或者使用.lrc）
3. 能操作歌单（计划配一个web页面，方便操作）

## 技术准备

我的核心需求是播放B站视频的音频部分，我准备先看看有没有现成的方法，如果没有，就根据B站视频的格式，找相关的播放方法。  
歌词准备自己写一个定时打印功能，一方面是我从网络上下载的.lrc文件，它的时间标注没有统一标准，有的到毫秒，有的只到秒；另一方面是功能不难，自己写更好修改。  
操作歌单就是比较进阶的需求了，可能会做一个网页吧，后面再说。

### 前期调研

网络上有一些下载视频或音频到本地的工具，但我不想下载，想在线播放。  
我找到了一种获取视频源的办法：模拟移动端调用，浏览器控制台-media，可以获取到可用的视频文件地址，可以在potplayer里直接播放。  
因为B站终归是视频网站，它没有道理把视频和音频拆开播放，那么我最开始设想的在线播放音频就不合适了。

为了方便描述，我们约定：

1. bv地址：
    - 形如`https://m.bilibili.com/video/BV1nk4y1P79g?p=22` 的链接
    - B站视频的公开地址
2. 源地址：
    - 形如
   ```text
   https://cn-hk-eq-01-14.bilivideo.com/upgcxcode/11/44/1190394411/1190394411-1-16.mp4
   ?e=ig8euxZM2rNcNbRVhwdVhwdlhWdVhwdVhoNvNC8BqJIzNbfq9rVEuxTEnE8L5F6VnEsSTx0vkX8fqJeYTj_lta53NCM=
   &uipk=5
   &nbs=1
   &deadline=1692035857
   &gen=playurlv2
   &os=bcache&oi=983052370
   &trid=00007c1bef8f8f0e424eb013d777b4aebf60h&mid=507453531
   &platform=html5
   &upsig=27967b89d0fa836e94ef28d0d279f9c4
   &uparams=e,uipk,nbs,deadline,gen,os,oi,trid,mid,platform
   &cdnid=68708&bvc=vod
   &nettype=0
   &bw=53263
   &logo=80000000
   ```
   的链接
    - B站视频源地址，可以直接通过potplayer等第三方播放器使用的地址

根据目前情况，整理解决方案流程：

1. 人工观看B站视频，找到一些好听的音源（这一步得到bv地址）
2. 获取视频源地址（模拟移动端调用上一步获得的地址，得到源地址）  
   可以看到源地址有一段`deadline`，像是请求时间+2h，并且经过尝试，只复制链接到.mp4是无法获取视频的  
   或许其他参数也有类似功能，所以我们需要在每次播放的时候都从bv地址请求一遍源地址
3. 播放视频源地址，目前有几个想法：调起第三方播放器、打开一个网页播放，或者可以找找有没有go写的mp4播放器  
   暂时计划打开一个网页播放视频，这样可以不用依赖第三方播放器，也因为我们本就计划写一个网页；  
   而且写网页还可以直接在网页上模拟移动端调用，go甚至可以只保存和传递歌单文件，剩下功能都在前端实现

### 编写demo

1. 编写go获取源地址demo，用go是因为我更熟悉
2. 编写网页播放视频demo

---

写完了demo，和预期一致，使用go获取源地址，再复制到前端，前端可以播放对应视频
现在整个解决方案已经确定可行，剩下就是把功能串联起来，然后优化代码了

## 代码编写

详细描述程序流程：

1. go从指定路径获取歌单文件，反序列化，保存在程序内存
    1. 歌单文件里每一首歌应该包括：id、歌曲名、bv地址、声音（用于设置视频声音，当有两个视频，一个声音很大、一个声音很小的时候，这个值就有用了）
    2. 歌单里保存的声音，实际上是一个偏移量；前端会有一个可设置的统一音量（localstorage），  
       真正设置视频声音时，是在前端统一音量的基础上偏移每首歌的**声音**值;  
       考虑到电脑设置和人的差异，可能我在我的电脑上用30%统一音量正好，而你就需要40%，所以我们这个值在前端处理
    3. 一个歌单里可能有很多个播放列表
2. 前端请求歌单
    1. 这里指的是歌单中全部歌曲的**id**和**歌曲名**，不包含地址
    2. 前端拿到歌曲列表就可以排序了，播放模式由前端实现
3. 前端根据歌曲id请求视频源地址
4. 后端根据bv地址获取视频源地址，向前端返回**视频源地址**和**声音**
5. 前端得到视频源地址，播放视频；根据声音设置视频音量
    1. 前端注册视频播放结束事件：事件中继续请求下一个视频源地址

第一版的歌单文件只有一个播放列表、前端也只允许顺序播放，但前端应考虑到多列表播放、播放顺序等问题。  
前端或许还可以有一个歌单修改页面，以期在界面上操作歌单，例如增删音乐、修改默认顺序等。  
为什么不能一起请求完所有视频源地址然后丢给前端？因为通过bv地址获取的源地址是存在有效期限制的，bv地址实际上对应一个html文件，我们要在其中找到源地址。

数据结构设计：

```go 
struct Music {
id
name
bv
volume
}

struct MusicBrief {
id
name
}

// 相应地，如果歌曲使用的是`MusicBrief`，那么播放列表的名字就叫`PlaylistBrief`、歌单的名字就叫`ListBrief`
struct Playlist {
id
name
Music[]
}

struct List {
Playlist[]
}

```

接口设计：

```text
1. getListBrief() => ListBrief
2. getOriginAddress(music_id) => Music
3. [next verison] getList() => List
4. [next version] modifyList(List) => nil，go要将前端传过来的List覆盖写入歌单文件，同时更新内存
```

下一步就是正式写代码了

---

编程过程中的思考：

1. 取消brief结构，因为功能简单，如果搞出完整版和简要版两套数据结构反而会显得代码很复杂
2. 统一管理http请求的go、ts数据结构，让它们都根据protocol buffer生成
    1. 我们实际上只用到了pb文件类型的序列化和反序列化，而没有用到grpc
    2. 其中pb生成go代码使用的官方工具，pb生成ts代码的工具是自己写的，在我的unnamed_plan库
    3. 生成的go代码无法简单通过json序列化、反序列化（因为有嵌套的结构体指针数组这种复杂类型），所以go数据结构部分自己写，
       但这样又导致了http请求的数据结构不是统一的，所以这部分应该怎么做呢？假设我在前后端之间使用http请求，在后端程序之间使用grpc。
       后端之间没啥好说的，就是使用pb对go做序列化和反序列化；主要是http请求，有以下几个问题：
        1. http请求是否要和grpc请求使用一套数据结构：从理论上讲是要分开的，但是在实际操作中，所谓的分开，基本上是相同代码copy了一遍
           因为从前端验证账号密码和从后端服务验证账号密码，其请求参数是一样的
           但还是分开好一些，从概念上做出区分，不容易造成混淆
2. 浏览器允许使用js控制播放声音：设置-隐私和安全-（更多内容设置）声音-添加网址

todo:

1. 计划在初始化过程中，检查歌单是否存在重复歌曲，包括单个播放列表（warning）和整个歌单（info）
2. 编写http请求go、ts数据结构统一工具：根据go代码生成ts代码，参考现在的根据pb代码生成ts代码工具
