# ResumeFilter

ResumeFilter is a GUI resume pre-screening program.

## Features

- 自动解压压缩包内的文件，可识别多种格式（.rar，.zip等）。⚠️*警告*：zip方式只支持固定UTF-8编码。

- 自动遍历压缩包内的目录。

- 提供关键字匹配功能，目前为抽取学校信息。

- 支持跨平台。

- 无需安装额外的依赖库。

## Build

- Windows

```ps
    PS > build.bat
```

- Linux

```bash
    $ build.sh
```

## Usage

- Windows

```ps
    resumefilter.exe \<压缩包文件\>
```

💡*Tip*：可以[在SendTo中创建一个快捷方式](https://devblogs.microsoft.com/oldnewthing/20170403-00/?p=95885)，以方便操作。

- Linux

```bash
    $ resumefilter <压缩包文件>
```

## Known Issues

- ⚠️*警告*：在关闭主窗口前，需要关闭所有已打开的PDF文件。

- 某些pdf因格式原因，处理较慢。导致窗口在加载时也非常缓慢。

- 由于提取信息使用的是比较简单的正则表达式方式，会导致信息显示不全。

- 目前只支持pdf形式的简历。

- 如果压缩包内有相同文件名的文件（在不同目录下），会产生冲突。此时，只会列举其中一个文件。