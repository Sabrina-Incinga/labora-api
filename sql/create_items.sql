CREATE DATABASE labora_proyect_1;

CREATE TABLE items(
id INT NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 ),
customer_name VARCHAR(100) NOT NULL,
order_date DATE NOT NULL,
product VARCHAR(255) NOT NULL,
quantity INT NOT NULL,
price NUMERIC NOT NULL,
PRIMARY KEY (id))