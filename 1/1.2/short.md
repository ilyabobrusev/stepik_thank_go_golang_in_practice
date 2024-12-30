```
package main

import "fmt"

func main() {
    fmt.Println("hello everyone")
}
```

```
go run hello.go
go build hello.go
./hello
```

```
Целые и дробные числа:

fmt.Println("1+1 =", 1+1)
// 1+1 = 2

fmt.Println("5.0/2.0 =", 5.0/2.0)
// 5.0/2.0 = 2.5

Логические значения и операторы:

fmt.Println(true && false)
// false

fmt.Println(true || false)
// true

fmt.Println(!true)
// false

Строки можно складывать через +:

fmt.Println("go" + "lang")
// golang
```


```
var объявляет переменную, = инициализирует конкретным значением:

var b bool = true
fmt.Println(b)
// true

var s string = "hello"
fmt.Println(s)
// hello

var i int = 42
fmt.Println(i)
// 42

var f float64 = 12.34
fmt.Println(f)
// 12.34

Можно объявить сразу несколько:

var one, two int = 1, 2
fmt.Println(one, two)
// 1 2

Если инициализировать переменную при объявлении, тип можно и не указывать. Go сам догадается:

var sunny = true
fmt.Printf("%v is %T\n", sunny, sunny)
// true is bool

fmt.Printf() подставляет переменные в строку по шаблону. %v — формат по умолчанию, %T — тип переменной.

Если не инициализировать переменную при объявлении, она получит нулевое значение (zero value). У каждого типа оно свое: int — 0, string — "", bool — false.

var num int
var str string
var ok bool

fmt.Printf("%#v %#v %#v\n", num, str, ok)
// 0 "" false

Оператор := объявляет и инициализирует переменную. var и тип можно не указывать:

food := "apple"  // var food string = "apple"
fmt.Println(food)
// apple
```

```
const объявляет константу:

const s string = "constant"
fmt.Println(s)
// constant

const n = 500000000
fmt.Println(n)
// 500000000

const ch = 'a'
fmt.Println(ch)
// 97
// это числовой код латинского символа 'a'
```

```
Циклы

for — единственная конструкция для циклов в Go. Вот несколько примеров.

С одним условием, аналог while в питоне:

i := 1
for i <= 3 {
    fmt.Println(i)
    i = i + 1
}
// 1
// 2
// 3

Классический for из трех частей (инициализация, условие, шаг):

for j := 7; j <= 9; j++ {
    fmt.Println(j)
}
// 7
// 8
// 9

Цикл от 0 до n-1 (Go 1.22+):

const n = 10
for i := range n {
    fmt.Print(i, " ")
}
// 0 1 2 3 4 5 6 7 8 9

Если хотим повторить определенное действие n раз, а конкретное значение счетчика цикла не интересно (Go 1.22+):

const n = 10
for range n {
    fmt.Print(".")
}
// ..........

 Бесконечный цикл, выполняется до break для выхода из цикла или return для выхода из функции:

for {
    fmt.Println("loop")
    break
}
// loop

continue переходит к следующей итерации цикла:

for n := 0; n <= 5; n++ {
    if n%2 == 0 {
        continue
    }
    fmt.Println(n)
}
// 1
// 3
// 5
```

```
If/else

Конструкция if-else в Go ведет себя без особых сюрпризов.

Вокруг условия не нужны круглые скобки, но фигурные скобки для веток обязательны:

if 7%2 == 0 {
    fmt.Println("7 is even")
} else {
    fmt.Println("7 is odd")
}
// 7 is odd

Можно использовать if без else:

if 8%4 == 0 {
    fmt.Println("8 is divisible by 4")
}
// 8 is divisible by 4

Единственный нюанс: перед условием может идти выражение. Объявленные в нем переменные доступны во всех ветках:

if num := 9; num < 0 {
    fmt.Println(num, "is negative")
} else if num < 10 {
    fmt.Println(num, "has 1 digit")
} else {
    fmt.Println(num, "has multiple digits")
}
// 9 has 1 digit
```

```
Switch

switch описывает условие с множеством веток.

В отличие от многих других языков, не нужно указывать break. Go выполняет только подходящую ветку и не проваливается в следующую:

i := 2
fmt.Print("Write ", i, " as ")
switch i {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
case 3:
    fmt.Println("three")
}
// Write 2 as two

Чтобы провалиться в следующую ветку, есть специальное ключевое слово fallthrough. Его можно использовать только в качестве последней инструкции в блоке:

i := 2
fmt.Print("Write ", i, " as ")
switch i {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
    fallthrough
case 3:
    fmt.Println("Bye-bye")
}
// Write 2 as two
// Bye-bye

В одной ветке можно указать несколько выражений. Ветка default сработает, если остальные не подошли:

// одной строкой инициализировали day
// и сразу сделали switch по нему
switch day := time.Now().Weekday(); day {
case time.Saturday, time.Sunday:
    fmt.Println(day, "is a weekend")
default:
    fmt.Println(day, "is a weekday")
}
// Tuesday is a weekday

Выражения в ветках не обязательно должны быть константами. switch может работать как if:

t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("It's before noon")
default:
    fmt.Println("It's after noon")
}
// It's before noon
```