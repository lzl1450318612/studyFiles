# Golang slice



### 声明方式

- var s []int
- s := *new([]int)
- s := []int{1,2,3,4,5}
- s := make([]int, 5, 10)
- s := array[1:5]

注：

第一种创建的是nil slice，len和cap都是0，和nil比较的结果是true

nil slice和空slice不一样，空slice的len和cap也是0，但是空slice结构体中的array字段存的是同一个地址0xc42003bda0，空slice和nil比较的结果是false



### 结构

```go
// runtime/slice.go
type slice struct {
	array unsafe.Pointer // 元素指针
	len   int // 长度 
	cap   int // 容量
}
```

<img src="/Users/bytedance/Library/Application Support/typora-user-images/image-20211011180657383.png" alt="image-20211011180657383" style="zoom:50%;float:left" />



### 创建和扩容

golang代码如下：

```go
// main.go
slice := make([]int, 5, 10) // 长度为5，容量为10
slice[2] = 2                // 索引为2的元素赋值为2
fmt.Println(slice)
```

这段代码涉及了slice比较核心的两个操作：**make、append**

使用  **go tool compile -S main.go **获取汇编代码如下：

```assembly
"".main STEXT size=252 args=0x0 locals=0x58 funcid=0x0
        0x0000 00000 (main.go:7)        TEXT    "".main(SB), ABIInternal, $88-0
        0x0000 00000 (main.go:7)        MOVQ    (TLS), CX
        0x0009 00009 (main.go:7)        CMPQ    SP, 16(CX)
        0x000d 00013 (main.go:7)        PCDATA  $0, $-2
        0x000d 00013 (main.go:7)        JLS     242
        0x0013 00019 (main.go:7)        PCDATA  $0, $-1
        0x0013 00019 (main.go:7)        SUBQ    $88, SP
        0x0017 00023 (main.go:7)        MOVQ    BP, 80(SP)
        0x001c 00028 (main.go:7)        LEAQ    80(SP), BP
        0x0021 00033 (main.go:7)        FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0021 00033 (main.go:7)        FUNCDATA        $1, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
        0x0021 00033 (main.go:7)        FUNCDATA        $2, "".main.stkobj(SB)
        0x0021 00033 (main.go:8)        LEAQ    type.int(SB), AX
        0x0028 00040 (main.go:8)        MOVQ    AX, (SP)
        0x002c 00044 (main.go:8)        XORPS   X0, X0
        0x002f 00047 (main.go:8)        MOVUPS  X0, 8(SP)
        0x0034 00052 (main.go:8)        PCDATA  $1, $0
        0x0034 00052 (main.go:8)        CALL    runtime.makeslice(SB)
        0x0039 00057 (main.go:8)        MOVQ    24(SP), AX
        0x003e 00062 (main.go:9)        LEAQ    type.int(SB), CX
        0x0045 00069 (main.go:9)        MOVQ    CX, (SP)
        0x0049 00073 (main.go:9)        MOVQ    AX, 8(SP)
        0x004e 00078 (main.go:9)        XORPS   X0, X0
        0x0051 00081 (main.go:9)        MOVUPS  X0, 16(SP)
        0x0056 00086 (main.go:9)        MOVQ    $1, 32(SP)
        0x005f 00095 (main.go:9)        NOP
        0x0060 00096 (main.go:9)        CALL    runtime.growslice(SB)
        0x0065 00101 (main.go:9)        MOVQ    40(SP), AX
        0x006a 00106 (main.go:9)        MOVQ    48(SP), CX
        0x006f 00111 (main.go:9)        MOVQ    56(SP), DX
        0x0074 00116 (main.go:9)        MOVQ    $2, (AX)
        0x007b 00123 (main.go:10)       MOVQ    AX, (SP)
        0x007f 00127 (main.go:9)        LEAQ    1(CX), AX
        0x0083 00131 (main.go:10)       MOVQ    AX, 8(SP)
        0x0088 00136 (main.go:10)       MOVQ    DX, 16(SP)
        0x008d 00141 (main.go:10)       CALL    runtime.convTslice(SB)
        0x0092 00146 (main.go:10)       MOVQ    24(SP), AX
        0x0097 00151 (main.go:10)       XORPS   X0, X0
        0x009a 00154 (main.go:10)       MOVUPS  X0, ""..autotmp_13+64(SP)
        0x009f 00159 (main.go:10)       LEAQ    type.[]int(SB), CX
        0x00a6 00166 (main.go:10)       MOVQ    CX, ""..autotmp_13+64(SP)
        0x00ab 00171 (main.go:10)       MOVQ    AX, ""..autotmp_13+72(SP)
        0x00b0 00176 (<unknown line number>)    NOP
        0x00b0 00176 ($GOROOT/src/fmt/print.go:274)     MOVQ    os.Stdout(SB), AX
        0x00b7 00183 ($GOROOT/src/fmt/print.go:274)     LEAQ    go.itab.*os.File,io.Writer(SB), CX
        0x00be 00190 ($GOROOT/src/fmt/print.go:274)     MOVQ    CX, (SP)
        0x00c2 00194 ($GOROOT/src/fmt/print.go:274)     MOVQ    AX, 8(SP)
        0x00c7 00199 ($GOROOT/src/fmt/print.go:274)     LEAQ    ""..autotmp_13+64(SP), AX
        0x00cc 00204 ($GOROOT/src/fmt/print.go:274)     MOVQ    AX, 16(SP)
        0x00d1 00209 ($GOROOT/src/fmt/print.go:274)     MOVQ    $1, 24(SP)
        0x00da 00218 ($GOROOT/src/fmt/print.go:274)     MOVQ    $1, 32(SP)
        0x00e3 00227 ($GOROOT/src/fmt/print.go:274)     CALL    fmt.Fprintln(SB)
        0x00e8 00232 (main.go:10)       MOVQ    80(SP), BP
        0x00ed 00237 (main.go:10)       ADDQ    $88, SP
        0x00f1 00241 (main.go:10)       RET
        0x00f2 00242 (main.go:10)       NOP
        0x00f2 00242 (main.go:7)        PCDATA  $1, $-1
        0x00f2 00242 (main.go:7)        PCDATA  $0, $-2
        0x00f2 00242 (main.go:7)        CALL    runtime.morestack_noctxt(SB)
        0x00f7 00247 (main.go:7)        PCDATA  $0, $-1
        0x00f7 00247 (main.go:7)        JMP     0
```

