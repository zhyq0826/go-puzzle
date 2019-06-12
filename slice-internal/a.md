输出如下

```
[a b new]
[a b new]
```

由于 str1 和 str2 共享底层 array，对 str2 的修改会影响 str1，append 操作会判断 str2 的 capcity 不够用，从而会重新申请新的空间，所以 append 只会 str2 和 str1 是用的两个不同的 array