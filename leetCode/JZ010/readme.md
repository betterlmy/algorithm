# 剑指 Offer II 010. 和为 k 的子数组
## 题目描述
给定一个**整数数组(非正整数,非有序)**和一个整数k，请找到该数组中和为k的**连续**子数组的个数。

```text
输入:nums = [1,1,1], k = 2
输出: 2
解释: 此题 [1,1] 与 [1,1] 为两种不同的情况
```

```text
输入:nums = [1,2,3], k = 3
输出: 2
```

## 关键点

1. 非正整数(存在负数)
2. 非有序(可能需要排序?)
3. 求连续的子数组

## 思路

### 枚举
$$O(n^2)$$的时间复杂度,遍历两次挨个求和对比

### 前缀和+哈希表优化
思路:将数据保存在哈希表中,减少计算量(重复计算的放在hashmap中)
之所以能放在哈希表中,是因为sum[1:5]=sum[1:2]+sum[2:5]

定义sum[i]=nums[0]+...+nums[i]  
sum[i] = sum[i-1]+nums[i]  
那么子数组(j<=i)的和[j,...,i]=nums[j]+...+nums[i] = sum[i]-sum[j-1]==k   
要使[j,,i]==k,则要求sum[i]-sum[j-1]==k ,即要求**sum[j-1]==sum[i]-k**即可  


所以考虑以i结尾的,和为k的连续子数组的个数只用统计有多少个前缀和等于sum[i]-k的sum[j-1]即可.
建立一个hashmap,以和为key,出现的次数为value,

https://leetcode.cn/problems/QTMn0o/solution/he-wei-k-de-zi-shu-zu-by-leetcode-soluti-1169/
