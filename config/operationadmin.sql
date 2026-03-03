/*
SQLyog Ultimate v13.1.1 (64 bit)
MySQL - 5.7.20-log : Database - operationadmin
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`operationadmin` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `operationadmin`;

/*Table structure for table `sys_dept` */

DROP TABLE IF EXISTS `sys_dept`;

CREATE TABLE `sys_dept` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(64) DEFAULT '' COMMENT '部门名称',
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父节点id',
  `tree_path` varchar(255) DEFAULT '' COMMENT '父节点id路径',
  `sort` int(11) DEFAULT '0' COMMENT '显示顺序',
  `status` tinyint(4) DEFAULT '1' COMMENT '状态(1:正常;0:禁用)',
  `deleted` tinyint(4) DEFAULT '0' COMMENT '逻辑删除标识(1:已删除;0:未删除)',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='部门表';

/*Data for the table `sys_dept` */

insert  into `sys_dept`(`id`,`name`,`parent_id`,`tree_path`,`sort`,`status`,`deleted`,`create_time`,`update_time`) values 
(1,'全部',0,'0',1,1,0,NULL,'2023-04-24 21:50:25'),
(2,'管理员组',1,'0,1',1,1,0,NULL,'2022-04-19 12:46:37'),
(3,'普通用户组',1,'0,1',1,1,0,NULL,'2022-04-19 12:46:37');

/*Table structure for table `sys_dict` */

DROP TABLE IF EXISTS `sys_dict`;

CREATE TABLE `sys_dict` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `type_code` varchar(64) DEFAULT NULL COMMENT '字典类型编码',
  `name` varchar(50) DEFAULT '' COMMENT '字典项名称',
  `value` varchar(50) DEFAULT '' COMMENT '字典项值',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态(1:正常;0:禁用)',
  `defaulted` tinyint(4) DEFAULT '0' COMMENT '是否默认(1:是;0:否)',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='字典数据表';

/*Data for the table `sys_dict` */

insert  into `sys_dict`(`id`,`type_code`,`name`,`value`,`sort`,`status`,`defaulted`,`remark`,`create_time`,`update_time`) values 
(1,'gender','男','1',1,1,0,NULL,'2019-05-05 13:07:52','2022-06-12 23:20:39'),
(2,'gender','女','2',2,1,0,NULL,'2019-04-19 11:33:00','2019-07-02 14:23:05'),
(3,'gender','未知','0',1,1,0,NULL,'2020-10-17 08:09:31','2020-10-17 08:09:31');

/*Table structure for table `sys_dict_type` */

DROP TABLE IF EXISTS `sys_dict_type`;

CREATE TABLE `sys_dict_type` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键 ',
  `name` varchar(50) DEFAULT '' COMMENT '类型名称',
  `code` varchar(50) DEFAULT '' COMMENT '类型编码',
  `status` tinyint(1) DEFAULT '0' COMMENT '状态(0:正常;1:禁用)',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `type_code` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='字典类型表';

/*Data for the table `sys_dict_type` */

insert  into `sys_dict_type`(`id`,`name`,`code`,`status`,`remark`,`create_time`,`update_time`) values 
(1,'性别','gender',1,NULL,'2019-12-06 19:03:32','2022-06-12 16:21:28');

/*Table structure for table `sys_domain` */

DROP TABLE IF EXISTS `sys_domain`;

CREATE TABLE `sys_domain` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `domain` varchar(64) NOT NULL COMMENT '域名',
  `keyword` varchar(64) NOT NULL COMMENT '关键词',
  `sort` int(11) NOT NULL COMMENT '显示顺序',
  `code` int(3) NOT NULL COMMENT '期望状态码',
  `status` tinyint(1) NOT NULL COMMENT '监控开关',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `deleted` tinyint(1) DEFAULT '0' COMMENT '逻辑删除标识(0-未删除；1-已删除)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_domain` */

