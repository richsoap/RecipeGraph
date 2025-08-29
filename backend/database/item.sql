-- Items table
-- 包含以下内容：物品ID、物品名称（支持中文）、创建时间、修改时间
CREATE TABLE IF NOT EXISTS items (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name TEXT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
