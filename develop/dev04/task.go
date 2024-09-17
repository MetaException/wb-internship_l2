package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fmt.Println(FindAnagram([]string{"пятак", "пятка", "тяпка", "тяпка", "листок", "слиток", "столик"}))
}

func FindAnagram(words []string) *map[string][]string {

	usedWords := make(map[string]struct{})
	keyMap := make(map[string]string)
	result := make(map[string][]string)

	for _, v := range words {
		lowered := strings.ToLower(v)
		if _, ok := usedWords[lowered]; !ok {
			toSort := []rune(lowered)
			slices.Sort(toSort)

			sortedWord := string(toSort)

			usedWords[lowered] = struct{}{}

			if _, ok := keyMap[sortedWord]; !ok {
				keyMap[sortedWord] = lowered
			}
			result[keyMap[sortedWord]] = append(result[keyMap[sortedWord]], lowered)
		}
	}

	for key, val := range result {
		if len(val) == 1 {
			delete(result, key)
			continue
		}
		sort.Strings(result[key])
	}

	return &result
}
