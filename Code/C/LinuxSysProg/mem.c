#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <errno.h>

typedef struct
{
    int age;
    char name[10];
} people;
int dynamicmemalloc()
{
    people *p = (people *)malloc(sizeof(people));
    if (!p)
    {
        perror("malloc people");
    }
    strcpy((char *)&(p->name), "tom");
    p->age = 10;
    printf("name:%s, age:%d\n", p->name, p->age);
    free(p);
    p = NULL;
    return 0;
}

int realloct()
{
    void *p = malloc(sizeof(int));
    void *q = realloc(p, sizeof(int) * 5);
    printf("old:%p new:%p\n", p, q); // p q 可能相等，可能不想等，具体要看连续内存是否足够
    free(q);
    // free(q);
    if (errno)
    {
        perror("free");
    }
    return 0;
}

int main()
{
    malloc(1);
    dynamicmemalloc();
    realloct();
}
