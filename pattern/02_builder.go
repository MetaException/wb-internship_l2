package pattern

import (
	"os"
	"time"
)

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Преимущества:
		* Позволяет создавать объекты пошагово
		* Позволяет использовать один и тот же код для разных продуктов
		* Изолирует код сборки продукта от его бизнес логики

	Недостатки:
		* Усложняет код из-за дополонительных структур
		* Клиент будет привязан к конкретным структурам строителей

	Пример:
		* Построение SQL запросов
		* Инициализация игровых персонажей
		* Создание объектов в тестировании
*/

type Logger struct {
	pattern      string
	rollingTime  time.Time
	dist         os.File
	limit        int
	minimumLevel string
}

type LoggerBuilder interface {
	SetRollingTime(time time.Time) LoggerBuilder
	SetPattern(pattern string) LoggerBuilder
	SetOutput(dist os.File) LoggerBuilder
	SetLogLimit(limit int) LoggerBuilder
	SetMinimumLevel(level string) LoggerBuilder
	Build() Logger
}

type ConcreteLoggerBuilder struct {
	logger *Logger
}

func NewConcreteLoggerBuilder() *ConcreteLoggerBuilder {
	return &ConcreteLoggerBuilder{
		logger: &Logger{},
	}
}

func (log *ConcreteLoggerBuilder) SetRollingTime(time time.Time) LoggerBuilder {
	log.logger.rollingTime = time
	return log
}

func (log *ConcreteLoggerBuilder) SetPattern(pattern string) LoggerBuilder {
	log.logger.pattern = pattern
	return log
}

func (log *ConcreteLoggerBuilder) SetOutput(dist os.File) LoggerBuilder {
	log.logger.dist = dist
	return log
}

func (log *ConcreteLoggerBuilder) SetLogLimit(limit int) LoggerBuilder {
	log.logger.limit = limit
	return log
}

func (log *ConcreteLoggerBuilder) SetMinimumLevel(level string) LoggerBuilder {
	log.logger.minimumLevel = level
	return log
}

func (b *ConcreteLoggerBuilder) Build() Logger {
	return *b.logger
}

/*
func main() {
	builder := NewConcreteLoggerBuilder()

	logger := builder.SetRollingTime(time.Now()).
		SetPattern("json").
		SetLogLimit(1024).
		SetMinimumLevel("INFO").
		Build()

	fmt.Println(logger)
}

*/
