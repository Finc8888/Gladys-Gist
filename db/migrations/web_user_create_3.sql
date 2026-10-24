DROP USER IF EXISTS 'web'@'%';
CREATE USER IF NOT EXISTS 'gladys'@'%';
GRANT SELECT, INSERT, UPDATE, DELETE ON gistbox.* TO 'gladys'@'%';
-- Important: Make sure to swap 'pass' with a password of your own choosing
ALTER USER 'gladys'@'%' IDENTIFIED BY 'web';
