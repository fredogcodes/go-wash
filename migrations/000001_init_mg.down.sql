-- CREATE USERS TABLE
CREATE TABLE "users" (
    "id" BIGSERIAL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE NOT NULL,
    "mobile" VARCHAR(255) NOT NULL UNIQUE NOT NULL,
    PRIMARY KEY ("id")
)

-- CREATE CARS TABLE
CREATE TABLE "cars" (
    "id" BIGSERIAL,
    "branch" VARCHAR(255) NOT NULL,
    "year" INTEGER NOT NULL,
    "user_id" BIGINT NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("user_id") REFERENCES "users" ("id")
)