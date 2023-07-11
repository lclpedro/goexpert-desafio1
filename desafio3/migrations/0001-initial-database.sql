CREATE TABLE orders (
    id CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(12,2) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO orders (id, name, price) VALUES
 ('d29247a7-4eae-4e21-8a3c-7294cc0f1d1d', 'Produto A', 19.99),
 ('7a131ab5-9b7d-4c8d-8e34-0a8e89f23aef', 'Produto B', 29.99),
 ('17a8a44f-7894-4866-96f8-6785a73a60fd', 'Produto C', 9.99),
 ('e4a7f231-9c74-47d4-aa7d-162f5d9d532b', 'Produto D', 15.99),
 ('d6d7e671-ae75-41de-82ae-3996181fe5a7', 'Produto E', 49.99),
 ('06e66e7c-514a-4529-af1a-c6681cb48627', 'Produto F', 12.99),
 ('f2e1590b-3c3c-43ff-9c1f-45fd82b1a9e4', 'Produto G', 24.99),
 ('c8d55a8e-3b61-45c4-99a6-5fe4a24a999e', 'Produto H', 8.99),
 ('4e74ebfc-3a79-45cc-bb9d-c6aef12a527d', 'Produto I', 39.99),
 ('92a9a245-8c3c-432e-93f3-0f99ef242968', 'Produto J', 17.99);