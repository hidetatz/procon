# Write your MySQL query statement below
SELECT apple.sale_date, (apple.sold_num - orange.sold_num) as diff from 
(
    select sale_date, fruit, sold_num from sales where fruit = 'apples'
) apple
left join
(
    select sale_date, fruit, sold_num from sales where fruit = 'oranges'
) orange
on apple.sale_date = orange.sale_date
