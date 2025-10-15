-- +goose Up
-- phpMyAdmin SQL Dump
-- version 5.1.3
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Aug 07, 2024 at 07:16 AM
-- Server version: 8.0.38
-- PHP Version: 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `offybox_v2`
--

-- --------------------------------------------------------

--
-- Table structure for table `core_area`
--

CREATE TABLE `core_area` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `pincode` int DEFAULT NULL,
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `city_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `state_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `country_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_category`
--

CREATE TABLE `core_category` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `parent_category_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sequence` int NOT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_city`
--

CREATE TABLE `core_city` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `state_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `country_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_company`
--

CREATE TABLE `core_company` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `website` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `mobile` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `incorporation_date` date DEFAULT NULL,
  `gst_no` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `address` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `area_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `city_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `state_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `country_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_country`
--

CREATE TABLE `core_country` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_employee`
--

CREATE TABLE `core_employee` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `mobile` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `email` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_employee_hierarchy`
--

CREATE TABLE `core_employee_hierarchy` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `employee_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `supervisor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_lookup_master`
--

CREATE TABLE `core_lookup_master` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `lu_key` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `lu_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `lu_value` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `group_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_module`
--

CREATE TABLE `core_module` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `system` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `parent_module_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `allowed_permission` json DEFAULT NULL,
  `url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sequence` int NOT NULL,
  `target` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_product`
--

CREATE TABLE `core_product` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `mrp` double NOT NULL,
  `uom` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `variant` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `barcode` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_product_price`
--

CREATE TABLE `core_product_price` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `area_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `city_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `state_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `country_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `start_date` date NOT NULL,
  `end_date` date DEFAULT NULL,
  `batch_no` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_role`
--

CREATE TABLE `core_role` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `role_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `data_access` enum('ALL','SUBORDINATES') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_role_module`
--

CREATE TABLE `core_role_module` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `role_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `module_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `allowed_permission` json DEFAULT NULL,
  `data_access` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_scheme`
--

CREATE TABLE `core_scheme` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `outlet_category_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `category_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `area_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `city_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `state_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `country_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `start_date` datetime DEFAULT NULL,
  `end_date` datetime DEFAULT NULL,
  `item_start` int DEFAULT NULL,
  `item_end` int DEFAULT NULL,
  `item_desc` double DEFAULT NULL,
  `item_disc_max` double DEFAULT NULL,
  `free_product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_state`
--

CREATE TABLE `core_state` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `country_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_tenant_configuration`
--

CREATE TABLE `core_tenant_configuration` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `data` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_territory`
--

CREATE TABLE `core_territory` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `territory_type_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `parent_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `location_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `location_type_value` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_territory_type`
--

CREATE TABLE `core_territory_type` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `level` int NOT NULL,
  `parent_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_user`
--

CREATE TABLE `core_user` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `first_name` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `last_name` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `username` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `mobile` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_user_role`
--

CREATE TABLE `core_user_role` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `role_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `is_primary` tinyint NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_user_territory`
--

CREATE TABLE `core_user_territory` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `territory_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `territory_type_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `core_user_territory_scope`
--

