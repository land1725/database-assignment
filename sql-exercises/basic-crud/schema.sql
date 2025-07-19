-- 创建students表
CREATE TABLE students
(
    id    INT AUTO_INCREMENT PRIMARY KEY,
    name  VARCHAR(50) NOT NULL,
    age   INT         NOT NULL,
    grade VARCHAR(20) NOT NULL
);
