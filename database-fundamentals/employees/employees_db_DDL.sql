DROP DATABASE IF EXISTS employees_db;
CREATE DATABASE employees_db;
USE employees_db;

DROP TABLE IF EXISTS `departments`;

CREATE TABLE `departments`(
`department_num` varchar(100) NOT NULL,
`name`varchar(100),
`location` varchar(100),
PRIMARY KEY (`department_num`)
);

INSERT INTO `departments` 
VALUES ('D-000-1', 'Software', 'Los Tigres'), ('D-000-2', 'Sistemas', 'Guadalupe'), ('D-000-4', 'Ventas', 'Plata');
DROP TABLE IF EXISTS `employees`;

CREATE TABLE `employees` (
  `emp_cod` varchar(100) NOT NULL UNIQUE,
  `first_name`varchar(100),
  `last_name` varchar(100),
  `position` varchar(100),
  `hire_date` date,
  `salary` int(10),
  `commission`int(10),
  `department_num`varchar(100),
  PRIMARY KEY (`emp_cod`),
  CONSTRAINT `employees_department_num_foreign`FOREIGN KEY (`department_num`) REFERENCES departments(department_num)
);

INSERT INTO `employees` 
VALUES ('E-0001','César', 'Piñero', 'Vendedor','2018-05-12', 80000, 15000, 'D-000-4');
