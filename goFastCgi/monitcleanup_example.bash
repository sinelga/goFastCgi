cd /home/juno/git/goFastCgi/goFastCgi

pgrep elabqueue || pgrep orphans || pgrep  cleanupspace || bin/cleanupspace
pgrep elabqueue || pgrep cleanupspace || pgrep orphans || bin/orphans -updated=800
/usr/bin/find www -type d -empty -delete

