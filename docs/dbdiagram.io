// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs

Table customers {
  id bigserial [primary key]
  fullname varchar(250) [not null]
  phone_number varchar(20) [not null]
  email varchar(30)
  created_at timestamp [default: `now()`]
}

Table tables {
  id bigserial [primary key]
  table_number int [not null]
  capacity int [not null]
  created_at timestamp [default: `now()`]
}

Table menu_items {
  id bigserial [primary key]
  name VARCHAR(100) [NOT NULL]
  description TEXT
  price DECIMAL(10, 2) [NOT NULL]
  available BOOLEAN [DEFAULT: TRUE]
  created_at timestamp [default: `now()`]
}


Table orders {
  id bigserial [primary key]
  customer_id int
  table_id int
  total_amount decimal(10,2)
  order_status varchar(50) [note: "Posible status: Pending, In Progress, Completed"]
  payment_method VARCHAR(50) [note: "Cash, Credit Card, etc."]
  payment_status VARCHAR(50) [DEFAULT: 'Unpaid', note: "Paid or Unpaid"]
  closed_at TIMESTAMP [default: null, note: "When the payment was completed"]
  delivery_address text
  created_at timestamp [default: `now()`]
}

Table order_items {
  id bigserial [primary key]
  order_id INT [NOT NULL]
  menu_item_id INT [NOT NULL]
  quantity INT [NOT NULL, DEFAULT: 1]
  price DECIMAL(10, 2) [NOT NULL, note: "price at the time of the order"]
  created_at timestamp [default: `now()`]
}

Table cash_closures {
  id bigserial [primary key]
  closed_at TIMESTAMP [DEFAULT: `now()`]
  total_revenue DECIMAL(10, 2) [note: "Total money collected during the session"]
  total_debt DECIMAL(10, 2) [note: "Total pending amount from unpaid orders"]
  created_at timestamp [default: `now()`]

}

Table cash_closure_orders {
  id bigserial [primary key]
  cash_closure_id INT [NOT NULL]
  order_id INT [NOT NULL]
}

Ref: orders.customer_id > customers.id
Ref: orders.table_id > tables.id
Ref: order_items.menu_item_id > menu_items.id
Ref: order_items.order_id > orders.id
Ref: cash_closure_orders.cash_closure_id > cash_closures.id
Ref: cash_closure_orders.order_id > orders.id
