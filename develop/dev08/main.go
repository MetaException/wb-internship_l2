package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import "dev08/cmd"

func main() {
	cmd.Execute()
}
