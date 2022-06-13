# 3 лабораторна робота

Для запуску треба створити файл config.json у корені проєкту з такою структурою:

```
{
  "db": {
    "name": "db-name",
		"user": "db-user",
		"password": "db-password",
		"host": "(addr:port)"
  },
  "apiUrl": "url for api"
}
```

- Сервер запускається командою
```shell script
go run ./cmd/server
```

- Клієнт запускається командою
```shell script
go run ./cmd/client
```