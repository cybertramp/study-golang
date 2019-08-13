/*

    This program run with test_cWithGo.go

 */

#include <stdio.h>
#include "test_cWithGo.h"

int main(int argc, char **argv){
    GoInt x = 12;
    GoInt y = 23;

    printf("About to call a Go function!\n");
    PrintMessage();

    GoInt p = Multiply(x,y);
    printf("Product: %d\n",(int)p);
    printf("It worked!\n");
    return 0;
}