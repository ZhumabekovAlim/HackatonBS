CREATE TABLE IF NOT EXISTS books
(
    id                  INT AUTO_INCREMENT PRIMARY KEY,
    isbn                BIGINT       NOT NULL,
    book_title          VARCHAR(255) NOT NULL,
    book_author         VARCHAR(255) NOT NULL,
    year_of_publication VARCHAR(4)   NOT NULL,
    publisher           VARCHAR(255) NOT NULL,
    image_url_s         VARCHAR(255) NOT NULL,
    image_url_m         VARCHAR(255) NOT NULL,
    image_url_l         VARCHAR(255) NOT NULL,
    book_status         INT          NOT NULL,
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
