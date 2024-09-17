package pattern

import (
	"fmt"
)

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Преимущества:
		* Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно вызывают
		* Позволяет реализовать простую отмену и повтор операций
		* Позволяет реализовать отложенный запуск операций
		* Позволяет собирать сложные команды из простых
		* Реализует принцип отрытости/закрытости
*/

type ICommand interface {
	Execute()
}

type PrintInfoCommand struct {
	info     string
	receiver *Receiver
}

func (c *PrintInfoCommand) Execute() {
	fmt.Println(c.info)
	c.receiver.Action()
}

func (c *PrintInfoCommand) SetReceiver(receiver *Receiver) {
	c.receiver = receiver
}

type PrintPICommand struct {
	pi       float64
	receiver *Receiver
}

func (c *PrintPICommand) Execute() {
	fmt.Printf("PI value: %f\n", c.pi)
	c.receiver.Action()
}

func (c *PrintPICommand) SetReceiver(receiver *Receiver) {
	c.receiver = receiver
}

type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Performing action in Receiver")
}

type Invoker struct {
	command ICommand
}

func (i *Invoker) SetCommand(command ICommand) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

/*
func main() {
	rec := &Receiver{}

	printInfoCommand := &PrintInfoCommand{
		info: "Hello 123",
	}
	printInfoCommand.SetReceiver(rec)

	printPICommand := &PrintPICommand{
		pi: 3.14159,
	}
	printPICommand.SetReceiver(rec)

	invoker := &Invoker{}

	// Выполнение команды PrintInfoCommand
	invoker.SetCommand(printInfoCommand)
	invoker.ExecuteCommand()

	// Выполнение команды PrintPICommand
	invoker.SetCommand(printPICommand)
	invoker.ExecuteCommand()
}
*/
