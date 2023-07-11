CREATE TABLE `accounts` (
    `id` varchar(36) NOT NULL,
    `name` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `users` (
    `id` varchar(36) NOT NULL,
    `account_id` varchar(36) NOT NULL,
    `name` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `password` char(60) NOT NULL,
    `role` varchar(36) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `projects` (
    `id` varchar(36) NOT NULL,
    `account_id` varchar(36) NOT NULL,
    `name` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `forms` (
    `id` varchar(36) NOT NULL,
    `project_id` varchar(36) NOT NULL,
    `name` varchar(255) NOT NULL,
    `properties` JSON NOT NULL,
    `button` JSON NOT NULL,
    `modules` JSON NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO accounts (id, name)
VALUES ('632f251f-58a5-45c0-b7da-ff7538e784ad', 'Test Account');

INSERT INTO users (id, account_id, name, email, password, role)
VALUES ('2c617a61-da45-4f78-aef8-05ea4a9ba59e', '632f251f-58a5-45c0-b7da-ff7538e784ad', 'Test User', 'test@getestimator.com', '$2a$10$rPcPIt0iWkbdBFfYAOQbPO6jQXbImovyXWhBEnn9QSOYpuHbPY7M2', 'admin');