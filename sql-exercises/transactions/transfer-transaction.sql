DROP PROCEDURE IF EXISTS transfer_money;
DELIMITER //
CREATE PROCEDURE transfer_money(IN from_id INT,
                                IN to_id INT,
                                IN money DECIMAL(10, 2))
BEGIN
    DECLARE v_balance DECIMAL(10, 2);
    DECLARE from_exists INT;
    DECLARE to_exists INT;
    SELECT COUNT(*) INTO from_exists FROM accounts WHERE id = from_id;
    SELECT COUNT(*) INTO to_exists FROM accounts WHERE id = from_id;
    IF from_exists != 0 AND to_exists != 0 AND money > 0 THEN
        START TRANSACTION;
        SELECT balance INTO v_balance FROM accounts WHERE id = from_id FOR UPDATE;
        IF v_balance >= money THEN
            UPDATE accounts SET balance = balance - money WHERE id = from_id;
            UPDATE accounts SET balance = balance + money WHERE id = to_id;
            INSERT INTO transactions(from_account_id, to_account_id, amount) VALUES (from_id, to_id, money);
            COMMIT;
        ELSE
            ROLLBACK;
        END IF;
    END IF;
END //
DELIMITER ;

CALL transfer_money(1, 2, 100);
