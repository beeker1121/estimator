CREATE TABLE `forms` (
    `id` varchar(36) NOT NULL,
    `modules` JSON NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;