其中最需要关注的是这几行：

```assembly
				0x0034 00052 (main.go:8)        CALL    runtime.makeslice(SB)
        0x0060 00096 (main.go:9)        CALL    runtime.growslice(SB)
        0x008d 00141 (main.go:10)       CALL    runtime.convTslice(SB)
        0x00e3 00227 ($GOROOT/src/fmt/print.go:274)     CALL    fmt.Fprintln(SB)
        0x00f2 00242 (main.go:7)        CALL    runtime.morestack_noctxt(SB)
```

#### runtime.makeslice:

```go
func makeslice(et *_type, len, cap int) unsafe.Pointer {
   mem, overflow := math.MulUintptr(et.size, uintptr(cap))
  
   if overflow || mem > maxAlloc || len < 0 || len > cap {
      mem, overflow := math.MulUintptr(et.size, uintptr(len))
      if overflow || mem > maxAlloc || len < 0 {
         panicmakeslicelen()
      }
      panicmakeslicecap()
   }

   return mallocgc(mem, et, true)
}


func MulUintptr(a, b uintptr) (uintptr, bool) {
   if a|b < 1<<(4*sys.PtrSize) || a == 0 {
      return a * b, false
   }
   overflow := b > MaxUintptr/a
   return a * b, overflow
}

```

可以看出 **MulUintptr** 函数的主要作用就是计算slice占用的空间

计算的方式就是用切片中元素大小和切片的容量相乘，如果容量大于最大可分配内存，overflow会返回true，makeslice就会报错

另外如果传入长度小于0或者长度小于容量，makeslice也会报错



