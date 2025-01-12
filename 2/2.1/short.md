```
Показать основной модуль и его зависимости:
go list -m all

Посмотреть все версии конкретного модуля:
go list -m -versions github.com/sahilm/fuzzy

Обновить конкретный модуль на последнюю patch-версию в пределах действующей minor-версии (например, 1.1.0 → 1.1.5):
go get -u=patch github.com/sahilm/fuzzy

Обновить конкретный модуль на последнюю minor-версию в пределах действующей major-версии (1.1.0 → 1.2.1):
go get -u github.com/sahilm/fuzzy

Удалить неактуальные зависимости из go.mod:
go mod tidy
```