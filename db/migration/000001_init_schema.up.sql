CREATE TABLE "users" (
                         "user_id" SERIAL PRIMARY KEY,
                         "username" VARCHAR UNIQUE,
                         "branch" VARCHAR,
                         "stk_location" VARCHAR,
                         "reset" VARCHAR,
                         "till_num" BIGINT,
                         "rights" JSONB,
                         "is_active" BOOLEAN
);

CREATE TABLE "suppliers" (
                             "supplier_id" SERIAL PRIMARY KEY,
                             "name" VARCHAR NOT NULL,
                             "address" VARCHAR NOT NULL,
                             "contact_number" VARCHAR NOT NULL,
                             "email" VARCHAR NOT NULL,
                             "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE "code_translator" (
                                   "master_code" VARCHAR,
                                   "link_code" VARCHAR UNIQUE NOT NULL,
                                   "pkg_qty" FLOAT NOT NULL DEFAULT 0,
                                   "discount" FLOAT NOT NULL DEFAULT 0,
                                   PRIMARY KEY ("master_code")
);

CREATE TABLE "department" (
                              "department_id" SERIAL PRIMARY KEY,
                              "category" VARCHAR,
                              "sub_category" VARCHAR,
                              "description" VARCHAR
);

CREATE TABLE "products" (
                            "product_id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
                            "name" VARCHAR NOT NULL,
                            "description" TEXT NOT NULL,
                            "category" VARCHAR NOT NULL,
                            "department_id" INT NOT NULL,
                            "supplier_id" BIGINT NOT NULL,
                            "cost" BIGINT NOT NULL,
                            "selling_price" BIGINT NOT NULL,
                            "wholesale_price" BIGINT NOT NULL,
                            "min_margin" FLOAT NOT NULL,
                            "quantity" BIGINT NOT NULL,
                            "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                            FOREIGN KEY ("department_id") REFERENCES "department" ("department_id"),
                            FOREIGN KEY ("supplier_id") REFERENCES "suppliers" ("supplier_id")
);

CREATE TABLE "customers" (
                             "customer_id" SERIAL PRIMARY KEY,
                             "name" VARCHAR NOT NULL,
                             "address" VARCHAR,
                             "contact_number" VARCHAR NOT NULL,
                             "email" VARCHAR,
                             "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE "sales_till" (
                              "till_num" BIGINT PRIMARY KEY,
                              "teller" VARCHAR,
                              "supervisor" VARCHAR,
                              "branch" VARCHAR,
                              "open_time" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                              "open_cash" FLOAT NOT NULL DEFAULT 0,
                              "close_time" TIMESTAMPTZ,
                              "close_cash" FLOAT,
                              "close_summary" JSONB
);

CREATE TABLE "sales" (
                         "receipt_num" BIGINT PRIMARY KEY,
                         "till_num" BIGINT,
                         "txn_time" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                         "product_id" BIGINT NOT NULL,
                         "item_name" VARCHAR,
                         "price" FLOAT NOT NULL DEFAULT 0,
                         "cost" FLOAT NOT NULL DEFAULT 0,
                         "quantity" FLOAT NOT NULL DEFAULT 1,
                         "vat_code" VARCHAR(1),
                         "hs_code" VARCHAR,
                         "VAT" FLOAT NOT NULL DEFAULT 0,
                         "batch_code" VARCHAR,
                         "serial_code" VARCHAR,
                         "serial_code_return" VARCHAR,
                         "served_by" VARCHAR,
                         "approved_by" VARCHAR,
                         "state" VARCHAR(30) NOT NULL DEFAULT 'active',
                         FOREIGN KEY ("product_id") REFERENCES "products" ("product_id")
);

