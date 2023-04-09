-- Create Table Member
CREATE TABLE members (
    id INT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    gender VARCHAR(255) NOT NULL,
    skin_type VARCHAR(255) NOT NULL,
    skin_color VARCHAR(255) NOT NULL
);

-- Create Table Product
CREATE TABLE products (
    id INT PRIMARY KEY,
    name_product VARCHAR(255) NOT NULL,
    price INT NOT NULL
);

-- Create Table Review Product
CREATE TABLE review_products (
    id INT PRIMARY KEY,
    product_id INT,
    member_id INT,
    desc_review VARCHAR(255),
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (member_id) REFERENCES members(id)
);

-- Create Table Like Review
CREATE TABLE like_reviews (
    id INT PRIMARY KEY,
    review_id INT,
    member_id INT,
    FOREIGN KEY (review_id) REFERENCES review_products(id),
    FOREIGN KEY (member_id) REFERENCES members(id)
);