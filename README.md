* сгенерировать сущности апи из файла open api
 - выполнять из goHomeSystem(root)
 - oapi-codegen -package=http -generate "types,spec,gin" noteBackend/api/http/note-api.yml > noteBackend/api/http/note-api.gen.go
 - oapi-codegen -package=http -generate "types,spec,gin" userService/api/http/user-service-api.yml > userService/api/http/user-service-api.gen.go


go get github.com/Zaytsev-Dmitry/home_system_open_api@v1.0.8
docker-compose build migrator --build-arg CONFIG_PATH_ARG=/app/configs/docker.yaml --build-arg APP_PROFILE_ARG=docker 