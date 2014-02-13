sqlite3 singo.db 'delete from keywords where locale="it_IT" and Themes="finance"'
sqlite3 singo.db 'delete from phrases where locale="it_IT" and Themes="finance"'
sqlite3 singo.db 'delete from keywords where locale="fi_FI" and Themes="finance"'
sqlite3 singo.db 'delete from phrases where locale="fi_FI" and Themes="finance"'
sqlite3 singo.db 'delete from keywords where locale="fi_FI" and Themes="porno"'
sqlite3 singo.db 'delete from phrases where locale="fi_FI" and Themes="porno"'
sqlite3 singo.db 'delete from keywords where locale="fi_FI" and Themes="fortune"'
sqlite3 singo.db 'delete from phrases where locale="fi_FI" and Themes="fortune"'


rm -fr markresources/fi_FI_finance/
rm -fr markresources/fi_FI_fortune/
rm -fr markresources/it_IT_finance/

bin/syncpushdomaindb 


