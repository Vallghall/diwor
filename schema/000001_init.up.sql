CREATE TABLE "users"
(
    "id"            serial UNIQUE PRIMARY KEY NOT NULL ,
    "name"          varchar        NOT NULL,
    "username"      varchar UNIQUE NOT NULL,
    "password_hash" varchar NOT NULL
);

CREATE TABLE "experiments"
(
    "id"          serial PRIMARY KEY NOT NULL UNIQUE,
    "user_id"     integer,
    "sample_id"   integer unique,
    "result_id"   integer unique,
    "started_at"  timestamp NOT NULL,
    "finished_at" timestamp NOT NULL
);

CREATE TABLE "samples"
(
    "id"        serial PRIMARY KEY NOT NULL UNIQUE,
    "algorithm" varchar NOT NULL,
    "mode"      varchar NOT NULL
);

CREATE TABLE "results"
(
    "id"       serial PRIMARY KEY NOT NULL UNIQUE,
    "plot_ref" varchar UNIQUE NOT NULL
);

ALTER TABLE "experiments"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "samples"
    ADD FOREIGN KEY ("id") REFERENCES "experiments" ("sample_id");

ALTER TABLE "experiments"
    ADD FOREIGN KEY ("result_id") REFERENCES "results" ("id");
