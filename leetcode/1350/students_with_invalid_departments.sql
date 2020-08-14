# Write your MySQL query statement below
SELECT 
  s.id, s.name 
from 
  Students s 
LEFT JOIN 
  Departments d 
ON 
  s.department_id = d.id
WHERE
  d.name is null;
