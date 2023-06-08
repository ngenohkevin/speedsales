CREATE TABLE "suppliers" (
                             "supplier_id" SERIAL PRIMARY KEY,
                             "name" VARCHAR NOT NULL,
                             "address" VARCHAR NOT NULL,
                             "contact_number" VARCHAR NOT NULL,
                             "email" VARCHAR NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
                            "product_id" SERIAL PRIMARY KEY,
                            "name" VARCHAR NOT NULL,
                            "description" TEXT NOT NULL,
                            "category" VARCHAR NOT NULL,
                            "supplier_id" INT NOT NULL,
                            "cost" bigint NOT NULL,
                            "selling_price" bigint NOT NULL,
                            "quantity" INT NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "customers" (
                             "customer_id" SERIAL PRIMARY KEY,
                             "name" VARCHAR NOT NULL,
                             "address" VARCHAR,
                             "contact_number" VARCHAR NOT NULL,
                             "email" VARCHAR,
                             "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sales" (
                         "sale_id" SERIAL PRIMARY KEY,
                         "product_id" INT NOT NULL,
                         "customer_id" INT NOT NULL,
                         "quantity" INT NOT NULL,
                         "sale_date" DATE NOT NULL,
                         "total_price" bigint NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "products" ADD FOREIGN KEY ("supplier_id") REFERENCES "suppliers" ("supplier_id");

ALTER TABLE "sales" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id");

ALTER TABLE "sales" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id");
