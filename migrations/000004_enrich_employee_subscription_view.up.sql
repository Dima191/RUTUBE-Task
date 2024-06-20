create view employee_subscriptions_by_name as
select target_id,
       target_employee.full_name     as target_full_name,
       target_employee.email         as target_email,
       target_employee.birth_date    as target_birth_date,
       subscriber_id,
       subscriber.full_name as subscriber_full_name,
       subscriber.email as subscriber_email
from employee_subscription
         inner join employee as target_employee on employee_subscription.target_id = target_employee.employee_id
         inner join employee as subscriber
                    on employee_subscription.subscriber_id = subscriber.employee_id;