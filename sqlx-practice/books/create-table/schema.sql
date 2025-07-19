-- 创建students表
CREATE TABLE if not exists books
(
    id     INT AUTO_INCREMENT PRIMARY KEY,
    title  VARCHAR(50)    NOT NULL,
    author VARCHAR(50)    NOT NULL,
    price  DECIMAL(10, 2) NOT NULL
);
-- 清空表（可选）
TRUNCATE TABLE books;

-- 插入20条随机书本数据（价格范围：10~200）
INSERT INTO books (title, author, price)
VALUES ('数据库系统概念', '王小明', ROUND(10 + RAND() * 190, 2)),
       ('Java编程思想', '李华', ROUND(10 + RAND() * 190, 2)),
       ('算法导论', '张伟', ROUND(10 + RAND() * 190, 2)),
       ('Python数据分析', '刘芳', ROUND(10 + RAND() * 190, 2)),
       ('C++ Primer', '赵强', ROUND(10 + RAND() * 190, 2)),
       ('深入浅出MySQL', '周杰', ROUND(10 + RAND() * 190, 2)),
       ('Go语言实战', '陈晨', ROUND(10 + RAND() * 190, 2)),
       ('Redis设计与实现', '吴磊', ROUND(10 + RAND() * 190, 2)),
       ('Linux命令行大全', '郑爽', ROUND(10 + RAND() * 190, 2)),
       ('计算机网络', '孙阳', ROUND(10 + RAND() * 190, 2)),
       ('操作系统精髓', '钱多多', ROUND(10 + RAND() * 190, 2)),
       ('前端开发艺术', '马云飞', ROUND(10 + RAND() * 190, 2)),
       ('人工智能基础', '马化腾', ROUND(10 + RAND() * 190, 2)),
       ('机器学习实战', '李彦宏', ROUND(10 + RAND() * 190, 2)),
       ('深度学习入门', '雷军', ROUND(10 + RAND() * 190, 2)),
       ('区块链技术指南', '周鸿祎', ROUND(10 + RAND() * 190, 2)),
       ('大数据处理', '丁磊', ROUND(10 + RAND() * 190, 2)),
       ('软件工程实践', '黄峥', ROUND(10 + RAND() * 190, 2)),
       ('计算机组成原理', '王兴', ROUND(10 + RAND() * 190, 2)),
       ('编译原理', '张一鸣', ROUND(10 + RAND() * 190, 2));

-- 查看插入的数据
SELECT *
FROM books;
