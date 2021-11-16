# Golang map



### 结构

<img src="/Users/bytedance/Library/Application Support/typora-user-images/image-20211018161707035.png" alt="image-20211018161707035" style="zoom:70%;" />

```go
// runtime/map.go
type hmap struct {
	count     int // 元素个数
	flags     uint8
  B         uint8  // 桶个数的对数，2^B = len(buckets)
	noverflow uint16 // 溢出桶个数
	hash0     uint32 // hash的种子，能为hash结果引入随机性，这个值在创建hash表时确定，在调用hash函数时作为参数传入

	buckets    unsafe.Pointer // 桶数组指针，如果count是0，则这个字段为nil
	oldbuckets unsafe.Pointer // map扩容时用来保存之前buckets的字段，大小是当前buckets的一半
	nevacuate  uintptr        // 用来指示扩容进度，小于此地址的 buckets 迁移完成

  // 可选字段，用来存储map额外信息
	extra *mapextra
}


type mapextra struct {
	
  // 如果map的key和value都不包含指针，那么GC的时候，就不会扫描map，可以节省性能
  // 但是map中又需要存储溢出桶指针，这样溢出桶才不会被GC回收
  // 所以把新老溢出桶存在了extra结构里
	overflow    *[]*bmap
	oldoverflow *[]*bmap

	// nextOverflow 包含空闲的 overflow bucket，这是预分配的 bucket
	nextOverflow *bmap
}

// 经过编译期间添加字段后的bmap
type bmap struct {
	topbits  [8]uint8
  // 源码中没有，是编译期间动态生成的
  keys     [8]keytype
  values   [8]valuetype
  pad      uintptr
  overflow uintptr
}

```

上面展示了map的大致结构示意图，可以看到key和value并不是一一对应的放置，这样是为了节省内存考虑

比如map[int64]int8，如果把key和value一一对应放置，那么读取的时候因为每个value是int8类型，1字节，那么读取内存块的时候相当于造成了浪费，而把key放在一起，value放在一起可以对齐内存，解决这个问题



