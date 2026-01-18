CREATE TABLE IF NOT EXISTS editions
(
    id   BLOB PRIMARY KEY CHECK (length(id) = 16),
    name TEXT not null UNIQUE
);

CREATE TABLE IF NOT EXISTS status_enum
(
    id   INTEGER,
    type TEXT NOT NULL,
    name TEXT NOT NULL,
    PRIMARY KEY (id, type, name),
    UNIQUE(type, name)
);

CREATE TABLE IF NOT EXISTS campaign
(
    id          BLOB PRIMARY KEY CHECK ( length(id) = 16 ),
    name        TEXT    NOT NULL,
    description TEXT,
    master_id   BLOB    NOT NULL CHECK ( length(master_id) = 16 ) REFERENCES users (id),
    edition_id  BLOB    NOT NULL CHECK ( length(edition_id) = 16 ) REFERENCES editions (id),
    status_id   integer NOT NULL DEFAULT 0 REFERENCES status_enum (id)
);

CREATE TABLE IF NOT EXISTS game_sessions
(
    date        TEXT    NOT NULL,
    campaign_id BLOB CHECK ( length(campaign_id) = 16 ) REFERENCES campaign (id),
    description TEXT,
    status      INTEGER not null default 0 references status_enum (id)
);


-- FILL
INSERT INTO editions (id, name)
VALUES (X'f96cd1bb88de42168ae854514b209152', 'e4'),
       (X'7a93e88884e3484582bd6482c5a67326', 'e5'),
       (X'2fcd619820eb40959fb34a6fab0effb4', 'homebrew');

insert INTO status_enum(id, type, name)
values (0, 'game', 'planed'),
       (1, 'game', 'inProgress'),
       (2, 'game', 'finished'),
       (3, 'game', 'canceled'),
       (0, 'campaign', 'preparing'),
       (1, 'campaign', 'inProgress'),
       (2, 'campaign', 'paused'),
       (3, 'campaign', 'finished'),
       (4, 'campaign', 'canceled');