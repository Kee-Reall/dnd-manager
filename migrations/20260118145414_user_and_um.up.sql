CREATE TABLE IF NOT EXISTS user_role(
    id INTEGER PRIMARY KEY,
    name TEXT not null
);

CREATE TABLE IF NOT EXISTS users(
    id BLOB PRIMARY KEY CHECK (length(id) = 16),
    name TEXT NOT NULL,
    role INTEGER NOT NULL REFERENCES user_role(id)
);

CREATE TABLE IF NOT EXISTS user_markers(
    user_id BLOB NOT NULL CHECK ( length(user_id) = 16) REFERENCES users(id) ON DELETE cascade,
    id TEXT NOT NULL,
    tag TEXT NOT NULL,
    PRIMARY KEY(id, tag)
);

insert into user_role(id, name)
values (0, 'NoRole'),
       (1, 'Player'),
       (2, 'Master'),
       (3, 'Admin');