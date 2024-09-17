package main

import (
	"errors"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	ErrIncorrectInput = errors.New("incorrect input")
)

func Unpack(toUnpack string) (string, error) {

	toAnalysis := []rune(toUnpack)

	var ret strings.Builder

	left, right := 0, 1

	for right <= len(toAnalysis) {

		leftIsLetter := unicode.IsLetter(toAnalysis[left])

		rightIsNumber := false
		if right < len(toAnalysis) {
			rightIsNumber = unicode.IsNumber(toAnalysis[right])
		}

		if leftIsLetter && rightIsNumber {
			toAppend := toAnalysis[left]
			toAppendCount := int(toAnalysis[right] - '0')

			for i := 0; i < toAppendCount; i++ {
				ret.WriteRune(toAppend)
			}
			left += 2
			right += 2
		} else if leftIsLetter && !rightIsNumber {
			ret.WriteRune(toAnalysis[left])
			left++
			right++
		} else {
			return "", ErrIncorrectInput
		}
	}
	return ret.String(), nil
}