CREATE TABLE `core_user_territory_scope` (
  `id` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `territory_id` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `territory_type_id` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `ent_category_user_scope`
--

CREATE TABLE `ent_category_user_scope` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `category_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `ent_distributor`
--

CREATE TABLE `ent_distributor` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `mobile` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `point_of_contact` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `ent_distributor_user`
--

CREATE TABLE `ent_distributor_user` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `ent_distributor_user_hierarchy`
--

CREATE TABLE `ent_distributor_user_hierarchy` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `distributor_user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `supervisor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `ent_outlet`
--

CREATE TABLE `ent_outlet` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `type` enum('L','O') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `mobile` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `credit_limit` double NOT NULL,
  `outstanding` double NOT NULL,
  `incorporation_date` date DEFAULT NULL,
  `created_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `outlet_category_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `ent_outlet_address`
--

CREATE TABLE `ent_outlet_address` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `outlet_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `address_type` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `address` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `area_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `city_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `state_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `country_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `landmark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `pincode` int NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `ent_outlet_category`
--

CREATE TABLE `ent_outlet_category` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `ent_outlet_contact_person`
--

CREATE TABLE `ent_outlet_contact_person` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `title` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `mobile` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `email` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `position` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `outlet_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `ent_outlet_license`
--

CREATE TABLE `ent_outlet_license` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `license_type` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `license_no` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `license_url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `outlet_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `ent_stock`
--

CREATE TABLE `ent_stock` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `warehouse_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `qty` int NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `batch_no` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `ent_stock_log`
--

CREATE TABLE `ent_stock_log` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `type` enum('IN','OUT') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `warehouse_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `stock_date` datetime NOT NULL,
  `qty` int NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `batch_no` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `ent_warehouse`
--

CREATE TABLE `ent_warehouse` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `sale_invoice`
--

CREATE TABLE `sale_invoice` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `outlet_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `item_count` int NOT NULL,
  `net_amount` double NOT NULL,
  `discount_amount` double NOT NULL,
  `tax_amount` double NOT NULL,
  `total_amount` double NOT NULL,
  `quotation_date` datetime NOT NULL,
  `created_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `remarks` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `latitude` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `longitude` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `outlet_address_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `file_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `order_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `paid_amount` double DEFAULT NULL,
  `unpaid_amount` double DEFAULT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `sale_invoice_item`
--

CREATE TABLE `sale_invoice_item` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `free_product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `qty` int NOT NULL,
  `price` double NOT NULL,
  `net_amount` double NOT NULL,
  `discount_amount` double NOT NULL,
  `tax_amount` double NOT NULL,
  `total_amount` double NOT NULL,
  `tax_percentage` double NOT NULL,
  `scheme_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `invoice_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `sale_invoice_scheme`
--

CREATE TABLE `sale_invoice_scheme` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `invoice_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `scheme_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `free_scheme_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `disc_qty` double DEFAULT NULL,
  `disc_value` double DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `sale_journey_plan`
--

CREATE TABLE `sale_journey_plan` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `outlet_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `assigned_date` datetime DEFAULT NULL,
  `closed_date` datetime DEFAULT NULL,
  `remarks` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `sale_order`
--

CREATE TABLE `sale_order` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `outlet_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `item_count` int NOT NULL,
  `net_amount` double NOT NULL,
  `discount_amount` double NOT NULL,
  `tax_amount` double NOT NULL,
  `total_amount` double NOT NULL,
  `quotation_date` datetime NOT NULL,
  `created_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `remarks` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `latitude` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `longitude` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `outlet_address_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `file_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `order_status` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `sale_order_item`
--

CREATE TABLE `sale_order_item` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `free_product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `qty` int NOT NULL,
  `free_qty` int DEFAULT NULL,
  `price` double NOT NULL,
  `net_amount` double NOT NULL,
  `discount_amount` double NOT NULL,
  `tax_amount` double NOT NULL,
  `total_amount` double NOT NULL,
  `tax_percentage` double NOT NULL,
  `scheme_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `order_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `sale_order_scheme`
--

CREATE TABLE `sale_order_scheme` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `order_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `scheme_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `free_scheme_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `disc_qty` double DEFAULT NULL,
  `disc_value` double DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `sale_participant`
--

CREATE TABLE `sale_participant` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `order_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `participant_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `task_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `sale_trip`
--

CREATE TABLE `sale_trip` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `start_date` datetime NOT NULL,
  `end_date` datetime DEFAULT NULL,
  `vehicle_number` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `vehicle_type` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `vehicle_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `driver_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `driver_proof` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `start_km` double DEFAULT NULL,
  `end_km` double DEFAULT NULL,
  `loaded_qty` double DEFAULT NULL,
  `returned_qty` double DEFAULT NULL,
  `damaged_qty` double DEFAULT NULL,
  `distributor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `sale_trip_item`
--

CREATE TABLE `sale_trip_item` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `trip_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `order_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `invoice_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `outlet_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `qty` int DEFAULT NULL,
  `free_qty` int DEFAULT NULL,
  `remarks` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `workflow`
--

CREATE TABLE `workflow` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `start_date` datetime DEFAULT NULL,
  `end_date` datetime DEFAULT NULL,
  `workflow_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `workflow_instance`
--

CREATE TABLE `workflow_instance` (
  `id` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `workflow_id` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `source_id` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `source_type` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `reference_id` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `reference_type` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `definition` json NOT NULL,
  `previous_instance_id` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `parent_instance_id` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `start_date` datetime NOT NULL,
  `end_date` datetime DEFAULT NULL,
  `active_task_id` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `completed_task_id` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `workflow_rule`
--

CREATE TABLE `workflow_rule` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `rule` json NOT NULL,
  `name` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `workflow_stage`
--

CREATE TABLE `workflow_stage` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `workflow_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `sequence` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `configuration` json DEFAULT NULL,
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `workflow_step`
--

CREATE TABLE `workflow_step` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `workflow_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `workflow_stage_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `step_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `configuration` json DEFAULT NULL,
  `ui_component` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `automatic_component` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `allocation_rule_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `display_rule_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `workflow_task`
--

CREATE TABLE `workflow_task` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `workflow_instance_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `workflow_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `workflow_step_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_role_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `assigned_date` datetime DEFAULT NULL,
  `closed_date` datetime DEFAULT NULL,
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `core_area`
--
ALTER TABLE `core_area`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_state_core_country1_idx` (`country_id`),
  ADD KEY `fk_core_city_core_state1_idx` (`state_id`),
  ADD KEY `fk_core_area_core_city1_idx` (`city_id`);

--
-- Indexes for table `core_category`
--
ALTER TABLE `core_category`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code_UNIQUE` (`code`),
  ADD KEY `fk_com_category_com_category1_idx` (`parent_category_id`),
  ADD KEY `fk_com_category_core_distributor1_idx` (`distributor_id`);

--
-- Indexes for table `core_city`
--
ALTER TABLE `core_city`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code_UNIQUE` (`code`),
  ADD KEY `fk_core_state_core_country1_idx` (`country_id`),
  ADD KEY `fk_core_city_core_state1_idx` (`state_id`);

--
-- Indexes for table `core_company`
--
ALTER TABLE `core_company`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_company_core_area1_idx` (`area_id`),
  ADD KEY `fk_company_core_city1_idx` (`city_id`),
  ADD KEY `fk_company_core_state1_idx` (`state_id`),
  ADD KEY `fk_company_core_country1_idx` (`country_id`);

--
-- Indexes for table `core_country`
--
ALTER TABLE `core_country`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code_UNIQUE` (`code`);

--
-- Indexes for table `core_employee`
--
ALTER TABLE `core_employee`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_employee_core_user1_idx` (`user_id`);

--
-- Indexes for table `core_employee_hierarchy`
--
ALTER TABLE `core_employee_hierarchy`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_employee_hierarchy_core_employee1_idx` (`employee_id`),
  ADD KEY `fk_core_employee_hierarchy_core_employee2_idx` (`supervisor_id`);

--
-- Indexes for table `core_lookup_master`
--
ALTER TABLE `core_lookup_master`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `lu_key_UNIQUE` (`lu_key`,`group_code`,`distributor_id`),
  ADD KEY `fk_core_lookup_master_core_distributor1_idx` (`distributor_id`);

--
-- Indexes for table `core_module`
--
ALTER TABLE `core_module`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_module_core_module_idx` (`parent_module_id`);

--
-- Indexes for table `core_product`
--
ALTER TABLE `core_product`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_com_product_core_distributor1_idx` (`distributor_id`);

--
-- Indexes for table `core_product_price`
--
ALTER TABLE `core_product_price`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_com_product_ price_com_product1_idx` (`product_id`),
  ADD KEY `fk_com_product_ price_core_area1_idx` (`area_id`),
  ADD KEY `fk_com_product_ price_core_city1_idx` (`city_id`),
  ADD KEY `fk_com_product_ price_core_state1_idx` (`state_id`),
  ADD KEY `fk_com_product_ price_core_country1_idx` (`country_id`),
  ADD KEY `fk_com_product_ price_core_user1_idx` (`created_by`);

--
-- Indexes for table `core_role`
--
ALTER TABLE `core_role`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_role_core_role1_idx` (`role_id`),
  ADD KEY `fk_core_role_core_distributor1_idx` (`distributor_id`);

--
-- Indexes for table `core_role_module`
--
ALTER TABLE `core_role_module`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_role_module_core_role1_idx` (`role_id`),
  ADD KEY `fk_core_role_module_core_module1_idx` (`module_id`);

--
-- Indexes for table `core_scheme`
--
ALTER TABLE `core_scheme`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_com_scheme_com_product1_idx` (`product_id`),
  ADD KEY `fk_com_scheme_com_category1_idx` (`category_id`),
  ADD KEY `fk_com_scheme_com_outlet_category1_idx` (`outlet_category_id`),
  ADD KEY `fk_com_scheme_core_area1_idx` (`area_id`),
  ADD KEY `fk_com_scheme_core_city1_idx` (`city_id`),
  ADD KEY `fk_com_scheme_core_state1_idx` (`state_id`),
  ADD KEY `fk_com_scheme_core_country1_idx` (`country_id`),
  ADD KEY `fk_com_scheme_com_product2_idx` (`free_product_id`),
  ADD KEY `fk_com_scheme_core_user1_idx` (`created_by`),
  ADD KEY `fk_com_scheme_core_distributor1_idx` (`distributor_id`);

--
-- Indexes for table `core_state`
--
ALTER TABLE `core_state`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code_UNIQUE` (`code`),
  ADD KEY `fk_core_state_core_country1_idx` (`country_id`);

--
-- Indexes for table `core_tenant_configuration`
--
ALTER TABLE `core_tenant_configuration`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code_UNIQUE` (`code`);

--
-- Indexes for table `core_territory`
--
ALTER TABLE `core_territory`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_territory_core_territory_type1_idx` (`territory_type_id`),
  ADD KEY `fk_core_territory_core_territory1_idx` (`parent_id`),
  ADD KEY `fk_core_territory_ent_distributor1_idx` (`distributor_id`);

--
-- Indexes for table `core_territory_type`
--
ALTER TABLE `core_territory_type`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_territory_type_core_territory_type1_idx` (`parent_id`),
  ADD KEY `fk_core_territory_type_ent_distributor1_idx` (`distributor_id`);

--
-- Indexes for table `core_user`
--
ALTER TABLE `core_user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email_UNIQUE` (`email`,`type`),
  ADD UNIQUE KEY `mobile_UNIQUE` (`type`,`mobile`);

--
-- Indexes for table `core_user_role`
--
ALTER TABLE `core_user_role`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_user_role_core_role1_idx` (`role_id`),
  ADD KEY `fk_core_user_role_core_user1_idx` (`user_id`);

--
-- Indexes for table `core_user_territory`
--
ALTER TABLE `core_user_territory`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_user_territory_core_user1_idx` (`user_id`),
  ADD KEY `fk_core_user_territory_core_territory1_idx` (`territory_id`),
  ADD KEY `fk_core_user_territory_core_territory_type1_idx` (`territory_type_id`);

--
-- Indexes for table `core_user_territory_scope`
--
ALTER TABLE `core_user_territory_scope`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_user_territory_core_user1_idx` (`user_id`),
  ADD KEY `fk_core_user_territory_core_territory1_idx` (`territory_id`),
  ADD KEY `fk_core_user_territory_core_territory_type1_idx` (`territory_type_id`);

--
-- Indexes for table `ent_category_user_scope`
--
ALTER TABLE `ent_category_user_scope`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `category_id_UNIQUE` (`category_id`,`user_id`),
  ADD KEY `fk_com_category_user_scope_com_category1_idx` (`category_id`),
  ADD KEY `fk_com_category_user_scope_core_user1_idx` (`user_id`);

--
-- Indexes for table `ent_distributor`
--
ALTER TABLE `ent_distributor`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code_UNIQUE` (`code`);

--
-- Indexes for table `ent_distributor_user`
--
ALTER TABLE `ent_distributor_user`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_distributor_user_core_distributor1_idx` (`distributor_id`),
  ADD KEY `fk_core_distributor_user_core_user1_idx` (`user_id`);

--
-- Indexes for table `ent_distributor_user_hierarchy`
--
ALTER TABLE `ent_distributor_user_hierarchy`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id_UNIQUE` (`distributor_user_id`,`supervisor_id`),
  ADD KEY `fk_core_distributor_user_hierarchy_core_distributor_user1_idx` (`distributor_user_id`),
  ADD KEY `fk_core_distributor_user_hierarchy_core_distributor1_idx` (`distributor_id`),
  ADD KEY `fk_core_distributor_user_hierarchy_core_distributor_user2_idx` (`supervisor_id`);

--
-- Indexes for table `ent_outlet`
--
ALTER TABLE `ent_outlet`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_outlet_core_user1_idx` (`created_by`),
  ADD KEY `fk_core_outlet_core_outlet_category1_idx` (`outlet_category_id`),
  ADD KEY `fk_com_outlet_core_distributor1_idx` (`distributor_id`);

--
-- Indexes for table `ent_outlet_address`
--
ALTER TABLE `ent_outlet_address`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_com_outlet_address_com_outlet1_idx` (`outlet_id`),
  ADD KEY `fk_com_outlet_address_core_area1_idx` (`area_id`),
  ADD KEY `fk_com_outlet_address_core_state1_idx` (`state_id`),
  ADD KEY `fk_com_outlet_address_core_city1_idx` (`city_id`),
  ADD KEY `fk_com_outlet_address_core_country1_idx` (`country_id`);

--
-- Indexes for table `ent_outlet_category`
--
ALTER TABLE `ent_outlet_category`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_com_outlet_category_core_distributor1_idx` (`distributor_id`);

--
-- Indexes for table `ent_outlet_contact_person`
--
ALTER TABLE `ent_outlet_contact_person`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_outlet_contact_person_core_outlet1_idx` (`outlet_id`);

--
-- Indexes for table `ent_outlet_license`
--
ALTER TABLE `ent_outlet_license`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_com_outlet_license_com_outlet1_idx` (`outlet_id`);

--
-- Indexes for table `ent_stock`
--
ALTER TABLE `ent_stock`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_com_stock_com_warehouse1_idx` (`warehouse_id`),
  ADD KEY `fk_com_stock_core_distributor1_idx` (`distributor_id`),
  ADD KEY `fk_com_stock_core_user1_idx` (`user_id`),
  ADD KEY `fk_com_stock_com_product1_idx` (`product_id`);

--
-- Indexes for table `ent_stock_log`
--
ALTER TABLE `ent_stock_log`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_com_stock_com_warehouse1_idx` (`warehouse_id`),
  ADD KEY `fk_com_stock_core_distributor1_idx` (`distributor_id`),
  ADD KEY `fk_com_stock_core_user1_idx` (`user_id`),
  ADD KEY `fk_com_stock_com_product1_idx` (`product_id`);

--
-- Indexes for table `ent_warehouse`
--
ALTER TABLE `ent_warehouse`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `sale_invoice`
--
ALTER TABLE `sale_invoice`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_com_quotation_com_outlet1_idx` (`outlet_id`),
  ADD KEY `fk_com_quotation_core_user1_idx` (`created_by`),
  ADD KEY `fk_com_quotation_com_outlet_address1_idx` (`outlet_address_id`),
  ADD KEY `fk_com_invoice_com_order1_idx` (`order_id`),
  ADD KEY `fk_com_invoice_core_distributor1_idx` (`distributor_id`);

--
-- Indexes for table `sale_invoice_item`
--
ALTER TABLE `sale_invoice_item`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_com_quotation_item_com_product1_idx` (`product_id`),
  ADD KEY `fk_com_quotation_item_com_product2_idx` (`free_product_id`),
  ADD KEY `fk_com_quotation_item_com_scheme1_idx` (`scheme_id`),
  ADD KEY `fk_com_invoice_item_com_invoice1_idx` (`invoice_id`);

--
-- Indexes for table `sale_invoice_scheme`
--
ALTER TABLE `sale_invoice_scheme`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_sales_order_scheme_com_scheme1_idx` (`scheme_id`),
  ADD KEY `fk_sales_order_scheme_com_product1_idx` (`product_id`),
  ADD KEY `fk_sales_order_scheme_com_scheme2_idx` (`free_scheme_id`),
  ADD KEY `fk_sales_invoice_scheme_sale_invoice1_idx` (`invoice_id`);

--
-- Indexes for table `sale_journey_plan`
--
ALTER TABLE `sale_journey_plan`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_com_journey_plan_com_outlet1_idx` (`outlet_id`),
  ADD KEY `fk_com_journey_plan_core_user1_idx` (`user_id`);

--
-- Indexes for table `sale_order`
--
ALTER TABLE `sale_order`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_com_quotation_com_outlet1_idx` (`outlet_id`),
  ADD KEY `fk_com_quotation_core_user1_idx` (`created_by`),
  ADD KEY `fk_com_quotation_com_outlet_address1_idx` (`outlet_address_id`),
  ADD KEY `fk_com_order_core_distributor1_idx` (`distributor_id`);

--
-- Indexes for table `sale_order_item`
--
ALTER TABLE `sale_order_item`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_com_quotation_item_com_product1_idx` (`product_id`),
  ADD KEY `fk_com_quotation_item_com_product2_idx` (`free_product_id`),
  ADD KEY `fk_com_quotation_item_com_scheme1_idx` (`scheme_id`),
  ADD KEY `fk_com_order_item_com_order1_idx` (`order_id`);

--
-- Indexes for table `sale_order_scheme`
--
ALTER TABLE `sale_order_scheme`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_sales_order_scheme_sale_order1_idx` (`order_id`),
  ADD KEY `fk_sales_order_scheme_com_scheme1_idx` (`scheme_id`),
  ADD KEY `fk_sales_order_scheme_com_product1_idx` (`product_id`),
  ADD KEY `fk_sales_order_scheme_com_scheme2_idx` (`free_scheme_id`);

--
-- Indexes for table `sale_participant`
--
ALTER TABLE `sale_participant`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `order_id_UNIQUE` (`order_id`,`user_id`,`distributor_id`),
  ADD KEY `fk_sales_participant_sale_order1_idx` (`order_id`),
  ADD KEY `fk_sales_participant_core_user1_idx` (`user_id`),
  ADD KEY `fk_sales_participant_ent_distributor1_idx` (`distributor_id`),
  ADD KEY `fk_sales_participant_workflow_task1_idx` (`task_id`);

--
-- Indexes for table `sale_trip`
--
ALTER TABLE `sale_trip`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_com_trip_core_user1_idx` (`user_id`),
  ADD KEY `fk_sale_trip_ent_distributor1_idx` (`distributor_id`);

--
-- Indexes for table `sale_trip_item`
--
ALTER TABLE `sale_trip_item`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_core_trip_item_com_trip1_idx` (`trip_id`),
  ADD KEY `fk_core_trip_item_com_order1_idx` (`order_id`),
  ADD KEY `fk_core_trip_item_com_invoice1_idx` (`invoice_id`),
  ADD KEY `fk_core_trip_item_com_outlet1_idx` (`outlet_id`),
  ADD KEY `fk_core_trip_item_com_product1_idx` (`product_id`);

--
-- Indexes for table `workflow`
--
ALTER TABLE `workflow`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `workflow_instance`
--
ALTER TABLE `workflow_instance`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_workflow_instance_workflow1_idx` (`workflow_id`),
  ADD KEY `fk_workflow_instance_workflow_instance1_idx` (`previous_instance_id`),
  ADD KEY `fk_workflow_instance_workflow_instance2_idx` (`parent_instance_id`),
  ADD KEY `fk_workflow_instance_workflow_task1_idx` (`active_task_id`),
  ADD KEY `fk_workflow_instance_workflow_task2_idx` (`completed_task_id`);

--
-- Indexes for table `workflow_rule`
--
ALTER TABLE `workflow_rule`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `workflow_stage`
--
ALTER TABLE `workflow_stage`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_workflow_stage_workflow1_idx` (`workflow_id`);

--
-- Indexes for table `workflow_step`
--
ALTER TABLE `workflow_step`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_workflow_step_workflow1_idx` (`workflow_id`),
  ADD KEY `fk_workflow_step_workflow_stage1_idx` (`workflow_stage_id`),
  ADD KEY `fk_workflow_step_workflow_rule1_idx` (`allocation_rule_id`),
  ADD KEY `fk_workflow_step_workflow_rule2_idx` (`display_rule_id`);

--
-- Indexes for table `workflow_task`
--
ALTER TABLE `workflow_task`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_workflow_task_workflow_instance1_idx` (`workflow_instance_id`),
  ADD KEY `fk_workflow_task_workflow1_idx` (`workflow_id`),
  ADD KEY `fk_workflow_task_workflow_step1_idx` (`workflow_step_id`),
  ADD KEY `fk_workflow_task_core_user1_idx` (`user_id`),
  ADD KEY `fk_workflow_task_core_user_role1_idx` (`user_role_id`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `core_area`
--
ALTER TABLE `core_area`
  ADD CONSTRAINT `fk_core_area_core_city1` FOREIGN KEY (`city_id`) REFERENCES `core_city` (`id`),
  ADD CONSTRAINT `fk_core_city_core_state10` FOREIGN KEY (`state_id`) REFERENCES `core_state` (`id`),
  ADD CONSTRAINT `fk_core_state_core_country100` FOREIGN KEY (`country_id`) REFERENCES `core_country` (`id`);

--
-- Constraints for table `core_category`
--
ALTER TABLE `core_category`
  ADD CONSTRAINT `fk_com_category_com_category1` FOREIGN KEY (`parent_category_id`) REFERENCES `core_category` (`id`),
  ADD CONSTRAINT `fk_com_category_core_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`);

--
-- Constraints for table `core_city`
--
ALTER TABLE `core_city`
  ADD CONSTRAINT `fk_core_city_core_state1` FOREIGN KEY (`state_id`) REFERENCES `core_state` (`id`),
  ADD CONSTRAINT `fk_core_state_core_country10` FOREIGN KEY (`country_id`) REFERENCES `core_country` (`id`);

--
-- Constraints for table `core_company`
--
ALTER TABLE `core_company`
  ADD CONSTRAINT `fk_company_core_area1` FOREIGN KEY (`area_id`) REFERENCES `core_area` (`id`),
  ADD CONSTRAINT `fk_company_core_city1` FOREIGN KEY (`city_id`) REFERENCES `core_city` (`id`),
  ADD CONSTRAINT `fk_company_core_country1` FOREIGN KEY (`country_id`) REFERENCES `core_country` (`id`),
  ADD CONSTRAINT `fk_company_core_state1` FOREIGN KEY (`state_id`) REFERENCES `core_state` (`id`);

--
-- Constraints for table `core_employee`
--
ALTER TABLE `core_employee`
  ADD CONSTRAINT `fk_core_employee_core_user1` FOREIGN KEY (`user_id`) REFERENCES `core_user` (`id`);

--
-- Constraints for table `core_employee_hierarchy`
--
ALTER TABLE `core_employee_hierarchy`
  ADD CONSTRAINT `fk_core_employee_hierarchy_core_employee1` FOREIGN KEY (`employee_id`) REFERENCES `core_employee` (`id`),
  ADD CONSTRAINT `fk_core_employee_hierarchy_core_employee2` FOREIGN KEY (`supervisor_id`) REFERENCES `core_employee` (`id`);

--
-- Constraints for table `core_lookup_master`
--
ALTER TABLE `core_lookup_master`
  ADD CONSTRAINT `fk_core_lookup_master_core_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`);

--
-- Constraints for table `core_module`
--
ALTER TABLE `core_module`
  ADD CONSTRAINT `fk_core_module_core_module` FOREIGN KEY (`parent_module_id`) REFERENCES `core_module` (`id`);

--
-- Constraints for table `core_product`
--
ALTER TABLE `core_product`
  ADD CONSTRAINT `fk_com_product_core_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`);

--
-- Constraints for table `core_product_price`
--
ALTER TABLE `core_product_price`
  ADD CONSTRAINT `fk_com_product_ price_com_product1` FOREIGN KEY (`product_id`) REFERENCES `core_product` (`id`),
  ADD CONSTRAINT `fk_com_product_ price_core_area1` FOREIGN KEY (`area_id`) REFERENCES `core_area` (`id`),
  ADD CONSTRAINT `fk_com_product_ price_core_city1` FOREIGN KEY (`city_id`) REFERENCES `core_city` (`id`),
  ADD CONSTRAINT `fk_com_product_ price_core_country1` FOREIGN KEY (`country_id`) REFERENCES `core_country` (`id`),
  ADD CONSTRAINT `fk_com_product_ price_core_state1` FOREIGN KEY (`state_id`) REFERENCES `core_state` (`id`),
  ADD CONSTRAINT `fk_com_product_ price_core_user1` FOREIGN KEY (`created_by`) REFERENCES `core_user` (`id`);

--
-- Constraints for table `core_role`
--
ALTER TABLE `core_role`
  ADD CONSTRAINT `fk_core_role_core_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`),
  ADD CONSTRAINT `fk_core_role_core_role1` FOREIGN KEY (`role_id`) REFERENCES `core_role` (`id`);

--
-- Constraints for table `core_role_module`
--
ALTER TABLE `core_role_module`
  ADD CONSTRAINT `fk_core_role_module_core_module1` FOREIGN KEY (`module_id`) REFERENCES `core_module` (`id`),
  ADD CONSTRAINT `fk_core_role_module_core_role1` FOREIGN KEY (`role_id`) REFERENCES `core_role` (`id`);

--
-- Constraints for table `core_scheme`
--
ALTER TABLE `core_scheme`
  ADD CONSTRAINT `fk_com_scheme_com_category1` FOREIGN KEY (`category_id`) REFERENCES `core_category` (`id`),
  ADD CONSTRAINT `fk_com_scheme_com_outlet_category1` FOREIGN KEY (`outlet_category_id`) REFERENCES `ent_outlet_category` (`id`),
  ADD CONSTRAINT `fk_com_scheme_com_product1` FOREIGN KEY (`product_id`) REFERENCES `core_product` (`id`),
  ADD CONSTRAINT `fk_com_scheme_com_product2` FOREIGN KEY (`free_product_id`) REFERENCES `core_product` (`id`),
  ADD CONSTRAINT `fk_com_scheme_core_area1` FOREIGN KEY (`area_id`) REFERENCES `core_area` (`id`),
  ADD CONSTRAINT `fk_com_scheme_core_city1` FOREIGN KEY (`city_id`) REFERENCES `core_city` (`id`),
  ADD CONSTRAINT `fk_com_scheme_core_country1` FOREIGN KEY (`country_id`) REFERENCES `core_country` (`id`),
  ADD CONSTRAINT `fk_com_scheme_core_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`),
  ADD CONSTRAINT `fk_com_scheme_core_state1` FOREIGN KEY (`state_id`) REFERENCES `core_state` (`id`),
  ADD CONSTRAINT `fk_com_scheme_core_user1` FOREIGN KEY (`created_by`) REFERENCES `core_user` (`id`);

--
-- Constraints for table `core_state`
--
ALTER TABLE `core_state`
  ADD CONSTRAINT `fk_core_state_core_country1` FOREIGN KEY (`country_id`) REFERENCES `core_country` (`id`);

--
-- Constraints for table `core_territory`
--
ALTER TABLE `core_territory`
  ADD CONSTRAINT `fk_core_territory_core_territory1` FOREIGN KEY (`parent_id`) REFERENCES `core_territory` (`id`),
  ADD CONSTRAINT `fk_core_territory_core_territory_type1` FOREIGN KEY (`territory_type_id`) REFERENCES `core_territory_type` (`id`),
  ADD CONSTRAINT `fk_core_territory_ent_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`);

--
-- Constraints for table `core_territory_type`
--
ALTER TABLE `core_territory_type`
  ADD CONSTRAINT `fk_core_territory_type_core_territory_type1` FOREIGN KEY (`parent_id`) REFERENCES `core_territory_type` (`id`),
  ADD CONSTRAINT `fk_core_territory_type_ent_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`);

--
-- Constraints for table `core_user_role`
--
ALTER TABLE `core_user_role`
  ADD CONSTRAINT `fk_core_user_role_core_role1` FOREIGN KEY (`role_id`) REFERENCES `core_role` (`id`),
  ADD CONSTRAINT `fk_core_user_role_core_user1` FOREIGN KEY (`user_id`) REFERENCES `core_user` (`id`);

--
-- Constraints for table `core_user_territory`
--
ALTER TABLE `core_user_territory`
  ADD CONSTRAINT `fk_core_user_territory_core_territory1` FOREIGN KEY (`territory_id`) REFERENCES `core_territory` (`id`),
  ADD CONSTRAINT `fk_core_user_territory_core_territory_type1` FOREIGN KEY (`territory_type_id`) REFERENCES `core_territory_type` (`id`),
  ADD CONSTRAINT `fk_core_user_territory_core_user1` FOREIGN KEY (`user_id`) REFERENCES `core_user` (`id`);

--
-- Constraints for table `core_user_territory_scope`
--
ALTER TABLE `core_user_territory_scope`
  ADD CONSTRAINT `fk_core_user_territory_core_territory10` FOREIGN KEY (`territory_id`) REFERENCES `core_territory` (`id`),
  ADD CONSTRAINT `fk_core_user_territory_core_territory_type10` FOREIGN KEY (`territory_type_id`) REFERENCES `core_territory_type` (`id`),
  ADD CONSTRAINT `fk_core_user_territory_core_user10` FOREIGN KEY (`user_id`) REFERENCES `core_user` (`id`);

--
-- Constraints for table `ent_category_user_scope`
--
ALTER TABLE `ent_category_user_scope`
  ADD CONSTRAINT `fk_com_category_user_scope_com_category1` FOREIGN KEY (`category_id`) REFERENCES `core_category` (`id`),
  ADD CONSTRAINT `fk_com_category_user_scope_core_user1` FOREIGN KEY (`user_id`) REFERENCES `core_user` (`id`);

--
-- Constraints for table `ent_distributor_user`
--
ALTER TABLE `ent_distributor_user`
  ADD CONSTRAINT `fk_core_distributor_user_core_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`),
  ADD CONSTRAINT `fk_core_distributor_user_core_user1` FOREIGN KEY (`user_id`) REFERENCES `core_user` (`id`);

--
-- Constraints for table `ent_distributor_user_hierarchy`
--
ALTER TABLE `ent_distributor_user_hierarchy`
  ADD CONSTRAINT `fk_core_distributor_user_hierarchy_core_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`),
  ADD CONSTRAINT `fk_core_distributor_user_hierarchy_core_distributor_user1` FOREIGN KEY (`distributor_user_id`) REFERENCES `ent_distributor_user` (`id`),
  ADD CONSTRAINT `fk_core_distributor_user_hierarchy_core_distributor_user2` FOREIGN KEY (`supervisor_id`) REFERENCES `ent_distributor_user` (`id`);

--
-- Constraints for table `ent_outlet`
--
ALTER TABLE `ent_outlet`
  ADD CONSTRAINT `fk_com_outlet_core_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`),
  ADD CONSTRAINT `fk_core_outlet_core_outlet_category1` FOREIGN KEY (`outlet_category_id`) REFERENCES `ent_outlet_category` (`id`),
  ADD CONSTRAINT `fk_core_outlet_core_user1` FOREIGN KEY (`created_by`) REFERENCES `core_user` (`id`);

--
-- Constraints for table `ent_outlet_address`
--
ALTER TABLE `ent_outlet_address`
  ADD CONSTRAINT `fk_com_outlet_address_com_outlet1` FOREIGN KEY (`outlet_id`) REFERENCES `ent_outlet` (`id`),
  ADD CONSTRAINT `fk_com_outlet_address_core_area1` FOREIGN KEY (`area_id`) REFERENCES `core_area` (`id`),
  ADD CONSTRAINT `fk_com_outlet_address_core_city1` FOREIGN KEY (`city_id`) REFERENCES `core_city` (`id`),
  ADD CONSTRAINT `fk_com_outlet_address_core_country1` FOREIGN KEY (`country_id`) REFERENCES `core_country` (`id`),
  ADD CONSTRAINT `fk_com_outlet_address_core_state1` FOREIGN KEY (`state_id`) REFERENCES `core_state` (`id`);

--
-- Constraints for table `ent_outlet_category`
--
ALTER TABLE `ent_outlet_category`
  ADD CONSTRAINT `fk_com_outlet_category_core_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`);

--
-- Constraints for table `ent_outlet_contact_person`
--
ALTER TABLE `ent_outlet_contact_person`
  ADD CONSTRAINT `fk_core_outlet_contact_person_core_outlet1` FOREIGN KEY (`outlet_id`) REFERENCES `ent_outlet` (`id`);

--
-- Constraints for table `ent_outlet_license`
--
ALTER TABLE `ent_outlet_license`
  ADD CONSTRAINT `fk_com_outlet_license_com_outlet1` FOREIGN KEY (`outlet_id`) REFERENCES `ent_outlet` (`id`);

--
-- Constraints for table `ent_stock`
--
ALTER TABLE `ent_stock`
  ADD CONSTRAINT `fk_com_stock_com_product1` FOREIGN KEY (`product_id`) REFERENCES `core_product` (`id`),
  ADD CONSTRAINT `fk_com_stock_com_warehouse1` FOREIGN KEY (`warehouse_id`) REFERENCES `ent_warehouse` (`id`),
  ADD CONSTRAINT `fk_com_stock_core_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`),
  ADD CONSTRAINT `fk_com_stock_core_user1` FOREIGN KEY (`user_id`) REFERENCES `core_user` (`id`);

--
-- Constraints for table `ent_stock_log`
--
ALTER TABLE `ent_stock_log`
  ADD CONSTRAINT `fk_com_stock_com_product10` FOREIGN KEY (`product_id`) REFERENCES `core_product` (`id`),
  ADD CONSTRAINT `fk_com_stock_com_warehouse10` FOREIGN KEY (`warehouse_id`) REFERENCES `ent_warehouse` (`id`),
  ADD CONSTRAINT `fk_com_stock_core_distributor10` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`),
  ADD CONSTRAINT `fk_com_stock_core_user10` FOREIGN KEY (`user_id`) REFERENCES `core_user` (`id`);

