-- tx1

begin transaction;
lock table accounts in access share mode;
select * from accounts order by balance limit 2;
commit transaction;

-- tx2

begin transaction;
insert into accounts (balance) values (10.0);
commit transaction;

-- tx3

begin transaction;
lock table accounts in access share mode;
select * from accounts order by balance limit 2;
commit transaction;