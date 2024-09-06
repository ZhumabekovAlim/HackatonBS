CREATE TABLE user_achievement
(
    id             INT AUTO_INCREMENT PRIMARY KEY,
    user_id        INT NOT NULL,
    achievement_id INT NOT NULL,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (achievement_id) REFERENCES achievements (id)
);
