CREATE TABLE addresses
(
    id         INT          NOT NULL AUTO_INCREMENT PRIMARY KEY,
    contact_id INT          NOT NULL,
    street VARCHAR(200),
    city VARCHAR(100),
    province VARCHAR(100),
    country VARCHAR(100) NOT NULL,
    postal_code VARCHAR(10),
    FOREIGN KEY (contact_id) REFERENCES contacts (id)
) ENGINE InnoDB;