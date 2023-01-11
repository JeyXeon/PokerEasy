CREATE TABLE IF NOT EXISTS account(
    account_id SERIAL PRIMARY KEY,
    account_user_name VARCHAR NOT NULL,
    money_balance BIGINT,
    connected_lobby_id INTEGER DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS lobby(
    lobby_id SERIAL PRIMARY KEY,
    lobby_name VARCHAR NOT NULL,
    players_amount INTEGER NOT NULL,
    creator_id INTEGER NOT NULL REFERENCES account
);
