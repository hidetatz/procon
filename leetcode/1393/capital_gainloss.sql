# Write your MySQL query statement below
select stock_name, sum(case when operation = 'Buy' then -price else price end) capital_gain_loss from stocks group by stock_name;
