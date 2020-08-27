# Write your MySQL query statement below
select u.unique_id, e.name from employees e left join employeeuni u on e.id = u.id;
