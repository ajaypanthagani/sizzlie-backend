CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT,
    user_id INT,
    recipe_id INT,
    quantity INT,
    price INT,
    address_id INT
);
