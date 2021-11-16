# Golang 基础



## Slice

### 声明方式

切片：

`s := make([]int, len, cap)`

数组：

`var a [len]int`

### slice扩容

当cap < 1024，每次x2

当cap >= 1024，每次x1.25

### 使用注意

- array长度为类型组成部分，不可把一个长度100的数组传参给一个参数类型为长度20的数组的方法

- array作为方法参数时会产生copy

- golang所有方法参数都是值传递

- append slice的时候，如果能预知容量，提前分配cap然后使用index赋值是最快方式

- 把slice传参进方法，然后修改slice的某个元素值，虽然是值传递，但是slice传参传的是指针，所以也会作用在原slice上

- slice如果没有发生扩容，修改会在原来的内存中；如果发生了扩容，修改会在新的内存中

  ```go
  func change(s []int) {
     s[0] = 1024
     s = append(s, 4)
  }
  
  func main() {
     s := []int{1, 2, 3}
  
     change(s)
  
     for _, a := range s {
        print(a)
     }
     // 输出结果是1024,2,3
  ```

  这是因为是值传递，传进去的slice扩容了，但是外面的slice没有，所以原slice还是到3为止，因为s[0]赋值为1024，所以是1，2，3

- 如果能确定slice的长度，可以先执行一次，让编译器去做下优化，例如：

  ```go
  func readSlice(s []int)  {
     _ = s[2]
  
     print(s[0])
     print(s[1])
     print(s[2])
  }
  ```

  比

  ```go
  func readSlice(s []int)  {
     print(s[0])
     print(s[1])
     print(s[2])
  }
  ```

  更快

  这种方式因为在生成的汇编代码中没有做数组越界校验的环节，所以速度较快，这种方式叫做Bounds Checking Elimination

## Map



### hash算法

- 程序启动时，会检测CPU是否支持AES，如果支持，直接使用AES进行hash（AES指令集）；如果不支持，使用memhash

- key经过hash后，会得到一个64bit位，分配落到哪个桶中只会用到倒数B（B是桶个数的对数，即2^B = len(buckets)）位，例如：

  ```text
  B = 5，即桶个数为2^5 = 32个
  10010111 | 000011110110110010001111001010100010010110010101010 │ 01010
  用到最后5位：01010 = 10
  所以会落到第10个桶中
  ```

### 使用注意

- map是指针，在方法内修改map的值，会影响到方法外，因为即使值传递copy map的指针，copy完了依然指向原来的map
- map写入读取都涉及到copy，所以如果往map里放大对象会性能不好，但是往map放指针又会对GC不好，所以要具体场景具体分析
- 在golang 1.5版本之后，如果map的key和value存的不是指针，在GC的时候就不会被扫描，可以利用这个特性节省性能
- map删除元素不会触发自动缩容



## Channel

channel是有锁的

channel底层是个环状buffer

channel的调用会触发go routine调度（因为channel里面有个recvg字段会挂着想要从channel阻塞获取数据的goroutine，等到channel有数据了直接给这些goroutine）

### 使用注意

- 有buffer的channel会在放入&交给goroutine的时候涉及到两次copy，没有buffer的channel只会在交给goroutine的时候涉及到一次copy

- for+select channel会造成死循环，因为select内的break只能break掉select，对外面的for不起作用，例如：

  ```go
  func f(s []int) {
     ch := make(chan int)
     go func() {
        ch <- 1
     }()
  
     for {
        select {
        case i := <-ch:
           print(i)
        default:
           // break的是select而不是for，for将会死循环
           break
        }
     }
  }
  ```

  如果想要break外层for，需要用break label的形式：

  ```go
  func f(s []int) {
     ch := make(chan int)
     go func() {
        ch <- 1
     }()
     
  Loop:
     for {
        select {
        case i := <-ch:
           print(i)
        default:
           // 这次break的是for了
           break Loop
        }
     }
  }
  ```





## init方法

### 执行顺序

1. import package
2. 初始化变量
3. 执行init方法

### 使用注意

- 可读性较差
- 其它包也有init的时候可能会引发不可预知的问题
- 避免使用init修改全局变量
- 避免依赖init的顺序
- 避免在init中进行IO调用



## Panic

### 效果

- 当前function终止执行
- 调用栈中任何有defer调用的方法会执行defer
- 程序会终止

### 使用场景

- 任何让程序无法继续的错误出现时
- 人为错误：例如给一个以指针为参数的方法传了空指针

### 处理Panic注意点

- 使用defer+recover
- 把defer逻辑放在panic之前
- 多个defer按先进后出的顺序执行



## 锁

### 如何提高锁的性能

