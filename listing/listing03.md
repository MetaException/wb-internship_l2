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
Тип interface хранит в себе тип и значение. При выводе err, будет выведено значение (nil), при сравнении будет, будет сравниваться и тип и значение. А тк. тип - *os.PathError и значение nil, то результат err == nil будет false.
```
