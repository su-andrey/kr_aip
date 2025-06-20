package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/su-andrey/kr_aip/config"
	"github.com/su-andrey/kr_aip/creators"
	"go.uber.org/zap"
)

var DB *pgxpool.Pool

func ConnectDB() {
	cfg := config.LoadConfig()

	if cfg.DbUrl == "" {
		config.Logger.Fatal("DB_URL не найден в .env файле") // логируем критические ошибки
	}

	// Контекст с таймаутом 5 секунд для подключения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Подключение к БД
	pool, err := pgxpool.New(ctx, cfg.DbUrl)
	if err != nil {
		config.Logger.Fatal("Ошибка подключения к базе данных:", zap.Error(err)) // логируем критические ошибки
	}

	DB = pool
	config.Logger.Info("✅ Успешное подключение к PostgreSQL") // Записываем сообщение об успехи

	creators.Migrate(DB) // проверяем существование таблиц и создаем их, если их нет
}
