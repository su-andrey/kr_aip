package creators

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePostsTable(DB *pgxpool.Pool) {
	ctx := context.Background()
	tx, err := DB.Begin(ctx)
	if err != nil {
		log.Fatal("Ошибка начала транзакции:", err)
	}
	defer tx.Rollback(ctx)

	var tableExists bool
	err = tx.QueryRow(ctx,
		"SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'posts');").
		Scan(&tableExists)
	if err != nil {
		log.Fatal("Ошибка проверки таблицы posts:", err)
	}

	if !tableExists {
		_, err = tx.Exec(ctx, `
			CREATE TABLE posts (
				id SERIAL PRIMARY KEY,
				category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
				author_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
				name TEXT NOT NULL,
				body TEXT NOT NULL,
				likes INTEGER DEFAULT 0,
				dislikes INTEGER DEFAULT 0
			);
		`)
		if err != nil {
			log.Fatal("Ошибка создания таблицы posts:", err)
		}

		err = tx.Commit(ctx)
		if err != nil {
			log.Fatal("Ошибка фиксации транзакции:", err)
		}

		log.Println("✅ Таблица posts успешно создана!")
	} else {
		tx.Rollback(ctx)
	}
}
