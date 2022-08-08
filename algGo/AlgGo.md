# AlgGo

### 八大排序

- 插入排序：
  - 原理：例如抽牌，每抽一张牌把它按顺序插入到对应的位置，倒着遍历数组
  - 具体操作：从数组的第二位开始向后遍历，每次遍历到一位之后，把这位和前一位比较，如果小于前一位，就往前交换位置

- 选择排序：
  - 原理：每次选择数组剩余数字中最小的一个放在当前遍历到的位置
  - 具体操作：从第一位开始遍历，每遍历到一位时，把剩余数字中最小的一个和这位的数字交换位置

- 归并排序：
  - 原理：把大数组平分为两个子数组，再把两个子数组平分为四个子数组......以此类推；当数组的数字个数是1的时候，自然是有序的；接下来把拆分成的子数组按顺序合并（类似合并两个有序数组的操作），最终合并成的大数组就是排序好的数组
    - 具体操作：递归地把数组一拆为二，二拆为四......，当拆到每个子数组只剩一个元素的时候，停止递归；合并操作可以转化为如何把两个有序数组合并成一个有序数组的问题，这个非常简单，使用两个指针，比较大小即可