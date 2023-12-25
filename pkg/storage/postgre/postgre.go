
package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"GoNews/pkg/storage"
)

// DB - структура, представляющая соединение с базой данных PostgreSQL.
type DB struct {
	conn *sql.DB
}

// NewDB создает новый экземпляр DB.
func NewDB(db *sql.DB) *DB {
	return &DB{conn: db}
}

// Posts возвращает все публикации.
func (db *DB) Posts() ([]storage.Post, error) {
	rows, err := db.conn.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []storage.Post
	for rows.Next() {
		var post storage.Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorName,
			&post.CreatedAt,
			&post.PublishedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// AddPost добавляет новую публикацию.
func (db *DB) AddPost(post storage.Post) error {
	_, err := db.conn.Exec("INSERT INTO posts (title, content, author_id, author_name, created_at, published_at) VALUES ($1, $2, $3, $4, $5, $6)",
		post.Title, post.Content, post.AuthorID, post.AuthorName, post.CreatedAt, post.PublishedAt)
	return err
}

// UpdatePost обновляет существующую публикацию.
func (db *DB) UpdatePost(post storage.Post) error {
	_, err := db.conn.Exec("UPDATE posts SET title=$1, content=$2, author_id=$3, author_name=$4, created_at=$5, published_at=$6 WHERE id=$7",
		post.Title, post.Content, post.AuthorID, post.AuthorName, post.CreatedAt, post.PublishedAt, post.ID)
	return err
}

// DeletePost удаляет публикацию по ID.
func (db *DB) DeletePost(id int) error {
	_, err := db.conn.Exec("DELETE FROM posts WHERE id=$1", id)
	return err
}
