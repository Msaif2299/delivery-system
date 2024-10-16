CREATE DATABASE IF NOT EXISTS deliveries;
CREATE TABLE IF NOT EXISTS `drivers` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `full_name` varchar(255) DEFAULT NULL,
    `license_number` varchar(20) DEFAULT NULL,
    `primary_phone_number` varchar(20) DEFAULT NULL,
    `primary_phone_country_code` varchar(5) DEFAULT NULL,
    `secondary_phone_number` varchar(20) DEFAULT NULL,
    `secondary_phone_country_code` varchar(5) DEFAULT NULL,
    `email` varchar(255) DEFAULT NULL,
    `status` tinyint(10) DEFAULT NULL,
    PRIMARY KEY(`id`)
);

-- INSERT INTO `drivers` (
--     full_name,
--     license_number,
--     primary_phone_country_code,
--     primary_phone_number,
--     secondary_phone_country_code,
--     secondary_phone_number,
--     email,
--     status
-- ) VALUES (
--     'Cardinal',
--     'JOISJHG980',
--     '+44',
--     '654789132',
--     '+91',
--     '885522256',
--     'test@test.com',
--     5
-- );