**mallocgc **的源码太长了，是用来分配对象空间的，小对象（<=32KB）的会通过每个P缓存的空闲列表分配，大对象会直接从堆内存分配，这个函数比较通用，很多分配对象空间的地方都复用了，这里涉及到go的内存分配，不展开了

这里使用 **mallocgc** 就是通过上面使用 **MulUintptr **函数计算出的空间大小为slice分配空间



#### runtime.growslice：

```go
// 参数为元素类型、原切片容量、append元素后的预计容量
func growslice(et *_type, old slice, cap int) slice {
   ...
  
   // 新容量初始化为原切片容量
   newcap := old.cap
   doublecap := newcap + newcap
   if cap > doublecap {
      // 如果双倍原切片容量依然不满足append元素后的预计容量，直接使用预计容量（一般会在同时append多个元素时发生）
      newcap = cap
   } else {
      if old.cap < 1024 {
         // 如果原切片容量 < 1024，新容量使用原切片的双倍
         newcap = doublecap
      } else {
         for 0 < newcap && newcap < cap {
            // 循环扩容，每次x1.25，直到满足预计容量
            newcap += newcap / 4
         }
      }
   }

   var overflow bool
   var lenmem, newlenmem, capmem uintptr

   switch {
   case et.size == 1:
      // 如果元素类型占用的空间是1字节，例如int8
      lenmem = uintptr(old.len)
      newlenmem = uintptr(cap)
      // 进行内存对齐操作,计算，下同
      capmem = roundupsize(uintptr(newcap))
      overflow = uintptr(newcap) > maxAlloc
      newcap = int(capmem)
   case et.size == sys.PtrSize:
      // 如果元素类型占用的空间和指针占用空间一样
      lenmem = uintptr(old.len) * sys.PtrSize
      newlenmem = uintptr(cap) * sys.PtrSize
      capmem = roundupsize(uintptr(newcap) * sys.PtrSize)
      overflow = uintptr(newcap) > maxAlloc/sys.PtrSize
      newcap = int(capmem / sys.PtrSize)
   case isPowerOfTwo(et.size):
      // 如果元素类型占用的空间是2的n次方，例如int32占用4字节
      var shift uintptr
      if sys.PtrSize == 8 {
         shift = uintptr(sys.Ctz64(uint64(et.size))) & 63
      } else {
         shift = uintptr(sys.Ctz32(uint32(et.size))) & 31
      }
      lenmem = uintptr(old.len) << shift
      newlenmem = uintptr(cap) << shift
      capmem = roundupsize(uintptr(newcap) << shift)
      overflow = uintptr(newcap) > (maxAlloc >> shift)
      newcap = int(capmem >> shift)
   default:
      lenmem = uintptr(old.len) * et.size
      newlenmem = uintptr(cap) * et.size
      capmem, overflow = math.MulUintptr(et.size, uintptr(newcap))
      capmem = roundupsize(capmem)
      newcap = int(capmem / et.size)
   }

   if overflow || capmem > maxAlloc {
      panic(errorString("growslice: cap out of range"))
   }

   var p unsafe.Pointer
   if et.ptrdata == 0 {
      // 根据容量申请内存，下同
      p = mallocgc(capmem, nil, false)

      memclrNoHeapPointers(add(p, newlenmem), capmem-newlenmem)
   } else {
      p = mallocgc(capmem, et, true)
      if lenmem > 0 && writeBarrier.enabled {
         bulkBarrierPreWriteSrcOnly(uintptr(p), uintptr(old.array), lenmem-et.size+et.ptrdata)
      }
   }
   // 原子地把原数组中的内容拷贝到新数组中，并使用p充当slice的新的指针
   memmove(p, old.array, lenmem)

   // 最后返回len不变，但是cap扩容了的slice
   return slice{p, old.len, newcap}
}
```



分析以上源码可知，growslice主要按顺序做了如下操作：

