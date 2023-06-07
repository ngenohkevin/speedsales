-- Create Suppliers table
CREATE TABLE suppliers (
                           supplier_id SERIAL PRIMARY KEY,
                           name VARCHAR,
                           address VARCHAR,
                           contact_number VARCHAR,
                           email VARCHAR
);

-- Create Products table
CREATE TABLE products (
                          product_id SERIAL PRIMARY KEY,
                          name VARCHAR,
                          description TEXT,
                          category VARCHAR,
                          supplier_id INT,
                          cost DECIMAL(10, 2),
                          selling_price DECIMAL(10, 2),
                          quantity INT,
                          FOREIGN KEY (supplier_id) REFERENCES suppliers(supplier_id)
);

-- Create Customers table
CREATE TABLE customers (
                           customer_id SERIAL PRIMARY KEY,
                           name VARCHAR,
                           address VARCHAR,
                           contact_number VARCHAR,
                           email VARCHAR
);

-- Create Sales table
CREATE TABLE sales (
                       sale_id SERIAL PRIMARY KEY,
                       product_id INT,
                       customer_id INT,
                       quantity INT,
                       sale_date DATE,
                       total_price DECIMAL(10, 2),
                       FOREIGN KEY (product_id) REFERENCES products(product_id),
                       FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);
