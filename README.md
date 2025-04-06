[![Note backend App](https://github.com/Zaytsev-Dmitry/home_system/actions/workflows/noteBackendAction.yml/badge.svg?branch=dev)](https://github.com/Zaytsev-Dmitry/home_system/actions/workflows/noteBackendAction.yml)

* сгенерировать сущности апи из файла open api
 - выполнять из goHomeSystem(root)
 - oapi-codegen -package=generatedApi -generate "types,spec,gin" noteBackendApp/api/docs/note-api.yml > noteBackendApp/api/docs/note-backend-api.gen.go
 - oapi-codegen -package=generatedApi -generate "types,spec,gin" authServer/api/docs/auth-server-api.yml > authServer/api/docs/auth-server-api.gen.go


go get github.com/Zaytsev-Dmitry/home_system_open_api@v1.0.8
