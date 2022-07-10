DROP TABLE IF EXISTS mc_block CASCADE;

CREATE TABLE mc_block
(
    id         BYTEA PRIMARY KEY NOT NULL,
    work_chain INTEGER           NOT NULL,
    shard      BYTEA             NOT NULL,
    seqno      BIGINT            NOT NULL,
    root_hash  BYTEA             NOT NULL,
    file_hash  BYTEA             NOT NULL
);