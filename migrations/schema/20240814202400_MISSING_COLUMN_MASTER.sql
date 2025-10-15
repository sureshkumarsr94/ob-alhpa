-- +goose Up

ALTER TABLE `ent_warehouse`
    ADD `distributor_id` VARCHAR(64) NULL  AFTER `description`,
    ADD FOREIGN KEY `fk_ent_warehouse_ent_distributor_idx`(`distributor_id`)
    REFERENCES `ent_distributor`(`id`) ON
DELETE
NO ACTION ON UPDATE NO ACTION;

ALTER TABLE `core_role` CHANGE
    `distributor_id` `distributor_id` VARCHAR (64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL;

ALTER TABLE `core_user`
    ADD `last_password_change` DATETIME NULL AFTER `associate_user_id`;

ALTER TABLE `sale_order_item` CHANGE `free_product_id` `free_product_id` VARCHAR (64) NULL;
ALTER TABLE `sale_order_item` CHANGE `discount_amount` `discount_amount` DOUBLE NULL;
