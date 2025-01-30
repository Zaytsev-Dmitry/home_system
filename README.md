* сгенерировать сущности апи из файла open api
 - выполнять из goHomeSystem(root)
 - oapi-codegen -package=generatedApi -generate "types,spec,gin" noteBackendApp/api/docs/note-api.yml > noteBackendApp/api/docs/note-backend-api.gen.go
 - oapi-codegen -package=generatedApi -generate "types,spec,gin" authServer/api/docs/auth-server-api.yml > authServer/api/docs/auth-server-api.gen.go

TODO:
1) Убрать в authBackend gorm и написать свой клиент используя (https://github.com/jackc/pgx или https://github.com/jmoiron/sqlx) - DONE
2) Допилить клиент для свагера
3) Создать проект для миграций и отдельно запускать его в докере
4) Вынос общего кода в отдельную либу (возможно стоит вынести все yml файлы и кодогенерацию)
5) Валидация данных в authBackend
6) Посмотреть на https://github.com/ogen-go/ogen возможно он лучше

Телеграмм бот
1) сохранение телеграм пользака через authserver - DONE
2) вынесение всех openapi в отдельную импортируемую либу - DONE
3) валидация данных при регистрации
4) регистрация client id + хождения во все сервисы с токеном
5) настраивание менюшек для каждого пользака
6) отловить везде ошибки
7) валидация данных в записках
8) реализация сохранения записок
9) удаление всех старых сообщений после того как пользак апрувнул регистрацию - DONE
10) вынести весь утилитный класс в библиотеку
11) Удалить /start после того как пользак апрувнул инфу - DONE
12) Добавить обучалку после регистрации пользака (слайдер - https://github.com/go-telegram/ui/blob/main/slider/readme.md) - DONE
13) Реализация /profile
14) добавление записки
15) добавить telegram bot backend в docker compose 
16) сделать меню для /notes - DONE
17) сделать меню для /profile - DONE
18) Написать свой пагинатор