1. growslice函数的三个参数分别是元素类型、原slice、原slice的cap加上新元素后所用的cap
2. 如果原cap x2仍然无法满足原slice加上新元素后所用的cap（这种情况会在一次append多个元素时发生），就使用原slice加上新元素后所用的cap；如果原cap x2可以满足满足原slice的cap加上新元素后所用的cap，判断原cap是否小于1024，如果小于，就确定用原cap x2；如果大于1024，就使用原cap x1.25
3. 初步确定cap之后，会根据元素的类型使用不同的计算方式计算需要分配的内存
4. 所有计算方式中都使用roundupsize函数进行了内存对齐的操作
5. 最后使用memmove将老slice中的元素拷贝到新slice中，并将slice指针更新成指向新slice



#### Q&A

##### 为什么内存对齐，怎么对齐的

操作系统的cpu不是一个字节一个字节访问内存的，是按2,4,8这样的字长来访问的，处理器从存储器子系统读取数据至寄存器，或者，写寄存器数据到存储器，传送的数据长度通常是字长。操作系统就可以一次定位到数据，这样会更加高效。无需多次读取、处理对齐运算等额外操作



在上层语言层面进行了内存对齐，那么用语言写出的代码在底层cpu执行时，就不会出现额外操作



内存在给程序使用时以8KB为单位分配，而go为对象分配内存空间的单位更加细化，一共有67种

下表反应了内存分配单位和浪费空间之间的关系：

| class | bytes/obj | bytes/span | objects | tail waste | max waste |
| :---: | --------: | ---------: | ------: | :--------: | :-------: |
|   1   |         8 |       8192 |    1024 |     0      |  87.50%   |
|   2   |        16 |       8192 |     512 |     0      |  43.75%   |
|   3   |        24 |       8192 |     341 |     0      |  29.24%   |
|   4   |        32 |       8192 |     256 |     0      |  46.88%   |
|   5   |        48 |       8192 |     170 |     32     |  31.52%   |
|   6   |        64 |       8192 |     128 |     0      |  23.44%   |
|   7   |        80 |       8192 |     102 |     32     |  19.07%   |
|   …   |         … |          … |       … |     …      |     …     |
|  67   |     32768 |      32768 |       1 |     0      |  12.50%   |

用class == 7的行举例子，每个对象最大80B大小，内存单元有8KB大，因为无法整除，所以尾部会浪费`8192%80 = 32B`

如果使用每个对象最大80B，那么对象至少也是65B（否则就分到64B的那种情况了），那么每个对象会浪费`80-65=15B`的空间

所以最多浪费((80-65)*102+32)/8192 = 19.07%的空间



进入刚才源码中的roundupsize函数：

```go
func roundupsize(size uintptr) uintptr {
   if size < _MaxSmallSize {
      if size <= smallSizeMax-8 {
         return uintptr(class_to_size[size_to_class8[divRoundUp(size, smallSizeDiv)]])
      } else {
         return uintptr(class_to_size[size_to_class128[divRoundUp(size-smallSizeMax, largeSizeDiv)]])
      }
   }
   if size+_PageSize < size {
      return size
   }
   return alignUp(size, _PageSize)
}
```

这个函数的名字的意思就是把size向上取整，可以看到逻辑比较简单，就是从几个数组中取固定值，继续去看这几个数组：

```go
var size_to_class8 = [smallSizeMax/smallSizeDiv + 1]uint8{0, 1, 2, 3, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 19, 19, 20, 20, 20, 20, 21, 21, 21, 21, 22, 22, 22, 22, 23, 23, 23, 23, 24, 24, 24, 24, 25, 25, 25, 25, 26, 26, 26, 26, 27, 27, 27, 27, 27, 27, 27, 27, 28, 28, 28, 28, 28, 28, 28, 28, 29, 29, 29, 29, 29, 29, 29, 29, 30, 30, 30, 30, 30, 30, 30, 30, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32}

var size_to_class128 = [(_MaxSmallSize-smallSizeMax)/largeSizeDiv + 1]uint8{32, 33, 34, 35, 36, 37, 37, 38, 38, 39, 39, 40, 40, 40, 41, 41, 41, 42, 43, 43, 44, 44, 44, 44, 44, 45, 45, 45, 45, 45, 45, 46, 46, 46, 46, 47, 47, 47, 47, 47, 47, 48, 48, 48, 49, 49, 50, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 53, 53, 54, 54, 54, 54, 55, 55, 55, 55, 55, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 58, 58, 58, 58, 58, 58, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 61, 61, 61, 61, 61, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67}

var class_to_size = [_NumSizeClasses]uint16{0, 8, 16, 24, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 208, 224, 240, 256, 288, 320, 352, 384, 416, 448, 480, 512, 576, 640, 704, 768, 896, 1024, 1152, 1280, 1408, 1536, 1792, 2048, 2304, 2688, 3072, 3200, 3456, 4096, 4864, 5376, 6144, 6528, 6784, 6912, 8192, 9472, 9728, 10240, 10880, 12288, 13568, 14336, 16384, 18432, 19072, 20480, 21760, 24576, 27264, 28672, 32768}
```

