CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author_id INTEGER NOT NULL,
    author_name VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL,
    published_at BIGINT NOT NULL
);
