# ************************************************************
# Sequel Pro SQL dump
# Version 5446
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.19)
# Database: blog_service
# Generation Time: 2021-10-13 15:24:24 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table blog_tag
# ------------------------------------------------------------

DROP TABLE IF EXISTS `blog_tag`;

CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_by` varchar(255) DEFAULT NULL,
  `modified_by` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `modified_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `is_del` tinyint(3) unsigned DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `state` tinyint(3) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `blog_tag` WRITE;
/*!40000 ALTER TABLE `blog_tag` DISABLE KEYS */;

INSERT INTO `blog_tag` (`id`, `created_by`, `modified_by`, `created_at`, `modified_at`, `deleted_at`, `is_del`, `name`, `state`)
VALUES
	(1,'jiang1','','2021-10-13 23:11:30','2021-10-13 23:11:30',NULL,0,'文学类',1),
	(2,'jiang1','','2021-10-13 23:18:03','2021-10-13 23:18:03',NULL,0,'历史类',1),
	(3,'jiang1','','2021-10-13 23:18:11','2021-10-13 23:18:11',NULL,0,'物理类',1),
	(4,'jiang1','','2021-10-13 23:18:22','2021-10-13 23:18:22',NULL,0,'化学类',1),
	(5,'jiang2','','2021-10-13 23:18:30','2021-10-13 23:18:30',NULL,0,'生物类',1),
	(6,'jiang3','jiang2','2021-10-13 23:19:42','2021-10-13 23:20:09','2021-10-13 23:22:59',1,'历史22',1);

/*!40000 ALTER TABLE `blog_tag` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
