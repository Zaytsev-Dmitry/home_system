* сгенерировать сущности апи из файла open api
 - cd .. to noteBackendApp
 - oapi-codegen -package=generatedApi -generate "types,spec,gin" api/openapi/note-api.yml > api/openapi/api.gen.go 