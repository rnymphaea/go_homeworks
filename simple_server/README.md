# Simple server

Simple server предсатвляет собой простейший сервер, позволяющий получать информацию о песнях / добавлять новые

## Запуск сервера
```bash
cd simple_server/
make
```

### Очистка локальной базы даных
```bash
cd simple_server
make cleandb
```

## Тестирование
**Доступно после запуска сервера :)**
- получение списка песен:
```bash
curl http://localhost:8080:server/songs
```
- добавление песни в базу данных:
```bash
curl -H 'Content-type: application/json' --data {"title": <название>, "artist": <автор>}
```
