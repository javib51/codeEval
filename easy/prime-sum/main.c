#include <stdio.h>

int main(int argc, const char * argv[]) {
  
  int sum, number,aux, result;
  result=0;

  for(number=1;sum<1000;number++){
   
    aux=number-1;
    while(aux>0 && number%aux!=0)aux--;
    
    if(aux==1){
      result+=number;
      sum++;
    }
  }
  printf("%d\n",result);
  return 0;  
} 
