CREATE TYPE "product_types" AS ENUM (
  'food',
  'goods'
);

CREATE TYPE "product_statusess" AS ENUM (
  'pending',
  'success'
);

CREATE TABLE "roles" (
  "id" integer PRIMARY KEY,
  "name" text,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "name" varchar,
  "email" varchar,
  "email_verified_at" timestamp,
  "password" varchar,
  "token" varchar,
  "role_id" integer,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "products" (
  "id" char(32) PRIMARY KEY,
  "unit_id" varchar(32),
  "delivered_by" text,
  "type" product_types,
  "qty" integer,
  "owner" varchar,
  "phone" varchar,
  "user_in" integer,
  "user_out" integer,
  "picked_by" text,
  "picked_at" timestamp,
  "deleted_at" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "units" (
  "id" char(32) PRIMARY KEY,
  "no" varchar,
  "email" varchar,
  "item_pending_qty" integer,
  "phone" varchar,
  "deleted_at" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

COMMENT ON COLUMN "products"."delivered_by" IS 'nama kurir pengantar paket';

COMMENT ON COLUMN "products"."owner" IS 'nama pemesan paket';

COMMENT ON COLUMN "products"."phone" IS 'nomor telpon pemesan paket';

COMMENT ON COLUMN "products"."user_in" IS 'user yang menerima paket';

COMMENT ON COLUMN "products"."user_out" IS 'user yang mengeluarkan paket';

COMMENT ON COLUMN "products"."picked_by" IS 'nama pengambil paket';

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("unit_id") REFERENCES "units" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("user_in") REFERENCES "users" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("user_out") REFERENCES "users" ("id");
