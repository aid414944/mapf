mapf
====
mapf是一个go语言使用的类似map的容器，和map一样使用key-value的形式操作数据，不一样的是mapf可以方便的将数据以json文本的形式保存到本地磁盘中，也可以用json文件来初始化mapf，适合用来实现软件的配置文件模块。mapf有两种工作模式，分别是：

· MODE_AUTOWRITE
    该模式下每次调用Put()方法，mapf都会自动写入到文件中

· MODE_NONWRITE
    在此模式下需要用户调用Flush()方法才会将mapf写入到文件
    
可以使用SwitchMode()方法来切换工作模式。

函数
===
```go
func New(path string)(m *Mapf, e error)
```