--
-- Constraints for table `sale_invoice`
--
ALTER TABLE `sale_invoice`
  ADD CONSTRAINT `fk_com_invoice_com_order1` FOREIGN KEY (`order_id`) REFERENCES `sale_order` (`id`),
  ADD CONSTRAINT `fk_com_invoice_core_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`),
  ADD CONSTRAINT `fk_com_quotation_com_outlet100` FOREIGN KEY (`outlet_id`) REFERENCES `ent_outlet` (`id`),
  ADD CONSTRAINT `fk_com_quotation_com_outlet_address100` FOREIGN KEY (`outlet_address_id`) REFERENCES `ent_outlet_address` (`id`),
  ADD CONSTRAINT `fk_com_quotation_core_user100` FOREIGN KEY (`created_by`) REFERENCES `core_user` (`id`);

--
-- Constraints for table `sale_invoice_item`
--
ALTER TABLE `sale_invoice_item`
  ADD CONSTRAINT `fk_com_invoice_item_com_invoice1` FOREIGN KEY (`invoice_id`) REFERENCES `sale_invoice` (`id`),
  ADD CONSTRAINT `fk_com_quotation_item_com_product100` FOREIGN KEY (`product_id`) REFERENCES `core_product` (`id`),
  ADD CONSTRAINT `fk_com_quotation_item_com_product200` FOREIGN KEY (`free_product_id`) REFERENCES `core_product` (`id`),
  ADD CONSTRAINT `fk_com_quotation_item_com_scheme100` FOREIGN KEY (`scheme_id`) REFERENCES `core_scheme` (`id`);

