-- MySQL dump 10.13  Distrib 5.7.25, for Linux (x86_64)
--
-- Host: localhost    Database: netdisk
-- ------------------------------------------------------
-- Server version	5.7.25

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `accounts_user`
--

DROP TABLE IF EXISTS `accounts_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `accounts_user` (
  `last_login` datetime(6) DEFAULT NULL,
  `is_superuser` tinyint(1) NOT NULL,
  `first_name` varchar(30) NOT NULL,
  `last_name` varchar(150) NOT NULL,
  `email` varchar(254) NOT NULL,
  `is_staff` tinyint(1) NOT NULL,
  `is_active` tinyint(1) NOT NULL,
  `date_joined` datetime(6) NOT NULL,
  `uuid` char(32) NOT NULL,
  `username` varchar(45) NOT NULL,
  `password` varchar(255) DEFAULT NULL,
  `create_time` datetime(6) NOT NULL,
  `update_time` datetime(6) NOT NULL,
  `avatar_url` varchar(255) DEFAULT NULL,
  `role` varchar(45) NOT NULL,
  `status` varchar(45) NOT NULL,
  `size_limit` bigint(20) NOT NULL,
  `total_size_limit` bigint(20) NOT NULL,
  `total_size` bigint(20) NOT NULL,
  `last_time` datetime(6) NOT NULL,
  `last_ip` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`uuid`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts_user`
--

