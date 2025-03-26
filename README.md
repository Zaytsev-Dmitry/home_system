* сгенерировать сущности апи из файла open api
 - выполнять из goHomeSystem(root)
 - oapi-codegen -package=generatedApi -generate "types,spec,gin" noteBackend/api/spec/openapi/note-api.yml > noteBackend/api/spec/note-api.gen.go
 - oapi-codegen -package=generatedApi -generate "types,spec,gin" authServer/api/docs/auth-server-api.yml > authServer/api/docs/auth-server-api.gen.go


go get github.com/Zaytsev-Dmitry/home_system_open_api@v1.0.8