CREATE TYPE "product_types" AS ENUM (
  'food',
  'goods'
);

CREATE TYPE "product_statusess" AS ENUM (
  'pending',
  'success'
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "roles" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()) NOT NULL,
  "deleted_at" timestamp DEFAULT null,
  "updated_at" timestamp DEFAULT null
);

CREATE TABLE "users" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "email_verified_at" timestamp DEFAULT null,
  "password" varchar NOT NULL,
  "token" varchar DEFAULT null,
  "role_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp DEFAULT null,
  "updated_at" timestamp DEFAULT null
);

CREATE TABLE "transactions" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "unit_id" uuid NOT NULL,
  "delivered_by" text NOT NULL,
  "type" product_types NOT NULL,
  "qty" integer NOT NULL,
  "owner" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "user_in_id" uuid NOT NULL,
  "user_out_id" uuid  DEFAULT null,
  "picked_by" text DEFAULT null,
  "picked_at" timestamp DEFAULT null,
  "deleted_at" timestamp DEFAULT null,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT null
);

CREATE TABLE "units" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "no" varchar NOT NULL,
  "email" varchar NOT NULL,
  "item_pending_qty" integer DEFAULT 0,
  "phone" varchar NOT NULL,
  "deleted_at" timestamp DEFAULT null,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT null
);

COMMENT ON COLUMN "transactions"."delivered_by" IS 'nama kurir pengantar paket';

COMMENT ON COLUMN "transactions"."owner" IS 'nama pemesan paket';

COMMENT ON COLUMN "transactions"."phone" IS 'nomor telpon pemesan paket';

COMMENT ON COLUMN "transactions"."user_in_id" IS 'user yang menerima paket';

COMMENT ON COLUMN "transactions"."user_out_id" IS 'user yang mengeluarkan paket';

COMMENT ON COLUMN "transactions"."picked_by" IS 'nama pengambil paket';

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("unit_id") REFERENCES "units" ("id");

ALTER TABLE "units"
ADD CONSTRAINT "no_unit_unique" UNIQUE (no);

ALTER TABLE "units"
ADD CONSTRAINT "email_unit_unique" UNIQUE (email);

ALTER TABLE "units"
ADD CONSTRAINT "phone_unit_unique" UNIQUE (phone);

ALTER TABLE "users"
ADD CONSTRAINT "email_user_unique" UNIQUE (email);

ALTER TABLE "users"
ADD CONSTRAINT "password_user_unique" UNIQUE (password);

ALTER TABLE "users"
ADD CONSTRAINT "token_user_unique" UNIQUE (token);