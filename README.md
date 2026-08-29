# Tests For Go Pet Shop

---
## Описание:
Реализованы unit-тесты для HTTP-хендлеров сервиса [Pet Shop](https://github.com/pavloging/go-pet-shop) 
Тесты представлены в двух версиях: 
- `v1` - ручной мок;
- `v2` - go:generate + mockery.
---
## Реализованные тесты
#### GetAllProducts

Проверяемые сценарии:
- успешное получение списка — `200 OK`;
- ошибка сервиса — `500 Internal Server Error`.

#### CreateProduct

Проверяемые сценарии:
- успешное создание продукта — `200 OK`;
- некорректный JSON - `400 Bad Request`;
- ошибка сервиса — `500 Internal Server Error`.

#### UpdateProduct

Проверяемые сценарии:
- успешное изменение продукта — `200 OK`;
- некорректный JSON - `400 Bad Request`;
- ошибка сервиса — `500 Internal Server Error`.

#### DeleteProduct 

Проверяемые сценарии:
- успешное удаление продукта — `200 OK`;
- пустой `id` - `400 Bad Request`;
- ошибка сервиса — `500 Internal Server Error`.
---
## Версии:
### Version 1 — ручной мок 
Использован ручной мок, реализующий интерфейс `Products`

Поведение методов мока задаётся вручную через функции:

- `GetAllProductsFunc`;

- `CreateProductFunc`;

- `UpdateProductFunc`;

- `DeleteProductFunc`.

### Version 2 — mockery
- удален ручной мок;
- добавлена директива `go:generate`  для генерации мока интерфейса `Products`;
- мок генерируется с помощью `mockery`;
- тесты используют сгенерированный мок.

---
## Запуск тестов
Для запуска тестов HTTP-хендлеров:

```
go test ./internal/handlers/product -v
```
Для запуска всех тестов: 
```
go test ./...
```
---
## Автор
[Nonlight](https://github.com/Nonlight)