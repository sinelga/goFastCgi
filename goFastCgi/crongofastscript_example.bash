#! /bin/bash
#chmod +x!! must be done

cd /home/juno/git/goFastCgi/goFastCgi
pgrep elabqueue && echo `date >>/tmp/elabqueue`
bin/paragraphshandler -locale=fi_FI -themes=finance -quant=500
pgrep elabqueue || bin/elabqueue
