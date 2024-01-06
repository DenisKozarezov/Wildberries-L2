Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
type EmptyInterface interface {
    type *Type => os.PathError
    data *Data => nil
}
1) Ответ: <nil> - поскольку выводит поле data, однако указатель на статический тип равен os.PathError
2) Ответ: false - поскольку сам по себе результат функции Foo не равен nil, а только поле data

```