 #!/bin/bash
  
split -d -l 2500  all_it_IT_finance.txt it_IT_finance

#split -d -b 455K all_it_IT_finance.txt it_IT_finance

 sed 1d allit_IT_finance.txt >it_IT_finance09
 
 rm allit_IT_finance.txt
 