--
-- Constraints for table `sale_invoice_scheme`
--
ALTER TABLE `sale_invoice_scheme`
  ADD CONSTRAINT `fk_sales_invoice_scheme_sale_invoice1` FOREIGN KEY (`invoice_id`) REFERENCES `sale_invoice` (`id`),
  ADD CONSTRAINT `fk_sales_order_scheme_com_product10` FOREIGN KEY (`product_id`) REFERENCES `core_product` (`id`),
  ADD CONSTRAINT `fk_sales_order_scheme_com_scheme10` FOREIGN KEY (`scheme_id`) REFERENCES `core_scheme` (`id`),
  ADD CONSTRAINT `fk_sales_order_scheme_com_scheme20` FOREIGN KEY (`free_scheme_id`) REFERENCES `core_scheme` (`id`);

--
-- Constraints for table `sale_journey_plan`
--
ALTER TABLE `sale_journey_plan`
  ADD CONSTRAINT `fk_com_journey_plan_com_outlet1` FOREIGN KEY (`outlet_id`) REFERENCES `ent_outlet` (`id`),
  ADD CONSTRAINT `fk_com_journey_plan_core_user1` FOREIGN KEY (`user_id`) REFERENCES `core_user` (`id`);

--
-- Constraints for table `sale_order`
--
ALTER TABLE `sale_order`
  ADD CONSTRAINT `fk_com_order_core_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`),
  ADD CONSTRAINT `fk_com_quotation_com_outlet10` FOREIGN KEY (`outlet_id`) REFERENCES `ent_outlet` (`id`),
  ADD CONSTRAINT `fk_com_quotation_com_outlet_address10` FOREIGN KEY (`outlet_address_id`) REFERENCES `ent_outlet_address` (`id`),
  ADD CONSTRAINT `fk_com_quotation_core_user10` FOREIGN KEY (`created_by`) REFERENCES `core_user` (`id`);