insert  into `sys_domain`(`id`,`domain`,`keyword`,`sort`,`code`,`status`,`create_time`,`update_time`,`deleted`) values 
(3,'http://dotsblog6.com','难受',1,202,1,'2023-03-31 15:53:33','2023-09-08 16:39:38',0),
(4,'http://sajkfaslfj.cosss','guanjian',0,201,1,'2023-09-08 16:19:48','2023-09-08 16:50:57',0);

/*Table structure for table `sys_group` */

DROP TABLE IF EXISTS `sys_group`;

CREATE TABLE `sys_group` (
  `id` tinyint(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '分组名称',
  `parent_id` tinyint(11) DEFAULT '0' COMMENT '父ID',
  `tree_path` varchar(128) DEFAULT '' COMMENT '树节点路径',
  `sort` tinyint(11) DEFAULT '1' COMMENT '排序',
  `status` tinyint(11) DEFAULT '1' COMMENT '状态',
  `folder_id` char(36) DEFAULT NULL COMMENT '文件夹ID',
  `deleted` tinyint(1) DEFAULT '0' COMMENT '逻辑删除(0:未删除,1:已删除)',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4;


/*Table structure for table `sys_host` */

DROP TABLE IF EXISTS `sys_host`;

CREATE TABLE `sys_host` (
  `id` tinyint(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `uid` char(36) NOT NULL DEFAULT '0' COMMENT 'grafana uid',
  `rule_uid` char(36) NOT NULL DEFAULT '0' COMMENT '告警策略 uid',
  `silence_uid` char(36) DEFAULT NULL COMMENT '静音策略 uid',
  `name` char(11) DEFAULT '' COMMENT '服务器名称',
  `ip_addr` char(15) NOT NULL COMMENT '主ip地址',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `group_id` tinyint(11) DEFAULT NULL COMMENT '分组ID',
  `remark` char(255) DEFAULT NULL COMMENT '备注',
  `is_alter` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否告警(0:否,1:是)',
  `flow_limit` tinyint(11) NOT NULL DEFAULT '5' COMMENT '流量告警值(单位:M)',
  `cpu_limit` tinyint(11) NOT NULL DEFAULT '80' COMMENT 'cpu告警峰值(单位:%)',
  `mem_limit` tinyint(11) NOT NULL DEFAULT '80' COMMENT '内存告警峰值(单位:%)',
  `load_limit` tinyint(11) NOT NULL DEFAULT '5' COMMENT '负载告警值',
  `disk_limit` tinyint(11) NOT NULL DEFAULT '80' COMMENT '硬盘告警值(单位:%)',
  `network_name` char(50) DEFAULT NULL COMMENT '网卡名称',
  `sort` int(11) NOT NULL DEFAULT '1' COMMENT '显示顺序',
  `delay_time` datetime DEFAULT NULL COMMENT '告警延迟时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除(0:未删除,1:已删除)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;


/*Table structure for table `sys_ipsec_domain` */

DROP TABLE IF EXISTS `sys_ipsec_domain`;

CREATE TABLE `sys_ipsec_domain` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `domain` varchar(64) NOT NULL COMMENT 'ipsec域名',
  `online_num` int(11) NOT NULL DEFAULT '0' COMMENT '在线人数',
  `alter_num` int(11) NOT NULL DEFAULT '30' COMMENT '告警人数',
  `sort` int(11) NOT NULL DEFAULT '1' COMMENT '排列顺序',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态',
  `rule_uid` char(36) DEFAULT NULL COMMENT '告警策略Uid',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `deleted` tinyint(1) DEFAULT '0' COMMENT '逻辑删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;


/*Table structure for table `sys_menu` */

DROP TABLE IF EXISTS `sys_menu`;

CREATE TABLE `sys_menu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) NOT NULL COMMENT '父菜单ID',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `type` tinyint(4) DEFAULT NULL COMMENT '菜单类型(1:菜单；2:目录；3:外链；4:按钮)',
  `path` varchar(128) DEFAULT '' COMMENT '路由路径(浏览器地址栏路径)',
  `component` varchar(128) DEFAULT NULL COMMENT '组件路径(vue页面完整路径，省略.vue后缀)',
  `perm` varchar(128) DEFAULT NULL COMMENT '权限标识',
  `visible` tinyint(1) NOT NULL DEFAULT '1' COMMENT '显示状态(1-显示;0-隐藏)',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `icon` varchar(64) DEFAULT '' COMMENT '菜单图标',
  `redirect_url` varchar(128) DEFAULT '' COMMENT '外链路径',
  `redirect` varchar(128) DEFAULT NULL COMMENT '跳转路径',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='菜单管理';

/*Data for the table `sys_menu` */

insert  into `sys_menu`(`id`,`parent_id`,`name`,`type`,`path`,`component`,`perm`,`visible`,`sort`,`icon`,`redirect_url`,`redirect`,`create_time`,`update_time`) values 
(1,0,'系统管理',2,'/system','Layout',NULL,1,1,'system','/system/user','/system/user','2021-08-28 09:12:21','2021-08-28 09:12:21'),
(2,1,'用户管理',1,'user','system/user/index',NULL,1,3,'user',NULL,NULL,'2021-08-28 09:12:21','2021-08-28 09:12:21'),
(3,1,'主机管理',1,'host','system/host/index',NULL,1,1,'host',NULL,NULL,'2021-08-28 09:12:21','2021-08-28 09:12:21'),
(4,1,'域名管理',1,'domain','system/domain/index',NULL,1,5,'domain',NULL,NULL,'2021-08-28 09:12:21','2021-08-28 09:12:21'),
(5,1,'部门管理',1,'dept','system/dept/index',NULL,1,4,'dept',NULL,NULL,'2021-08-28 09:12:21','2021-08-28 09:12:21'),
(6,1,'监控点管理',1,'node','system/node/index',NULL,1,6,'node',NULL,NULL,'2021-08-28 09:12:21','2021-08-28 09:12:21'),
(7,1,'分组管理',1,'group','system/group/index',NULL,1,2,'group',NULL,NULL,'2023-04-25 17:37:37','2023-04-25 17:37:38'),
(8,0,'统计管理',2,'/statistics','Layout',NULL,1,2,'server','/statistics/server','/statistics/server','2024-06-28 16:40:03','2024-06-28 16:40:05'),
(9,8,'服务器管理',1,'server','statistics/server/index',NULL,1,7,'server',NULL,NULL,'2023-10-25 17:00:44','2023-10-25 17:00:47'),
(10,8,'项目管理',1,'project','statistics/project/index',NULL,1,8,'project',NULL,NULL,'2023-10-26 11:18:08','2023-10-26 11:18:11'),
(11,8,'机房管理',1,'room','statistics/room/index',NULL,1,9,'room',NULL,NULL,'2023-10-26 11:19:17','2023-10-26 11:19:19'),
(12,8,'地区设置',1,'zone','statistics/zone/index',NULL,1,10,'zone','',NULL,'2023-10-26 15:33:46','2023-10-26 15:33:49'),
(13,8,'计费管理',1,'order','statistics/order/index',NULL,1,11,'order','',NULL,'2024-06-28 16:54:41','2024-06-28 16:54:43'),
(14,0,'监控管理',2,'/monitor','Layout',NULL,1,3,'monitor','/monitor/user','/monitor/user','2024-07-02 16:03:19','2024-07-02 16:03:21'),
(15,14,'IPsec管理',1,'ipsec','monitor/ipsec/index',NULL,1,14,'ipsec','',NULL,'2024-07-02 16:04:46','2024-07-02 16:04:48'),
(31,2,'用户新增',4,'',NULL,'sys:user:add',1,1,'','','','2022-10-23 11:04:08','2022-10-23 11:04:11'),
(32,2,'用户编辑',4,'',NULL,'sys:user:edit',1,2,'','','','2022-10-23 11:04:08','2022-10-23 11:04:11'),
(33,2,'用户删除',4,'',NULL,'sys:user:delete',1,3,'','','','2022-10-23 11:04:08','2022-10-23 11:04:11'),
(36,0,'组件封装',2,'/demo','Layout',NULL,1,10,'menu','','','2022-10-31 09:18:44','2022-10-31 09:18:47'),
(37,36,'富文本编辑器',1,'editor','demo/editor',NULL,1,1,'','','',NULL,NULL),
(38,36,'上传组件',1,'uploader','demo/uploader',NULL,1,2,'','','','2022-11-20 23:16:30','2022-11-20 23:16:32'),
(39,36,'图标选择器',1,'icon-selector','demo/icon-selector',NULL,1,3,'','','','2022-11-20 23:16:30','2022-11-20 23:16:32'),
(40,0,'接口',2,'/api','Layout',NULL,1,7,'api','','','2022-02-17 22:51:20','2022-02-17 22:51:20'),
(41,40,'接口文档',1,'apidoc','demo/apidoc',NULL,1,1,'api','','','2022-02-17 22:51:20','2022-02-17 22:51:20'),
(42,0,'被墙跳转',2,'/jump','Layout',NULL,1,11,'api','/jump/free','/jump/free','2023-03-28 21:52:36','2023-03-28 21:52:38'),
(43,42,'免费跳转',1,'free','jump/free/index',NULL,1,11,'api',NULL,NULL,'2023-03-28 21:54:22','2023-03-28 21:54:25'),
(44,0,'设置',2,'/settings','Layout',NULL,1,12,'api','/settings/global','/settings/global','2023-08-24 14:38:26','2023-08-24 14:38:28'),
(45,44,'全局设置',1,'global','settings/global/index',NULL,1,12,'api','',NULL,'2023-08-24 14:38:30','2023-08-24 14:38:33');

/*Table structure for table `sys_node` */

DROP TABLE IF EXISTS `sys_node`;

CREATE TABLE `sys_node` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `node_ip` varchar(15) NOT NULL COMMENT '监控点IP',
  `node_name` varchar(255) NOT NULL COMMENT '监控点名称',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `sort` int(11) NOT NULL COMMENT '排序',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `deleted` tinyint(1) DEFAULT '0' COMMENT '逻辑删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;


/*Table structure for table `sys_node_domain` */

DROP TABLE IF EXISTS `sys_node_domain`;

CREATE TABLE `sys_node_domain` (
  `id` tinyint(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `node_id` tinyint(11) NOT NULL COMMENT '监控点id',
  `domain_id` tinyint(11) NOT NULL COMMENT '域名id',
  `rule_id` varchar(36) DEFAULT NULL COMMENT '告警规则id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `sys_node_domain` */

/*Table structure for table `sys_operation_log` */

DROP TABLE IF EXISTS `sys_operation_log`;

CREATE TABLE `sys_operation_log` (
  `id` tinyint(11) NOT NULL AUTO_INCREMENT,
  `detail` text NOT NULL,
  `create_time` datetime NOT NULL,
  `server_id` tinyint(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_operation_log` */

/*Table structure for table `sys_order` */

DROP TABLE IF EXISTS `sys_order`;

CREATE TABLE `sys_order` (
  `id` tinyint(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '计费方名称',
  `parent_id` tinyint(11) DEFAULT '0' COMMENT '父ID',
  `tree_path` varchar(128) DEFAULT '' COMMENT '树节点路径',
  `sort` tinyint(11) DEFAULT '1' COMMENT '排序',
  `status` tinyint(11) DEFAULT '1' COMMENT '状态',
  `deleted` tinyint(1) DEFAULT '0' COMMENT '逻辑删除(0:未删除,1:已删除)',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_order` */

insert  into `sys_order`(`id`,`name`,`parent_id`,`tree_path`,`sort`,`status`,`deleted`,`create_time`,`update_time`) values 
(1,'全部计费',0,'0',1,1,0,'2024-07-01 15:37:28','2024-07-01 15:37:28');

/*Table structure for table `sys_project` */

DROP TABLE IF EXISTS `sys_project`;

CREATE TABLE `sys_project` (
  `id` tinyint(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '分组名称',
  `parent_id` tinyint(11) DEFAULT '0' COMMENT '父ID',
  `tree_path` varchar(128) DEFAULT '' COMMENT '树节点路径',
  `sort` tinyint(11) DEFAULT '1' COMMENT '排序',
  `status` tinyint(11) DEFAULT '1' COMMENT '状态',
  `deleted` tinyint(1) DEFAULT '0' COMMENT '逻辑删除(0:未删除,1:已删除)',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_project` */

insert  into `sys_project`(`id`,`name`,`parent_id`,`tree_path`,`sort`,`status`,`deleted`,`create_time`,`update_time`) values 
(1,'全部项目',0,'0',1,1,0,'2023-10-26 14:30:17','2023-10-26 14:30:19');

/*Table structure for table `sys_role` */

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '角色名称',
  `code` varchar(32) DEFAULT NULL COMMENT '角色编码',
  `sort` int(11) DEFAULT NULL COMMENT '显示顺序',
  `status` tinyint(1) DEFAULT '1' COMMENT '角色状态(1-正常；0-停用)',
  `data_scope` tinyint(4) DEFAULT NULL COMMENT '数据权限(0-所有数据；1-部门及子部门数据；2-本部门数据；3-本人数据)',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除标识(0-未删除；1-已删除)',
  `create_time` datetime DEFAULT NULL COMMENT '更新时间',
  `update_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色表';

/*Data for the table `sys_role` */

insert  into `sys_role`(`id`,`name`,`code`,`sort`,`status`,`data_scope`,`deleted`,`create_time`,`update_time`) values 
(1,'超级管理员','ROOT',1,1,0,0,'2021-05-21 14:56:51','2018-12-23 16:00:00'),
(2,'系统管理员','ADMIN',2,1,1,0,'2021-03-25 12:39:54',NULL),
(3,'普通用户','GUEST',3,1,2,0,'2021-05-26 15:49:05','2019-05-05 16:00:00');

/*Table structure for table `sys_role_menu` */

DROP TABLE IF EXISTS `sys_role_menu`;

CREATE TABLE `sys_role_menu` (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色和菜单关联表';

/*Data for the table `sys_role_menu` */

insert  into `sys_role_menu`(`role_id`,`menu_id`) values 
(2,1),
(2,2),
(2,3),
(2,4),
(2,5),
(2,6),
(2,11),
(2,12),
(2,19),
(2,18),
(2,17),
(2,13),
(2,14),
(2,15),
(2,16),
(2,9),
(2,10),
(2,37),
(2,20),
(2,21),
(2,22),
(2,23),
(2,24),
(2,32),
(2,33),
(2,39),
(2,34),
(2,26),
(2,30),
(2,31),
(2,36),
(2,38),
(2,39),
(2,40),
(2,41),
(3,43),
(3,42),
(2,7),
(2,44),
(2,45),
(2,8),
(2,9),
(2,10),
(2,11);

/*Table structure for table `sys_room` */

DROP TABLE IF EXISTS `sys_room`;

CREATE TABLE `sys_room` (
  `id` tinyint(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '分组名称',
  `parent_id` tinyint(11) DEFAULT '0' COMMENT '父ID',
  `tree_path` varchar(128) DEFAULT '' COMMENT '树节点路径',
  `sort` tinyint(11) DEFAULT '1' COMMENT '排序',
  `status` tinyint(11) DEFAULT '1' COMMENT '状态',
  `deleted` tinyint(1) DEFAULT '0' COMMENT '逻辑删除(0:未删除,1:已删除)',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_room` */

insert  into `sys_room`(`id`,`name`,`parent_id`,`tree_path`,`sort`,`status`,`deleted`,`create_time`,`update_time`) values 
(1,'全部机房',0,'0',1,1,0,'2023-10-26 14:37:22','2023-10-26 14:37:24');

/*Table structure for table `sys_server` */

DROP TABLE IF EXISTS `sys_server`;

CREATE TABLE `sys_server` (
  `id` tinyint(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `ip_addr` text NOT NULL COMMENT 'ip地址',
  `zone_id` tinyint(11) NOT NULL COMMENT '地域id',
  `project_id` tinyint(11) NOT NULL COMMENT '项目id',
  `room_id` tinyint(11) NOT NULL COMMENT '机房id',
  `order_id` tinyint(15) NOT NULL COMMENT '计费',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `remark` text COMMENT '备注',
  `sort` int(11) NOT NULL DEFAULT '1' COMMENT '显示顺序',
  `all_ip` text COMMENT '拆分为单个IP',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除(0:未删除,1:已删除)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;


/*Table structure for table `sys_user` */

DROP TABLE IF EXISTS `sys_user`;

CREATE TABLE `sys_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL COMMENT '用户名',
  `nickname` varchar(64) DEFAULT NULL COMMENT '昵称',
  `gender` tinyint(1) DEFAULT '1' COMMENT '性别((1:男;2:女))',
  `password` varchar(100) DEFAULT NULL COMMENT '密码',
  `dept_id` int(11) DEFAULT NULL COMMENT '部门ID',
  `avatar` varchar(255) DEFAULT '' COMMENT '用户头像',
  `mobile` varchar(20) DEFAULT NULL COMMENT '联系方式',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '用户状态((1:正常;0:禁用))',
  `email` varchar(128) DEFAULT NULL COMMENT '用户邮箱',
  `deleted` tinyint(1) DEFAULT '0' COMMENT '逻辑删除标识(0:未删除;1:已删除)',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `login_name` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户信息表';

/*Data for the table `sys_user` */

insert  into `sys_user`(`id`,`username`,`nickname`,`gender`,`password`,`dept_id`,`avatar`,`mobile`,`status`,`email`,`deleted`,`create_time`,`update_time`) values 
(1,'root','有来技术',0,'123456',1,'https://s2.loli.net/2022/04/07/gw1L2Z5sPtS8GIl.gif','17621590365',1,'youlaitech@163.com',0,NULL,'2023-04-24 21:49:32'),
(2,'admin','系统管理员',1,'$2a$10$1LT9eQM0APDQdDYzjYiCreBjMFktQx09xBRS7k0bPuHp2EbyxJOEy',2,'https://s2.loli.net/2022/04/07/gw1L2Z5sPtS8GIl.gif','17621210366',1,'',0,'2019-10-10 13:41:22','2022-07-31 12:39:30'),
(3,'test1','测试小用户',1,'$2a$10$xVWsNOhHrCxh5UbpCE7/HuJ.PAOKcYAqRxD2CO2nVnJS.IAXkr5aq',3,'https://s2.loli.net/2022/04/07/gw1L2Z5sPtS8GIl.gif','17621210365',0,'youlaitech@163.com',0,'2021-06-05 01:31:29','2023-03-28 20:31:08');

/*Table structure for table `sys_user_role` */

DROP TABLE IF EXISTS `sys_user_role`;

CREATE TABLE `sys_user_role` (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户和角色关联表';

/*Data for the table `sys_user_role` */

insert  into `sys_user_role`(`user_id`,`role_id`) values 
(1,1),
(2,2),
(3,3),
(19,3),
(22,3);

/*Table structure for table `sys_zone` */

DROP TABLE IF EXISTS `sys_zone`;

CREATE TABLE `sys_zone` (
  `id` tinyint(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '分组名称',
  `parent_id` tinyint(11) DEFAULT '0' COMMENT '父ID',
  `tree_path` varchar(128) DEFAULT '' COMMENT '树节点路径',
  `sort` tinyint(11) DEFAULT '1' COMMENT '排序',
  `status` tinyint(11) DEFAULT '1' COMMENT '状态',
  `deleted` tinyint(1) DEFAULT '0' COMMENT '逻辑删除(0:未删除,1:已删除)',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_zone` */

insert  into `sys_zone`(`id`,`name`,`parent_id`,`tree_path`,`sort`,`status`,`deleted`,`create_time`,`update_time`) values 
(1,'全部地区',0,'0',1,1,0,'2023-10-26 15:42:36','2024-07-01 14:47:32');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
