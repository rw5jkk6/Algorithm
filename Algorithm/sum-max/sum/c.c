/*
*/
#include <stdio.h>


int n1 = 1000000;

int sum(int n[n1]){
	int s = 0;
	int i;
	for (i=0; i < n1; i++){
		s+=n[i];
	}
	return s;
}


int main(void){
	int n[n1];
	int i, num;
	for (i = 0; i < n1; i++){
		n[i] = i;
	}
	
	num = sum(n);
	printf("%d\n", num);
}
