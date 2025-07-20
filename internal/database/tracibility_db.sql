/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `product_traces`;
CREATE TABLE `product_traces` (
  `trace_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `product_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `product_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `event_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `event_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `description` text,
  `location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `user_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `actor` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `tx_hash` varchar(255) DEFAULT NULL,
  `tx_status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `product_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `product_name` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `description` text,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `ipfs_hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `status` varchar(50) DEFAULT NULL,
  `creator` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `user_name` varchar(100) DEFAULT NULL,
  `wallet_address` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `user_role` varchar(50) DEFAULT NULL,
  `password` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `product_traces` (`trace_id`, `product_id`, `product_name`, `event_time`, `event_type`, `description`, `location`, `user_id`, `actor`, `tx_hash`, `tx_status`) VALUES
('42cea427-8334-42d6-a018-fd42800530fb', 'SP-1RQWKNLJ8', '', NULL, '', '', 'Hà Nội', '', '0x5ecE7f92DaDf614D1Be62897C55A8e5A0dB9d089', '0x772e4d1cc52750e7c444848fa280520ae4273301cf9b1dad272c3bcf720954d7', 'SUCCESS'),
('c224e4cd-e22f-44f9-a216-8669002fd181', 'SP-PFQ4R5Q87', '', NULL, 'Received', '', 'Hà Nội', '', '0x5ecE7f92DaDf614D1Be62897C55A8e5A0dB9d089', '0xdfa350b3319db1dd4b33cf1016059f0992e3bb3ae9d420e99c6d235e4dfb61cc', 'SUCCESS'),
('203848a2-28c2-430a-b0c0-0a5c04d5ca85', 'SP-N445TKZMC', '', NULL, 'Created', '', 'Hà Nội', '', '0x5ecE7f92DaDf614D1Be62897C55A8e5A0dB9d089', '0xdac69a90323341ca1cdbe2e56d04a78c2d748209285baa7ae8f50ca61af6f83a', 'SUCCESS'),
('f7c6dc11-549b-44a4-a43b-3ef87664529a', 'SP-N445TKZMC', '', NULL, 'Shipped', '', 'Hải Phòng', '', '0x5ecE7f92DaDf614D1Be62897C55A8e5A0dB9d089', '0x43696f3d4b64b79f7d9f764d34b705302d14ac1a4e0abb7747218cc799571e5c', 'SUCCESS'),
('48a66fa2-5a51-4425-b68a-ee75cc22d2ab', 'SP-N445TKZMC', '', NULL, 'Sold', '', 'Hải Phòng', '', '0x5ecE7f92DaDf614D1Be62897C55A8e5A0dB9d089', '0xb365097e77440898f09832be2ddf787c93348971558f4c2f7a467cbc1452c1b1', 'SUCCESS'),
('f327d852-7647-4190-b0c3-ea166b1461af', 'SP-N445TKZMC', '', NULL, 'Returned', '', 'Hà Nội', '', '0x5ecE7f92DaDf614D1Be62897C55A8e5A0dB9d089', '0xd744c559cf06f6a6838c5ae16d265eea2d378d766f00906c132fb4973e036aa7', 'SUCCESS'),
('a0a6cba9-3871-4081-8fe3-37657294aac2', 'SP-PFQ4R5Q87', '', '2025-07-20 10:54:27', 'Received', '', 'Hà Nội', '', '0x5ecE7f92DaDf614D1Be62897C55A8e5A0dB9d089', '0xcb5641252b96d398beff5907a31a5e91b6325d322dd416c50476666bdb0a18bc', 'SUCCESS');
INSERT INTO `products` (`product_id`, `product_name`, `description`, `created_at`, `updated_at`, `ipfs_hash`, `status`, `creator`, `location`) VALUES
('SP-1RQWKNLJ8', 'Trà sữa Hipp', '', '2025-07-19 14:47:19', '2025-07-19 14:47:19', '', '', '0x5ecE7f92DaDf614D1Be62897C55A8e5A0dB9d089', NULL),
('SP-OSI84I7GG', 'Trà sữa Olong', '', '2025-07-20 10:15:02', '2025-07-20 10:15:02', '', 'Created', '0x5ecE7f92DaDf614D1Be62897C55A8e5A0dB9d089', 'Hải Phòng'),
('SP-N445TKZMC', 'Saltmax', '', '2025-07-20 10:32:25', '2025-07-20 10:32:25', '', 'Created', '0x5ecE7f92DaDf614D1Be62897C55A8e5A0dB9d089', 'Hà Nội');
INSERT INTO `users` (`user_name`, `wallet_address`, `user_role`, `password`) VALUES
('lamnt27', '0x5ecE7f92DaDf614D1Be62897C55A8e5A0dB9d089', 'ADMIN', '$2a$10$cdYC7cAoyihpftC/G6Iyceheh4eShSdueI65k7NWC5H2i3c2ZMGHu');


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;