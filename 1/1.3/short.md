```
Массивы

Массив в Go — это нумерованная последовательность элементов. Длина массива известна заранее и зафиксирована.

var arr [5]int
fmt.Println("empty:", arr)
// empty: [0 0 0 0 0]

arr[4] = 100
fmt.Println("set:", arr)
// set: [0 0 0 0 100]
fmt.Println("get:", arr[4])
// get: 100

fmt.Println("len:", len(arr))
// len: 5

arr := [5]int{1, 2, 3, 4, 5}
fmt.Println("init:", arr)
// init: [1 2 3 4 5]

var arr [2][3]int
for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
        arr[i][j] = i + j
    }
}
fmt.Println("2d:", arr)
// 2d: [[0 1 2] [1 2 3]]
```

``
Срезы

Срез (slice) — ключевая структура данных в Go. Это массив изменяемой длины, как list в питоне или Array в js. Обычно в программах на Go оперируют именно срезами, «чистые» массивы встречаются намного реже.

s := make([]string, 3)
fmt.Printf("empty: %#v\n", s)
// empty: []string{"", "", ""}
Шаблон %#v возвращает «внутреннее представление» значения, примерно как repr() в питоне.

s[0] = "a"
s[1] = "b"
s[2] = "c"

fmt.Println("set:", s)
// set: [a b c]

fmt.Println("get:", s[2])
// get: c

len() возвращает длину среза:

fmt.Println("len:", len(s))
// len: 3

s := []string{"a", "b", "c"}
fmt.Println("init:", s)
// init: [a b c]

В отличие от массива, в срез можно добавлять новые элементы через встроенную функцию append(). Функция возвращает новый срез:

s := []string{"a", "b", "c"}

fmt.Println("src:", s)
// src: [a b c]

s = append(s, "d")
s = append(s, "e", "f")
fmt.Println("upd:", s)
// upd: [a b c d e f]

src := []string{"a", "b", "c", "d", "e", "f"}
dst := make([]string, len(src))

copy(dst, src)
fmt.Println("copy:", dst)
// copy: [a b c d e f]

Срезы поддерживают... срезы (отсюда их название). Выражение slice[from:to] вернет срез от элемента с индексом from включительно до элемента с индексом to не включительно:

s := []string{"a", "b", "c", "d", "e", "f"}

sl1 := s[2:5]
fmt.Println("sl1:", sl1)
// sl1: [c d e]

Этот срез включает все элементы, кроме s[5]:

sl2 := s[:5]
fmt.Println("sl2:", sl2)
// sl2: [a b c d e]

А этот срез включает элементы от s[2] и до конца:

sl3 := s[2:]
fmt.Println("sl3:", sl3)
// sl3: [c d e f]
```