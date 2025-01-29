* сгенерировать сущности апи из файла open api
 - выполнять из goHomeSystem(root)
 - oapi-codegen -package=generatedApi -generate "types,spec,gin" noteBackendApp/api/docs/note-api.yml > noteBackendApp/api/docs/note-backend-api.gen.go
 - oapi-codegen -package=generatedApi -generate "types,spec,gin" authServer/api/docs/auth-server-api.yml > authServer/api/docs/auth-server-api.gen.go

TODO:
1) Убрать gorm и написать свой клиент используя (https://github.com/jackc/pgx или https://github.com/jmoiron/sqlx) - DONE
2) Допилить клиент для свагера
3) Создать проект для миграций и отдельно запускать его в докере
4) Вынос общего кода в отдельную либу (возможно стоит вынести все yml файлы и кодогенерацию)
5) Валидация данных в handlers
6) Посмотреть на https://github.com/ogen-go/ogen возможно он лучше

Телеграмм бот
1) сохранение телеграм пользака через authserver - DONE
2) вынесение всех openapi в отдельную импортируемую либу - DONE
3) валидация данных
4) регистрация client id + хождения во все сервисы с токеном
5) настраивание менюшек для каждого пользака
6) отловить везде ошибки
7) валидация данных
8) реализация сохранения записок
9) удаление всех старых сообщений после того как пользак апрувнул регистрацию - DONE
10) вынести весь утилитный класс в библиотеку
11) Удалить /start после того как пользак апрувнул инфу