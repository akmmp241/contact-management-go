CREATE TABLE contacts
(
    id         INT          NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id    INT          NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name  VARCHAR(255),
    email      VARCHAR(200),
    phone      VARCHAR(20),
    FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE InnoDB;