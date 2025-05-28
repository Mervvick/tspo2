-- scripts/init-db.sql
CREATE DATABASE digital_market;
CREATE USER app WITH ENCRYPTED PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE digital_market TO app;
