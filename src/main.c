#include <stdio.h>
#include <unistd.h>

int main(void) {
  while (1) {
    printf("Hello from raptor digital twin!\n");
    fflush(stdout);
    sleep(1);
  }
}

