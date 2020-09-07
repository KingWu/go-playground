CREATE SCHEMA "app";

CREATE TABLE "app"."todo" (
  "id" serial4,
  "text" varchar(255),
  "user_id" int4,
  PRIMARY KEY ("id")
);

CREATE TABLE "app"."user" (
  "id" serial4,
  "name" varchar(255),
  PRIMARY KEY ("id")
);

ALTER TABLE "app"."todo" ADD CONSTRAINT "fk_todo_user_1" FOREIGN KEY ("user_id") REFERENCES "app"."user" ("id");
