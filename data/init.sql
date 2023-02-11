USE payment;

DROP TABLE IF EXISTS debit_transactions;
DROP TABLE IF EXISTS payment_transactions;
DROP TABLE IF EXISTS accounting_reports;

CREATE TABLE debit_transactions (
    id int NOT NULL AUTO_INCREMENT,
    amount DECIMAL(13, 2) NOT NULL,
    bank_id varchar(255) NOT NULL,
    process_id varchar(255) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

CREATE TABLE payment_transactions (
    id int NOT NULL AUTO_INCREMENT,
    amount DECIMAL(13, 2) NOT NULL,
    bank_id varchar(255) NOT NULL,
    process_id varchar(255) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

CREATE TABLE accounting_reports (
    id int NOT NULL AUTO_INCREMENT,
    bank_id varchar(255) NOT NULL,
    report_type varchar(255) NOT NULL,
    total DECIMAL(13, 2) NOT NULL,
    PRIMARY KEY (id)
);
