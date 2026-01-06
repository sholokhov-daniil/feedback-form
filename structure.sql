-- ======================
-- auth_types
-- ======================
CREATE TABLE IF NOT EXISTS auth_types (
    id SERIAL PRIMARY KEY,
    code VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT DEFAULT
);

-- ======================
-- users
-- ======================
CREATE TABLE IF NOT EXISTS users (
    id            SERIAL PRIMARY KEY,
    active        BOOLEAN NOT NULL DEFAULT TRUE,
    name          VARCHAR(255) NOT NULL,
    date_create   TIMESTAMP NOT NULL DEFAULT NOW(),
    date_update   TIMESTAMP NOT NULL DEFAULT NOW()
);

-- ======================
-- user_auth
-- ======================
CREATE TABLE IF NOT EXISTS user_auth (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,                    -- связь с таблицей users
    auth_type_id INT NOT NULL,               -- связь с auth_types
    identifier VARCHAR(255) NOT NULL,       -- login / client_id / subject
    secret_hash TEXT NOT NULL,               -- password / token hash
    active BOOLEAN NOT NULL DEFAULT TRUE,    -- флаг активности
    expires_at TIMESTAMP,                    -- дата истечения токена/сессии
    date_create TIMESTAMP NOT NULL DEFAULT NOW(),
    date_update TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_userauth_user
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    CONSTRAINT fk_userauth_authtype
        FOREIGN KEY (auth_type_id) REFERENCES auth_types(id) ON DELETE CASCADE
);

-- ======================
-- forms
-- ======================
CREATE TABLE IF NOT EXISTS forms (
    id            VARCHAR(255) PRIMARY KEY,
    active        BOOLEAN NOT NULL DEFAULT TRUE,
    user_id       INTEGER NOT NULL,
    date_create   TIMESTAMP NOT NULL DEFAULT NOW(),
    date_update   TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_forms_owner
        FOREIGN KEY (owner)
        REFERENCES users (id)
        ON DELETE CASCADE
);

-- ======================
-- field_types
-- ======================
CREATE TABLE IF NOT EXISTS field_types (
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(255) NOT NULL
);

-- ======================
-- fields
-- ======================
CREATE TABLE IF NOT EXISTS fields (
    id            VARCHAR(255) PRIMARY KEY,
    form_id       VARCHAR(255) NOT NULL,
    code          VARCHAR(255) NOT NULL,
    active        BOOLEAN NOT NULL DEFAULT TRUE,
    name          VARCHAR(255) NOT NULL,
    type          INTEGER NOT NULL,
    settings      TEXT NOT NULL DEFAULT '',
    date_create   TIMESTAMP NOT NULL DEFAULT NOW(),
    date_update   TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_fields_form
        FOREIGN KEY (form_id)
        REFERENCES forms (id)
        ON DELETE CASCADE,

    CONSTRAINT fk_fields_type
        FOREIGN KEY (type)
        REFERENCES field_types (id)
);