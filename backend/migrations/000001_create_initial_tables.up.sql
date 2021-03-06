DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS currencies CASCADE;
DROP TABLE IF EXISTS providers CASCADE;
DROP TABLE IF EXISTS instruments CASCADE;
DROP TABLE IF EXISTS markets CASCADE;
DROP TABLE IF EXISTS registers CASCADE;
DROP TABLE IF EXISTS portfolios CASCADE;
DROP TABLE IF EXISTS assets CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;
-- CREATE EXTENSION IF NOT EXISTS postgis;
-- CREATE EXTENSION IF NOT EXISTS postgis_topology;

CREATE TABLE users
(
    user_id    UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    name       VARCHAR(32)              NOT NULL CHECK ( name <> '' ),
    email      VARCHAR(64) UNIQUE       NOT NULL CHECK ( email <> '' ),
    password   VARCHAR(250)             NOT NULL CHECK ( octet_length(password) <> 0 ),
    role       VARCHAR(10)              NOT NULL DEFAULT 'user',
    avatar     VARCHAR(512),
    confirmed  UUID                              DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE currencies
(
    currency_id UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    title       VARCHAR(10) UNIQUE       NOT NULL CHECK ( title <> '' ),
    description VARCHAR(100)             NOT NULL CHECK ( description <> '' ),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE providers
(
    provider_id UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    title       VARCHAR(100) UNIQUE      NOT NULL CHECK ( title <> '' ),
    description VARCHAR(250)                      DEFAULT '',
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE instruments
(
    instrument_id UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    title         VARCHAR(100) UNIQUE      NOT NULL CHECK ( title <> '' ),
    description   VARCHAR(250)             NOT NULL CHECK ( description <> '' ),
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE markets
(
    market_id     UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    title         VARCHAR(250)             NOT NULL CHECK ( title <> '' ),
    ticker        VARCHAR(50) UNIQUE       NOT NULL CHECK ( ticker <> '' ),
    content       TEXT                              DEFAULT '',
    image_url     VARCHAR(1024) CHECK ( image_url <> '' ),
    currency_id   UUID                     NOT NULL REFERENCES currencies (currency_id) ON DELETE CASCADE,
    instrument_id UUID                     NOT NULL REFERENCES instruments (instrument_id) ON DELETE CASCADE,
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE registers
(
    register_id UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    identify    VARCHAR(250) UNIQUE      NOT NULL CHECK ( identify <> '' ),
    provider_id UUID                     NOT NULL REFERENCES providers (provider_id) ON DELETE CASCADE,
    market_id   UUID                     NOT NULL REFERENCES markets (market_id) ON DELETE CASCADE,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE portfolios
(
    portfolio_id UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    title        VARCHAR(250),
    active       BOOLEAN                  NOT NULL DEFAULT FALSE,
    user_id      UUID                     NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    currency_id  UUID                     NOT NULL REFERENCES currencies (currency_id) ON DELETE CASCADE,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE assets
(
    asset_id     UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    amount       INTEGER                  NOT NULL,
    quantity     INTEGER                  NOT NULL,
    portfolio_id UUID                     NOT NULL REFERENCES portfolios (portfolio_id) ON DELETE CASCADE,
    market_id    UUID                     NOT NULL REFERENCES markets (market_id) ON DELETE CASCADE,
    notation_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS markets_title_id_idx ON markets (title);
CREATE INDEX IF NOT EXISTS markets_ticker_id_idx ON markets (ticker);
CREATE INDEX IF NOT EXISTS registers_identify_id_idx ON registers (identify);

INSERT INTO currencies (currency_id, title, description)
VALUES ('e6dffe5f-af39-44c4-a9f2-4938cceb7f7c', 'RUB', '??????????'),
       ('fd39d16f-db12-4aa2-80d6-a2917dbc8715', 'USD', '????????????'),
       ('9c093338-0079-45af-80b7-c58c991d4535', 'EUR', '????????'),
       ('3f909b14-f18b-4b8b-95b3-19a2fcf1f9d7', 'CNY', '????????');

INSERT INTO instruments (instrument_id, title, description)
VALUES ('2ca3707d-03b6-4f12-8f1a-6c8ec522ac95', 'STOCK', '??????????'),
       ('2bbc7edd-8f11-4625-846a-8a98c89e0a29', 'BOND', '??????????????????'),
       ('99a91a87-24eb-4202-af0d-104309a42487', 'ETF', '?????????? ETF'),
       ('83b6e4ef-0feb-4935-9544-a81d06506d76', 'CURRENCY', '????????????'),
       ('cc376387-4f0b-4688-88e6-02c3af93a646', 'CRYPTO', '????????????????????????');

INSERT INTO providers (provider_id, title, description)
VALUES ('514edc8f-0921-468e-95f4-2284cba5b7bb', 'Tinkoff', '???????????????? ????????????????????'),
       ('ba93ed83-8687-41cf-8741-edf79548e7df', 'Binance', '???????????????????????????? ?????????? Binance');
