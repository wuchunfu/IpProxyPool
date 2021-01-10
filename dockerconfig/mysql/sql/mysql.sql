
CREATE DATABASE IF NOT EXISTS `test`;
USE TEST;

CREATE TABLE IF NOT EXISTS `proxy_ip` (
  `proxy_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `proxy_host` varchar(255) NOT NULL,
  `proxy_port` int(11) NOT NULL,
  `proxy_type` varchar(64) NOT NULL,
  `proxy_location` varchar(255) DEFAULT NULL,
  `proxy_speed` int(20) NOT NULL DEFAULT '0',
  `proxy_source` varchar(64) NOT NULL,
  `create_time` varchar(50) NOT NULL,
  `update_time` varchar(50) NOT NULL,
  PRIMARY KEY (`proxy_id`),
  UNIQUE KEY `UNIQUE_HOST_PORT` (`proxy_host`,`proxy_port`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;