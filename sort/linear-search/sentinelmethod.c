/*
*/
#include <stdio.h>
#include <stdlib.h> 

int search(int a[], int n, int key){
  int i = 0;
  a[n] = key;
  while (1) {
    if (a[i] == key)
        break;
    i++;
  }
  return i == n ? -1 : i;
}


int main(void){
  int idx;
  
  int num[] = {6, 4, 3, 1, 2};
  idx = search(num, 5, 1);

  if (idx == -1){
    puts("fail");
  }else{
    printf("%d=>success\n", idx);
  }
  return 0;
}
