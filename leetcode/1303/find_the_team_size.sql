# Write your MySQL query statement below
select employee_id, (select count(team_id) from employee where e.team_id = team_id) as team_size from employee e;
