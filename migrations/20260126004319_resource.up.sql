CREATE TABLE resource
(
    id          BLOB PRIMARY KEY CHECK ( length(id) = 16 ),
    media_id    TEXT NOT NULL,
    TITLE       TEXT NOT NULL,
    edition_id  BLOB REFERENCES editions (id) ON DELETE NO ACTION,
    campaign_id BLOB REFERENCES campaign (id) ON DELETE NO ACTION,

    CHECK (
        (
            edition_id IS NOT NULL
            AND campaign_id IS NULL
            AND length(edition_id) = 16
        )
        OR
        (
            edition_id IS NULL
            AND campaign_id IS NOT NULL
            AND length(campaign_id) = 16)
    )
);