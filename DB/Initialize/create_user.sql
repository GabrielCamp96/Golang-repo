CREATE USER IF NOT EXISTS 'golang'@'localhost' IDENTIFIED BY 'golang';

GRANT ALL PRIVILEGES on devbook.* TO 'golang'@'localhost';