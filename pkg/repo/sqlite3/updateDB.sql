CREATE TABLE if not exists examen_ended (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL DEFAULT (0),
  key TEXT NOT NULL DEFAULT '',
  date TEXT NOT NULL DEFAULT '',
  created TEXT DEFAULT (strftime('%Y-%m-%dT%H:%M','now', 'localtime'))
);
