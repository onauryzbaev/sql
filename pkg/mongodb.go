package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"GoNews/pkg/storage"
)

// DB - структура, представляющая соединение с базой данных MongoDB.
type DB struct {
	conn *mongo.Client
}

// NewDB создает новый экземпляр DB.
func NewDB(client *mongo.Client) *DB {
	return &DB{conn: client}
}

// Posts возвращает все публикации.
func (db *DB) Posts() ([]storage.Post, error) {
	collection := db.conn.Database("your_database_name").Collection("posts")

	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var posts []storage.Post
	err = cursor.All(context.Background(), &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// AddPost добавляет новую публикацию.
func (db *DB) AddPost(post storage.Post) error {
	collection := db.conn.Database("your_database_name").Collection("posts")

	_, err := collection.InsertOne(context.Background(), post)
	return err
}

// UpdatePost обновляет существующую публикацию.
func (db *DB) UpdatePost(post storage.Post) error {
	collection := db.conn.Database("your_database_name").Collection("posts")

	filter := bson.M{"_id": post.ID}
	update := bson.M{"$set": post}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

// DeletePost удаляет публикацию по ID.
func (db *DB) DeletePost(id int) error {
	collection := db.conn.Database("your_database_name").Collection("posts")

	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.Background(), filter)
	return err
}
