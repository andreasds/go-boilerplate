DROP TABLE IF EXISTS `foo_items`;
DROP TABLE IF EXISTS `foo`;

CREATE TABLE IF NOT EXISTS `foo` (
    id BINARY(16) NOT NULL,
    receipt_name VARCHAR(255) NOT NULL,
    receipt_notes TEXT NULL,
    total_price DECIMAL(14,2) NOT NULL,
    is_active TINYINT(1) NOT NULL DEFAULT 1,
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by BINARY(16) NOT NULL,
    updated_time TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    updated_by BINARY(16) NULL DEFAULT NULL,
    removed_time TIMESTAMP NULL DEFAULT NULL,
    removed_by BINARY(16) NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX `idx_search_name` (`receipt_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `foo_item` (
    id BINARY(16) NOT NULL,
    foo_id BINARY(16) NOT NULL,
    item_name VARCHAR(255) NOT NULL,
    item_notes TEXT NULL,
    item_price DECIMAL(14,2) NOT NULL,
    is_active TINYINT(1) NOT NULL DEFAULT 1,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_foo_item_foo_id` FOREIGN KEY (`foo_id`)
        REFERENCES `foo` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
    INDEX `idx_search_item_name` (`item_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `foo`
    (`id`, `receipt_name`, `receipt_notes`, `total_price`, `is_active`,
    `created_time`, `created_by`, `updated_time`, `updated_by`, `removed_time`,
    `removed_by`)
VALUES
    (
        UUID_TO_BIN('f5610c25-91ed-4c18-9beb-4272a2e7ca90'),
        'Traditional Market',
        'monthly needs',
        37000,
        1,
        NOW(),
        UUID_TO_BIN('aa7141f2-67a1-4a5a-bce1-441b6716bddf'),
        NULL,
        NULL,
        NULL,
        NULL
    ),
    (
        UUID_TO_BIN('5b1272ab-8bdc-4288-9d43-d7508adf5eef'),
        'Carrefour',
        'buy snacks',
        10000,
        1,
        DATE_SUB(NOW(), INTERVAL 2 DAY),
        UUID_TO_BIN('aa7141f2-67a1-4a5a-bce1-441b6716bddf'),
        DATE_SUB(NOW(), INTERVAL 2 HOUR),
        UUID_TO_BIN('9f8d340b-3341-4e83-b23e-cc1059007f98'),
        NULL,
        NULL
    );

INSERT INTO `foo_item`
    (`id`, `foo_id`, `item_name`, `item_notes`, `item_price`, `is_active`)
VALUES
    (
        UUID_TO_BIN('bb13fb47-5ab1-4b3b-ae0b-f4d5e880b56b'),
        UUID_TO_BIN('f5610c25-91ed-4c18-9beb-4272a2e7ca90'),
        'Vegetable',
        'Carrot',
        25000,
        1
    ),
    (
        UUID_TO_BIN('9d12396b-3e19-49ce-a6db-6ca6395b5ed3'),
        UUID_TO_BIN('f5610c25-91ed-4c18-9beb-4272a2e7ca90'),
        'Fish',
        'Tuna',
        12000,
        1
    ),
    (
        UUID_TO_BIN('a07f7a99-4b8d-458c-9de2-139106245b2b'),
        UUID_TO_BIN('5b1272ab-8bdc-4288-9d43-d7508adf5eef'),
        'Cheetos',
        '',
        10000,
        1
    );
