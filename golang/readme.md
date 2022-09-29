[toc]

# 基本语法

## Chapter1

go特性：
开源
静态类型，编译型语言：强类型、语法简洁。
跨平台：编译成多平台可执行文件。
自动垃圾回收：编程不需要考虑释放内存。
并发编程：原生的并发模型（goroutine+channel）。
完善的构建工具：很多命令工具
多编程范式：函数式编程、面向对象编程（不太一样，也有接口和实现的概念，不过用嵌入代替继承）
统一的代码风格：
高效编程和运行：
丰富的标准库：标准库很多开箱即用的API，编写系统级程序、Web程序、分布式程序很方便。

go目录：

- pkg：平台相关的目录，归档，go install的库，将编译成平台相关的归档文件。
- go：放go本身、标准工具、标准库的源码文件
- test：go测试和验证相关的文件

go推荐将go目录放到 /usr/local 下面

设置环境变量 GOROOT，环境变量增加 $GOROOT/bin，可以使用 go 的工具命令

Go推崇软件工程理念，高度强调代码和项目的规范统一，工程结构比较严谨和完善

工程结构：

工作区就是特定工程结构的目录：
src：以代码包的形式组织和保存源码。
pkg：go install安装的包的代码归档文件（扩展名.a），区别于GOROOT/pkg，因为存的是用户代码
bin：可执行文件目录

命令源码文件和库源码文件的区别：
命令源码文件属于main包，可独立启动运行（go run），属于程序入口。
库源码文件是指某个代码包中的普通源码文件。

测试源码文件：_test.go结尾，功能测试和基准测试。


GOPATH是工作区的目录路径，可以有多个。

go get会下载到工作区中。

go文件需要utf-8编码。

代码包是go中编译和安装的基本单元。

声明、导入、指定命名，. _
包初始化函数 init()


初始化顺序
全局变量 > 包初始化函数 => main函数执行

go标准命令和工具

## Chapter2

基本构成

标识符（identifier）
关键字（keyword）
字面量（literal）
分隔符（delimiter）
操作符（operator）
表达式

预定义标识符
	内置类型/基本数据类型+接口类型，error
	内置函数
	常量：true、false、iota、nil

所有内建函数的名称、空标识符（_）一般用到导入包和变量接收语句（绕过编译器检查）

关键字/保留字

25个关键字，go chan select 三个个并发编程有关。

type 用于类型声明

任何类型都是空接口(interface{})的实现类型

字面量

符合数据类型的结构体，定义变量

表达式

选择表达式
断言表达式
切片表达式
调用表达式
索引表达式

断言失败引起运行时异常

v1.(int)

解决方式：

var res,err := interface{}(v1).(int)

基本类型

1. 简单类型
byte(uint8) rune(int32) int bool float nil

2. string
英文，每个字符占 1 byte，和 ASCII 编码是一样的，非常节省空间，如果是中文，一般占3字节。

字面量：字符串的长度（底层字节序列的字节数）在编译期间就能确定，有两种：
原生字符串字面量（反引号`），所见即所得。
解释型字符串字面量（”），会解析转义字符。

高级类型

数组（array）：len、cap
[5]int{1,2,3,4,5}
[…]int{1,2,3,4,5}
[5]int{}

切片（slice）：底层数组结构，指针、长度、容量
[]int{1,2,3,4,5}
[]int{}
make([]int, 100)
append()

字典（map）/散列表/哈希表，实现了关联数组的数据结构。
var a = map[int]string{}
make([int]string)
delete(a, 2)

变量与内置数据类型


流程控制


函数


结构体

延迟函数 延迟调用 defer 

recover()
panic()

变量类型，并发安全？

# 模块（module）和包（package）

一个模块是一些Go package的集合。

module：
package：

导入第三方包
require go.uber.org/zap v1.23.0

导入本地包（测试用？）

1. 第一种

```
/example
     /mymath
	  /mymathtest.go
	  /go.mod
```

go.mod
require mymath v0.0.0
replace mymath => ./mymath

import “mymath”
mymath.Calc()

2. 第二种

```
/example
     /mymath2
	  /mymathtest2.go
```

import “example/mymath2”

mymath2.Calc()

# 单元测试（unit test）

测试文件和源代码文件放在一起。

```
mypackage
    file.go
    file_test.go
    file2.go
    file2_test.go
