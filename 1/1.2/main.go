package main

import "fmt"

func main() {
    fmt.Println("hello everyone")
}

//
package main

import (
	"fmt"
	"time"
)

func main() {
	// считываем временной отрезок из os.Stdin
	// гарантируется, что значение корректное
	// не меняйте этот блок
	var s string
	fmt.Scan(&s)
	d, _ := time.ParseDuration(s)

	// выведите исходное значение
	// и количество минут в нем
	// в формате "исходное = X min"
	// используйте метод .Minutes() объекта d
	fmt.Println(s, "=", d.Minutes(), "min")
}

//
package main

import (
	"fmt"
	"math"
)

func main() {
	// объявите переменные x1, y1, x2, y2 типа float64
	var x1, y1, x2, y2 float64

	// считываем числа из os.Stdin
	// гарантируется, что значения корректные
	// не меняйте этот блок
	fmt.Scan(&x1, &y1, &x2, &y2)

	// рассчитайте d по формуле эвклидова расстояния
	// используйте math.Pow(x, 2) для возведения в квардрат
	// используйте math.Sqrt(x) для извлечения корня
    d := math.Sqrt(math.Pow(x1 - x2, 2) + math.Pow(y1 - y2, 2))

	// выводим результат в os.Stdout
	fmt.Println(d)
}

//
package main

import (
	"fmt"
)

func main() {
	var source string
	var times int
	// гарантируется, что значения корректные
	fmt.Scan(&source, &times)

	// возьмите строку `source` и повторите ее `times` раз
	// запишите результат в `result`
    var result string
	for range times {
		result += source
	}

	fmt.Println(result)
}

//
package main

import (
	"fmt"
)

func main() {
	var code, lang string
	fmt.Scan(&code)

	// определите полное название языка по его коду
	// и запишите его в переменную `lang`
	// var lang string
	switch code {
	case "en":
		lang = "English"
	case "fr":
		lang = "French"
	case "ru", "rus":
		lang = "Russian"
	default:
		lang = "Unknown"
	}

    fmt.Println(lang)
}

//
