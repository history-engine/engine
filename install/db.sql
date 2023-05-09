CREATE TABLE `page` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `user_id` bigint(21) UNSIGNED NOT NULL COMMENT '用户id',
    `unique_id` char(32) NOT NULL DEFAULT '' COMMENT '页面唯一id',
    `version` int(11) UNSIGNED NOT NULL DEFAULT '1' COMMENT '版本',
    `title` varchar(300) NOT NULL DEFAULT '' COMMENT '页面标题',
    `url` varchar(2048) NOT NULL DEFAULT '' COMMENT '原始地址',
    `full_path` varchar(500) NOT NULL DEFAULT '' COMMENT '完整本地文件地址',
    `full_size` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '文件大小',
    `lite_path` varchar(500) NOT NULL DEFAULT '' COMMENT '提取后文件地址',
    `lite_size` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '提取后文件大小',
    `indexed_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后索引时间',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '入库时间',
    `updated_at` datetime NOT NULL COMMENT '最后更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='页面';

CREATE TABLE `user` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `username` varchar(50) NOT NULL COMMENT '用户名',
    `email` varchar(100) NOT NULL COMMENT '邮箱',
    `password` char(32) NOT NULL COMMENT '密码',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL COMMENT '最后更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';

ALTER TABLE `page`
    ADD PRIMARY KEY (`id`),
    ADD UNIQUE KEY `user` (`user_id`),
    ADD UNIQUE KEY `page_version` (`unique_id`,`version`);

ALTER TABLE `user`
    ADD PRIMARY KEY (`id`),
    ADD UNIQUE KEY `username` (`username`),
    ADD UNIQUE KEY `email` (`email`);

