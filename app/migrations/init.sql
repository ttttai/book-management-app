CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (name, email) VALUES 
('Taro Yamada', 'taro.yamada@example.com'),
('Saki Abe', 'saki.abe@example.com'),
('Taishi Yamamoto', 'taishi.yamamoto@example.com');