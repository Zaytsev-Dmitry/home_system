package main

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"migrator/config"
	"path/filepath"
)

func main() {
	cfg := config.LoadConfig()
	for _, database := range cfg.Databases {
		db, err := sql.Open(
			"postgres",
			"postgres://"+database.User+":"+database.Password+"@"+cfg.App.DbUrl+"/"+database.Name+"?sslmode=disable",
		)
		if err != nil {
			log.Fatal(err)
		}
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			log.Fatal(err)
		}

		migrationsPath, err := filepath.Abs("../migrator/migrations/" + database.DirectoryName)
		if err != nil {
			log.Fatal("Директория с миграциями не найдена:", err)
		}
		m, err := migrate.NewWithDatabaseInstance(
			"file://"+migrationsPath,
			"postgres", driver)
		if err != nil {
			log.Fatal(err)
		}

		// Выводим текущую версию перед миграцией
		version, dirty, err := m.Version()
		if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
			log.Fatal("Ошибка получения версии миграции:", err)
		}
		log.Print("------------------------------------------------------")
		log.Printf("Текущая версия БД: %d (dirty: %v)", version, dirty)

		log.Println("Начинаем выполнение миграций...")
		err = m.Up()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal("Ошибка выполнения миграций:", err)
		}

		// Выводим версию после миграции
		version, dirty, err = m.Version()
		if err != nil {
			log.Fatal("Ошибка получения версии миграции после обновления:", err)
		}
		log.Printf("Миграция завершена. Текущая версия БД: %d (dirty: %v)", version, dirty)
		log.Println("Миграции успешно применены для базы данных:", database.Name)
	}
}
