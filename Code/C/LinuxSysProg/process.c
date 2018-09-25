#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>

int getpid_and_ppid()
{
    printf("parent pid:%d pid:%d\n", getpid(), getppid());
    return 0;
}

int execlearn()
{
    // 使用vi替换当前进程，有t.txt参数，所以打开此文件, 第二个参数代表启动程序名字
    int ret = execl("/usr/local/bin/vi", "vi", "t.txt", NULL); //mac
    if (ret == -1)
    {
        ret = execl("/usr/bin/vi", "vi", "t.txt", NULL); // linux
    }
    if (ret == -1)
    {
        perror("execl");
        return 1;
    }
    return 0;
}
void printatexitf(void)
{
    printf("atexit success at pid:%d\n",getpid());
}
int forklearn()
{
    atexit(printatexitf);
    pid_t pid = fork();
    if (pid > 0)
    {
        printf("cur:%d, parent id:%d\n", getpid(), pid);
        exit(EXIT_FAILURE);
    }
    else if (!pid)
    {
        printf("cur:%d, id:%d\n", getpid(), pid);
    }
    else if (pid == -1)
    {
        perror("fork");
    }
    return 0;
}

int main()
{
    // getpid_and_ppid();
    // execlearn();
    forklearn();
}
