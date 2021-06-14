#include <stdio.h>

#define SIZE 4

void show(int []);  // (int*) でもOK
void swap(int*, int*);


int a[] = {4, 1, 3, 2};

int main(){
    int i, j;
    for (i = 0; i < SIZE; i++){
        for (j = 0; j < SIZE - 1; j++){
            if (a[j] > a[j + 1]){
                swap(&a[j], &a[j+1]);
            }
        }
    }
    show(a);
    return 0;
}

void show(int* array){ //array []intはダメ
    int i;
    for (i=0;i<SIZE;i++){
        printf("%d ", array[i]);
    }
    printf("\n");
}

void swap(int* n1, int* n2){
    int tmp = *n1;
    *n1 = *n2;
    *n2 = tmp;
}
