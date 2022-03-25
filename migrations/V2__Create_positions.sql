CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

do $$
begin
for i in 1..100000 loop
        insert into positions (id,account_id,order_id,open_price,close_price,take_profit,stop_loss,symbol,guaranteed_stop_loss,state,quantity, leverage,side)
        values (uuid_generate_v4(),'1234','',0,0,28,0,'silver',false,'OPEN',0.1,false,'BUY');
end loop;
end $$