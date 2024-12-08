# \l , \c snippetdb, \dt : check tables

CREATE TABLE todos (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    title VARCHAR(50) NOT NULL,
    status VARCHAR(15) NOT NULL DEFAULT 'INCOMPLETE',
    expires_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE todos
    ADD COLUMN archive BOOLEAN NOT NULL DEFAULT FALSE;

CREATE INDEX idx_todos_created ON todos(created_at);

INSERT INTO todos (title, status, expires_at) VALUES 
    ('Learn SQL', 'INCOMPLETE', '2024-11-18 15:45:00'),
    ('Build Todo App', 'INCOMPLETE', '2024-12-01 12:00:00'),
    ('Review PostgreSQL Guide', 'COMPLETE', '2024-11-20 09:30:00');


SELECT * FROM todos;