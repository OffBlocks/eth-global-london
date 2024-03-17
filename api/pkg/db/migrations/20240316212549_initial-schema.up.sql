-- create "accounts" table
CREATE TABLE "public"."accounts" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "identifier" text NOT NULL,
  "identifier_type" text NOT NULL,
  "wallet_id" text NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);
-- create index "idx_accounts_created_at" to table: "accounts"
CREATE INDEX "idx_accounts_created_at" ON "public"."accounts" ("created_at");
-- create index "idx_accounts_identifier" to table: "accounts"
CREATE INDEX "idx_accounts_identifier" ON "public"."accounts" ("identifier");
-- create index "idx_accounts_updated_at" to table: "accounts"
CREATE INDEX "idx_accounts_updated_at" ON "public"."accounts" ("updated_at");
-- create index "idx_accounts_wallet_id" to table: "accounts"
CREATE INDEX "idx_accounts_wallet_id" ON "public"."accounts" ("wallet_id");
-- create "payments" table
CREATE TABLE "public"."payments" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "amount" text NOT NULL,
  "currency" text NOT NULL,
  "source_wallet_id" text NULL,
  "destination" uuid NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);
-- create index "idx_payments_created_at" to table: "payments"
CREATE INDEX "idx_payments_created_at" ON "public"."payments" ("created_at");
-- create index "idx_payments_destination" to table: "payments"
CREATE INDEX "idx_payments_destination" ON "public"."payments" ("destination");
-- create index "idx_payments_updated_at" to table: "payments"
CREATE INDEX "idx_payments_updated_at" ON "public"."payments" ("updated_at");