所以我们可以发现，go就是提前计算了一个表，然后在内存对齐的过程中向上取整，找到最适合的空间大小



##### 在append操作中，扩容完成后老slice对象改为新slice对象会不会发生元素拷贝，导致性能损耗

要看具体用法

如果是`s = append(s, 1, 2, 3)`，那么append后的slice会覆盖原slice，因为append关键字会以不同的方式展开，append的调用源码可以看`/cmd/compile/internal/gc/ssa.go`：

```go
// /cmd/compile/internal/gc/ssa.go
func (s *state) append(n *Node, inplace bool) *ssa.Value {
  
   // If inplace is false, process as expression "append(s, e1, e2, e3)":
   ptr, len, cap := s
   newlen := len + 3
   if newlen > cap {
       ptr, len, cap = growslice(s, newlen)	
       newlen = len + 3 // recalculate to avoid a spill
   }
   ...
   return makeslice(ptr, newlen, cap)
  
--------------------------------------------------------------------------------------------
   
  // If inplace is true, process as statement "s = append(s, e1, e2, e3)":
   a := &s
   ptr, len, cap := s
   newlen := len + 3
   if uint(newlen) > uint(cap) {
      newptr, len, newcap = growslice(ptr, len, cap, newlen)
      vardef(a)       // if necessary, advise liveness we are writing a new a
      *a.cap = newcap // write before ptr to avoid a spill
      *a.ptr = newptr // with write barrier
   }
   ...
```

可以看到如果是覆盖式用法，直接复用原slice对象

这里不要误会，growslice中，底层数组的元素拷贝在扩容中必然会发生，这里的优化只是针对上层slice对象拷贝做的



### 访问slice

在使用slice的时候，最常用的操作是：获取len、获取cap、根据index访问元素

这部分的源码在`cmd/compile/internal/gc/ssa.go`中可以找到，ssa就是将go语言转化成汇编的步骤

```go
func (s *state) expr(n *Node) *ssa.Value {
	switch n.Op {
  case OLEN, OCAP:
		switch {
		case n.Left.Type.IsSlice():
			op := ssa.OpSliceLen
			if n.Op == OCAP {
				op = ssa.OpSliceCap
			}
			return s.newValue1(op, types.Types[TINT], s.expr(n.Left))
		...
		}
	...
	case OINDEX:
		switch {
		case n.Left.Type.IsSlice():
			p := s.addr(n, false)
			return s.load(n.Left.Type.Elem(), p)
		...
		}
	...
	}
}
```

从源码中大体可以可以看出，访问slice cap和len的操作都在ssa阶段转化为了OpSliceCap和OpSliceLen操作，这些操作最终会被编译成汇编执行；直接使用下标访问的操作会被直接转化为对地址的访问



### 参考

- https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-array-and-slice/#323-%E8%AE%BF%E9%97%AE%E5%85%83%E7%B4%A0
- https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-memory-allocator/
- https://qcrao.com/2019/04/02/dive-into-go-slice/
- https://jodezer.github.io/2017/05/golangSlice%E7%9A%84%E6%89%A9%E5%AE%B9%E8%A7%84%E5%88%99
- https://juejin.cn/post/6844903812331732999
- https://zhuanlan.zhihu.com/p/53413177
- https://segmentfault.com/a/1190000037781000
