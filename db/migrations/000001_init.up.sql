
CREATE TABLE "tables" (
  "id" bigserial PRIMARY KEY,
  "table_number" int NOT NULL,
  "capacity" int NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "menu_items" (
  "id" bigserial PRIMARY KEY,
  "name" VARCHAR(100) NOT NULL,
  "description" TEXT,
  "price" DECIMAL(10,2) NOT NULL,
  "available" BOOLEAN DEFAULT true,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "customer_name" varchar(250) NOT NULL,
  "customer_phone" varchar(20),
  "customer_email" varchar(30),
  "table_id" int,
  "total_amount" decimal(10,2),
  "order_status" varchar(50),
  "payment_method" VARCHAR(50),
  "payment_status" VARCHAR(50) DEFAULT 'UNPAID',
  "closed_at" TIMESTAMP DEFAULT null,
  "delivery_address" text,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "order_items" (
  "id" bigserial PRIMARY KEY,
  "order_id" INT NOT NULL,
  "menu_item_id" INT NOT NULL,
  "quantity" INT NOT NULL DEFAULT 1,
  "price" DECIMAL(10,2) NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "cash_closures" (
  "id" bigserial PRIMARY KEY,
  "closed_at" TIMESTAMP DEFAULT (now()),
  "total_revenue" DECIMAL(10,2),
  "total_debt" DECIMAL(10,2),
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "cash_closure_orders" (
  "id" bigserial PRIMARY KEY,
  "cash_closure_id" INT NOT NULL,
  "order_id" INT NOT NULL
);

COMMENT ON COLUMN "orders"."order_status" IS 'Posible status: Pending, In Progress, Completed';

COMMENT ON COLUMN "orders"."payment_method" IS 'Cash, Credit Card, etc.';

COMMENT ON COLUMN "orders"."payment_status" IS 'Paid or Unpaid';

COMMENT ON COLUMN "orders"."closed_at" IS 'When the payment was completed';

COMMENT ON COLUMN "order_items"."price" IS 'price at the time of the order';

COMMENT ON COLUMN "cash_closures"."total_revenue" IS 'Total money collected during the session';

COMMENT ON COLUMN "cash_closures"."total_debt" IS 'Total pending amount from unpaid orders';


ALTER TABLE "orders" ADD FOREIGN KEY ("table_id") REFERENCES "tables" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("menu_item_id") REFERENCES "menu_items" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "cash_closure_orders" ADD FOREIGN KEY ("cash_closure_id") REFERENCES "cash_closures" ("id");

ALTER TABLE "cash_closure_orders" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
