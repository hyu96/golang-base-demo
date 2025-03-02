CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    total_amount FLOAT NOT NULL,
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items (
     id SERIAL PRIMARY KEY,
     order_id INT NOT NULL,
     product_id INT NOT NULL,
     quantity INT NOT NULL,
     price FLOAT NOT NULL,
     FOREIGN KEY (order_id) REFERENCES orders(id)
);

INSERT INTO orders (customer_id, total_amount, status) VALUES
   (1, 100.50, 'pending'),
   (2, 250.00, 'processing'),
   (3, 50.75, 'completed');

INSERT INTO order_items (order_id, product_id, quantity, price) VALUES
    (1, 10, 2, 50.25),
    (1, 11, 1, 0),
    (2, 20, 5, 50.00),
    (3, 30, 1, 50.75)