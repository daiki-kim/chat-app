-- TODO:
--Change id's "SERIAL" to "INT AUTO_INCREMENT PRIMARY KEY" when to use MySQL
-- Add "ON UPDATE CURRENT_TIMESTAMP" to rooms_updated_at "timestamp" when to use MySQL

-- users テーブルを作成
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

-- rooms テーブルを作成
CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    owner_id INT NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES users(id)
);

-- messages テーブルを作成
CREATE TABLE messages (
    id SERIAL PRIMARY KEY, 
    room_id INT NOT NULL, 
    sender_id INT NOT NULL, 
    content TEXT NOT NULL, 
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (sender_id) REFERENCES users(id)
);

-- room_members テーブルを作成
CREATE TABLE room_members (
    room_id INT NOT NULL,
    user_id INT NOT NULL,
    PRIMARY KEY (room_id, user_id),
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
;