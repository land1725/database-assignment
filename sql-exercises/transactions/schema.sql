-- 创建accounts表
CREATE TABLE if not exists accounts
(
    id      INT AUTO_INCREMENT PRIMARY KEY,
    balance DECIMAL(10, 2) NOT NULL
);

-- 创建transactions表
CREATE TABLE if not exists transactions
(
    id              INT AUTO_INCREMENT PRIMARY KEY,
    from_account_id INT            NOT NULL,
    to_account_id   INT            NOT NULL,
    amount          DECIMAL(10, 2) NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO accounts (id, balance)
VALUES (1, 500),
       (2, 300)