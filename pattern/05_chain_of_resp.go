package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Преимущества:
		* Уменьшает зависимость между клиентом и обработчиками
		* Реализует принцип единственной обязанности
		* Реализует принцип открытости/закрытости

	Недостатки:
		* Запрос может остаться никем не обработанным

	Применение:
		* Система обработка запросов
*/

type Handler interface {
	SetNext(Handler)
	Handle(request string)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) Handle(request string) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

type LogHandler struct {
	BaseHandler
}

func (h LogHandler) Handle(request string) {
	if request == "ABC" {
		fmt.Println("log handler is working")
	} else {
		h.BaseHandler.Handle(request)
	}
}

type AuthHandler struct {
	BaseHandler
}

func (h AuthHandler) Handle(request string) {
	if request == "auth" {
		fmt.Println("auth is working")
	} else {
		h.BaseHandler.Handle(request)
	}
}

/*
func main() {
	logHandler := &LogHandler{}
	authHandler := &AuthHandler{}

	logHandler.SetNext(authHandler)

	logHandler.Handle("ABC")
	logHandler.Handle("auth")
	logHandler.Handle("qwerty")
}
*/
