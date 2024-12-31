```
var str string = "Hello, World!"

str := "Hello, World!"

var str string
str = "Hello, World!"
```

```
var a string = "Hello "
var b string = "world!"

var c string = a + b  // Hello world!	
```

```
var s string = "Basketball"
a := string(s[0])  // "B"
b := string(s[9])  // "l"
var count int = len(s)      // 10 
fmt.Println("Последний символ равен" + string(s[count - 1]))  // 'l'
```

