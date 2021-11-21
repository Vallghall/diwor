CREATE TABLE "users"
(
    "id"            integer UNIQUE PRIMARY KEY,
    "name"          varchar        NOT NULL,
    "experiment_id" integer unique,
    "password_hash" varchar UNIQUE NOT NULL
);

CREATE TABLE "experiments"
(
    "id"          integer PRIMARY KEY,
    "user_id"     integer,
    "sample_id"   integer unique,
    "result_id"   integer unique,
    "started_at"  timestamp NOT NULL,
    "finished_at" timestamp NOT NULL
);

CREATE TABLE "samples"
(
    "id"        integer PRIMARY KEY,
    "algorithm" varchar NOT NULL,
    "mode"      varchar NOT NULL
);

CREATE TABLE "results"
(
    "id"       integer PRIMARY KEY,
    "plot_ref" varchar UNIQUE NOT NULL
);

ALTER TABLE "experiments"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("experiment_id");

ALTER TABLE "samples"
    ADD FOREIGN KEY ("id") REFERENCES "experiments" ("sample_id");

ALTER TABLE "experiments"
    ADD FOREIGN KEY ("result_id") REFERENCES "results" ("id");
