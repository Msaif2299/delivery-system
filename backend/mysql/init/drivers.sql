CREATE DATABASE IF NOT EXISTS deliveries;
CREATE TABLE IF NOT EXISTS `drivers` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `full_name` varchar(255) DEFAULT NULL,
    `license_number` varchar(20) UNIQUE NOT NULL,
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

INSERT INTO `drivers` (
    full_name, 
    license_number, 
    primary_phone_number, 
    primary_phone_country_code, 
    secondary_phone_number, 
    secondary_phone_country_code, 
    email, 
    status
) VALUES 
('Bbxndxqxw Qnyclf','0DXAMBL3ZB','4529235181','+111','1181291757','+023','dq35k@kyaex.fto',0),
('Dmcdznjoe Uaqlaz','Q3964VMCSM','1233569643','+492','2449871718','+389','bphcn@zepmh.gmh',4),
('Ksrtvvgit Mzsgdk','TJ0J0U0V9H','7670867264','+005','4523997619','+929','wuvqb@arpks.ohq',0),
('Eqpmacmrz Phpujn','172UJ5K4Y8','2915625663','+566','9978348456','+664','dgvxa@hqhic.sgu',8),
('Ezymdoefk Bhhwhr','KNLSGPVWR4','0352444757','+724','9278892731','+823','rw54l@lknei.ksi',2),
('Jpetmyzal Ofotfe','GA633LU9XP','4787282275','+245','1581629841','+351','a0tqn@oxqpb.fjv',9),
('Lsztyovtu Bzukmt','L97Y4ZWOLV','3339826917','+264','6353125448','+737','is27h@tpsxf.san',6),
('Gsozlwntx Ijdzyh','5VGWYKWB53','9988201203','+745','5106794769','+088','jkq64@xhpsv.pma',2),
('Hhsjmydqx Huihul','AAJTVL70QR','7182482013','+926','0801660556','+989','xggi0@fvvls.xrv',7),
('Gdlhehmpi Lfhaus','3KNUB10SUU','8143879473','+322','2761819443','+874','umyep@rkjvm.yqq',4);