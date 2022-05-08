CREATE DATABASE IF NOT EXISTS clean_architecture;

USE clean_architecture;

DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS phones;
DROP TABLE IF EXISTS recommendations;

CREATE TABLE students (
    id int auto_increment primary key,
    cpf varchar(11) not null unique,
    name varchar(50) not null,
    email varchar(50) not null
) ENGINE=INNODB;

CREATE TABLE phones (
    id int auto_increment primary key,
    ddd varchar(3) not null,
    number varchar(9) not null,
    student_id int not null,
    FOREIGN KEY(student_id) REFERENCES students(id)
) ENGINE=INNODB;

CREATE TABLE recommendations (
    id int auto_increment primary key,
    student_indicator int not null,
    student_indicated int not null,
    date_recommendation timestamp default current_timestamp(),
    FOREIGN KEY(student_indicated) REFERENCES students(id),
    FOREIGN KEY(student_indicator) REFERENCES students(id)
) ENGINE=INNODB;