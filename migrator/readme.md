## Install golang-migrate
```bash
    brew install golang-migrate
```
## Generate new migration
```bash
    migrate create -ext sql -seq {migration_name}
```
## Run in docker local machine
```bash
    docker-compose build migrator --build-arg CONFIG_PATH_ARG=/app/config/docker.yaml --build-arg APP_PROFILE=docker
    docker-compose build up migrator
```