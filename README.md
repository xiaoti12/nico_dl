针对新版nico，自动下载普通/频道会限视频的程序（需自行解决网络问题）

基础使用：`.\nico_dl.exe https://www.nicovideo.jp/watch/sm000000000 `

目前最高支持720p。1080p视频存在鉴权问题，暂时无法直接下载

---

# 准备工作

## 保存cookies
### 获取cookies
1. **F12获取** 

    打开F12进行请求监控，打开nico网页，选择某个带有cookies的请求，复制所有字段（这里借用了黄师傅的图）
![image.png](https://s2.loli.net/2024/02/02/fOrUe51BluiNX8c.png)
2. **插件获取**

    通过`EditThisCookies`插件，在设置中选择导出方式，然后导出到剪贴板

![image.png](https://s2.loli.net/2024/02/02/CZiAgIMLPB4dmvq.png)

![image.png](https://s2.loli.net/2024/02/02/SJTrvaftKP2LY69.png)

### 保存为文件

在程序目录下新建`xxx.txt`的文件，将复制的cookies内容粘贴进文件。文件名不限，**保证为`txt`后缀**即可。当使用插件获取cookies时，需要将前面的注释内容（带`//`的内容）去掉

```
// Semicolon separated Cookie File
// This file was generated by EditThisCookie
// Details: http://www.iet
// Example: http://www.tutorialspoint.com/
nicosid=yyyy; _ss_pp_id=xxxx;
```

## 保存m3u8文件（可选）

使用猫抓，下载主m3u8文件，放到程序目录下。文件可重命名，**保证后缀为`m3u8`**。

![image.png](https://s2.loli.net/2024/02/02/O1ZE2xKjWzbltvp.png)

# 运行程序

## 通过视频链接下载

可以用两种方式提供下载链接，目前仅测试过一般视频（sm开头）和频道视频（so开头）

```
.\nico_dl.exe https://www.nicovideo.jp/watch/smXXXXXXX
.\nico_dl.exe smXXXXXXXX
```

运行完成后，在程序目录会生成下载好的`mp4`文件

## 通过M3U8文件下载

在目录下准备好`m3u8`后缀文件后，通过以下方式下载

```
.\nico_dl.exe --m3u8
```

# 环境与依赖

## 工具

程序额外使用了如下工具，一并放在了压缩包中：
- N_m3u8DL-CLI_v3.0.2.exe
- ffmpeg.exe

如果系统中已经配置了`ffmpeg`，可以删去程序目录下该文件。

## 运行环境

本程序由Golang编写并编译，理论上可以编译生成Linux和Mac等平台的二进制版本。但由于所依赖的`N_m3u8DL-CLI_v3.0.2.exe`只支持windows平台，所以目前只发布windows版本。

`N_m3u8DL-CLI`的作者已发布跨平台的`N_m3u8DL-RE`工具，后续如有需要可以迁移测试不同平台的运行。

# 后续优化

- [ ] 视频文件保存为实际名称
- [ ] 支持1080p视频下载
- [ ] 文件目录非法字符校验
- [ ] 重构代码（目前基本是用脚本的方式去写的，又不是不能用.jpg）









 

 