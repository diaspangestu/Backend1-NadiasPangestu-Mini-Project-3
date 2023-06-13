-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Jun 09, 2023 at 02:07 AM
-- Server version: 5.7.39
-- PHP Version: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bootcampbri_minpro1`
--

-- --------------------------------------------------------

--
-- Table structure for table `actors`
--

CREATE TABLE `actors` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role_id` int(10) UNSIGNED NOT NULL,
  `is_verified` enum('true','false') DEFAULT NULL,
  `is_actived` enum('true','false') DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `actors`
--

INSERT INTO `actors` (`id`, `username`, `password`, `role_id`, `is_verified`, `is_actived`, `created_at`, `updated_at`) VALUES
(1, 'superadmin', '$2a$12$p0FofjcjfCtfoFIfbdYz6.TMaoPDOn7QirI4AYawoCOJL8mozJku6', 1, 'true', 'true', '2023-06-08 08:46:29', '2023-06-08 08:46:29');

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` bigint(20) NOT NULL,
  `first_name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `avatar` varchar(100) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `first_name`, `last_name`, `email`, `avatar`, `created_at`, `updated_at`) VALUES
(1, 'john1', 'doe1', 'johndoe1@gmail.com', 'john1.jpg', '2023-06-05 23:22:31', '2023-06-05 23:22:31'),
(3, 'john3', 'doe3', 'johndoe3@gmail.com', 'john3.jpg', '2023-06-05 23:23:15', '2023-06-05 23:23:15'),
(31, 'Michael', 'Lawson', 'michael.lawson@reqres.in', 'https://reqres.in/img/faces/7-image.jpg', '2023-06-07 03:00:35', '2023-06-07 03:00:35'),
(32, 'Lindsay', 'Ferguson', 'lindsay.ferguson@reqres.in', 'https://reqres.in/img/faces/8-image.jpg', '2023-06-07 03:00:35', '2023-06-07 03:00:35'),
(33, 'Tobias', 'Funke', 'tobias.funke@reqres.in', 'https://reqres.in/img/faces/9-image.jpg', '2023-06-07 03:00:35', '2023-06-07 03:00:35'),
(34, 'Byron', 'Fields', 'byron.fields@reqres.in', 'https://reqres.in/img/faces/10-image.jpg', '2023-06-07 03:00:35', '2023-06-07 03:00:35'),
(35, 'George', 'Edwards', 'george.edwards@reqres.in', 'https://reqres.in/img/faces/11-image.jpg', '2023-06-07 03:00:35', '2023-06-07 03:00:35'),
(36, 'Rachel', 'Howell', 'rachel.howell@reqres.in', 'https://reqres.in/img/faces/12-image.jpg', '2023-06-07 03:00:35', '2023-06-07 03:00:35'),
(38, 'john6', 'doe6', 'johndoe6@gmailcom', 'john6.jpg', '2023-06-08 08:43:11', '2023-06-08 08:43:11');

-- --------------------------------------------------------

--
-- Table structure for table `register_approvals`
--

CREATE TABLE `register_approvals` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `admin_id` bigint(20) UNSIGNED NOT NULL,
  `super_admin_id` bigint(20) UNSIGNED NOT NULL,
  `status` varchar(25) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` int(10) UNSIGNED NOT NULL,
  `role_name` varchar(15) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `role_name`) VALUES
(1, 'superadmin'),
(2, 'admin');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `actors`
--
ALTER TABLE `actors`
  ADD PRIMARY KEY (`id`),
  ADD KEY `role_id` (`role_id`);

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `register_approvals`
--
ALTER TABLE `register_approvals`
  ADD PRIMARY KEY (`id`),
  ADD KEY `admin_id` (`admin_id`),
  ADD KEY `super_admin_id` (`super_admin_id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `actors`
--
ALTER TABLE `actors`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=39;

--
-- AUTO_INCREMENT for table `register_approvals`
--
ALTER TABLE `register_approvals`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `actors`
--
ALTER TABLE `actors`
  ADD CONSTRAINT `actors_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

--
-- Constraints for table `register_approvals`
--
ALTER TABLE `register_approvals`
  ADD CONSTRAINT `register_approvals_ibfk_1` FOREIGN KEY (`admin_id`) REFERENCES `actors` (`id`),
  ADD CONSTRAINT `register_approvals_ibfk_2` FOREIGN KEY (`super_admin_id`) REFERENCES `actors` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS ROLE;
