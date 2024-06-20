create table employee_subscription(
                                      target_id     bigint references employee (employee_id),
                                      subscriber_id bigint references employee (employee_id)
);