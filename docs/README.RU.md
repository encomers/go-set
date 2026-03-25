# go-set

Мощная и производительная реализация множества (Set) на Go с поддержкой дженериков (Go 1.23+). Подходит как для простых задач, так и для функционального стиля программирования и конкурентных сценариев.

## Обзор

`go-set` предоставляет удобный и быстрый способ работы с множествами в Go. Библиотека построена на основе встроенных map и обеспечивает высокую производительность без лишних зависимостей.

**Библиотека не требует внешних зависимостей.**

---

## Возможности

### Базовые операции

* Дженерики: `Set[T comparable]`
* O(1) в среднем:

  * `Add`
  * `Remove`
  * `Contains`
* Минимальные аллокации
* Простая и понятная API

### Функциональный стиль

* `Map`
* `Filter`
* `Reduce`
* `FlatMap`
* `Any`, `All`, `Count`
* `Partition`

### Итераторы (Go 1.23+)

* Поддержка `iter.Seq`
* Ленивые операции:

  * `Iter`
  * `FilterIter`
  * `MapIter`
  * `FlatMapIter`

### Операции над множествами

* `Union`
* `Intersection`
* `Difference`
* `SymmetricDifference`
* `IsSubset`
* `IsSuperset`
* `IsDisjoint`
* `Equals`

### Утилиты

* `Min`, `Max`, `Sum`
* `Sort`
* `ToSlice`, `Copy`, `Clear`

### Конкурентность

* `SyncSet` — потокобезопасная версия
* Использует `sync.RWMutex`
* Подходит для многопоточных сценариев

---

## Установка

```bash
go get github.com/encomers/go-set
```

---

## Требования

* Go **1.23+**

---

## Базовый пример

```go
package main

import (
    "fmt"
    set "github.com/encomers/go-set"
)

func main() {
    s := set.NewSet(1, 2, 3)

    s.Add(4, 5)
    s.Remove(2)

    fmt.Println(s.Contains(3)) // true
    fmt.Println(s.Size())      // 4
}
```

---

## Функциональные операции

```go
s := set.NewSet(1, 2, 3, 4, 5)

// Фильтрация
evens := s.Filter(func(v int) bool {
    return v%2 == 0
})

// Преобразование
mapped := s.Map(func(v int) int {
    return v * 10
})

// Агрегация
sum := s.Reduce(0, func(acc, v int) int {
    return acc + v
})
```

---

## Итерация

```go
for v := range s.Iter() {
    fmt.Println(v)
}
```

### Ленивые итераторы

```go
for v := range s.FilterIter(func(v int) bool {
    return v > 2
}) {
    fmt.Println(v)
}
```

---

## Операции над множествами

```go
a := set.NewSet(1, 2, 3)
b := set.NewSet(3, 4, 5)

union := a.Union(b)
intersection := a.Intersection(b)
difference := a.Difference(b)
```

---

## Утилиты

```go
min := set.Min(s)
max := set.Max(s)
sum := set.Sum(s)
```

---

## Потокобезопасное использование

```go
s := set.NewSyncSet[int]()

s.Add(1, 2, 3)

if s.Contains(2) {
    fmt.Println("exists")
}
```

### Особенности SyncSet

* Безопасен для конкурентного доступа
* Использует read/write lock
* Некоторые методы возвращают обычный Set

---

## Производительность

Бенчмарки (Apple M4, arm64):

```
BenchmarkSetAdd                         ~96 ns/op
BenchmarkSetContains                    ~4.7 ns/op
BenchmarkSetRemove                      ~21 ns/op

BenchmarkSyncSetAdd                     ~140 ns/op
BenchmarkSyncSetContains                ~7.2 ns/op
BenchmarkSyncSetConcurrentAdd           ~129 ns/op
BenchmarkSyncSetConcurrentContains      ~88 ns/op
```

---

## Важно

* `Set` не потокобезопасен
* Используйте `SyncSet` при работе с goroutine
* Порядок элементов не гарантирован
* Не изменяйте множество во время итерации

---

## Лицензия

MIT
