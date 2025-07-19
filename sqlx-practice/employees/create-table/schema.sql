-- 创建students表
CREATE TABLE if not exists employees
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(50)    NOT NULL,
    department VARCHAR(50)    NOT NULL,
    salary     DECIMAL(10, 2) NOT NULL
);
-- 清空表（可选）
TRUNCATE TABLE employees;

-- 插入20条随机员工数据
INSERT INTO employees (name, department, salary)
VALUES ('张三', '研发部', ROUND(3000 + RAND() * 47000, 2)),
       ('李四', '市场部', ROUND(3000 + RAND() * 47000, 2)),
       ('王五', '财务部', ROUND(3000 + RAND() * 47000, 2)),
       ('赵六', '人事部', ROUND(3000 + RAND() * 47000, 2)),
       ('钱七', '行政部', ROUND(3000 + RAND() * 47000, 2)),
       ('孙八', '销售部', ROUND(3000 + RAND() * 47000, 2)),
       ('周九', '客服部', ROUND(3000 + RAND() * 47000, 2)),
       ('吴十', '技术部', ROUND(3000 + RAND() * 47000, 2)),
       ('小明', '产品部', ROUND(3000 + RAND() * 47000, 2)),
       ('小红', '设计部', ROUND(3000 + RAND() * 47000, 2)),
       ('小刚', '运营部', ROUND(3000 + RAND() * 47000, 2)),
       ('小丽', '法务部', ROUND(3000 + RAND() * 47000, 2)),
       ('马化云', '总裁办', ROUND(30000 + RAND() * 20000, 2)),
       ('刘强西', '投资部', ROUND(30000 + RAND() * 20000, 2)),
       ('王石头', '战略部', ROUND(30000 + RAND() * 20000, 2)),
       ('雷小军', '创新部', ROUND(30000 + RAND() * 20000, 2)),
       ('丁三石', '国际部', ROUND(30000 + RAND() * 20000, 2)),
       ('张朝阳', '公关部', ROUND(30000 + RAND() * 20000, 2)),
       ('李彦红', 'AI实验室', ROUND(30000 + RAND() * 20000, 2)),
       ('黄征', '电商部', ROUND(30000 + RAND() * 20000, 2));

-- 查看生成的数据
SELECT *
FROM employees;
