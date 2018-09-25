#include <stdio.h>
#include <sys/epoll.h>
#include <fcntl.h>

int epollt()
{
    int epfd;
    struct epoll_event event;
    int ret;
    int fd = open("t.txt", O_RDWR);
    if (fd == -1)
    {
        perror("open");
        return 1;
    }
    epfd = epoll_create(100);
    if (epfd < 0)
    {
        perror("epoll create");
    }
    event.data.fd = fd;
    event.events = EPOLLIN | EPOLLOUT;
    ret = epoll_ctl(epfd, EPOLL_CTL_ADD, fd, &event);
    if (ret)
    {
        perror("epoll_ctl?");
    }
    return 0;
}

int main()
{
    printf("epoll\n");
    epollt();
}
