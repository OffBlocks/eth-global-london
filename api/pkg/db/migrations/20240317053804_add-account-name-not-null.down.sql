-- reverse: modify "accounts" table
ALTER TABLE "public"."accounts" ALTER COLUMN "name" DROP NOT NULL;
