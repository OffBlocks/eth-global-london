-- reverse: create index "idx_payments_updated_at" to table: "payments"
DROP INDEX "public"."idx_payments_updated_at";
-- reverse: create index "idx_payments_destination" to table: "payments"
DROP INDEX "public"."idx_payments_destination";
-- reverse: create index "idx_payments_created_at" to table: "payments"
DROP INDEX "public"."idx_payments_created_at";
-- reverse: create "payments" table
DROP TABLE "public"."payments";
-- reverse: create index "idx_accounts_wallet_id" to table: "accounts"
DROP INDEX "public"."idx_accounts_wallet_id";
-- reverse: create index "idx_accounts_updated_at" to table: "accounts"
DROP INDEX "public"."idx_accounts_updated_at";
-- reverse: create index "idx_accounts_identifier" to table: "accounts"
DROP INDEX "public"."idx_accounts_identifier";
-- reverse: create index "idx_accounts_created_at" to table: "accounts"
DROP INDEX "public"."idx_accounts_created_at";
-- reverse: create "accounts" table
DROP TABLE "public"."accounts";
