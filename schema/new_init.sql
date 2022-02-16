CREATE TABLE "users"
(
    "id"            serial UNIQUE PRIMARY KEY NOT NULL ,
    "name"          varchar        NOT NULL,
    "username"      varchar UNIQUE NOT NULL,
    "password_hash" varchar NOT NULL
);

CREATE TABLE "experiments"
(
    "id"           serial PRIMARY KEY NOT NULL UNIQUE,
    "user_id"      integer,
    "algorithm_type" varchar NOT NULL,
    "results"  jsonb NOT NULL,
    "started_at"   timestamp NOT NULL,
    "finished_at"  timestamp NOT NULL
);

ALTER TABLE "experiments"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");