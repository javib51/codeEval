#include <stdio.h>
#include <stdlib.h>

int main(int argc, const char * argv[]) {
  
  FILE* file=fopen(argv[1],"r");
  char* line=NULL,sentence[100],characters[100],result[100];
  size_t bufsize = 41;
  int i,j,m;
  
  while(getline(&line,&bufsize,file)>0){
    m=0;
    j=0;
    memset(result, 0, 20);
    memset(sentence, 0, 20);
    memset(characters, 0, 20);
    for(i=0;line[i]!='\0';i++){
      if(line[i]=='\n')
	continue;
     
      if(line[i]==','){
	m=1;
	sentence[j]='\0';
	j=0;
	continue;
      }
      if(m==0){
	sentence[j++]=line[i];
      }else
	characters[j++]=line[i];	
    }
    characters[j]='\0';

    
   
    m=-1;
    for(i=0;sentence[i]!='\0';i++){
      j=0;
     
      while(sentence[i]!=characters[j]&&characters[j]!='\0')
	j++;
	
      if(sentence[i]!=characters[j] || sentence[i]==' ')
	result[++m]=sentence[i];
    }
    
    //result[m]='\0';
    printf("%s\n",result);
  }
    
  return 0;  
} 
