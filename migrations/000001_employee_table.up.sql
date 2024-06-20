create table employee(
                         employee_id   bigint primary key,
                         full_name     varchar(32) not null,
                         birth_date    date        not null,
                         email         varchar(32) not null unique,
                         hash_password text        not null
);