--
-- Constraints for table `sale_order_item`
--
ALTER TABLE `sale_order_item`
  ADD CONSTRAINT `fk_com_order_item_com_order1` FOREIGN KEY (`order_id`) REFERENCES `sale_order` (`id`),
  ADD CONSTRAINT `fk_com_quotation_item_com_product10` FOREIGN KEY (`product_id`) REFERENCES `core_product` (`id`),
  ADD CONSTRAINT `fk_com_quotation_item_com_product20` FOREIGN KEY (`free_product_id`) REFERENCES `core_product` (`id`),
  ADD CONSTRAINT `fk_com_quotation_item_com_scheme10` FOREIGN KEY (`scheme_id`) REFERENCES `core_scheme` (`id`);

--
-- Constraints for table `sale_order_scheme`
--
ALTER TABLE `sale_order_scheme`
  ADD CONSTRAINT `fk_sales_order_scheme_com_product1` FOREIGN KEY (`product_id`) REFERENCES `core_product` (`id`),
  ADD CONSTRAINT `fk_sales_order_scheme_com_scheme1` FOREIGN KEY (`scheme_id`) REFERENCES `core_scheme` (`id`),
  ADD CONSTRAINT `fk_sales_order_scheme_com_scheme2` FOREIGN KEY (`free_scheme_id`) REFERENCES `core_scheme` (`id`),
  ADD CONSTRAINT `fk_sales_order_scheme_sale_order1` FOREIGN KEY (`order_id`) REFERENCES `sale_order` (`id`);

