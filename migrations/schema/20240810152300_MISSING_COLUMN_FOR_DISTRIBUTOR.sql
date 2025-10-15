-- +goose Up

ALTER TABLE `core_user` ADD `associate_user_id` VARCHAR(64) NULL AFTER `code`;

ALTER TABLE `ent_distributor_user`
    ADD `role_id` VARCHAR(64) NULL  AFTER `user_id`,
    ADD FOREIGN KEY `fk_ent_distributor_user_core_role1_idx`(`role_id`)
    REFERENCES `core_role`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

