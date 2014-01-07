apt-get autoremove
apt-get clean
apt-get autoclean

 echo 524288 > /proc/sys/fs/inotify/max_user_watches
  vi /etc/sysctl.conf
  fs.inotify.max_user_watches=524288
  sysctl -p


bin/newdomain -locale=fi_FI -themes=porno -domain=test.com -expire=600
bin/cleanupspace -hits=20 -created=20

git gc --aggressive --prune=now

select count() from paragraphs where Siteid is null;

cp  singo.db  singo.db.last
rm singo.db-journal singo.db
echo .dump | sqlite3 singo.db >backup.sql

## Backup
echo '.dump' | sqlite3 gofast.db | gzip -c >gofast.dump.gz
rm gofast.db
zcat gofast.dump.gz | sqlite3 gofast.db
##


create index pridind on sentences (Prid);
create index siteind on paragraphs (Siteid);

PRAGMA writable_schema=ON;
to alter table 
Then do a manual UPDATE of the sqlite_master table to insert
an "id INTEGER PRIMARY KEY" into the SQL for the table definition.

Auto-VACCUM
SQLite Auto-VACUUM does not do the same as VACUUM rather it only moves free pages to the end of the database thereby reducing the database size. By doing so it can significantly fragment the database while VACUUM ensures defragmentation. So Auto-VACUUM just keeps the database small.

You can enable/disable SQLite auto-vacuuming by the following pragmas running at SQLite prompt:

sqlite> PRAGMA auto_vacuum = NONE;  -- 0 means disable auto vacuum
sqlite> PRAGMA auto_vacuum = INCREMENTAL;  -- 1 means enable incremental vacuum
sqlite> PRAGMA auto_vacuum = FULL;  -- 2 means enable full auto vacuum



sqlite3 singo.db 'delete from keywords where locale="it_IT" and Themes="finance"'
sqlite3 singo.db 'insert into keywords (Locale,Themes,Keyword) select Locale,Themes,Keyword from it_IT_finance_keywords where Block=0'
sqlite3 singo.db 'delete from phrases where locale="it_IT" and Themes="finance"'
sqlite3 singo.db 'insert into phrases (Phrase, Locale,Themes) select distinct Keyword as Phrase,"it_IT" as Locale,"finance" as Themes from it_IT_finance_phrases'


sqlite3 singo.db 'delete from keywords where locale="fi_FI" and Themes="fortune"'
sqlite3 singo.db 'insert into keywords (Locale,Themes,Keyword) select Locale,Themes,Keyword from fi_FI_fortune_keywords where Block=0'
sqlite3 singo.db 'delete from phrases where locale="fi_FI" and Themes="fortune"'
sqlite3 singo.db 'insert into phrases (Phrase, Locale,Themes) select distinct Keyword as Phrase,"fi_FI" as Locale,"fortune" as Themes from fi_FI_fortune_phrases'


!! delete ," space rom all_fi_FI_phrases !!

sqlite3 singo.db 'delete from keywords where locale="fi_FI" and Themes="porno"'
sqlite3 singo.db 'delete from phrases where locale="fi_FI" and Themes="porno"'



sqlite3 -init backup.sql singo.db

GOBIN=./bin GOPATH=$PWD:/home/juno/junoworkspace/gows go install -a src/app.go
