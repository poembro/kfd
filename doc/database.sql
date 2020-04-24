

## mysql表
```
-- ----------------------------
-- Table structure for uid
-- ----------------------------
DROP TABLE IF EXISTS `uid`;
CREATE TABLE `uid`
(
    `id`          bigint(20) unsigned              NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `business_id` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT '业务id',
    `max_id`      bigint(20) unsigned              NOT NULL DEFAULT '0' COMMENT '最大id',
    `step`        int(10) unsigned                 NOT NULL DEFAULT '1000' COMMENT '步长',
    `description` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '描述',
    `create_time` datetime                         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime                         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_business_id` (`business_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='分布式自增主键';

-- ----------------------------
-- Records of uid
-- ----------------------------
BEGIN;
INSERT INTO `uid`
VALUES (1, 'device_id', 1580, 5, '设备id', '2019-10-15 16:42:05', '2019-12-24 14:23:13');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `app_id` bigint(20) unsigned NOT NULL COMMENT 'app_id',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户id',
  `nickname` varchar(20) COLLATE utf8mb4_bin NOT NULL COMMENT '昵称',
  `sex` tinyint(4) NOT NULL COMMENT '性别，0:未知；1:男；2:女',
  `avatar_url` varchar(50) COLLATE utf8mb4_bin NOT NULL COMMENT '用户头像链接',
  `extra` varchar(1024) COLLATE utf8mb4_bin NOT NULL COMMENT '附加属性',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `account` varchar(11) COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT 'Hash索引只能使用=或<=>(NULL比较符号),不能使用范围查询',
  `password` varchar(200) COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  `visible` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_app_id_user_id` (`app_id`,`user_id`),
  UNIQUE KEY `account` (`account`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户';

 
CREATE TABLE `user_third` (
  `user_id` bigint(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户UID',
  `typeid` tinyint(1) NOT NULL DEFAULT '1' COMMENT '第三方类型 1-Facebook',
  `openid` varchar(40) NOT NULL DEFAULT '' COMMENT 'openid',
  PRIMARY KEY (`user_id`,`typeid`),
  KEY `openid` (`openid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='第三方登陆表';

CREATE TABLE `goods_category` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT COMMENT '类型id',
  `name` varchar(60) NOT NULL DEFAULT '' COMMENT '类型名',
  `visible` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态 1-显示 0-删除',
  `dateline` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `updatetime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `sortid` int(10) NOT NULL DEFAULT '0' COMMENT '排序id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品类型表';

CREATE TABLE `googs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称',
  `description` varchar(300) NOT NULL DEFAULT '描述' COMMENT '描述',
  `category_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '分类id',
  `is_on_sale` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '上架状态 1-上架 0-下架',
  `thumb` text NOT NULL COMMENT '商品缩略图',
  `goods_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '商品价格',
  `market_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '市场价格',
  `sells_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '销售数',
  `comments_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '评价数',
  `favorites_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '收藏数',
  `is_hot` tinyint(1) NOT NULL DEFAULT '0' COMMENT '推荐，首页显示',
  `goods_sort` int(10) NOT NULL COMMENT '排序字段',
  `inventory` int(10) unsigned NOT NULL DEFAULT '0',
  `visible` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否显示 1-是 0-否',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `goods_sn` varchar(22) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品表';
