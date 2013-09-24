

cp  singo.db  singo.db.last
rm singo.db-journal singo.db
echo .dump | sqlite3 singo.db >backup.sql

sqlite3 -init backup.sql singo.db


GOBIN=./bin GOPATH=$PWD:/home/juno/junoworkspace/gows go install -a src/app.go