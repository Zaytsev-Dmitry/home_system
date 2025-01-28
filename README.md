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
6) Продолжить пилить tg бот (сохранение телеграм пользака через authserver, регистрация client id telegram client в keycloak)
7) Посмотреть на https://github.com/ogen-go/ogen возможно он лучше