- 减少持有锁的时间
- 减小锁粒度
- 读写分离
- 使用原子操作

### 使用注意

- 不要拷贝mutex，会导致使用的不是同一把锁

- go的mutex不是可重入锁，反复调用会死锁，例如：

  ```go
  func f() {
  
     var m sync.Mutex
     
     m.Lock()
     defer m.Unlock()
     
     m.Lock() // 因为不是可重入锁，所以要等待自己解锁，才能重新加上，就死锁了
     defer m.Unlock()
     
  }
  ```

  

### 锁的进化史

##### 单核CPU：

- 屏蔽CPU中断（保证指令执行不会被打断，保证原子性）
- 使用CAS指令，提高效率

##### 多核CPU：

- 锁内存总线（效率太低，影响其他核读写）
- MESI协议，降低锁粒度，从硬件结构上优化，从锁内存总线变为锁CacheLine

##### Go Mutex实现：

- 正常模式：

  有一个等待队列

  新来的goroutine先抢占锁，如果失败加入等待队列

  当等待队列中的goroutine被唤醒，再次加锁失败，发现自己已经等待超过1ms，则把Mutex转化为饥饿模式

- 饥饿模式：

  按队列顺序先进先出

  如果队列中某个goroutine发现自己加锁成功后自己是队列中的最后一个/等待时间<1ms，就会重新切换回正常模式





## Golang调度



### 什么是调度

指按照一定的算法在适当的时候挑选出合适的goroutine，放到CPU上运行

伪代码如下：

```go
// 创建N个线程，每个线程执行调度函数
for i :=0; i < N; i++ {
  createThread(run(){
    schedule()
  })
}

func schedule() {
  // 调度主循环
  for {
    // 根据算法找出一个需要运行的goroutine
    g := findGroutine()
    // 运行goroutine
    run(g)
    // 如果运行遇到阻塞点，保存状态（例如读取channel遇到阻塞了）
    saveStatus(g)
  }
   
}
```



### GM模型

<img src="/Users/bytedance/Library/Application Support/typora-user-images/image-20210916154107655.png" alt="image-20210916154107655" style="zoom:30%;float:left" />

存在的问题：

- 存在全局互斥锁，每个线程去获取G的时候都要阻塞全局队列
- G的传递问题，当第二个G运行的时候，创建了一个新的G，新的G被另一个线程运行会导致隔离性被破坏
- M持有内存缓存，当G运行阻塞的时候，导致M持有的内存无法释放，造成浪费

计算机科学领域有句话：**任何问题都能通过增加中间层解决：**

### GPM模型









### 调度流程

<img src="/Users/bytedance/Library/Application Support/typora-user-images/image-20210917154428806.png" alt="image-20210917154428806" style="zoom:30%;float:left;" />

- 当写出go func这样的代码，就会创建G goroutine
- M会从P的本地队列、全局队列拉取G来执行，如果执行完了没活干，也会别的P的本地队列偷取G来执行
- M执行调度函数的逻辑主要是执行
- sysmon是后台监控goroutine，当发现有goroutine执行时间过长，就会抢占



##### 初始化

伪代码如下：

```go
func init() {
	// 初始化栈空间，复用管理链表
	stackInit()

	// P初始个数为CPU核数
	PCount := CPUCount

	// P个数不超过 GOMAXPROCS 配置的值
	if PCount > GOMAXPROCS {
		PCount = GOMAXPROCS
	}

	// 调整P的数量
	procResize(PCount)
}
```

```go
// 把 P 进行去余少补操作
func procResize(PCount int32) {
	// 申请新的 P
	for i := 0; i < PCount; i++ {
		// P是使用数组存储的
		P := allP[i]
		if P == nil {
			// 如果数组中i位置P不存在，就new一个
			P = newP(P)
			// 原子的把P存到P数组中
			atomicStore(allP[i], P)
		}

		// 为每个P分配mcache, go采用TCmalloc内存分配，会预先分配空闲内存等待goroutine使用，不够了再去内存中继续分配
		if P.cache == nil {
			P.cache = allocMCache()
		}
	}

	// 释放多余的P
	for i := 0; i < GOMAXPROCS; i++ {
		P := allP[i]

		if P.hasRunq {
			// 本地任务转移到全局任务队列
			putIntoGloabalRunq(P.runq)
		}

		// 释放当前P绑定的mcache
		freeMCache(P)

		...
	}

}
```





### 协作式调度&抢占式调度

协作式调度：依靠被调度方主动弃权

抢占式调度：调度器强制将被调度方中断

<img src="/Users/bytedance/Library/Application Support/typora-user-images/image-20210917164403489.png" alt="image-20210917164403489" style="zoom:50%;" />

