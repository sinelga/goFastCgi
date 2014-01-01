#!/bin/bash                               
 
inotifywait -m -q -e OPEN  -r "www/" --format "%w%f" --excludei '\.(jpg|png|gif|ico|log|sql|pdf|php|swf|ttf|eot|woff|)$' |
	while read file
	do
		if [[ $file =~ \.(gz)$ ]];
		then
			#gzip -f -c -1 $file > $file.gz
			echo $file
		fi
	done
