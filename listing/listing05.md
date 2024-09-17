Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Программа выведет "error". Тк. интерфейс хранит в себе тип и значение. Чтобы сработало условие err == nil, и тип и значение должны быть nil. Но в нашем случае тип будет *customError, тк. реализует интерфейс Error, хоть значение и nil.
```
