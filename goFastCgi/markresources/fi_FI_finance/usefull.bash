 #!/bin/bash
  
split -d -l 10000   all_fi_FI_finance.txt fi_FI_finance

split -d -b 5000K all_fi_FI_finance.txt fi_FI_finance
 
