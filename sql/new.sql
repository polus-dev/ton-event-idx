DROP TABLE IF EXISTS mc_block CASCADE;
DROP TABLE IF EXISTS block CASCADE;
DROP TABLE IF EXISTS contract CASCADE;
DROP TABLE IF EXISTS event CASCADE;

CREATE TABLE mc_block
( ------------ work_chain always -1 (master chain)
    id         UUID      NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    shard      BYTEA     NOT NULL,
    seqno      BIGINT    NOT NULL
);

CREATE TABLE block
(
    id          UUID      NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    mc_block_id UUID      NOT NULL REFERENCES mc_block (id),

    work_chain  INTEGER   NOT NULL,
    shard       BYTEA     NOT NULL,
    seqno       BIGINT    NOT NULL
);


CREATE TABLE contract
(
    id         SERIAL UNIQUE PRIMARY KEY NOT NULL,
    work_chain BYTEA                     NOT NULL,
    address    BYTEA                     NOT NULL
);

CREATE TABLE event
(
    id          UUID      NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),

    block_id    UUID      NOT NULL REFERENCES block (id),
    contract_id INTEGER   NOT NULL REFERENCES contract (id),

    body        BYTEA     NOT NULL
);