--
-- Constraints for table `sale_participant`
--
ALTER TABLE `sale_participant`
  ADD CONSTRAINT `fk_sales_participant_core_user1` FOREIGN KEY (`user_id`) REFERENCES `core_user` (`id`),
  ADD CONSTRAINT `fk_sales_participant_ent_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`),
  ADD CONSTRAINT `fk_sales_participant_sale_order1` FOREIGN KEY (`order_id`) REFERENCES `sale_order` (`id`),
  ADD CONSTRAINT `fk_sales_participant_workflow_task1` FOREIGN KEY (`task_id`) REFERENCES `workflow_task` (`id`);

--
-- Constraints for table `sale_trip`
--
ALTER TABLE `sale_trip`
  ADD CONSTRAINT `fk_com_trip_core_user1` FOREIGN KEY (`user_id`) REFERENCES `core_user` (`id`),
  ADD CONSTRAINT `fk_sale_trip_ent_distributor1` FOREIGN KEY (`distributor_id`) REFERENCES `ent_distributor` (`id`);

--
-- Constraints for table `sale_trip_item`
--
ALTER TABLE `sale_trip_item`
  ADD CONSTRAINT `fk_core_trip_item_com_invoice1` FOREIGN KEY (`invoice_id`) REFERENCES `sale_invoice` (`id`),
  ADD CONSTRAINT `fk_core_trip_item_com_order1` FOREIGN KEY (`order_id`) REFERENCES `sale_order` (`id`),
  ADD CONSTRAINT `fk_core_trip_item_com_outlet1` FOREIGN KEY (`outlet_id`) REFERENCES `ent_outlet` (`id`),
  ADD CONSTRAINT `fk_core_trip_item_com_product1` FOREIGN KEY (`product_id`) REFERENCES `core_product` (`id`),
  ADD CONSTRAINT `fk_core_trip_item_com_trip1` FOREIGN KEY (`trip_id`) REFERENCES `sale_trip` (`id`);

