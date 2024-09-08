CREATE TABLE user_book
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    user_id     INT          NOT NULL,
    book_id     INT          NOT NULL,
    date_from   VARCHAR(255) NOT NULL,
    date_to     VARCHAR(255) NOT NULL,
    date_return VARCHAR(255),
    extension_number INT NOT NULL ,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (book_id) REFERENCES books (id)
);
