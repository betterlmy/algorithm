# 链表

链表(Linked List)是一种物理存储结构上非连续,非顺序的存储结构,数据的逻辑顺序是是通过链表中的指针链接次序所实现的,由一系列节点组成.

> 结点=节点

![img](linkedlist.assets/d5d5bee4be28326ba3c28373808a62cd.jpg)

数组需要一块连续的物理内存进行数据存储,而链表不需要,通过指针将内存块串联起来.

## 常见的链表

链表结构五花八门,最常见的链表结构包括[单链表](#单链表),[双向链表](#双向链表)和[循环链表](#循环链表).

### 单链表

单链表是最简单最常用的链表.

![img](linkedlist.assets/b93e7ade9bb927baad1348d9a806ddeb.jpg)

单链表中的每个节点(Node)除了要存储数据之外还需要存储一个指针,这个指针指向下一个节点,记录着下一个节点的地址,我们称其为`后继指针(next)`.最后一个后继指针指向NULL(Go语言中的nil).

在单链表中有两个特殊的Node,分别是`头节点(Head Node)`和`尾节点(Tail Node)`. 

> 头节点用来存储链表的基地址,通过从基地址开始遍历可以得到整条链表.
>
> 尾节点是链表的最后一个节点,当遍历到尾节点时,证明对当前的链表遍历已经结束.

与数组相同,链表也支持查找,插入和删除操作.

在链表中插入或者删除一个数据我们只需要修改Node的指针内容即可.时间复杂度为O(1).

![img](linkedlist.assets/452e943788bdeea462d364389bd08a17.jpg)

有利有弊,链表想要访问第k个元素,需要从头进行遍历,寻址k+1次即可找到第k个元素.平均时间复杂度为O(n).



### 循环链表

循环链表是一种特殊的单链表,特殊在tail node,在循环链表中,尾节点的指针并非指向null,而是指向了head node,如下图表示

![img](linkedlist.assets/86cb7dc331ea958b0a108b911f38d155.jpg)

### 双向链表

单链表只有一个方向,而双向链表则支持两个方向,每个节点包括一个`前驱指针prev`,`后继指针next`和一个`数据data`.

![img](linkedlist.assets/cbc8ab20276e2f9312030c313a9ef70b.jpg)

在实际的软件开发中,从链表中删除一个数据无非两种情况:

* 删除链表中data=x的node
* 删除给定指针指向的node

对于第一种情况,无论是单链表还是双向链表,为了找到指定节点,都需要从头开始遍历

对于第二种情况,我们已经找到了要删除的节点,假设我们需要使用单链表,我们需要将要被删除的节点的前置节点的指针指向被删除的下一个节点,**但是我们没有办法得到前置节点的地址**,只能重新开始遍历,时间复杂度为O(n).

此时双向链表直接定位到要删除的节点,根据前驱指针得到前置节点,完成后续操作即可,时间复杂度O(1).

同理 当我们需要在指定元素前插入某个节点,双向链表的时间复杂度仍为O(1).

双向链表是经典的**空间换时间**的思想体现.

### 双向循环链表

![img](linkedlist.assets/d1665043b283ecdf79b157cfc9e5ed91.jpg)

## 链表与数组性能对比

![img](linkedlist.assets/4f63e92598ec2551069a0eef69db7168.jpg)

## 