LOCK TABLES `accounts_user` WRITE;
/*!40000 ALTER TABLE `accounts_user` DISABLE KEYS */;
INSERT INTO `accounts_user` VALUES ('2021-06-18 14:35:02.445170',1,'','','',1,1,'2021-05-22 18:39:47.533408','a295f63bbd1e4eb9866f687b0b39332f','admin','argon2$argon2i$v=19$m=512,t=2,p=2$MW5RWUhJR09ZUzRk$RWuIHEbA50A3+58AF2JvkA','2021-05-22 18:39:47.554607','2021-05-22 18:39:47.554635',NULL,'ADMINISTRATOR','OK',268435456,-1,0,'2021-05-22 18:39:47.554702',NULL),('2021-06-16 03:19:44.744165',0,'','','',0,1,'2021-05-22 18:39:47.604910','b494dea08b9049158272836fc98b6d4e','user0','argon2$argon2i$v=19$m=512,t=2,p=2$SG9jM2lwdm5GQkMy$7NbFCWg08fPq8udTTv7u6g','2021-05-22 18:39:47.608648','2021-05-22 19:18:26.261109','','USER','OK',10240000,-1,0,'2021-05-22 19:18:26.261127',NULL);
/*!40000 ALTER TABLE `accounts_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `accounts_user_groups`
--

DROP TABLE IF EXISTS `accounts_user_groups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `accounts_user_groups` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` char(32) NOT NULL,
  `group_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `accounts_user_groups_user_id_group_id_59c0b32f_uniq` (`user_id`,`group_id`),
  KEY `accounts_user_groups_group_id_bd11a704_fk_auth_group_id` (`group_id`),
  CONSTRAINT `accounts_user_groups_group_id_bd11a704_fk_auth_group_id` FOREIGN KEY (`group_id`) REFERENCES `auth_group` (`id`),
  CONSTRAINT `accounts_user_groups_user_id_52b62117_fk_accounts_user_uuid` FOREIGN KEY (`user_id`) REFERENCES `accounts_user` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts_user_groups`
--

LOCK TABLES `accounts_user_groups` WRITE;
/*!40000 ALTER TABLE `accounts_user_groups` DISABLE KEYS */;
/*!40000 ALTER TABLE `accounts_user_groups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `accounts_user_user_permissions`
--

DROP TABLE IF EXISTS `accounts_user_user_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `accounts_user_user_permissions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` char(32) NOT NULL,
  `permission_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `accounts_user_user_permi_user_id_permission_id_2ab516c2_uniq` (`user_id`,`permission_id`),
  KEY `accounts_user_user_p_permission_id_113bb443_fk_auth_perm` (`permission_id`),
  CONSTRAINT `accounts_user_user_p_permission_id_113bb443_fk_auth_perm` FOREIGN KEY (`permission_id`) REFERENCES `auth_permission` (`id`),
  CONSTRAINT `accounts_user_user_p_user_id_e4f0a161_fk_accounts_` FOREIGN KEY (`user_id`) REFERENCES `accounts_user` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts_user_user_permissions`
--

LOCK TABLES `accounts_user_user_permissions` WRITE;
/*!40000 ALTER TABLE `accounts_user_user_permissions` DISABLE KEYS */;
/*!40000 ALTER TABLE `accounts_user_user_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `auth_group`
--

DROP TABLE IF EXISTS `auth_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `auth_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(150) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `auth_group`
--

LOCK TABLES `auth_group` WRITE;
/*!40000 ALTER TABLE `auth_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `auth_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `auth_group_permissions`
--

DROP TABLE IF EXISTS `auth_group_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `auth_group_permissions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_group_permissions_group_id_permission_id_0cd325b0_uniq` (`group_id`,`permission_id`),
  KEY `auth_group_permissio_permission_id_84c5c92e_fk_auth_perm` (`permission_id`),
  CONSTRAINT `auth_group_permissio_permission_id_84c5c92e_fk_auth_perm` FOREIGN KEY (`permission_id`) REFERENCES `auth_permission` (`id`),
  CONSTRAINT `auth_group_permissions_group_id_b120cbf9_fk_auth_group_id` FOREIGN KEY (`group_id`) REFERENCES `auth_group` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `auth_group_permissions`
--

LOCK TABLES `auth_group_permissions` WRITE;
/*!40000 ALTER TABLE `auth_group_permissions` DISABLE KEYS */;
/*!40000 ALTER TABLE `auth_group_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `auth_permission`
--

DROP TABLE IF EXISTS `auth_permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `auth_permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `content_type_id` int(11) NOT NULL,
  `codename` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_permission_content_type_id_codename_01ab375a_uniq` (`content_type_id`,`codename`),
  CONSTRAINT `auth_permission_content_type_id_2f476e4b_fk_django_co` FOREIGN KEY (`content_type_id`) REFERENCES `django_content_type` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `auth_permission`
--

LOCK TABLES `auth_permission` WRITE;
/*!40000 ALTER TABLE `auth_permission` DISABLE KEYS */;
INSERT INTO `auth_permission` VALUES (1,'Can add permission',1,'add_permission'),(2,'Can change permission',1,'change_permission'),(3,'Can delete permission',1,'delete_permission'),(4,'Can view permission',1,'view_permission'),(5,'Can add group',2,'add_group'),(6,'Can change group',2,'change_group'),(7,'Can delete group',2,'delete_group'),(8,'Can view group',2,'view_group'),(9,'Can add content type',3,'add_contenttype'),(10,'Can change content type',3,'change_contenttype'),(11,'Can delete content type',3,'delete_contenttype'),(12,'Can view content type',3,'view_contenttype'),(13,'Can add session',4,'add_session'),(14,'Can change session',4,'change_session'),(15,'Can delete session',4,'delete_session'),(16,'Can view session',4,'view_session'),(17,'Can add site',5,'add_site'),(18,'Can change site',5,'change_site'),(19,'Can delete site',5,'delete_site'),(20,'Can view site',5,'view_site'),(21,'Can add log entry',6,'add_logentry'),(22,'Can change log entry',6,'change_logentry'),(23,'Can delete log entry',6,'delete_logentry'),(24,'Can view log entry',6,'view_logentry'),(25,'Can add Token',7,'add_token'),(26,'Can change Token',7,'change_token'),(27,'Can delete Token',7,'delete_token'),(28,'Can view Token',7,'view_token'),(29,'Can add user',8,'add_user'),(30,'Can change user',8,'change_user'),(31,'Can delete user',8,'delete_user'),(32,'Can view user',8,'view_user'),(33,'Can add 站点偏好',9,'add_preference'),(34,'Can change 站点偏好',9,'change_preference'),(35,'Can delete 站点偏好',9,'delete_preference'),(36,'Can view 站点偏好',9,'view_preference'),(37,'Can add 文件对象',10,'add_matter'),(38,'Can change 文件对象',10,'change_matter'),(39,'Can delete 文件对象',10,'delete_matter'),(40,'Can view 文件对象',10,'view_matter'),(41,'Can add 空间',11,'add_project'),(42,'Can change 空间',11,'change_project'),(43,'Can delete 空间',11,'delete_project'),(44,'Can view 空间',11,'view_project'),(45,'Can add 空间授权',12,'add_permission'),(46,'Can change 空间授权',12,'change_permission'),(47,'Can delete 空间授权',12,'delete_permission'),(48,'Can view 空间授权',12,'view_permission');
/*!40000 ALTER TABLE `auth_permission` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `authtoken_token`
--

DROP TABLE IF EXISTS `authtoken_token`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `authtoken_token` (
  `key` varchar(40) NOT NULL,
  `created` datetime(6) NOT NULL,
  `user_id` char(32) NOT NULL,
  PRIMARY KEY (`key`),
  UNIQUE KEY `user_id` (`user_id`),
  CONSTRAINT `authtoken_token_user_id_35299eff_fk_accounts_user_uuid` FOREIGN KEY (`user_id`) REFERENCES `accounts_user` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authtoken_token`
--

LOCK TABLES `authtoken_token` WRITE;
/*!40000 ALTER TABLE `authtoken_token` DISABLE KEYS */;
/*!40000 ALTER TABLE `authtoken_token` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `django_admin_log`
--

DROP TABLE IF EXISTS `django_admin_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `django_admin_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `action_time` datetime(6) NOT NULL,
  `object_id` longtext,
  `object_repr` varchar(200) NOT NULL,
  `action_flag` smallint(5) unsigned NOT NULL,
  `change_message` longtext NOT NULL,
  `content_type_id` int(11) DEFAULT NULL,
  `user_id` char(32) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `django_admin_log_content_type_id_c4bce8eb_fk_django_co` (`content_type_id`),
  KEY `django_admin_log_user_id_c564eba6_fk_accounts_user_uuid` (`user_id`),
  CONSTRAINT `django_admin_log_content_type_id_c4bce8eb_fk_django_co` FOREIGN KEY (`content_type_id`) REFERENCES `django_content_type` (`id`),
  CONSTRAINT `django_admin_log_user_id_c564eba6_fk_accounts_user_uuid` FOREIGN KEY (`user_id`) REFERENCES `accounts_user` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `django_admin_log`
--

LOCK TABLES `django_admin_log` WRITE;
/*!40000 ALTER TABLE `django_admin_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `django_admin_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `django_content_type`
--

DROP TABLE IF EXISTS `django_content_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `django_content_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_label` varchar(100) NOT NULL,
  `model` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `django_content_type_app_label_model_76bd3d3b_uniq` (`app_label`,`model`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `django_content_type`
--

LOCK TABLES `django_content_type` WRITE;
/*!40000 ALTER TABLE `django_content_type` DISABLE KEYS */;
INSERT INTO `django_content_type` VALUES (8,'accounts','user'),(6,'admin','logentry'),(2,'auth','group'),(1,'auth','permission'),(7,'authtoken','token'),(3,'contenttypes','contenttype'),(10,'matter','matter'),(12,'permission','permission'),(9,'preference','preference'),(11,'project','project'),(4,'sessions','session'),(5,'sites','site');
/*!40000 ALTER TABLE `django_content_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `django_migrations`
--

DROP TABLE IF EXISTS `django_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `django_migrations` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `applied` datetime(6) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `django_migrations`
--

LOCK TABLES `django_migrations` WRITE;
/*!40000 ALTER TABLE `django_migrations` DISABLE KEYS */;
INSERT INTO `django_migrations` VALUES (1,'contenttypes','0001_initial','2021-05-22 18:39:45.107354'),(2,'contenttypes','0002_remove_content_type_name','2021-05-22 18:39:45.248871'),(3,'auth','0001_initial','2021-05-22 18:39:45.412943'),(4,'auth','0002_alter_permission_name_max_length','2021-05-22 18:39:45.822246'),(5,'auth','0003_alter_user_email_max_length','2021-05-22 18:39:45.832118'),(6,'auth','0004_alter_user_username_opts','2021-05-22 18:39:45.841709'),(7,'auth','0005_alter_user_last_login_null','2021-05-22 18:39:45.852906'),(8,'auth','0006_require_contenttypes_0002','2021-05-22 18:39:45.857074'),(9,'auth','0007_alter_validators_add_error_messages','2021-05-22 18:39:45.868148'),(10,'auth','0008_alter_user_username_max_length','2021-05-22 18:39:45.878226'),(11,'auth','0009_alter_user_last_name_max_length','2021-05-22 18:39:45.888229'),(12,'auth','0010_alter_group_name_max_length','2021-05-22 18:39:45.963321'),(13,'auth','0011_update_proxy_permissions','2021-05-22 18:39:45.973159'),(14,'accounts','0001_initial','2021-05-22 18:39:46.105357'),(15,'accounts','0002_auto_20210322_0534','2021-05-22 18:39:46.455752'),(16,'accounts','0003_auto_20210404_2233','2021-05-22 18:39:46.474599'),(17,'admin','0001_initial','2021-05-22 18:39:46.519947'),(18,'admin','0002_logentry_remove_auto_add','2021-05-22 18:39:46.678973'),(19,'admin','0003_logentry_add_action_flag_choices','2021-05-22 18:39:46.691578'),(20,'authtoken','0001_initial','2021-05-22 18:39:46.736634'),(21,'authtoken','0002_auto_20160226_1747','2021-05-22 18:39:46.919910'),(22,'matter','0001_initial','2021-05-22 18:39:46.962606'),(23,'matter','0002_matter_file','2021-05-22 18:39:47.022920'),(24,'permission','0001_initial','2021-05-22 18:39:47.066714'),(25,'permission','0002_auto_20210404_0258','2021-05-22 18:39:47.129229'),(26,'permission','0003_auto_20210405_0324','2021-05-22 18:39:47.138115'),(27,'preference','0001_initial','2021-05-22 18:39:47.182009'),(28,'preference','0002_auto_20210322_1631','2021-05-22 18:39:47.191679'),(29,'preference','0003_auto_20210322_1633','2021-05-22 18:39:47.204699'),(30,'project','0001_initial','2021-05-22 18:39:47.243703'),(31,'project','0002_auto_20210327_0144','2021-05-22 18:39:47.252563'),(32,'sessions','0001_initial','2021-05-22 18:39:47.291023'),(33,'sites','0001_initial','2021-05-22 18:39:47.355289'),(34,'sites','0002_alter_domain_unique','2021-05-22 18:39:47.396124'),(35,'sites','0003_set_site_domain_and_name','2021-05-22 18:39:47.428879'),(36,'sites','0004_init_superuser','2021-05-22 18:39:47.574841'),(37,'sites','0005_init_pemission','2021-05-22 18:39:47.617305');
/*!40000 ALTER TABLE `django_migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `django_session`
--

DROP TABLE IF EXISTS `django_session`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `django_session` (
  `session_key` varchar(40) NOT NULL,
  `session_data` longtext NOT NULL,
  `expire_date` datetime(6) NOT NULL,
  PRIMARY KEY (`session_key`),
  KEY `django_session_expire_date_a5c62663` (`expire_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `django_session`
--

LOCK TABLES `django_session` WRITE;
/*!40000 ALTER TABLE `django_session` DISABLE KEYS */;
INSERT INTO `django_session` VALUES ('3yuszn9lb9215kd5jcj36bbc9nzxjd08','MmRlMWVmZTcyNWIxMTVlMGJjYWFmNGExYzY5Njg3MDAyYTZlNTNjZDp7Il9hdXRoX3VzZXJfaWQiOiJiNDk0ZGVhMC04YjkwLTQ5MTUtODI3Mi04MzZmYzk4YjZkNGUiLCJfYXV0aF91c2VyX2JhY2tlbmQiOiJkamFuZ28uY29udHJpYi5hdXRoLmJhY2tlbmRzLk1vZGVsQmFja2VuZCIsIl9hdXRoX3VzZXJfaGFzaCI6IjU1NzU1Y2QwNjdlMTYxOGZjMmQzOTE4MWY1Y2RkMWY3MzU5NjMxYzcifQ==','2021-06-30 03:19:44.872254'),('8mxvigndgjmxuxt58rowt44nkm7udzef','NmU0MjQ4OTk1NzhhYjBiZGQzYTcyNjI5NmE1Y2MzYzg5MWEwODZkNDp7Il9hdXRoX3VzZXJfaWQiOiJhMjk1ZjYzYi1iZDFlLTRlYjktODY2Zi02ODdiMGIzOTMzMmYiLCJfYXV0aF91c2VyX2JhY2tlbmQiOiJkamFuZ28uY29udHJpYi5hdXRoLmJhY2tlbmRzLk1vZGVsQmFja2VuZCIsIl9hdXRoX3VzZXJfaGFzaCI6IjAwYmY4YTg1MmMzYWFkZmU3N2EyMTBmY2EyZjk4MDFjMTA0ODEyOGQifQ==','2021-06-15 12:12:36.666749'),('ek7fn9vzo2oiwxkjbo6kumspphlfamwk','MmRlMWVmZTcyNWIxMTVlMGJjYWFmNGExYzY5Njg3MDAyYTZlNTNjZDp7Il9hdXRoX3VzZXJfaWQiOiJiNDk0ZGVhMC04YjkwLTQ5MTUtODI3Mi04MzZmYzk4YjZkNGUiLCJfYXV0aF91c2VyX2JhY2tlbmQiOiJkamFuZ28uY29udHJpYi5hdXRoLmJhY2tlbmRzLk1vZGVsQmFja2VuZCIsIl9hdXRoX3VzZXJfaGFzaCI6IjU1NzU1Y2QwNjdlMTYxOGZjMmQzOTE4MWY1Y2RkMWY3MzU5NjMxYzcifQ==','2021-06-05 19:18:35.638274'),('g8145p57vkdkeoz8sazbq5mywbpfambt','MmRlMWVmZTcyNWIxMTVlMGJjYWFmNGExYzY5Njg3MDAyYTZlNTNjZDp7Il9hdXRoX3VzZXJfaWQiOiJiNDk0ZGVhMC04YjkwLTQ5MTUtODI3Mi04MzZmYzk4YjZkNGUiLCJfYXV0aF91c2VyX2JhY2tlbmQiOiJkamFuZ28uY29udHJpYi5hdXRoLmJhY2tlbmRzLk1vZGVsQmFja2VuZCIsIl9hdXRoX3VzZXJfaGFzaCI6IjU1NzU1Y2QwNjdlMTYxOGZjMmQzOTE4MWY1Y2RkMWY3MzU5NjMxYzcifQ==','2021-06-24 03:22:34.718772'),('nx9d5wltllmpudohtr5lntu08l33iwan','MmRlMWVmZTcyNWIxMTVlMGJjYWFmNGExYzY5Njg3MDAyYTZlNTNjZDp7Il9hdXRoX3VzZXJfaWQiOiJiNDk0ZGVhMC04YjkwLTQ5MTUtODI3Mi04MzZmYzk4YjZkNGUiLCJfYXV0aF91c2VyX2JhY2tlbmQiOiJkamFuZ28uY29udHJpYi5hdXRoLmJhY2tlbmRzLk1vZGVsQmFja2VuZCIsIl9hdXRoX3VzZXJfaGFzaCI6IjU1NzU1Y2QwNjdlMTYxOGZjMmQzOTE4MWY1Y2RkMWY3MzU5NjMxYzcifQ==','2021-06-05 19:21:52.656948'),('x3n8iuoxnous5u0cswfhagvzmdty2vnc','NmU0MjQ4OTk1NzhhYjBiZGQzYTcyNjI5NmE1Y2MzYzg5MWEwODZkNDp7Il9hdXRoX3VzZXJfaWQiOiJhMjk1ZjYzYi1iZDFlLTRlYjktODY2Zi02ODdiMGIzOTMzMmYiLCJfYXV0aF91c2VyX2JhY2tlbmQiOiJkamFuZ28uY29udHJpYi5hdXRoLmJhY2tlbmRzLk1vZGVsQmFja2VuZCIsIl9hdXRoX3VzZXJfaGFzaCI6IjAwYmY4YTg1MmMzYWFkZmU3N2EyMTBmY2EyZjk4MDFjMTA0ODEyOGQifQ==','2021-07-02 14:35:02.489261');
/*!40000 ALTER TABLE `django_session` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `django_site`
--

DROP TABLE IF EXISTS `django_site`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `django_site` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain` varchar(100) NOT NULL,
  `name` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `django_site_domain_a2e37b91_uniq` (`domain`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `django_site`
--

LOCK TABLES `django_site` WRITE;
/*!40000 ALTER TABLE `django_site` DISABLE KEYS */;
INSERT INTO `django_site` VALUES (1,'miya.com','icloud');
/*!40000 ALTER TABLE `django_site` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `matter`
--

DROP TABLE IF EXISTS `matter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `matter` (
  `uuid` varchar(36) NOT NULL,
  `puuid` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `username` varchar(45) NOT NULL,
  `user_uuid` varchar(36) NOT NULL,
  `md5` varchar(45) DEFAULT NULL,
  `size` bigint(20) NOT NULL,
  `dir` tinyint(1) NOT NULL,
  `privacy` tinyint(1) NOT NULL,
  `path` varchar(1024) DEFAULT NULL,
  `update_time` datetime(6) NOT NULL,
  `create_time` datetime(6) NOT NULL,
  `times` int(11) NOT NULL,
  `file` varchar(100) NOT NULL,
  PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `matter`
--

LOCK TABLES `matter` WRITE;
/*!40000 ALTER TABLE `matter` DISABLE KEYS */;
INSERT INTO `matter` VALUES ('006e0fba-c988-4c21-a640-bfb375fff19d','3e941119-2931-4fd1-9777-17f28927f1f2','apple-touch-icon.png','admin','a295f63b-bd1e-4eb9-866f-687b0b39332f','',8589,0,1,'/media/admin/root/images/favicon_io%20(1)/apple-touch-icon.png','2021-05-22 19:17:45.787687','2021-05-22 19:17:45.763795',0,'admin/root/images/favicon_io (1)/apple-touch-icon.png'),('3e941119-2931-4fd1-9777-17f28927f1f2','5cfa8798-fe3e-4ffa-a0ba-b9afd88003f5','favicon_io (1)','admin','a295f63b-bd1e-4eb9-866f-687b0b39332f','',0,1,0,'/images/favicon_io (1)','2021-05-22 19:17:45.631477','2021-05-22 19:17:45.631510',0,''),('53baace6-80de-43db-92cf-67b211e433c1','root','video','admin','a295f63b-bd1e-4eb9-866f-687b0b39332f','',0,1,0,'/video','2021-05-22 19:17:08.086907','2021-05-22 19:17:08.086950',0,''),('5cfa8798-fe3e-4ffa-a0ba-b9afd88003f5','root','images','admin','a295f63b-bd1e-4eb9-866f-687b0b39332f','',0,1,0,'/images','2021-05-22 19:17:17.168983','2021-05-22 19:17:17.169016',0,''),('b3b0a693-5dd6-4ad5-be9b-885006e91e2b','root','游戏','user0','b494dea0-8b90-4915-8272-836fc98b6d4e','',0,1,0,'/游戏','2021-05-22 19:16:41.097230','2021-05-22 19:16:41.097265',0,''),('b6051e39-fc46-49fb-a703-7def803bb47d','root','software','admin','a295f63b-bd1e-4eb9-866f-687b0b39332f','',0,1,0,'/software','2021-05-22 19:17:12.104503','2021-05-22 19:17:12.104535',0,''),('b7543112-a18e-4e90-bdf5-f543a851ebb9','root','图片','user0','b494dea0-8b90-4915-8272-836fc98b6d4e','',0,1,0,'/图片','2021-05-22 19:16:33.208234','2021-05-22 19:16:33.208266',0,''),('bd4c49dd-06c9-48f4-acba-d9fd87351b2c','3e941119-2931-4fd1-9777-17f28927f1f2','favicon-32x32.png','admin','a295f63b-bd1e-4eb9-866f-687b0b39332f','',1240,0,1,'/media/admin/root/images/favicon_io%20(1)/favicon-32x32.png','2021-05-22 19:17:45.816111','2021-05-22 19:17:45.800171',0,'admin/root/images/favicon_io (1)/favicon-32x32.png'),('bf8acaae-68b6-442d-beb9-7fc37b428748','root','软件','user0','b494dea0-8b90-4915-8272-836fc98b6d4e','',0,1,0,'/软件','2021-05-22 19:16:38.584716','2021-05-22 19:16:38.584750',0,''),('c827c372-05da-45ca-b883-7aac95acfac2','3e941119-2931-4fd1-9777-17f28927f1f2','android-chrome-512x512.png','admin','a295f63b-bd1e-4eb9-866f-687b0b39332f','',35666,0,1,'/media/admin/root/images/favicon_io%20(1)/android-chrome-512x512.png','2021-05-22 19:17:45.813629','2021-05-22 19:17:45.806141',0,'admin/root/images/favicon_io (1)/android-chrome-512x512.png'),('d2381a79-583c-4498-8c16-bfdb4280abdc','3e941119-2931-4fd1-9777-17f28927f1f2','favicon.ico','admin','a295f63b-bd1e-4eb9-866f-687b0b39332f','',15406,0,1,'/media/admin/root/images/favicon_io%20(1)/favicon.ico','2021-05-22 19:17:45.749436','2021-05-22 19:17:45.738915',0,'admin/root/images/favicon_io (1)/favicon.ico'),('d7fc3d21-6842-4691-9612-4650950315a5','3e941119-2931-4fd1-9777-17f28927f1f2','android-chrome-192x192.png','admin','a295f63b-bd1e-4eb9-866f-687b0b39332f','',9438,0,1,'/media/admin/root/images/favicon_io%20(1)/android-chrome-192x192.png','2021-05-22 19:17:45.746448','2021-05-22 19:17:45.731609',0,'admin/root/images/favicon_io (1)/android-chrome-192x192.png'),('da116798-b955-4c8e-817c-51ffc229dda5','root','深入REACT技术栈+陈屹著_犀牛前端部落www.pipipi.net.pdf','user0','b494dea0-8b90-4915-8272-836fc98b6d4e','',10035163,0,1,'/media/user0/root/%E6%B7%B1%E5%85%A5REACT%E6%8A%80%E6%9C%AF%E6%A0%88%E9%99%88%E5%B1%B9%E8%91%97_%E7%8A%80%E7%89%9B%E5%89%8D%E7%AB%AF%E9%83%A8%E8%90%BDwww.pipipi.net.pdf','2021-05-22 19:19:08.898735','2021-05-22 19:19:08.894908',0,'user0/root/深入REACT技术栈陈屹著_犀牛前端部落www.pipipi.net.pdf'),('db0fb3b8-2583-4c58-b984-923075f0253e','root','视频','user0','b494dea0-8b90-4915-8272-836fc98b6d4e','',0,1,0,'/视频','2021-05-22 19:16:29.574927','2021-05-22 19:16:29.574960',0,''),('e51b7a25-e413-4037-b15e-4eb5825044eb','5cfa8798-fe3e-4ffa-a0ba-b9afd88003f5','favico','admin','a295f63b-bd1e-4eb9-866f-687b0b39332f','',0,1,0,'/images/favico','2021-05-22 19:17:37.529179','2021-05-22 19:17:37.529229',0,''),('e5886bbf-e88a-472b-b5c5-efece64f2ba4','3e941119-2931-4fd1-9777-17f28927f1f2','favicon-16x16.png','admin','a295f63b-bd1e-4eb9-866f-687b0b39332f','',555,0,1,'/media/admin/root/images/favicon_io%20(1)/favicon-16x16.png','2021-05-22 19:17:45.734001','2021-05-22 19:17:45.687848',0,'admin/root/images/favicon_io (1)/favicon-16x16.png'),('e8d5f061-384a-4d2d-8c39-e60dc84141cf','3e941119-2931-4fd1-9777-17f28927f1f2','site.webmanifest','admin','a295f63b-bd1e-4eb9-866f-687b0b39332f','',263,0,1,'/media/admin/root/images/favicon_io%20(1)/site.webmanifest','2021-05-22 19:17:45.773757','2021-05-22 19:17:45.771306',0,'admin/root/images/favicon_io (1)/site.webmanifest'),('fdafe2ea-1212-4b67-8e67-93e714a09c52','root','music','admin','a295f63b-bd1e-4eb9-866f-687b0b39332f','',0,1,0,'/music','2021-05-22 19:17:04.123300','2021-05-22 19:17:04.123340',0,'');
/*!40000 ALTER TABLE `matter` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `permission`
--

DROP TABLE IF EXISTS `permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `permission` (
  `created_by` varchar(255) NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  `deleted_by` varchar(255) DEFAULT NULL,
  `deleted_at` datetime(6) DEFAULT NULL,
  `uuid` varchar(36) NOT NULL,
  `username` varchar(45) NOT NULL,
  `project_uuid` char(32) DEFAULT NULL,
  `role` varchar(45) NOT NULL,
  PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permission`
--

LOCK TABLES `permission` WRITE;
/*!40000 ALTER TABLE `permission` DISABLE KEYS */;
INSERT INTO `permission` VALUES ('','2021-05-22 18:39:47.609553','2021-05-22 18:39:47.609597',NULL,0,NULL,NULL,'6c3db44a-64db-4490-8306-82afed54c67e','user0','96265589f6424377831307cd2ec2a0ef','USER'),('','2021-05-22 18:39:47.601381','2021-05-22 18:39:47.601426',NULL,0,NULL,NULL,'8521b4e1-22f5-4a2d-9612-656cb1634a3f','admin',NULL,'ADMINISTRATOR');
/*!40000 ALTER TABLE `permission` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `preference`
--

DROP TABLE IF EXISTS `preference`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `preference` (
  `uuid` varchar(36) NOT NULL,
  `name` varchar(45) NOT NULL,
  `update_time` datetime(6) NOT NULL,
  `create_time` datetime(6) NOT NULL,
  `logo_url` varchar(255) DEFAULT NULL,
  `favicon_url` varchar(255) DEFAULT NULL,
  `copyright` varchar(1024) DEFAULT NULL,
  `record` varchar(1024) DEFAULT NULL,
  `download_dir_max_size` bigint(20) NOT NULL,
  `download_dir_max_num` bigint(20) NOT NULL,
  `default_total_size_limit` bigint(20) NOT NULL,
  `allow_register` tinyint(1) NOT NULL,
  PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `preference`
--

LOCK TABLES `preference` WRITE;
/*!40000 ALTER TABLE `preference` DISABLE KEYS */;
INSERT INTO `preference` VALUES ('15b5ecdc-e29a-48e8-9471-7f6b05308182','netdisk','2021-05-22 18:39:47.602821','2021-05-22 18:39:47.602852','','','','',-1,-1,-1,0);
/*!40000 ALTER TABLE `preference` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `project`
--

DROP TABLE IF EXISTS `project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `project` (
  `created_by` varchar(255) NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  `deleted_by` varchar(255) DEFAULT NULL,
  `deleted_at` datetime(6) DEFAULT NULL,
  `uuid` varchar(36) NOT NULL,
  `name` varchar(64) NOT NULL,
  `description` longtext NOT NULL,
  PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `project`
--

LOCK TABLES `project` WRITE;
/*!40000 ALTER TABLE `project` DISABLE KEYS */;
INSERT INTO `project` VALUES ('','2021-05-22 18:39:47.604476','2021-05-22 18:39:47.604506',NULL,0,NULL,NULL,'96265589-f642-4377-8313-07cd2ec2a0ef','Demo','demo空间');
/*!40000 ALTER TABLE `project` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-06-29 21:04:35
