package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	Преимущества:
		* Упрощает добавление операций, работающих со слонжными структурами объектов
		* Объединяет родственные опреации в одном классе
		* Посетитетль может накапливать состояние при обходе структуры элементов

	Недостатки:
		* Паттерн не оправдан, если иерархия элементов часто меняется
		* Может привести к нарушению инкапсуляции элемнтов

	Пример:
		* Создание отчётов о разных типах транзакции
		* Обработка разных типов файлов
		* Анализ документов разных форматов
*/

type IDocumentVisitor interface {
	VisitPdf(pdfDoc PDFDocument)
	VisitWord(wordDoc WordDocument)
}

type DocumentInfoPrinter struct {
}

func (sc DocumentInfoPrinter) VisitPDF(pdfDoc PDFDocument) {
	// Получение и вывод информации о pdf
}

func (sc DocumentInfoPrinter) VisitWord(pdfWord WordDocument) {
	// Получение и вывод информации о word
}

type DocumentExporter struct {
}

func (sc DocumentExporter) VisitPDF(pdfDoc PDFDocument) {
	// Экспорт pdf
}

func (sc DocumentExporter) VisitWord(pdfWord WordDocument) {
	// Экспорт word
}

type IDocument interface {
	Accept(visitor IDocumentVisitor)
}

type PDFDocument struct {
}

func (pdfDoc PDFDocument) Accept(visitor IDocumentVisitor) {
	visitor.VisitPdf(pdfDoc)
}

type WordDocument struct {
}

func (wordDoc WordDocument) Accept(visitor IDocumentVisitor) {
	visitor.VisitWord(wordDoc)
}

/*
func main() {
	smths := make([]IAbstractSmth, 0)

	smthCalc := smthCalc{}
	smthCalc1 := smthCalc1{}

	for _, smth := range smths {
		smth.Accept(smthCalc1)
		smth.Accept(smthCalc)
	}
}
*/
