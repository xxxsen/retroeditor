Retroeditor
=== 

一个用于管理游戏的小工具

## 使用

### 限制

1. 目前只能在linux下使用, mac/windows都没测试过。
2. 工具对于单文件的游戏能很好的进行管理, 例如nes/snes这些, 但是对于以目录形式提供的游戏目前处理起来还是有各种问题, 例如scummvm, dos 这些。

**NOTE:**
- **不要对多文件游戏使用重建Gamelist.xml功能, 不然容易搞炸。**
- **虽然在自己600+G的游戏上测试各种能力都没什么问题, 但难免会有小bug, 如果有问题可以提issue**


### 具体例子

```shell
${bin} --root=${你的roms目录}
```

例如:

```shell
./retroeditor --root=/home/test/roms
```

## 编译

### 依赖包安装

- qt5: sudo apt-get install qt5-default
- go: sudo apt-get install golang

### 编译

进入到源码目录后, 运行 `go build -o retroeditor` 即可