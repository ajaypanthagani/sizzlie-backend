CREATE TABLE ingredients (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    description VARCHAR(100),
    thumbnail VARCHAR(1000) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
