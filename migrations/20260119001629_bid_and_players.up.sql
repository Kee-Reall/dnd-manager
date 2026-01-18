CREATE TABLE players
(
    id          BLOB primary key CHECK ( length(id) = 16 ),
    campaign_id BLOB CHECK ( length(campaign_id) = 16 ) REFERENCES campaign (id) on delete cascade,
    user_id     BLOB CHECK ( length(user_id) = 16 ) REFERENCES users (id) on delete cascade,
    character   TEXT NOT NULL
);

CREATE TABLE bid
(
    id          BLOB primary key CHECK ( length(id) = 16 ),
    user_id     BLOB CHECK ( length(user_id) = 16 ) REFERENCES users (id) on delete cascade,
    campaign_id BLOB CHECK ( length(campaign_id) = 16 ) REFERENCES campaign (id) on delete cascade,
    status      integer not null default 0,
    text        TEXT             default null

);