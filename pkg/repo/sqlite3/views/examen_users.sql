drop view IF EXISTS examen_users;
CREATE VIEW examen_users as select 
  ifnull(ROUND(JULIANDAY(SUBSTRING(us.value,7,4)||'-'||SUBSTRING(us.value,4,2)||'-'||SUBSTRING(us.value,1,2)) - JULIANDAY('now')+ 0.5),9999) as before_date_days,
  ifnull((select us2.value from user_states us2 join state_key sk2 on sk2."key" = us2."key" and sk2.is_intro = 1 where us2.user_id = t.id),'') as ident,
  ifnull((select SUBSTRING(us2.value,7,4)||'-'||SUBSTRING(us2.value,4,2)||'-'||SUBSTRING(us2.value,1,2) from user_states us2 where us2.user_id = keys.id and us2.key=keys.ex),'') as date_ordered,
  t.username as name, 
  t.id as id, 
  keys.ex as examen, 
  ifnull(us.value,'') as date, 
  t.masters as masters_list 
from (select t.id as id,  sk.key as ex from telebotusers t inner join state_key sk on sk.is_examen = 1) keys
  left join telebotusers t on t.id = keys.id
  left join user_states us on us.user_id = keys.id and us.key = keys.ex
order by 1
