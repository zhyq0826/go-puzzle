struct 中的方法不能被继承，案例中的访问的方法实际上是被内嵌的 People 这个实例所持有，所以输出是：
```shell
showA
showB
```