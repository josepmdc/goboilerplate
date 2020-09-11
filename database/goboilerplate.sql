CREATE TABLE IF NOT EXISTS users (
  id            UUID PRIMARY KEY,
  github_id     VARCHAR(1024) UNIQUE,
  email         VARCHAR NOT NULL,
  username      VARCHAR NOT NULL,
  password      VARCHAR,
  full_name     VARCHAR,
  avatar_url    VARCHAR,
  admin         BOOLEAN NOT NULL DEFAULT false,
  score         INT NOT NULL DEFAULT 0,
  created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS permissions (
  id            UUID PRIMARY KEY,
  name          VARCHAR,
  description   VARCHAR
);

CREATE TABLE IF NOT EXISTS user_permissions (
  user_id         UUID REFERENCES users(id),
  permission_id   UUID REFERENCES permissions(id)
);

