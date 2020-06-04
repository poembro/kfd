/*
Navicat MySQL Data Transfer

Source Server         : sh-cdb-dw5p0dpo.sql.tencentcdb.com:61740
Source Server Version : 50718
Source Host           : sh-cdb-dw5p0dpo.sql.tencentcdb.com:61740
Source Database       : db_scmj

Target Server Type    : MYSQL
Target Server Version : 50718
File Encoding         : 65001

Date: 2020-06-03 20:10:50
*/

SET FOREIGN_KEY_CHECKS=0;
-- ----------------------------
-- Table structure for `kfd_goods`
-- ----------------------------
DROP TABLE IF EXISTS `kfd_goods`;
CREATE TABLE `kfd_goods` (
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='商品表';

-- ----------------------------
-- Records of kfd_goods
-- ----------------------------
INSERT INTO `kfd_goods` VALUES ('1', '1', '描述', '1', '11', '1', '1.00', '1.00', '11', '1', '1', '1', '11', '1', '11', '2020-06-02 23:09:58', '2020-06-02 23:09:58', '1');

-- ----------------------------
-- Table structure for `kfd_goods_category`
-- ----------------------------
DROP TABLE IF EXISTS `kfd_goods_category`;
CREATE TABLE `kfd_goods_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名称',
  `sort_id` int(10) NOT NULL COMMENT '排序字段',
  `visible` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否显示 1-是 0-否',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='ren表';

-- ----------------------------
-- Records of kfd_goods_category
-- ----------------------------
INSERT INTO `kfd_goods_category` VALUES ('1', '分类', '1', '1', '2020-06-03 13:01:16', '2020-06-03 13:01:16');

-- ----------------------------
-- Table structure for `kfd_order`
-- ----------------------------
DROP TABLE IF EXISTS `kfd_order`;
CREATE TABLE `kfd_order` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单id',
  `order_sn` varchar(20) NOT NULL DEFAULT '' COMMENT '订单号',
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户uid',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '订单状态 0-待付款 1-已付款 2-待配送  3-配送中  4-已完成 5-已取消',
  `payid` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '付款方式 1-建行',
  `paytime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '付款时间',
  `ispay` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否付款 1-是  0-否',
  `shiptime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '发货时间',
  `finished_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '订单完成时间',
  `refund_msg` varchar(100) NOT NULL DEFAULT '' COMMENT '退款提示',
  `cancel_type` int(10) NOT NULL DEFAULT '0' COMMENT '取消原因 1-手机号码不是本人 2-用户要求取消 3-其它',
  `cancel_remark` varchar(200) NOT NULL DEFAULT '' COMMENT '取消备注',
  `realname` varchar(20) NOT NULL DEFAULT '' COMMENT '姓名',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号码',
  `shopname` varchar(255) NOT NULL,
  `address` varchar(200) NOT NULL DEFAULT '' COMMENT '地址',
  `total_price` bigint(10) NOT NULL DEFAULT '0' COMMENT '总价格',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '买家留言',
  `score` tinyint(1) NOT NULL DEFAULT '5' COMMENT '分数',
  `visible` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否显示 1-是 0-否',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_sn` (`order_sn`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='订单表';

-- ----------------------------
-- Records of kfd_order
-- ----------------------------
INSERT INTO `kfd_order` VALUES ('1', 'O_123456', '2561', '1', '1', '0', '0', '0', '0', '', '0', '', '', '', 'xxx店铺', '', '0', '', '5', '1', '2020-06-03 16:11:26', '2020-06-03 16:29:36');

-- ----------------------------
-- Table structure for `kfd_order_goods`
-- ----------------------------
DROP TABLE IF EXISTS `kfd_order_goods`;
CREATE TABLE `kfd_order_goods` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单id',
  `order_sn` varchar(20) NOT NULL DEFAULT '' COMMENT '订单号',
  `goods_id` int(10) NOT NULL DEFAULT '0' COMMENT '商品ID',
  `refund_msg` varchar(100) NOT NULL DEFAULT '' COMMENT '退款提示',
  `cancel_type` int(10) NOT NULL DEFAULT '0' COMMENT '取消原因 1-手机号码不是本人 2-用户要求取消 3-其它',
  `cancel_remark` varchar(200) NOT NULL DEFAULT '' COMMENT '取消备注',
  `goods_name` varchar(100) NOT NULL DEFAULT '' COMMENT '商品名称',
  `thumb` varchar(200) NOT NULL DEFAULT '' COMMENT '缩略图',
  `goods_num` int(10) NOT NULL DEFAULT '1' COMMENT '数量',
  `goods_price` bigint(10) NOT NULL DEFAULT '0' COMMENT '价格',
  `goods_spec` int(10) NOT NULL DEFAULT '0' COMMENT '规格',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '订单状态 0-待付款 1-已付款 2-待配送  3-配送中  4-已完成 5-已取消',
  `visible` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否显示 1-是 0-否',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `order_sn` (`order_sn`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 COMMENT='订单商品表';

-- ----------------------------
-- Records of kfd_order_goods
-- ----------------------------
INSERT INTO `kfd_order_goods` VALUES ('1', 'O_123456', '1', '', '0', '', '商品名字2', '', '1', '0', '0', '0', '1', '2020-06-03 16:11:50', '2020-06-03 16:12:54');
INSERT INTO `kfd_order_goods` VALUES ('2', 'O_123456', '2', 'x', '0', '', '商品名字1', '', '1', '0', '0', '0', '1', '2020-06-03 16:12:46', '2020-06-03 16:12:46');

-- ----------------------------
-- Table structure for `kfd_uid`
-- ----------------------------
DROP TABLE IF EXISTS `kfd_uid`;
CREATE TABLE `kfd_uid` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `business_id` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT '业务id',
  `max_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '最大id',
  `step` int(10) unsigned NOT NULL DEFAULT '1000' COMMENT '步长',
  `description` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '描述',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_business_id` (`business_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='分布式自增主键';

-- ----------------------------
-- Records of kfd_uid
-- ----------------------------
INSERT INTO `kfd_uid` VALUES ('1', 'device_id', '2760', '5', '设备id', '2019-10-15 16:42:05', '2020-06-03 20:06:46');

-- ----------------------------
-- Table structure for `kfd_user`
-- ----------------------------
DROP TABLE IF EXISTS `kfd_user`;
CREATE TABLE `kfd_user` (
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
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户';

-- ----------------------------
-- Records of kfd_user
-- ----------------------------
INSERT INTO `kfd_user` VALUES ('29', '1', '2361', '昵称', '1', '/#', '#', '2020-04-04 18:24:58', '2020-04-04 18:24:58', '13000000000', 'd26b8235af26d2b3b21d654cd564afef', '0');
INSERT INTO `kfd_user` VALUES ('30', '1', '2362', 'poembro', '1', 'https://wx.qlogo.cn/mmopen/vi_32/GPXCzVBricXuTaH3H', '#', '2020-04-04 18:33:53', '2020-04-04 18:33:53', '', '', '0');
INSERT INTO `kfd_user` VALUES ('31', '1', '2561', '昵称', '1', '/#', '#', '2020-06-03 16:16:12', '2020-06-03 16:16:12', '13200000001', 'f3b6980ed0155be02a45885a3349faeb', '0');

-- ----------------------------
-- Table structure for `kfd_user_third`
-- ----------------------------
DROP TABLE IF EXISTS `kfd_user_third`;
CREATE TABLE `kfd_user_third` (
  `user_id` bigint(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户UID',
  `typeid` tinyint(1) NOT NULL DEFAULT '1' COMMENT '第三方类型 1-Facebook',
  `openid` varchar(40) NOT NULL DEFAULT '' COMMENT 'openid',
  PRIMARY KEY (`user_id`,`typeid`),
  KEY `openid` (`openid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='第三方登陆表';

-- ----------------------------
-- Records of kfd_user_third
-- ----------------------------
INSERT INTO `kfd_user_third` VALUES ('2362', '1', 'o6d0s5CSClLxMfnTpYrX55U3Cdyk');
