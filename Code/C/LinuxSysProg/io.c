#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <fcntl.h>
#include <errno.h>
#include <unistd.h>
#include <limits.h>
#include <string.h>
#include <time.h>
#include <stddef.h>
#include <sys/select.h>
#include <sys/poll.h>
#include <sys/uio.h>
#include <sys/mman.h>
// #include <sys/epoll.h>
#include <aio.h>
// #include <sys/stat.h>

#ifndef SSIZE_MAX
// ubuntu 18.04 nodef
#define SSIZE_MAX 1000
#endif
#define TIMEOUT 5 // in second
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

// !!!! #include <sys/select.h>
void selectl()
{
    fd_set rfds;
    struct timeval tv;
    int retval;

    /* Watch stdin (fd 0) to see when it has input. */
    FD_ZERO(&rfds);
    FD_SET(STDIN_FILENO, &rfds);

    /* Wait up to five seconds. */
    tv.tv_sec = 5;
    tv.tv_usec = 0;

    retval = select(1, &rfds, NULL, NULL, &tv);
    /* Don't rely on the value of tv now! */
    if (retval == -1)
        perror("select()");
    else if (retval)
        printf("Data is available now.\n");
    /* FD_ISSET(0, &rfds) will be true. */
    else
        printf("No data within five seconds.\n");

    exit(EXIT_SUCCESS);
}

int selectvsleep()
{
    struct timeval tv;
    tv.tv_usec = 0;
    tv.tv_usec = 500; // sleep 500 microsecondes 在linux和mac的单位不一样？
    select(0, NULL, NULL, NULL, &tv);
    return 0;
}

int polll()
{
    struct pollfd fds[2];
    int ret;

    fds[0].fd = STDIN_FILENO;
    fds[0].events = POLLIN;
    fds[1].fd = STDOUT_FILENO;
    fds[1].events = POLLOUT;

    ret = poll(fds, 2, TIMEOUT * 1000);
    if (ret == -1)
    {
        perror("poll");
        return 1;
    }
    if (!ret)
    {
        printf("poll timeout %d s", TIMEOUT);
        return 0;
    }

    if (fds[0].revents & POLLIN)
    {
        printf("%d is readable\n", fds[0]);
    }
    if (fds[1].revents & POLLOUT)
    {
        printf("%d is writeable\n", fds[1]);
    }
    return 0;
}

#define LINE_MAX 10
// fopen
int stdfileprocess()
{
    FILE *fd = fopen("t.txt", "r+");
    char ch = fgetc(fd);
    printf("###%c %d###\n", ch, EOF);
    char buf[LINE_MAX];
    if (fgets(buf, LINE_MAX, fd))
    {
        printf("fgets |%d|%s|\n", LINE_MAX, buf);
    }
    fputc('c', fd);
    fputs("write string\n", fd);
    fclose(fd);

    FILE *in, *out;
    struct pirate
    {
        char name[100];
        unsigned long booty;
        unsigned int beard_len;
    } p, blockbeard = {"Edward Teach", 950, 48};
    out = fopen("data.bat", "w");
    if (!out)
    {
        perror("fopen");
        return 1;
    }
    if (!fwrite(&blockbeard, sizeof(struct pirate), 1, out))
    {
        perror("fwrite");
        return 1;
    }
    if (fclose(out))
    {
        perror("fclose");
        return 1;
    }
    in = fopen("data.bat", "r");
    if (!in)
    {
        perror("fopen");
    }
    if (!fread(&p, sizeof(struct pirate), 1, in))
    {
        perror("fread");
        return 1;
    }
    if (fclose(in))
    {
        perror("fclose");
        return 1;
    }
    printf("name=%s booty=%lu beard_len=%u\n", p.name, p.booty, p.beard_len);
    return 0;
}

int readvwriteviov()
{
    struct iovec iov[3];
    ssize_t nr;
    int fd, i;
    char *buf[] = {
        "aaaaaaaaaaaa.",
        "bbbbbbbbbbbb.",
        "cccccccccccc.",
    };
    fd = open("iov.bat", O_WRONLY | O_CREAT, 777);
    if (fd == -1)
    {
        perror("open");
        return 1;
    }
    for (i = 0; i < 3; i++)
    {
        iov[i].iov_base = buf[i];
        iov[i].iov_len = strlen(buf[i] + 1);
    }
    nr = writev(fd, iov, 3);
    if (nr == -1)
    {
        perror("writev");
        return 1;
    }
    printf("write %d byte\n", nr);
    if (close(fd))
    {
        perror("close");
    }

    // read
    char a[13], b[13], c[13];
    a[12] = '\0'; // 补充最后的一个\0
    b[12] = '\0';
    c[12] = '\0';

    fd = open("iov.bat", O_RDONLY);
    if (fd == -1)
    {
        perror("open");
        return 1;
    }

    iov[0].iov_base = a;
    iov[0].iov_len = sizeof(a) - 1; // 保留最后的\0
    iov[1].iov_base = b;
    iov[1].iov_len = sizeof(b) - 1;
    iov[2].iov_base = c;
    iov[2].iov_len = sizeof(c) - 1;

    nr = readv(fd, iov, 3);
    if (nr == -1)
    {
        perror("readv");
        return 1;
    }
    printf("write %d byte\n", nr);
    if (close(fd))
    {
        perror("close");
    }
    printf("%s\n", a);
    return 0;
}

int mmapt()
{
    int fd = open("t.txt", O_RDWR);
    void *p;
    p = mmap(0, 1, PROT_READ, MAP_SHARED, fd, 0);
    if (p == MAP_FAILED)
    {
        perror("mmap");
        return 1;
    }
    printf("%c\n", ((char *)p)[45]);
}
int main(int argc, char *argv[], char *env[])
{
    writesome();
    readsome();
    // selectl();
    // selectvsleep();
    polll();
    stdfileprocess();
    readvwriteviov();
    mmapt();
}
