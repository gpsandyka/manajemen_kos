drop type "stay_type";
drop type "payment_type";
drop type "pay_type";
drop type "room_status";

CREATE TYPE "stay_type" AS ENUM (
  'TEMPORARY',
  'PERMANENT'
);


CREATE TYPE "payment_type" AS ENUM (
  'GOPAY',
  'OVO',
  'SHOPEEPAY',
  'CASH'
);

CREATE TYPE "pay_type" AS ENUM (
  'ROOMS',
  'FOODS'
);

CREATE TYPE "room_status" AS ENUM (
  'UNOCCUPIED',
  'OCCUPIED',
  'NOTREADY'
);

CREATE TYPE "problem_type" AS ENUM (
  'TIDAKPENTINGTIDAKGAWAT',
  'PENTINGTIDAKGAWAT',
  'PENTINGGAWAT'
);

drop table residents CASCADE;
drop table rooms CASCADE;
drop table invoices ;
drop table admins ;

CREATE TABLE "residents" (
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "id" int PRIMARY KEY NOT NULL,
  "stay_type" stay_type NOT NULL,
  "bills" int,
  "phone_number" bigint,
  "full_name" varchar NOT NULL,
  "occupation" varchar,
  "address" varchar,
  "nik_number" bigint NOT NULL,
  "username" varchar UNIQUE,
  "password" varchar
);

CREATE TABLE "rooms" (
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "id" int PRIMARY KEY NOT NULL,
  "resident_id" int,
  "floor_level" int,
  "room_type" varchar,
  "price" int NOT NULL,
  "start_occupied_date" timestamp NOT NULL,
  "end_occupied_date" timestamp,
  "status" room_status NOT NULL,
  "door_code" varchar,
  "qr_api_code" varchar
);

CREATE TABLE "invoices" (
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "id" int PRIMARY KEY NOT NULL,
  "resident_id" int NOT NULL,
  "room_id" int NOT NULL,
  "total_price" int NOT NULL,
  "pay_date" timestamp NOT NULL,
  "payment_type" payment_type NOT NULL,
  "pay_type" pay_type NOT NULL,
  "description" varchar
);

CREATE TABLE "problems" (
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "id" int PRIMARY KEY NOT NULL,
  "room_id" int,
  "problem_type" problem_type NOT NULL,
  "description" varchar
);

CREATE TABLE "notifications" (
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "id" int PRIMARY KEY NOT NULL,
  "title" varchar,
  "description" varchar
);

CREATE TABLE "admins" (
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "id" int PRIMARY KEY NOT NULL,
  "admin_name" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

ALTER TABLE "invoices" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "invoices" ADD FOREIGN KEY ("resident_id") REFERENCES "residents" ("id");

ALTER TABLE "rooms" ADD FOREIGN KEY ("resident_id") REFERENCES "residents" ("id");

ALTER TABLE "problems" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");
