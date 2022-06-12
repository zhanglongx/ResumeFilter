# ResumeFilter

ResumeFilter is a resume GUI pre-screening program.

# Feature

- 自动解压压缩包内的文件，可识别多种格式（.rar，.zip等）。

- 自动遍历压缩包内的目录。

- 提供关键字匹配功能，目前为抽取学校信息。

- 支持跨平台。

- 无需安装额外的依赖库。

# Build

- Windows

```ps
    PS > go build -ldflags -H=windowsgui .
```

- Linux

```bash
    $ go build .
```

# Known Issues

- 如果压缩包内有相同文件名的文件（在不同目录下），会产生冲突。此时，只会列举其中一个文件。

- ⚠️*警告*：在关闭主窗口前，需要关闭所有已打开的PDF文件。