CREATE TABLE "salestrace" (
                              "sale_id" SERIAL PRIMARY KEY,
                              "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                              "receipt_num" BIGINT NOT NULL,
                              "till_num" BIGINT,
                              "smart_card" INT NOT NULL DEFAULT 0,
                              "customer_id" BIGINT NOT NULL,
                              "quantity" BIGINT NOT NULL,
                              "sale_time" TIMESTAMPTZ NOT NULL,
                              "total_price" BIGINT NOT NULL,
                              "customer_num" BIGINT,
                              "cash_paid" FLOAT NOT NULL DEFAULT 0,
                              "payment_summary" JSONB,
                              "change" FLOAT NOT NULL DEFAULT 0,
                              "state" VARCHAR(30),
                              FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id"),
                              FOREIGN KEY ("till_num") REFERENCES "sales_till" ("till_num"),
                              FOREIGN KEY ("receipt_num") REFERENCES "sales" ("receipt_num")
);

CREATE TABLE "order_log" (
                             "order_time" TIMESTAMPTZ,
                             "order_num" BIGINT UNIQUE,
                             "poster" VARCHAR,
                             "approver" VARCHAR,
                             "customer_id" BIGINT,
                             "retailer" VARCHAR,
                             FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id")
);

CREATE TABLE "dispatch_log" (
                                "dispatch_time" TIMESTAMPTZ,
                                "dispatch_id" BIGSERIAL PRIMARY KEY,
                                "poster" VARCHAR,
                                "approver" VARCHAR,
                                "order_num" BIGINT,
                                "vehicle" VARCHAR,
                                FOREIGN KEY ("order_num") REFERENCES "order_log" ("order_num")
);

CREATE TABLE "receive_log" (
                               "receive_time" TIMESTAMPTZ,
                               "receive_id" BIGSERIAL PRIMARY KEY,
                               "poster" VARCHAR,
                               "approver" VARCHAR,
                               "dispatch_id" BIGINT,
                               "order_num" BIGINT,
                               FOREIGN KEY ("dispatch_id") REFERENCES "dispatch_log" ("dispatch_id"),
                               FOREIGN KEY ("order_num") REFERENCES "order_log" ("order_num")
);

CREATE TABLE "acquisition_items_trail" (
                                           "item_code" BIGINT PRIMARY KEY NOT NULL,
                                           "item_name" VARCHAR NOT NULL,
                                           "order_quantity" BIGINT,
                                           "dispatch_quantity" BIGINT,
                                           "receive_quantity" BIGINT,
                                           "order_id" BIGINT,
                                           "dispatch_id" BIGINT,
                                           "receive_id" BIGINT,
                                           "price" BIGINT,
                                           "cost" BIGINT,
                                           "vat_code" VARCHAR,
                                           "vat" FLOAT,
                                           FOREIGN KEY ("order_id") REFERENCES "order_log" ("order_num"),
                                           FOREIGN KEY ("dispatch_id") REFERENCES "dispatch_log" ("dispatch_id"),
                                           FOREIGN KEY ("receive_id") REFERENCES "receive_log" ("receive_id")
);

CREATE TABLE "vehicle" (
                           "vehicle_id" SERIAL PRIMARY KEY NOT NULL,
                           "vehicle_name" VARCHAR,
                           "registration_num" VARCHAR NOT NULL,
                           "vin_num" VARCHAR,
                           "manufacture_date" DATE,
                           "mileage" FLOAT,
                           "last_mileage_read" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                           "insurance_expiry" DATE NOT NULL,
                           "checklist" JSONB,
                           "state" VARCHAR
);

CREATE TABLE "vehicle_activity_log" (
                                        "activity_id" BIGSERIAL PRIMARY KEY NOT NULL,
                                        "activity_time" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                        "done_by" VARCHAR,
                                        "vehicle_id" INT NOT NULL,
                                        "approved_by" VARCHAR,
                                        "checklist" JSONB,
                                        "notation" VARCHAR,
                                        FOREIGN KEY ("vehicle_id") REFERENCES "vehicle" ("vehicle_id"),
                                        FOREIGN KEY ("done_by") REFERENCES "users" ("username")
);
