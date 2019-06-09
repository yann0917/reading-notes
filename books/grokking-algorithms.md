# 算法图解

## 算法简介

大 O 表示法指出了算法的运行时间的增速

* O(1), 也叫 `常量时间`,
* O(log n)，也叫 `对数时间` ，这样的算法包括二分查找。
* O(n)，也叫 `线性时间` ，这样的算法包括简单查找。
* O(n * log n)，这样的算法包括快速排序，归并排序。
* O(n^2)，这样的算法包括选择排序，冒泡排序。
* O(n!)，这样的算法包括旅行商问题的解决方案——一种非常慢的算法。

## 选择排序

数组和链表的时间复杂度

|      | array | list |
|------|-------|------|
| 读取 | O(1)  | O(n) |
| 插入 | O(n)  | O(1) |
| 删除 | O(n)  | O(1) |

* 链表只能顺序访问，数组支持随机访问
* 链表的优势在于插入和删除，数组的优势在于读取

## 递归

递归只是让解决方案更清晰，并没有性能上的优势。每个递归函数有两个部分：基线条件和递归条件，基线条件是跳出递归的条件，避免进入死循环；递归条件指的是函数调用自己。

**尾递归**：

尾递归就是把当前的运算结果（或路径）放在参数里传给下层函数，深层函数所面对的不是越来越简单的问题，而是越来越复杂的问题，因为参数里带有前面若干步的运算路径。

尾递归的实现依赖于编译器的帮助，并非所有的语言都实现了尾递归。

## 快速排序

分而治之（D&C）工作原理：

1. 找出简单的基线条件；
2. 确定如何缩小问题的规模，使其符合基线条件；

每一轮挑选一个基准元素(pivot)，并让其他比它大的元素移动到一边，比它小的元素移动到另一边，从而把数列拆解成了两个部分。然后对这两个部分进行快速排序。
快速排序的性能高度依赖于选择的基准值，平均运行时间为O(n * log n)，最糟糕的情况下，运行时间为O(n^2)

## 散列表

## 广度优先搜索

## 迪克斯特拉算法

## 贪婪算法

## 动态规划

## K 最近邻算法