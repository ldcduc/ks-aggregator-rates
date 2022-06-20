CREATE TABLE "kyberswap" (
    "id" serial PRIMARY KEY,
    "pair" varchar,
    "input_token" varchar,
    "output_token" varchar,
    "input_amount" varchar,
    "output_amount" varchar,
    "total_gas" double precision,
    "gas_price_gwei" varchar,
    "gas_usd" double precision,
    "timestamp" bigint
);

CREATE TABLE "paraswap" (
    "id" serial PRIMARY KEY,
    "pair" varchar,
    "block_number" integer,
    "src_token" varchar,
    "src_amount" varchar,
    "dest_token" varchar,
    "dest_amount" varchar,
    "timestamp" bigint
);

CREATE TABLE "oneinch" (
    "id" serial PRIMARY KEY,
    "pair" varchar,
    "from_token_address" varchar,
    "to_token_address" varchar,
    "from_token_amount" varchar,
    "to_token_amount" varchar,
    "estimated_gas" double precision,
    "timestamp" bigint
);

CREATE TABLE "zerox" (
    "id" serial PRIMARY KEY,
    "pair" varchar,
    "buy_token_address" varchar,
    "sell_token_address" varchar,
    "buy_amount" varchar,
    "sell_amount" varchar,
    "timestamp" bigint
);
