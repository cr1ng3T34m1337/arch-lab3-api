# 3 лабораторна робота

Для запуску треба створити файл `config.json` у корені проєкту з такою структурою:

```json
{
  "db": {
    "name": "db-name",
    "user": "db-user",
    "password": "db-password",
    "host": "(addr:port)"
  },
  "api": {
    "addr": "url-address",
    "port": "port to run on"
  }
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