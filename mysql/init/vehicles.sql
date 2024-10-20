CREATE TABLE IF NOT EXISTS `vehicles` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `license_plate` varchar(20) UNIQUE NOT NULL,
    `type` varchar(100) DEFAULT NULL,
    `make` varchar(100) DEFAULT NULL,
    `model` varchar(100) DEFAULT NULL,
    `year` int(4) DEFAULT NULL,
    `capacity_kg` int(10) DEFAULT NULL,
    `driver_id` int(10) DEFAULT NULL,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`driver_id`) REFERENCES drivers(`id`)
);

INSERT INTO `vehicles` (
    `license_plate`, 
    `type`, 
    `make`, 
    `model`, 
    `year`, 
    `capacity_kg`
) VALUES 
('YW0SNYPS3E','UWCDMQLCEU','lpmlxkawmq','xxunn',7073,68974),    
('PDDEJ4CR88','NTQMLTPZSR','ccsneweajb','irwkr',1380,27355),
('HBT0UQQJT6','QJDKCWOHIH','slnnmsdlrm','qjjkm',7566,59646),
('ZI2ICSA62H','BNNAZASXDO','naitmuplpw','ltiid',361,62149),
('QM5YGZYZ8F','DTVDKFVVSJ','nukcvizjeh','ggdyx',28,50089),
('2LI03BZE2E','MESZDOLEZV','wibxmbemqu','uibsz',543,8616),
('IAV6V9GI2C','RJUYNMBYID','wvrnmkfmym','hetsl',5855,63452),
('WSMYMHDR0J','QJRCBBCBTS','ubuimeagdn','mhavw',9405,58382),
('1I878U80XK','KIODVMFIUA','ocruupvixe','ahhiu',7801,32868),
('0E6Q9U9XFG','XCSCGFPLUL','tpwcnufqeg','bouey',3154,56370);