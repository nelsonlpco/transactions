CREATE TABLE IF NOT EXISTS account(
    id binary(16) NOT NULL PRIMARY KEY,
    document_number VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS operation_type(
    id binary(16) NOT NULL PRIMARY KEY, 
    description VARCHAR(100) NOT NULL UNIQUE, 
    operation INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions(
    id binary(16) NOT NULL PRIMARY KEY,
    account_id binary(16) NOT NULL,
    operation_type_id binary(16) NOT NULL,
    amount DOUBLE NOT NULL,
    event_date date NOT NULL,
    CONSTRAINT transaction_account_fk FOREIGN KEY (account_id) REFERENCES account(id),
    CONSTRAINT transaction_operation_type_fk FOREIGN KEY (operation_type_id) REFERENCES operation_type(id)
);
