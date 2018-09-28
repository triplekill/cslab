#include <stdio.h>
#include <pthread.h>
#include <semaphore.h>
// POSIX互斥锁

int mutex()
{
    // https://docs.oracle.com/cd/E19253-01/819-7051/6n919hpag/index.html
    pthread_mutex_t mp = PTHREAD_MUTEX_INITIALIZER;
    pthread_mutex_init(&mp, NULL);
    int i = 1;

    pthread_mutex_lock(&mp);
    i++;
    pthread_mutex_unlock(&mp);
    return 0;
}

int sem()
{
    sem_t sem;
    int pshared = 0; // 共享选项
    int value = 1;
    int ret=0;

    ret += sem_init(&sem, pshared, value);
    printf("%d %d %d\n", sem, value, ret);
    ret += sem_wait(&sem);
    printf("%d %d %d\n", sem, value, ret);
    ret += sem_post(&sem);
    printf("%d %d %d\n", sem, value, ret);
    return 0;
}

int main()
{
    int ret = 0;
    ret = mutex();
    if (ret)
    {
        return ret;
    }
    ret = sem();
    if (ret)
    {
        return ret;
    }
    return ret;
}