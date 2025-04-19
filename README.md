# Go Task Service

Go Task Service — это простой HTTP API для управления долгими I/O bound задачами. Сервис позволяет создавать задачи, выполнение которых занимает несколько минут, и получать их результат по завершении.

## Структура проекта

```
go-task-service
├── cmd
│   └── main.go          # Точка входа приложения
├── internal
│   ├── api
│   │   └── handler.go   # HTTP-обработчики API
│   ├── tasks
│   │   ├── manager.go   # Логика управления задачами
│   │   └── types.go     # Структуры данных задач
│   └── storage
│       └── memory.go    # In-memory хранилище задач (не используется напрямую в текущей реализации)
├── go.mod               # Модуль и зависимости
├── go.sum               # Контрольные суммы зависимостей
└── README.md            # Документация
```

## Быстрый старт

### Требования

- Go 1.18 или новее

### Установка

1. Клонируйте репозиторий:
   ```
   git clone https://github.com/yourusername/go-task-service.git
   cd go-task-service
   ```

2. Установите зависимости:
   ```
   go mod tidy
   ```

### Запуск сервиса

Запустите сервис командой:

```
go run cmd/main.go
```

Сервер стартует на `localhost:8082`.

## API

### Создать задачу

- **POST** `/tasks`
- Тело запроса: `{}` (параметры пока не используются)
- Ответ:
  ```json
  {
    "task_id": "some-uuid"
  }
  ```

### Получить результат задачи

- **GET** `/tasks/{id}`
- Ответ:
  ```json
  {
    "id": "some-uuid",
    "result": "Task result for ID: some-uuid"
  }
  ```
  Если задача не найдена — HTTP 404.

## Примеры использования

1. Создать задачу:
   ```
   curl -X POST http://localhost:8082/tasks -d '{}'
   ```

2. Получить результат:
   ```
   curl http://localhost:8082/tasks/<task_id>
   ```# API_io_bounds
