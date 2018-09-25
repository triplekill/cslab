#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>

int getpid_and_ppid(){
    printf("parent pid:%d pid:%d\n", getpid(), getppid());
    return 0;
}

int main()
{
    getpid_and_ppid();
}
