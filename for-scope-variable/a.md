输出如下:
```
0
1
2
map[zhou:0xc000086020 li:0xc000086020 wang:0xc000086020]
map[li:0xc000088018 wang:0xc000088030 zhou:0xc000088000]
```

for 循环是块变量作用域，stu 引用的是最后一个变量实体，所以 &stu 实际上取的是最后一个的地址，应该改用 index 的方式取