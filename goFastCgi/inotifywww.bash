#!/bin/bash                               
 
inotifywait -m -q -e CLOSE_NOWRITE,CLOSE -r  "www/" --format "-htmlfile=%w%f -event=%e" --excludei '\.(jpg|png|gif|ico|log|sql|pdf|php|swf|ttf|eot|woff|)$' |
	while read file
	do

#		echo $file
		bin/goinotify $file 
#		if [[ $file =~ \.(gz)$ ]];
#		then
#			#gzip -f -c -1 $file > $file.gz
#			echo $file
#			bin/goinotify -htmlfile=$file
#		fi
	done
