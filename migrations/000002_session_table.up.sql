create table session(
                        session_id    int generated always as identity primary key,
                        employee_id   bigint references employee (employee_id) not null unique,
                        refresh_token text                                  not null unique,
                        expires_at    timestamp                                  not null
);