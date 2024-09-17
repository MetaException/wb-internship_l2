package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Преимущества:
		* Избавляет класс от привязки к конкретным классам продуктов
		* Выделяет код производства продуктов в одно место, упрощая поддержку кода
		* Упрощает добавление новых продуктов в программу
		* Реализует принцип открытости/закрытости

	Недостатки:
		* Может привести к созданию больших параллельных иерархий классов, тк. для каждого класса продукта надо создать свой подкласс создателя

*/

type IParser interface {
	Parse(html string)
}

type TimeTableParser struct{}

func (p TimeTableParser) Parse(html string) {
	fmt.Println("Parsing timetable:", html)
}

type PersonParser struct{}

func (p PersonParser) Parse(html string) {
	fmt.Println("Parsing person:", html)
}

type Creator interface {
	FactoryMethod() IParser
}

type TimeTableParserCreator struct{}

func (c TimeTableParserCreator) FactoryMethod() IParser {
	return TimeTableParser{}
}

type PersonParserCreator struct{}

func (c PersonParserCreator) FactoryMethod() IParser {
	return PersonParser{}
}

/*
func main() {
	ttCreator := TimeTableParserCreator{}
	timeTableParser := ttCreator.FactoryMethod()
	timeTableParser.Parse("HTML")

	pCreator := PersonParserCreator{}
	personParser := pCreator.FactoryMethod()
	personParser.Parse("HTML")
}
*/
