#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

int main(int argc, const char * argv[]) {
  
  FILE* file=fopen(argv[1],"r");
  char* line=NULL,word[20],mask[20];
  size_t bufsize = 41;
  int i,j,m;
  
  while(getline(&line,&bufsize,file)>0){
    m=0;
    j=0;
    for(i=0;line[i]!='\0';i++){
      if(line[i]=='\n')
	continue;
     
      if(line[i]==' '){
	m=1;
	word[j]='\0';
	j=0;
	continue;
      }
      if(m==0){
	word[j++]=line[i];
      }else
	mask[j++]=line[i];	
    }
     
    for(i=0;word[i]!='\0';i++){
      if(mask[i]=='1')
	word[i]=toupper(word[i]);
      else
	word[i]=tolower(word[i]);

    }
    printf("%s\n",word);
    
  }
  return 0;  
} 
