# Bili-Start-Live-Helper

众所周知，B站限制5000粉在Web端进行开播，强制用户使用直播姬进行开播

该工具可以让你绕开该限制，直接用其他推流软件进行推流

主要功能有

- GUI界面，对小白更友好
- 二维码登录，不用你去手动获取cookies
- 一键开播
- 更改直播标题和分区

## 编译

鉴于目前macos版本签名问题还未修复，可自行编译使用

编译环境需要：
- wails v2.10.0
- NodeJS v22.*
- pnpm v10

```shell
brew install wails
brew install node@22
npm i -g pnpm

git clone https://github.com/mewmewlab/bili-start-live-helper
cd bili-start-live-helper
wails build
```
或者下载Release版本安装后，打开终端执行`sudo xattr -r -d com.apple.quarantine bili-start-live-helper文件路径`命令绕过签名校验
