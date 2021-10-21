/*
Navicat MySQL Data Transfer

Source Server         : 192.168.99.100
Source Server Version : 50650
Source Host           : 192.168.99.100:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50650
File Encoding         : 65001

Date: 2021-10-21 13:03:39
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `casbin_rule`
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=647 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('381', 'g', '1', '40', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('399', 'g', '4', '39', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('645', 'p', '39', '/admin/changeOwnInfo', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('646', 'p', '39', '/admin/changeOwnPassword', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('641', 'p', '39', '/admin/getRules', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('643', 'p', '39', '/admin/getUsers', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('637', 'p', '39', '/admin/index', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('636', 'p', '39', '/admin/main', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('639', 'p', '39', '/admin/menusView', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('640', 'p', '39', '/admin/rulesView', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('642', 'p', '39', '/admin/usersView', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('644', 'p', '39', '/admin/userView', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('638', 'p', '39', 'javascript:void(0)', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('634', 'p', '40', '/admin/changeOwnInfo', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('635', 'p', '40', '/admin/changeOwnPassword', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('630', 'p', '40', '/admin/changePassword/:id', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('614', 'p', '40', '/admin/createBaseMenu', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('615', 'p', '40', '/admin/createChildMenu/:parent_id', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('620', 'p', '40', '/admin/createRule', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('627', 'p', '40', '/admin/createUser', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('622', 'p', '40', '/admin/deleteRule/:id', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('628', 'p', '40', '/admin/deleteUser/:id', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('619', 'p', '40', '/admin/getRules', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('626', 'p', '40', '/admin/getUsers', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('611', 'p', '40', '/admin/index', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('610', 'p', '40', '/admin/main', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('613', 'p', '40', '/admin/menusView', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('623', 'p', '40', '/admin/ruleRbacViwe/:id', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('618', 'p', '40', '/admin/rulesView', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('632', 'p', '40', '/admin/setUserRules', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('616', 'p', '40', '/admin/updateMenu/:id', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('624', 'p', '40', '/admin/updateRbac', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('621', 'p', '40', '/admin/updateRule/:id', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('629', 'p', '40', '/admin/updateUser/:id', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('631', 'p', '40', '/admin/userRuleView/:id', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('625', 'p', '40', '/admin/usersView', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('633', 'p', '40', '/admin/userView', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('617', 'p', '40', '/deleteMenu/:id', 'get|post', '', '', '');
INSERT INTO `casbin_rule` VALUES ('612', 'p', '40', 'javascript:void(0)', 'get|post', '', '', '');

-- ----------------------------
-- Table structure for `sys_menus`
-- ----------------------------
DROP TABLE IF EXISTS `sys_menus`;
CREATE TABLE `sys_menus` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `parent_id` int(10) NOT NULL DEFAULT '0' COMMENT '父菜单ID',
  `path` varchar(191) NOT NULL COMMENT '路由path',
  `name` varchar(191) NOT NULL COMMENT '路由name',
  `hidden` tinyint(1) NOT NULL COMMENT '是否在列表隐藏',
  `sort` int(10) NOT NULL DEFAULT '1' COMMENT '排序标记',
  `is_view` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否视图',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`),
  KEY `idx_parent_id` (`parent_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
INSERT INTO `sys_menus` VALUES ('14', '2021-09-28 21:18:34', '2021-09-28 21:18:34', null, '29', '/admin/menusView', '菜单管理', '0', '8', '1');
INSERT INTO `sys_menus` VALUES ('15', '2021-09-29 15:45:26', '2021-10-13 15:51:03', null, '14', '/admin/createBaseMenu', '新增基础菜单', '0', '8', '0');
INSERT INTO `sys_menus` VALUES ('16', '2021-09-29 20:39:05', '2021-09-30 16:24:02', null, '29', '/admin/rulesView', '角色管理', '0', '2', '1');
INSERT INTO `sys_menus` VALUES ('17', '2021-09-29 20:49:06', '2021-09-29 21:18:12', '2021-09-30 14:29:23', '0', '/admin/createMenu', '新增菜单2', '0', '10', '1');
INSERT INTO `sys_menus` VALUES ('18', '2021-09-29 20:55:42', '2021-09-30 14:58:01', null, '16', '/admin/getRules', '角色列表信息', '0', '3', '0');
INSERT INTO `sys_menus` VALUES ('19', '2021-09-29 21:10:02', '2021-09-29 21:10:02', '2021-09-30 14:53:13', '0', '/admin/createMenu', '新增菜单222', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('20', '2021-09-29 21:12:16', '2021-10-18 19:12:41', null, '16', '/admin/updateRule/:id', '更新角色', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('21', '2021-09-30 14:47:24', '2021-09-30 14:47:24', null, '14', '/admin/createChildMenu/:parent_id', '创建子菜单', '0', '2', '0');
INSERT INTO `sys_menus` VALUES ('22', '2021-09-30 14:53:50', '2021-10-19 18:08:13', null, '0', '/admin/index', '框架首页', '0', '15', '0');
INSERT INTO `sys_menus` VALUES ('23', '2021-09-30 14:55:38', '2021-09-30 14:55:38', '2021-09-30 14:56:11', '15', '/admin/updateMenu/:id', '更新菜单', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('24', '2021-09-30 14:56:26', '2021-09-30 14:56:58', null, '14', '/admin/updateMenu/:id', '更新菜单', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('25', '2021-09-30 14:56:51', '2021-09-30 14:56:51', null, '14', '/deleteMenu/:id', '删除菜单', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('26', '2021-09-30 14:57:38', '2021-09-30 14:57:52', null, '16', '/admin/createRule', '创建角色', '0', '2', '0');
INSERT INTO `sys_menus` VALUES ('27', '2021-09-30 15:00:42', '2021-09-30 15:00:42', null, '16', '/admin/deleteRule/:id', '删除角色', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('28', '2021-09-30 15:01:25', '2021-10-12 16:44:50', null, '29', '/admin/ruleRbacViwe/:id', '权限管理', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('29', '2021-09-30 15:49:53', '2021-10-13 15:46:54', null, '0', 'javascript:void(0)', '系统管理', '0', '0', '1');
INSERT INTO `sys_menus` VALUES ('30', '2021-09-30 16:31:19', '2021-09-30 16:31:19', '2021-09-30 17:28:12', '14', '/admin/menusView	', '菜单列表', '0', '0', '1');
INSERT INTO `sys_menus` VALUES ('31', '2021-09-30 17:08:18', '2021-09-30 17:26:26', '2021-09-30 17:26:38', '0', '/admin/index', '一级菜单', '1', '0', '1');
INSERT INTO `sys_menus` VALUES ('32', '2021-09-30 17:08:33', '2021-09-30 17:08:44', '2021-09-30 17:08:49', '31', '/', '二级菜单', '0', '2', '1');
INSERT INTO `sys_menus` VALUES ('33', '2021-10-13 15:48:44', '2021-10-13 15:48:44', null, '29', '/admin/usersView', '用户管理', '0', '0', '1');
INSERT INTO `sys_menus` VALUES ('34', '2021-10-13 15:50:23', '2021-10-13 15:50:23', null, '33', '/admin/getUsers', '用户列表信息', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('35', '2021-10-13 15:52:35', '2021-10-13 15:52:35', null, '28', '/admin/updateRbac', '设置角色权限', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('36', '2021-10-14 11:46:59', '2021-10-14 11:46:59', null, '33', '/admin/createUser', '新建用户', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('37', '2021-10-14 18:50:02', '2021-10-14 18:50:02', null, '33', '/admin/deleteUser/:id', '删除用户', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('38', '2021-10-14 18:50:55', '2021-10-14 18:50:55', null, '33', '/admin/updateUser/:id', '编辑用户', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('39', '2021-10-15 10:34:03', '2021-10-15 10:34:03', null, '33', '/admin/changePassword/:id', '重置管理员密码', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('40', '2021-10-15 10:34:24', '2021-10-15 10:34:24', '2021-10-19 16:28:07', '33', '/admin/setAdmin/:id', '设置超管', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('41', '2021-10-15 13:09:16', '2021-10-15 13:09:16', null, '33', '/admin/userRuleView/:id', '用户配置角色页', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('42', '2021-10-15 13:10:43', '2021-10-15 13:40:25', null, '33', '/admin/setUserRules', '分配角色', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('43', '2021-10-19 15:47:18', '2021-10-19 16:07:30', null, '33', '/admin/userView', '个人信息页', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('44', '2021-10-19 15:48:17', '2021-10-19 15:48:17', null, '33', '/admin/changeOwnInfo', '更新个人信息', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('45', '2021-10-19 16:31:55', '2021-10-19 16:31:55', null, '33', '/admin/changeOwnPassword', '重置个人密码', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('46', '2021-10-19 16:44:27', '2021-10-19 16:44:27', '2021-10-19 16:46:34', '22', '/admin/main', 'main页', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('47', '2021-10-19 17:10:51', '2021-10-19 17:10:51', '2021-10-19 18:08:22', '22', '/admin/main', 'main页', '0', '0', '0');
INSERT INTO `sys_menus` VALUES ('48', '2021-10-19 18:08:43', '2021-10-19 18:08:43', null, '0', '/admin/main', '后台首页', '0', '40', '1');

-- ----------------------------
-- Table structure for `sys_rules`
-- ----------------------------
DROP TABLE IF EXISTS `sys_rules`;
CREATE TABLE `sys_rules` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `role_name` char(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '角色名称(组名)',
  `role_desc` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '角色描述',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用：0-禁用，1-启用',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COMMENT='用户角色表';

-- ----------------------------
-- Records of sys_rules
-- ----------------------------
INSERT INTO `sys_rules` VALUES ('37', '菜单管理员', '菜单管理员', '1', '2021-09-26 09:57:34', '2021-10-18 19:29:00', '2021-10-19 16:04:12');
INSERT INTO `sys_rules` VALUES ('38', '小程序管理员', '小程序管理全部权限', '1', '2021-09-28 16:23:05', '2021-09-28 16:23:05', '2021-10-19 16:27:33');
INSERT INTO `sys_rules` VALUES ('39', '基础权限', '最基本的角色', '1', '2021-10-19 16:01:54', '2021-10-19 16:01:54', null);
INSERT INTO `sys_rules` VALUES ('40', '超级管理员', '全部权限', '1', '2021-10-19 16:02:29', '2021-10-19 16:02:29', null);

-- ----------------------------
-- Table structure for `sys_users`
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `real_name` varchar(255) NOT NULL COMMENT '真实姓名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '使用状态 0 禁用 1 使用中',
  `type` tinyint(1) NOT NULL DEFAULT '2' COMMENT '1 超级管理员 2 普通管理员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES ('1', '0000-00-00 00:00:00', '2021-10-19 16:28:34', null, 'admin', '管理员', 'e64b78fc3bc91bcbc7dc232ba8ec59e0', '1', '1');
INSERT INTO `sys_users` VALUES ('2', '2021-10-14 11:57:38', '2021-10-14 11:57:38', '2021-10-14 18:53:58', '白翀', '', 'e10adc3949ba59abbe56e057f20f883e', '1', '2');
INSERT INTO `sys_users` VALUES ('3', '2021-10-14 18:36:56', '2021-10-14 18:36:56', '2021-10-14 18:54:04', 'chonghua', '白', '1692fcfff3e01e7ba8cffc2baadef5f5', '1', '2');
INSERT INTO `sys_users` VALUES ('4', '2021-10-14 18:46:01', '2021-10-19 16:32:39', null, 'tester', '测试专用', '6aca4053faeb36da0cf94aca1bd89d77', '1', '2');
INSERT INTO `sys_users` VALUES ('5', '2021-10-19 15:54:49', '2021-10-19 15:54:49', '2021-10-19 16:00:55', 'chonghua', '张英强', 'fbbb701101c1ae93a7efb4713a71ab76', '1', '2');
