drop view IF EXISTS mission_users;
CREATE VIEW mission_users as select 
	(select SUBSTRING(m.start,7,4)||'-'||SUBSTRING(m.start,4,2)||'-'||SUBSTRING(m.start,1,2) <= strftime('%Y-%m-%d','now', 'localtime')
	and SUBSTRING(m.end,7,4)||'-'||SUBSTRING(m.end,4,2)||'-'||SUBSTRING(m.end,1,2) >= strftime('%Y-%m-%d','now', 'localtime')
	) as in_progress,
	(select SUBSTRING(m.start,7,4)||'-'||SUBSTRING(m.start,4,2)||'-'||SUBSTRING(m.start,1,2) > strftime('%Y-%m-%d','now', 'localtime')) as in_future,
	(select SUBSTRING(m.end,7,4)||'-'||SUBSTRING(m.end,4,2)||'-'||SUBSTRING(m.end,1,2) < strftime('%Y-%m-%d','now', 'localtime')) as in_past,
	SUBSTRING(m.start,7,4)||'-'||SUBSTRING(m.start,4,2)||'-'||SUBSTRING(m.start,1,2) as start_ordered,
	SUBSTRING(m.end,7,4)||'-'||SUBSTRING(m.end,4,2)||'-'||SUBSTRING(m.end,1,2) as end_ordered,
	ifnull(ROUND(JULIANDAY(SUBSTRING(m.start,7,4)||'-'||SUBSTRING(m.start,4,2)||'-'||SUBSTRING(m.start,1,2)) - JULIANDAY('now')+ 0.5),9999) as before_start_days,
	ifnull(ROUND(JULIANDAY(SUBSTRING(m.end,7,4)||'-'||SUBSTRING(m.end,4,2)||'-'||SUBSTRING(m.end,1,2)) - JULIANDAY('now')+ 0.5),9999) as before_end_days,
	ifnull((select us2.value from user_states us2 join state_key sk2 on sk2."key" = us2."key" and sk2.is_intro = 1 where us2.user_id = t.id),'') as ident,
	t.username as name, 
	ifnull((select group_concat('['||us."key"||']', " ") from user_states us where us.user_id = t.id
and SUBSTRING(us.value,7,4)||'-'||SUBSTRING(us.value,4,2)||'-'||SUBSTRING(us.value,1,2) >= SUBSTRING(m.start,7,4)||'-'||SUBSTRING(m.start,4,2)||'-'||SUBSTRING(m.start,1,2)
and SUBSTRING(us.value,7,4)||'-'||SUBSTRING(us.value,4,2)||'-'||SUBSTRING(us.value,1,2) <= SUBSTRING(m.end,7,4)||'-'||SUBSTRING(m.end,4,2)||'-'||SUBSTRING(m.end,1,2)
GROUP by us.user_id),'') as conflicts,
	t.id as user_id,
	m.id as mission_id,
	m.place as place,
	m.start as start,
	m.end as end,
	m.active as active,
	t.masters as masters_list 
from telebotusers t 
join missions m on m.recepient_id = t.id
where JulianDay(SUBSTRING(m.start,7,4)||'-'||SUBSTRING(m.start,4,2)||'-'||SUBSTRING(m.start,1,2)) not null
	  and JulianDay(SUBSTRING(m.end,7,4)||'-'||SUBSTRING(m.end,4,2)||'-'||SUBSTRING(m.end,1,2)) not null
order by SUBSTRING(m.start,7,4)||'-'||SUBSTRING(m.start,4,2)||'-'||SUBSTRING(m.start,1,2), SUBSTRING(m.end,7,4)||'-'||SUBSTRING(m.end,4,2)||'-'||SUBSTRING(m.end,1,2)
;
