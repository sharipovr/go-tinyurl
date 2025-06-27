-- Table «short» ids
CREATE TABLE urls (
    id          VARCHAR(16) PRIMARY KEY,      -- short id
    original    TEXT        NOT NULL,         -- full URL
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Table of visits (for the future)
CREATE TABLE visits (
    id          BIGSERIAL PRIMARY KEY,
    url_id      VARCHAR(16) NOT NULL REFERENCES urls(id) ON DELETE CASCADE,
    visited_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    referrer    TEXT,
    user_agent  TEXT
);

CREATE INDEX idx_visits_url_id ON visits (url_id);
