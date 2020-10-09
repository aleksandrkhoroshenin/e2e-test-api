# Описание задачи

REST сервис, по удалению клиентов из базы по префиксу 

```json
{
  "count": 3,
  "ids": [1,2,3]
}
```  
где `count` - число удаленных клиентов, `ids` - список их идентификаторов.

# Настройка проекта

docker-compose up --force-recreate

docker-compose down -v

Установка golang - https://golang.org/doc/install

Все необходимые команды и параметры подключения к БД находятся в `Makefile`

- Запуск миграции для создания таблицы
```shell script
make migrate
``` 
- Запуск сервиса
```shell script
make run
```
- Откатить миграцию
```shell script
make migrate_down
```

# TODO::
Добавить в компоуз интеграционные тесты