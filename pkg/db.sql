CREATE TABLE tbl_user (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    create_ts timestamp without time zone default current_timestamp,
    update_ts timestamp without time zone default current_timestamp
);

INSERT INTO tbl_user(username, password) VALUES('admin', '$2a$12$ibhAdTkqdPgBvM2nyiLlp.W0sMWW9BdhL5jil1p4qE6PUXG2imWD.');
