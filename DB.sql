/*
SQLyog Ultimate v13.1.1 (64 bit)
MySQL - 10.4.22-MariaDB : Database - indodev
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`indodev` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `indodev`;

/*Table structure for table `departments` */

DROP TABLE IF EXISTS `departments`;

CREATE TABLE `departments` (
  `DepartmentCode` varchar(10) NOT NULL,
  `DepartmentName` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`DepartmentCode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `departments` */

insert  into `departments`(`DepartmentCode`,`DepartmentName`) values 
('BOD00001','Board of Direcast');

/*Table structure for table `divisions` */

DROP TABLE IF EXISTS `divisions`;

CREATE TABLE `divisions` (
  `DivisionCode` varchar(10) NOT NULL,
  `DivisionName` varchar(255) DEFAULT NULL,
  `DepartmentCode` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`DivisionCode`),
  KEY `DepartmentCode` (`DepartmentCode`),
  CONSTRAINT `divisions_ibfk_1` FOREIGN KEY (`DepartmentCode`) REFERENCES `departments` (`DepartmentCode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `divisions` */

insert  into `divisions`(`DivisionCode`,`DivisionName`,`DepartmentCode`) values 
('DVS00001','Information Technology','BOD00001');

/*Table structure for table `employees` */

DROP TABLE IF EXISTS `employees`;

CREATE TABLE `employees` (
  `EmployeeID` varchar(10) NOT NULL,
  `EmployeeName` varchar(255) DEFAULT NULL,
  `RoleCode` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`EmployeeID`),
  KEY `RoleCode` (`RoleCode`),
  CONSTRAINT `employees_ibfk_1` FOREIGN KEY (`RoleCode`) REFERENCES `roles` (`RoleCode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `employees` */

insert  into `employees`(`EmployeeID`,`EmployeeName`,`RoleCode`) values 
('EMP001','Endah','ITSDERP');

/*Table structure for table `roles` */

DROP TABLE IF EXISTS `roles`;

CREATE TABLE `roles` (
  `RoleCode` varchar(10) NOT NULL,
  `RoleName` varchar(255) DEFAULT NULL,
  `DivisionCode` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`RoleCode`),
  KEY `DivisionCode` (`DivisionCode`),
  CONSTRAINT `roles_ibfk_1` FOREIGN KEY (`DivisionCode`) REFERENCES `divisions` (`DivisionCode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `roles` */

insert  into `roles`(`RoleCode`,`RoleName`,`DivisionCode`) values 
('ITSDERP','ERP Development','DVS00001');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
