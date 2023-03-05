# golang
```
go mod init com.xcrj.goleetcode
go mod tidy
go.mod用来标识一个模块，一个项目由1~n个模块组成，基于模块目录来引用包
一个目录下面只能有一个main方法
package是逻辑概念，不同于java必须物理上的目录结构
跨包访问：方法首字母大写 小写
目录结构参考：https://github.com/golang-standards/project-layout
cmd中存放main.go, package main, func main
internal中存放内部代码，其他模块不能import的代码
```