为什么bmap的结构要加上注释指明是在编译期间添加字段后的结构，是因为在运行期间，bmap结构体其实不止包含 `tophash` 字段，因为map中可能存储不同类型的键值对，而且 Go 也不支持泛型，所以键值对占据的内存空间大小只能在编译时进行推导。bmap中的其他字段在运行时也都是通过计算内存地址的方式访问的，我们能根据编译期间的 [`cmd/compile/internal/gc.bmap`](https://draveness.me/golang/tree/cmd/compile/internal/gc.bmap) 函数重建它的结构



### hash方式

```go
// src/runtime/alg.go
func alginit() {
   // Install AES hash algorithms if the instructions needed are present.
   if (GOARCH == "386" || GOARCH == "amd64") &&
      cpu.X86.HasAES && // AESENC
      cpu.X86.HasSSSE3 && // PSHUFB
      cpu.X86.HasSSE41 { // PINSR{D,Q}
      initAlgAES()
      return
   }
   if GOARCH == "arm64" && cpu.ARM64.HasAES {
      initAlgAES()
      return
   }
   getRandomData((*[len(hashkey) * sys.PtrSize]byte)(unsafe.Pointer(&hashkey))[:])
   hashkey[0] |= 1 // make sure these numbers are odd
   hashkey[1] |= 1
   hashkey[2] |= 1
   hashkey[3] |= 1
}

func initAlgAES() {
   useAeshash = true
   // Initialize with random data so hash collisions will be hard to engineer.
   getRandomData(aeskeysched[:])
}
```

目前的一些cpu可以支持aes指令，从硬件上做这种操作速度很快，因此，对于有aes支持的cpu，使用 aes进行hash，否则使用 memhash

这里有一些常用hash算法性能对比：https://blog.cyeam.com/hash/2018/05/28/hash-method



go中表示类型的结构体：

```go
type _type struct {
    size       uintptr
    ptrdata    uintptr // size of memory prefix holding all pointers
    hash       uint32
    tflag      tflag
    align      uint8
    fieldalign uint8
    kind       uint8
    alg        *typeAlg // 注意这个
    gcdata    *byte
    str       nameOff
    ptrToThis typeOff
}

// src/runtime/alg.go
type typeAlg struct {
    // (ptr to object, seed) -> hash
    hash func(unsafe.Pointer, uintptr) uintptr
    // (ptr to object A, ptr to object B) -> ==?
    equal func(unsafe.Pointer, unsafe.Pointer) bool
}
```

类型结构体中alg字段就和hash有关，可以看到类型是typeAlg，而typeAlg中，hash字段的函数用来计算类型的hash值，equal函数用来计算判断两个类型是否hash相等



例如，string类型的hash和equal函数如下：

```go
func strhash(a unsafe.Pointer, h uintptr) uintptr {
    x := (*stringStruct)(a)
    return memhash(x.str, h, uintptr(x.len))
}
func strequal(p, q unsafe.Pointer) bool {
    return *(*string)(p) == *(*string)(q)
}
```



### 创建

#### runtime.makemap:

```go
// makemap implements Go map creation for make(map[k]v, hint).
// If the compiler has determined that the map or the first bucket
// can be created on the stack, h and/or bucket may be non-nil.
// If h != nil, the map can be created directly in h.
// If h.buckets != nil, bucket pointed to can be used as the first bucket.
func makemap(t *maptype, hint int, h *hmap) *hmap {
  
  // 计算hash占用的内存是否溢出或者超出能分配的最大值
	mem, overflow := math.MulUintptr(uintptr(hint), t.bucket.size)
	if overflow || mem > maxAlloc {
		hint = 0
	}

	// 初始化hmap对象
	if h == nil {
		h = new(hmap)
	}
  // 获取hash函数随机种子
	h.hash0 = fastrand()

	// 根据hint计算需要的最少的桶数量
	B := uint8(0)
	for overLoadFactor(hint, B) {
		B++
	}
	h.B = B

	// 如果B == 0，那么buckets就会在给map赋值的时候再分配
	if h.B != 0 {
		var nextOverflow *bmap
    // 创建用来保存桶的array
		h.buckets, nextOverflow = makeBucketArray(t, h.B, nil)
    // 如果有预分配的溢出桶，就把它保存在extra中
		if nextOverflow != nil {
			h.extra = new(mapextra)
			h.extra.nextOverflow = nextOverflow
		}
	}

	return h
}



func makeBucketArray(t *maptype, b uint8, dirtyalloc unsafe.Pointer) (buckets unsafe.Pointer, nextOverflow *bmap) {
	base := bucketShift(b)
	nbuckets := base
	// 当桶的数量小于 2^4 时，由于数据较少、使用溢出桶的可能性较低，省略创建的过程以减少额外开销
	if b >= 4 {
    // 当桶的数量大于等于2^4，会额外创建2^(B-4)个溢出桶
		nbuckets += bucketShift(b - 4)
		sz := t.bucket.size * nbuckets
    // 熟悉的内存对齐操作，对bucket占用空间做对齐
		up := roundupsize(sz)
		if up != sz {
      // 如果内存对齐后和改变了占用空间不同，就重新计算一下桶个数
			nbuckets = up / t.bucket.size
		}
	}

  // 如果桶的个数经过上边的操作之后有更改，就重新修改溢出桶的指针指向
	if base != nbuckets {
		nextOverflow = (*bmap)(add(buckets, base*uintptr(t.bucketsize)))
		last := (*bmap)(add(buckets, (nbuckets-1)*uintptr(t.bucketsize)))
		last.setoverflow(t, (*bmap)(buckets))
	}
	return buckets, nextOverflow
}

```

可以看到 `makemap` 函数返回的结果是一个指针，这里也可以看出和slice的不同，makeslice函数返回的是一个slice结构体，所以，当map作为函数参数时，函数内部的操作会作用到函数外部的map；而slice则不会，这是因为go是值传递，调用方法是穿的都是参数的copy，map指针经过copy后，指向的还是原来map的地址



### 查询

key经过hash算法计算后，会得到一个hash值，例如：

`1001011100001111011011001000111100101010001001011001010101001010`

判断元素要落到哪个桶的时候，只会用到最后B个bit位

假设B = 5，也就是有2^5 = 32个桶，那么

`10010111 | 000011110110110010001111001010100010010110010101010 │ 01010`

最后5位是 `01010` ，是10，所以这个元素会落到第10个桶中

前8位就是tophash，用这8位寻找应该放在桶中哪个位置



整体过程大概如下图所示：

<img src="/Users/bytedance/Library/Application Support/typora-user-images/image-20211018170914316.png" alt="image-20211018170914316" style="zoom:67%;" />





```go
// mapaccess1是使用 x := map[key]的时候调用的函数
// mapaccess2是使用 x, ok := map[key]的时候调用的函数
// 区别只是多了bool返回值的处理操作
func mapaccess1(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
   ...
  
   // 如果map是空，直接返回零值
   if h == nil || h.count == 0 {
      if t.hashMightPanic() {
         // 这个是go解决的一个小bug，可以看https://github.com/golang/go/issues/23734
         t.hasher(key, 0) 
      }
      return unsafe.Pointer(&zeroVal[0])
   }
  // 读写冲突
   if h.flags&hashWriting != 0 {
      throw("concurrent map read and map write")
   }
   // 计算hash值
   hash := t.hasher(key, uintptr(h.hash0))
   // 因为B是2的对数，所以 m=B-1 在二进制上是全1，所以 hash&m 可以达到取hash值的最后B位的效果
   m := uintptr(1)<<h.B - 1
   // b就是桶的地址，查询的时候，把
   b := (*bmap)(add(h.buckets, (hash&m)*uintptr(t.bucketsize)))
   // 如果oldbuckets不是nil，说明正在扩容
   if c := h.oldbuckets; c != nil {
      // 如果不是同size扩容（下面说）
      if !h.sameSizeGrow() {
         // 新桶数量是老的2倍，所以B要+1，所以根据B算出来的m要右移1位
         m >>= 1
      }
      // 求出key在老map中的桶地址
      oldb := (*bmap)(add(c, (hash&m)*uintptr(t.bucketsize)))
      // 如果此元素没有移动到新桶，就在老桶里找
      if !evacuated(oldb) {
         b = oldb
      }
   }
   // 计算高8位的tophash，计算方式是右移56位，只取高8位
   top := tophash(hash)
bucketloop:
   // 遍历桶和溢出桶链表，查询元素
   for ; b != nil; b = b.overflow(t) {
      // 
      for i := uintptr(0); i < bucketCnt; i++ {
         // 如果tophash不匹配，继续在桶里找下一个位置比对
         if b.tophash[i] != top {
            // 这个位置的tophash是0，并且没有别的tophash位或溢出桶
            if b.tophash[i] == emptyRest {
               break bucketloop
            }
            continue
         }
         // 如果tophash值匹配，就定位到key的位置
         k := add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
         if t.indirectkey() {
            k = *((*unsafe.Pointer)(k))
         }
         // 如果key相等
         if t.key.equal(key, k) {
            // 找到value的位置
            e := add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.elemsize))
            if t.indirectelem() {
               e = *((*unsafe.Pointer)(e))
            }
            // 返回value
            return e
         }
      }
   }
   // 如果啥都没找到，返回零值
   return unsafe.Pointer(&zeroVal[0])
}
```

这里说一下key和value的定位方式：

```go
// key 定位
k := add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
// value 定位
v := add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.valuesize))
```

dataOffset是key相对于map起始地址的偏移，可以对照着刚才画的示意图看，那么：

**一个key的位置=dataOffset+它前面所有key的大小**

**一个value的位置=dataOffset+全部key的大小+它前面所有value的大小**



### 扩容



#### loadFactor

我们都知道如果一个map的元素过多，不作处理的话，hash碰撞概率太高，大部分元素都落到一个桶里，效率很低，最差情况直接退化成链表，这种就需要进行扩容来解决。那么衡量这种情况的指标，就是loadFactor：

```go
loadFactor := count / (2^B)
```

#### 触发时机

下面是map赋值过程的部分源码

扩容会在赋值/删除元素的函数中触发

```go
// runtime/map.go
func mapassign(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
   ...

   // if(loadFactor超过阈值 || 溢出桶太多) && 没有处在扩容过程中，那么进行扩容操作
   if !h.growing() && (overLoadFactor(h.count+1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
      hashGrow(t, h)
      goto again // Growing the table invalidates everything, so try again
   }
  
   ...
  
   again:
	 	bucket := hash & bucketMask(h.B)
   	// 当map正在处于扩容状态时，写入/删除值时都会触发 growWork 增量扩容
	 	if h.growing() {
		 	growWork(t, h, bucket)
	 	}
  
	 ...
}

const (
  loadFactorNum = 13
	loadFactorDen = 2
)

func overLoadFactor(count int, B uint8) bool {
  //loadFactor超出6.5
	return count > bucketCnt && uintptr(count) > loadFactorNum*(bucketShift(B)/loadFactorDen)
}

func tooManyOverflowBuckets(noverflow uint16, B uint8) bool {
	if B > 15 {
		B = 15
	}
  // 溢出桶个数>=2^B时，算多的
	return noverflow >= uint16(1)<<(B&15)
}
```

1. loadFactor超过阈值6.5
2. 溢出桶数量过多

#### 扩容策略

1. 对于loadFactor超过阈值6.5的情况，显然是元素数量太多，已经要把所有的桶都装满了
2. 对于溢出桶数量过多的情况，可能是因为先插入了很多元素，导致使用了很多溢出桶，然后又删除了很多元素，造成元素过于分散，很多桶都没装满，这种情况也需要扩容，分配一个新的桶，把老桶中的元素都迁移到新桶中来，让元素存储更紧密，效率提高

针对不同的扩容策略，go在源码中实现了两种两种情况下的扩容



#### 核心源码

可以看下hashGrow的源码：

```go
func hashGrow(t *maptype, h *hmap) {
   bigger := uint8(1)
   // 如果是有很多溢出桶的情况触发的扩容，那么桶数量不增加，只是rehash元素
   if !overLoadFactor(h.count+1, h.B) {
      bigger = 0
      // 使用等量扩容标志
      h.flags |= sameSizeGrow
   }
   // 暂存老的桶
   oldbuckets := h.buckets
  // 分配新的桶，如果bigger是1，那么2^(B+1)相当于桶个数翻倍
   newbuckets, nextOverflow := makeBucketArray(t, h.B+bigger, nil)

   // 使用buckets字段的迭代器使用标志清零，并转移到oldbuckets字段
   flags := h.flags &^ (iterator | oldIterator)
   if h.flags&iterator != 0 {
      flags |= oldIterator
   }
   // 提交扩容
   h.B += bigger
   h.flags = flags
   h.oldbuckets = oldbuckets
   h.buckets = newbuckets
   // 迁移进度为0
   h.nevacuate = 0
   // 溢出桶个数为0
   h.noverflow = 0
   
   ...
}
```

可以看出hashGrow函数中主要是申请新的桶空间，以及修改了一些状态标志

真正的迁移操作在函数 `growWork` 中：

```go
func growWork(t *maptype, h *hmap, bucket uintptr) {
   // make sure we evacuate the oldbucket corresponding
   // to the bucket we're about to use
   evacuate(t, h, bucket&h.oldbucketmask())

   // 如果h正在扩容，那么获取目前扩容进度，然后从当前进度再迁移一个桶，加速扩容
   if h.growing() {
      evacuate(t, h, h.nevacuate)
   }
}
```



核心函数：

```go
type evacDst struct {
	b *bmap          // 迁移目标桶
	i int            // key/value在桶中的index
	k unsafe.Pointer // key地址
	e unsafe.Pointer // value地址
}

func evacuate(t *maptype, h *hmap, oldbucket uintptr) {
   // 定位老桶的地址
   b := (*bmap)(add(h.oldbuckets, oldbucket*uintptr(t.bucketsize)))
   // 老桶的数量，2^B
   newbit := h.noldbuckets()
   // 如果没有被迁移过
   if !evacuated(b) {

      // 用来分别表示两种扩容策略中要迁移到的目标地址，如果是不是等量扩容的话，容量变为原来的两倍，那么用
      // X来描述原来的所有桶，用Y来描述后来新加的所有桶
      var xy [2]evacDst
      // 预设使用x，也就是等量扩容策略
      x := &xy[0]
      x.b = (*bmap)(add(h.buckets, oldbucket*uintptr(t.bucketsize)))
      x.k = add(unsafe.Pointer(x.b), dataOffset)
      x.e = add(x.k, bucketCnt*uintptr(t.keysize))

      if !h.sameSizeGrow() {
         // 如果不是等量扩容策略，使用y进行迁移
         y := &xy[1]
         // 这里迁移到的新桶中的桶序号增加了2^B
         y.b = (*bmap)(add(h.buckets, (oldbucket+newbit)*uintptr(t.bucketsize)))
         y.k = add(unsafe.Pointer(y.b), dataOffset)
         y.e = add(y.k, bucketCnt*uintptr(t.keysize))
      }
			// 遍历所有桶，包括溢出桶
      for ; b != nil; b = b.overflow(t) {
         // 获取key和value的起始地址
         k := add(unsafe.Pointer(b), dataOffset)
         e := add(k, bucketCnt*uintptr(t.keysize))
         // 遍历桶内所有cell
         for i := 0; i < bucketCnt; i, k, e = i+1, add(k, uintptr(t.keysize)), add(e, uintptr(t.elemsize)) {
            // 获取当前cell的tophash值
            top := b.tophash[i]
            if isEmpty(top) {
               // 如果tophash是空，就标记这个cell被迁移过了
               b.tophash[i] = evacuatedEmpty
               continue
            }
            if top < minTopHash {
               throw("bad map state")
            }
            k2 := k
            if t.indirectkey() {
               k2 = *((*unsafe.Pointer)(k2))
            }
            var useY uint8
            if !h.sameSizeGrow() {
               // 如果不是等量扩容，就需要rehash
               hash := t.hasher(k2, uintptr(h.hash0))
               // 如果有goroutine正在遍历map
               if h.flags&iterator != 0 && !t.reflexivekey() && !t.key.equal(k2, k2) {
                  // 如果key是NaN值，两次hash的值不同，那么只通过tophash的最低一位判断分配到X还是Y中
                  useY = top & 1
                  top = tophash(hash)
               } else {
                  if hash&newbit != 0 {
                     useY = 1
                  }
               }
            }

            if evacuatedX+1 != evacuatedY || evacuatedX^1 != evacuatedY {
               throw("bad evacuatedN")
            }

            b.tophash[i] = evacuatedX + useY // evacuatedX + 1 == evacuatedY
            dst := &xy[useY]                 // evacuation destination

            // 如果目标桶内cell个数达到8个了，满了
            if dst.i == bucketCnt {
               // 移动目的地变为下一个溢出桶
               dst.b = h.newoverflow(t, dst.b)
               // index从0重新开始计数
               dst.i = 0
               // 重新计算k和v的位置
               dst.k = add(unsafe.Pointer(dst.b), dataOffset)
               dst.e = add(dst.k, bucketCnt*uintptr(t.keysize))
            }
            // 规避边界检查，优化性能
            dst.b.tophash[dst.i&(bucketCnt-1)] = top
            if t.indirectkey() {
               *(*unsafe.Pointer)(dst.k) = k2 // copy pointer
            } else {
               typedmemmove(t.key, dst.k, k) // copy elem
            }
            if t.indirectelem() {
               *(*unsafe.Pointer)(dst.e) = *(*unsafe.Pointer)(e)
            } else {
               typedmemmove(t.elem, dst.e, e)
            }
            // index计数+1
            dst.i++
            // key和value位置都后移一个
            dst.k = add(dst.k, uintptr(t.keysize))
            dst.e = add(dst.e, uintptr(t.elemsize))
         }
      }
      // 如果没有goroutine在使用老桶，就清理掉，缓解GC压力
      if h.flags&oldIterator == 0 && t.bucket.ptrdata != 0 {
         b := add(h.oldbuckets, oldbucket*uintptr(t.bucketsize))
         // 只清除dataOffset之后的部分，保留tophash部分，来指示迁移状态
         ptr := add(b, dataOffset)
         n := uintptr(t.bucketsize) - dataOffset
         memclrHasPointers(ptr, n)
      }
   }

   if oldbucket == h.nevacuate {
      advanceEvacuationMark(h, t, newbit)
   }
}

func advanceEvacuationMark(h *hmap, t *maptype, newbit uintptr) {
  // 迁移进度+1
	h.nevacuate++
	// 提前向后看1024个桶，这样可以加快迁移速度，1024是go实验得出
	stop := h.nevacuate + 1024
	if stop > newbit {
		stop = newbit
	}
  // 寻找没有迁移的桶
	for h.nevacuate != stop && bucketEvacuated(t, h, h.nevacuate) {
		h.nevacuate++
	}
  // 如果迁移进度等于老桶数量，代表迁移完成
	if h.nevacuate == newbit {
    // 清除老桶
		h.oldbuckets = nil
		
		if h.extra != nil {
      // 清除老的溢出桶
			h.extra.oldoverflow = nil
		}
    // 清除正在扩容的标志
		h.flags &^= sameSizeGrow
	}
}
```

- 扩容会去掉不必要的溢出桶，紧凑数据，提高效率
- 如果是双倍扩容，相当于原来的1倍桶裂变成2倍桶，原来取hash值最后B位判断落到哪个桶，现在就是取最后B+1位进行判断



### 删除

删除的源码比较简单，很多都和前面的重复，也是先获取hash值，然后两层for循环找到元素，最后清除，这里只列出部分源码

```go
func mapdelete(t *maptype, h *hmap, key unsafe.Pointer) {
   ...
   // 添加写状态标记
   h.flags ^= hashWriting

	 bucket := hash & bucketMask(h.B)
   // 如果map正在扩容，加速扩容
	 if h.growing() {
		 growWork(t, h, bucket)
	 }
  
   ...
  
search:
   for ; b != nil; b = b.overflow(t) {
      for i := uintptr(0); i < bucketCnt; i++ {
        
         ...
        
         // 清除key
         if t.indirectkey() {
            *(*unsafe.Pointer)(k) = nil
         } else if t.key.ptrdata != 0 {
            memclrHasPointers(k, t.key.size)
         }
         // 清除value
         e := add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.elemsize))
         if t.indirectelem() {
            *(*unsafe.Pointer)(e) = nil
         } else if t.elem.ptrdata != 0 {
            memclrHasPointers(e, t.elem.size)
         } else {
            memclrNoHeapPointers(e, t.elem.size)
         }
        
         ...
        
      notLast:
         // 元素个数count值-1
         h.count--
         ...
      }
   }

   if h.flags&hashWriting == 0 {
      throw("concurrent map writes")
   }
   // 清除写状态标记
   h.flags &^= hashWriting
}
```



### 遍历

map的遍历依赖map迭代器进行操作，迭代器的结构如下：

```go
type hiter struct {
   // key指针
   key         unsafe.Pointer
   // value指针
   elem        unsafe.Pointer
   // map类型
   t           *maptype
   // map起始地址
   h           *hmap
   // 初始指向的桶
   buckets     unsafe.Pointer // bucket ptr at hash_iter initialization time
   // 当前指向的桶
   bptr        *bmap          // current bucket
   // 新老溢出桶
   overflow    *[]*bmap       // keeps overflow buckets of hmap.buckets alive
   oldoverflow *[]*bmap       // keeps overflow buckets of hmap.oldbuckets alive
   // 遍历起始地址
   startBucket uintptr        // bucket iteration started at
   // 遍历起始cell编号
   offset      uint8          // intra-bucket offset to start from during iteration (should be big enough to hold bucketCnt-1)
   // 是否从头遍历
   wrapped     bool           // already wrapped around from end of bucket array to beginning
   // B
   B           uint8
   // 当前cell序号
   i           uint8
   // 当前的桶
   bucket      uintptr
   // 扩容需要检查的桶
   checkBucket uintptr
}
```

接下来来看源码

首先会对迭代器进行初始化操作，看mapiterinit函数：

```go
func mapiterinit(t *maptype, h *hmap, it *hiter) {
   
   ...

   // 生成随机数，从随机位置开始遍历
   r := uintptr(fastrand())
   if h.B > 31-bucketCntBits {
      r += uintptr(fastrand()) << 31
   }
   // 从哪个桶开始遍历
   it.startBucket = r & bucketMask(h.B)
   // 从桶的哪个cell开始遍历
   it.offset = uint8(r >> h.B & (bucketCnt - 1))
   
   ...

   // 开始遍历
   mapiternext(it)
}
```

可以看到map的遍历都是从随机位置开始的，也就是说即使map没有改变过，多次遍历的顺序也是不一样的

遍历的核心代码mapiternext函数比较简单，重复度较高，这里不全部列出了，我们都能猜到就是双层循环遍历每个桶和其中的cell，其中值得注意的是当扩容操作进行中进行遍历会做的处理，例如，扩容前B=1，有2个桶；扩容后B=2，有4个桶，但是扩容刚刚进行了一半，这时进行遍历：

<img src="/Users/bytedance/Library/Application Support/typora-user-images/image-20211021142922997.png" alt="image-20211021142922997" style="zoom:67%;" />

假设随机之后，startBucket从最后一个桶开始，那么遍历的流程就是：

1. 遍历11号桶中所有cell，返回遍历元素
2. 遍历00号桶所有cell，因为map处在扩容过程中，00号桶的位置应该由原来的1号桶裂变迁移得到，但是还未完成，于是回到oldBuckets字段中找到原来的1号桶进行遍历操作，这里只会取出将会rehash到00号桶的元素，对将会rehash的10号桶的元素不会取到结果集中
3. 遍历01号桶
4. 遍历10号桶，做和2种相同的操作
5. 返回结果集



### 参考

- https://www.bookstack.cn/read/qcrao-Go-Questions/map-map%20%E7%9A%84%E9%81%8D%E5%8E%86%E8%BF%87%E7%A8%8B%E6%98%AF%E6%80%8E%E6%A0%B7%E7%9A%84.md
- https://juejin.cn/post/6844903728697327630
- https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap/
- https://qcrao.com/2019/05/22/dive-into-go-map/#%E5%88%9B%E5%BB%BA-map
