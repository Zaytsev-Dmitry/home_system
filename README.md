* сгенерировать сущности апи из файла open api
 - выполнять из goHomeSystem(root)
 - oapi-codegen -package=http -generate "types,spec,gin" noteBackend/api/http/note-api.yml > noteBackend/api/http/note-api.gen.go
 - oapi-codegen -package=http -generate "types,spec,gin" authBackend/api/http/auth-server-api.yml > authBackend/api/http/auth-server-api.gen.go


go get github.com/Zaytsev-Dmitry/home_system_open_api@v1.0.8