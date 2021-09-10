## Golang Slice

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



