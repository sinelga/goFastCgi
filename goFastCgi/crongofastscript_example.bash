#! /bin/bash
#chmod +x!! must be done
#*/5 * * * * /home/juno/git/goFastCgi/goFastCgi/crongofastscript.bash

cd /home/juno/git/goFastCgi/goFastCgi
#pgrep elabqueue && echo `date >>/tmp/elabqueue`
#bin/paragraphshandler -locale=fi_FI -themes=finance -quant=500
pgrep elabqueue || bin/elabqueue

pgrep elabqueue || bin/cleanupspace -hits=1 -created=500

#bin/newdomain -locale=fi_FI -themes=porno -domain=tissit.tv -expire=3600