```

使用标准库：testing
单元测试的参数是 *testing.T
基准测试(Benchmark)的参数是 *testing.B
TestMain 的参数是 *testing.M 类型

运行go test 自动运行当前 package 下的所有测试用例。
-v 查看详细信息，显示每个用例的测试结果。
-cover 参数查看覆盖率。
-run 执行指定文件

go test -run TestSubtest -v
go test -run TestSubtest/aa -v

测试中使用的是 t.Error/t.Errorf
子测试使用的是 t.Fatal/t.Fatalf
区别在于前者遇错不停，还会继续执行其他的测试用例，后者遇错即停。

使用子测试好处：
新建测试用例简单
测试代码可读性好，能直观地看到每个子测试的参数和期待的返回值。
用例失败时，报错信息的格式比较统一，测试报告易于阅读。

# 反射（reflect）

获取目标对象信息：
reflect.TypeOf()
reflect.ValueOf()

利用反射修改对象，必须传指针类型变量。
使用反射动态调用对象。

# 错误和异常处理（error handling）

error
错误处理是业务的一部分，需要我们在程序中处理。
比如：连接数据库失败，连接网络失败
Go内建了一个error接口类型作为go的错误标准处理。

exception
通常是预料之外的问题。
比如：空指针引用，下标越界，向空map添加键值等。

人为制造被自动触发的异常，比如：数组越界，向空map添加键值对等。 手工触发异常并终止异常，比如：连接数据库失败主动panic。 
panic（恐慌）
对于真正意外的情况，那些表示不可恢复的程序错误，不可恢复才使用panic。
对于其他的错误情况，我们应该是期望使用error来进行判定。 
理论上panic只存在于server启动阶段，比如config文件解析失败，端口监听失败等等，所有业务逻辑禁止主动panic，所有异步的goroutine都要用recover去兜底处理。

Go处理错误的三种方式：
case1：直接输出错误，如果业务逻辑不是很清楚的场景。
case2：先暂存错误到结构体内，等业务处理完成，统一处理错误。适合代码很少去改动的场景，类似标准库。
case3：函数式编程延迟运行。适合比较复杂的场景，复杂到抽象成一种设计模式。


错误要被日志记录；
应用程序处理错误，保证100%完整性；
之后不再报告当前错误（错误只被处理一次）。 
最佳实践：https://mp.weixin.qq.com/s/o4k9Bu1X6KTK8Mvv9ufJPQ

# 并发编程（goroutine）

执行体/协程（goroutinue）

两种并发编程风格：
channel
共享内存多线程：

￼

启动一个协程：
go func(){
	//…
}()

并发协程：

1. 多个并发协程之间不需要通信的场景，使用 sync.WaitGroup，等待所有并发协程执行结束。
场景：并发下载 N 个资源。

var wg sync.WaitGroup
wg.Add(1)
wg.Done()
wg.Wait()

time go run sync.go

2. 需要通信的场景：channel（信道/管道）
使用 channel ，可以在goroutinue之间传递消息。阻塞等待goroutinue返回消息。

数据流动：<- 这个符号，箭头指向谁，数据就往那个方向流动。
数据接收方<- 数据发送方
数据接收方 := <-数据发送方

创建信道：
ch := make(chan int)

定义信道缓冲大小为10：
var ch = make(chan string, 10)

关闭信道：
close(ch)

写入数据到缓冲信道
ch <- url

从缓冲信道中取出数据
msg := <-ch


main函数所在的goroutine叫做 main goroutine，如果这个主goroutine退出了，那么整个Go程序就会跟着一起退出。

channel 有三种形式：
双向：make(chan int)
只读：make(<-chan int)
只写：make(chan<- int)

读取channel中的数据是会被阻塞住的。

3. 用select解决channel阻塞

可处理一个或多个 channel 的发送与接收
同时有多个可用的 channel 时按随机顺序处理
可用空的 select 来阻塞 main 函数
可设置超时

4. 锁（lock）

互斥（mutex）
sync.Mutex
锁的作用就是，把数据锁定，持有锁的人才能执行代码，其他人都是等待。

sync.RWMutex，读写锁和 Mutex 差不多，多了两个方法：
RLock()
RUnlock()
写操作之间会互相阻塞，写操作和读操作之间会互相阻塞，但是读操作之间不会互相阻塞。

sync/atomic
这个库提供了CPU层级的原子操作实现

# IDE

VS Code

变量显示
调用堆栈显示

开启调试/继续(F5)：
调试下一步(F10)：逐过程执行
单步跳入(F11)：逐语句执行
单步跳出(Shift F11)：
重新调试(Ctrl + Shift + F5)：
结束调试(Shift + F5)：

command+k+0 折叠代码
command+k+j 展开代码

# 问题

go run main.go 

报错：
fork/exec /var/folders/c_/p7dbmlgd6vxd6lrpph41fp1c0000gn/T/go-build2845153107/b001/exe/main: exec format error

解决：go env 打印环境变量发现GOOS不是macOS，修改一下解决了。
go env -w GOOS=darwin