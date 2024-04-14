-- +goose Up

CREATE TABLE books (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    volume INT NOT NULL default 0,
    category TEXT NOT NULL,
    author TEXT NOT NULL,
    published_at TIMESTAMP NOT NULL,
    publisher TEXT NOT NULL,
    finished BOOL NOT NULL default false,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    constraint check_category check (category in ('Manga', 'Light Novel', 'Art Book', 'Other')) 
);

-- +goose Down
DROP TABLE books;