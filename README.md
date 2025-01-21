* сгенерировать сущности апи из файла open api
 - cd .. to noteBackendApp
 - oapi-codegen -package=generatedApi -generate "types,spec,gin" api/docs/note-api.yml > api/docs/api.gen.go 