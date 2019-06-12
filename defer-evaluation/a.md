A deferred function's arguments are evaluated when the defer statement is evaluated.
deferred 的函数参数再 defer 时已经开始计算了，所以 calc("10", a, b) 先开始计算