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
    docker-compose build migrator && docker-compose up -d migrator && docker image prune -a -f
```