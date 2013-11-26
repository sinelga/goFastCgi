#! /bin/bash

DISKSPACE=`df -H /dev/vda | sed '1d' | awk '{print $5}' | cut -d'%' -f1`

# Disk capacity threshold
ALERT=70
if [ ${DISKSPACE} -ge ${ALERT} ]; then
    echo "Still enough disk space....${DISKSPACE}% capacity."
    exit
fi