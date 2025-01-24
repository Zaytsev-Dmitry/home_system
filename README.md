* сгенерировать сущности апи из файла open api
 - выполнять из goHomeSystem(root)
 - oapi-codegen -package=generatedApi -generate "types,spec,gin" noteBackendApp/api/docs/note-api.yml > noteBackendApp/api/docs/note-backend-api.gen.go
 - oapi-codegen -package=generatedApi -generate "types,spec,gin" authServer/api/docs/auth-server-api.yml > authServer/api/docs/auth-server-api.gen.go

TODO:
1) Регистрация пользака в [authServer](authServer) + сохранение пользака в базу
2) Допилить клиент для свагера
3) Клиент для БД (gorm) + сделать на интерфейсах
4) Создание базы для notebackend (на миграциях)
5) Вынос общего кода в отдельную либу (возможно стоит вынести все yml файлы и кодогенерацию)
6) Телеграмм бот для создания записок
7) Валидация данных в handlers