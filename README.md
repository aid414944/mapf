mapf
==========
mapf是一个go语言使用的类似map的容器，和map一样使用key-value的形式操作数据，不一样的是mapf可以方便的将数据以json文本的形式保存到本地磁盘中，也可以用json文件来初始化mapf，适合用来实现软件的配置文件模块。mapf有两种工作模式，分别是：

· MODE_AUTOWRITE
    该模式下每次调用Put()方法，mapf都会自动写入到文件中

· MODE_NONWRITE
    在此模式下需要用户调用Flush()方法才会将mapf写入到文件
    
默认MODE_AUTOWRITE模式，可以使用SwitchMode()方法切换工作模式。

函数
----------
```golang
func New(path string)(m *Mapf, e error)
```
    ### 新建一个Mapf，并以path参数指定的json文件初始化它。成功则返回新建Mapf的指针，失败则返回nil，具体失败原因可查看e。

方法
----------
```golang
func (m *Mapf)Flush()(e error)
```
    ### 将Mapf写入磁盘，如果成功e为nil。
    
```golang
func (m *Mapf)SwitchMode(mode int)
```    
    ### 切换工作模式。mode可选值有：
        ### · MODE_AUTOWRITE
            ###该模式下每次调用Put()方法，mapf都会自动写入到文件中。
        ### · MODE_NONWRITE
            ###在此模式下需要用户调用Flush()方法才会将mapf写入到文件。
            
```golang
func (m *Mapf)Put(k string, v interface{}) error
```    
    ### 压入一个key-value，key的类型必须是string。
    
```golang    
func (m *Mapf)Get(k string)(v interface{}, ok bool)
```
    ### 根据key返回一个value，如果key不存在ok将为false
