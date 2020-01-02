#include <stdio.h>
#include "libsample.h"

int main() {
    void * sample = NewSampleC("hello world");
    PrintSampleC(sample);
    CloseSampleC(sample);
    printf("End of program\n");
}