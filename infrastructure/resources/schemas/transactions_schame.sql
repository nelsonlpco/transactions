CREATE TABLE account(
    id INTEGER NOT NULL PRIMARY KEY
    document_number VARCHAR(100) NOT NULL
);

CREATE TABLE operation_type(
    id INTEGER NOT NULL PRIMARY KEY,
    description VARCHAR(100) NOT NULL,
    operation ENUM('debito', 'credito')
);

CREATE TABLE transaction(
    id INTEGER NOT NULL PRIMARY KEY,
    account_id INTEGER NOT NULL,
    operation_type_id INTEGER NOT NULL,
    ammount DOUBLE,
    event_date date NOT NULL,
    CONSTRAINT transaction_account_fk FOREIGN KEY (account_id) REFERENCES account(id),
    CONSTRAINT transaction_operation_type_fk FOREIGN KEY (operation_type_id) REFERENCES operation_type(id)
);