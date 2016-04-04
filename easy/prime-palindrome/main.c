#include <stdio.h>

int main(int argc, const char * argv[]) {
  
  int sum, number,aux;
 
  for(number=999;sum==0;number--){
   
    sum=0;
    aux=number/2;
    while(aux>0 && number%aux!=0)aux--;
    
    if(aux==1){
      int nsplit[2];
      if(number>=100 && number%10==((number-(number%100))/100))
	  sum=number;
      
      else  if (number>=10 && number%10==((number-(number%10))/10))
	sum=number;
      else if(number<10)
	sum=number;
      
    }
  }
  printf("%d\n",sum);
  return 0;  
} 
