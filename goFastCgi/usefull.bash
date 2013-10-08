
select count() from paragraphs where Siteid is null;

cp  singo.db  singo.db.last
rm singo.db-journal singo.db
echo .dump | sqlite3 singo.db >backup.sql

## Backup
echo '.dump' | sqlite3 singo.db | gzip -c >singo.dump.gz

sqlite3 singo.db 'select * from keywords where locale="it_IT" and Themes="finance"' 

sqlite3 singo.db 'select Locale,Themes,Keyword from it_IT_finance_keywords where Block=0'

sqlite3 singo.db 'delete from keywords where locale="it_IT" and Themes="finance"'

sqlite3 singo.db 'insert into keywords ( Locale,Themes,Keyword) select Locale,Themes,Keyword from it_IT_finance_keywords where Block=0'

sqlite3 singo.db 'delete from phrases where locale="it_IT" and Themes="finance"'
sqlite3 singo.db 'insert into phrases (Phrase, Locale,Themes) select distinct Keyword as Phrase,"it_IT" as Locale,"finance" as Themes from it_IT_finance_phrases'



sqlite3 -init backup.sql singo.db

GOBIN=./bin GOPATH=$PWD:/home/juno/junoworkspace/gows go install -a src/app.go