--
-- Constraints for table `workflow_instance`
--
ALTER TABLE `workflow_instance`
  ADD CONSTRAINT `fk_workflow_instance_workflow1` FOREIGN KEY (`workflow_id`) REFERENCES `workflow` (`id`),
  ADD CONSTRAINT `fk_workflow_instance_workflow_instance1` FOREIGN KEY (`previous_instance_id`) REFERENCES `workflow_instance` (`id`),
  ADD CONSTRAINT `fk_workflow_instance_workflow_instance2` FOREIGN KEY (`parent_instance_id`) REFERENCES `workflow_instance` (`id`),
  ADD CONSTRAINT `fk_workflow_instance_workflow_task1` FOREIGN KEY (`active_task_id`) REFERENCES `workflow_task` (`id`),
  ADD CONSTRAINT `fk_workflow_instance_workflow_task2` FOREIGN KEY (`completed_task_id`) REFERENCES `workflow_task` (`id`);

--
-- Constraints for table `workflow_stage`
--
ALTER TABLE `workflow_stage`
  ADD CONSTRAINT `fk_workflow_stage_workflow1` FOREIGN KEY (`workflow_id`) REFERENCES `workflow` (`id`);

--
-- Constraints for table `workflow_step`
--
ALTER TABLE `workflow_step`
  ADD CONSTRAINT `fk_workflow_step_workflow1` FOREIGN KEY (`workflow_id`) REFERENCES `workflow` (`id`),
  ADD CONSTRAINT `fk_workflow_step_workflow_rule1` FOREIGN KEY (`allocation_rule_id`) REFERENCES `workflow_rule` (`id`),
  ADD CONSTRAINT `fk_workflow_step_workflow_rule2` FOREIGN KEY (`display_rule_id`) REFERENCES `workflow_rule` (`id`),
  ADD CONSTRAINT `fk_workflow_step_workflow_stage1` FOREIGN KEY (`workflow_stage_id`) REFERENCES `workflow_stage` (`id`);

--
-- Constraints for table `workflow_task`
--
ALTER TABLE `workflow_task`
  ADD CONSTRAINT `fk_workflow_task_core_user1` FOREIGN KEY (`user_id`) REFERENCES `core_user` (`id`),
  ADD CONSTRAINT `fk_workflow_task_core_user_role1` FOREIGN KEY (`user_role_id`) REFERENCES `core_user_role` (`id`),
  ADD CONSTRAINT `fk_workflow_task_workflow1` FOREIGN KEY (`workflow_id`) REFERENCES `workflow` (`id`),
  ADD CONSTRAINT `fk_workflow_task_workflow_instance1` FOREIGN KEY (`workflow_instance_id`) REFERENCES `workflow_instance` (`id`),
  ADD CONSTRAINT `fk_workflow_task_workflow_step1` FOREIGN KEY (`workflow_step_id`) REFERENCES `workflow_step` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
