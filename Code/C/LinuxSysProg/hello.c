#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <fcntl.h>
#include <errno.h>
#include <unistd.h>
#include <limits.h>
#include <string.h>

#define BUF_SIZE 8192
#define PROCESS_FILENAME "t.txt"
char buffer[BUF_SIZE];

// 正确处理read
int readsome()
{

    int fd = open(PROCESS_FILENAME, O_RDONLY);
    if (fd == -1)
    {
        printf("open file error");
    }
    ssize_t len = 100;
    char *buf = buffer;
    int ret;
    if (len > SSIZE_MAX)
    {
        len = SSIZE_MAX;
    }
    while (len != 0 && ((ret = read(fd, buf, len)) != 0))
    {
        if (ret == -1)
        {
            if (errno == EINTR) //被信号打断
            {
                continue;
            }
            if (errno == EAGAIN) //noblock
            {
                // resubmit later
                printf("#####EAGAIN");
                continue; //???
            }
            else
            {
                // error
                continue; //????!!!!
            }
            perror("read");
            break;
        }
        len -= ret;
        buf += ret;
    }
    printf("%s\n", buffer);
    lseek(fd, (off_t)45, SEEK_SET);
    read(fd, buffer, 1);
    printf("||%c,%d||\n", buffer[0], buffer[0] == '4');

    ftruncate(fd, 1);
    close(fd);
    return 0;
}

int writesome()
{
    const char *buf = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
    int fd = open(PROCESS_FILENAME, O_WRONLY);

    ssize_t ret, len = strlen(buf);
    if (len > SSIZE_MAX) // 如果超过SSIZE_MAX 行为未定义的
    {
        len = SSIZE_MAX;
    }
    while (len != 0 && (ret = write(fd, buf, len)) != -1)
    {
        if (ret == -1)
        {
            if (errno == EINTR)
            {
                continue;
            }
            perror("error");
            break;
        }
        len -= ret;
        buf += ret;
    }
    return 0;
}

int main(int argc, char *argv[], char *env[])
{
    writesome();
    readsome();
}
