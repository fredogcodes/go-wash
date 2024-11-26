-- CREATE USERS TABLE
CREATE TABLE users (
    "id" UUID,
    "first_name" VARCHAR(50) NOT NULL,
    "last_name" VARCHAR(50) NOT NULL,
    "email" VARCHAR(50) NOT NULL UNIQUE,
    "mobile" VARCHAR(50) NOT NULL UNIQUE,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

--CREATE CARS TABLE
CREATE TABLE cars (
    "id" UUID,
    "branch" VARCHAR(50) NOT NULL,
    "year" VARCHAR(50) NOT NULL,
    "user_id" UUID NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

--CREATE FOREIGN KEY
ALTER TABLE cars ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id);