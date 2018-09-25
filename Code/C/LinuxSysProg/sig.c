#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>
#include <unistd.h>
#include <signal.h>
#include <assert.h>
#include <db.h>

static void singint_handler(int signo)
{
    printf("Caught SIGINT\n");
    exit(EXIT_SUCCESS);
}
// 可以实现优雅退出
int siglearn()
{
    if(signal(SIGINT, singint_handler)==SIG_ERR){
        fprintf(stderr, "Cannot handle SIGINT\n");
        exit(EXIT_SUCCESS);
    }
    raise(SIGINT);// 进程给自己发送信号 等价于 kill(getpid(), signo)
    return 0;
}

int main()
{
    siglearn();
}
