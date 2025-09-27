drop view IF EXISTS master_users;
CREATE VIEW master_users as select 
  t.id as id,
  t.is_admin as is_admin,
  ifnull((select us5.value from user_states us5 join state_key sk5 on sk5."key" = us5."key" and sk5.is_intro = 1 where us5.user_id = t.id),'') as ident,
  IFNULL((select group_concat(t2.id) from telebotusers t2 where instr(t2.masters, (select us2.value from user_states us2 join state_key sk2 on sk2."key" = us2."key" and sk2.is_intro = 1 where us2.user_id = t.id)) > 0), '') as ids_subordinate,
  IFNULL((select group_concat((select us4.value from user_states us4 join state_key sk4 on sk4."key" = us4."key" and sk4.is_intro = 1 where us4.user_id = t3.id)) from telebotusers t3 where instr(t3.masters, (select us3.value from user_states us3 join state_key sk3 on sk3."key" = us3."key" and sk3.is_intro = 1 where us3.user_id = t.id)) > 0), '') as idents_subordinate
from telebotusers t 
where
	t.is_admin > 0
	OR (select count(*) from telebotusers t2 where instr(t2.masters, (select us2.value from user_states us2 join state_key sk2 on sk2."key" = us2."key" and sk2.is_intro = 1 where us2.user_id = t.id)) > 0) > 0
;
