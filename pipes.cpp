#define _GNU_SOURCE
#include <sys/types.h>
#include <sys/wait.h>
#include <stdio.h>
#include <sched.h>
#include <signal.h>
#include <unistd.h>
#include <assert.h>

int main()
{
    int contact[2];
    pipe(contact);

    int fd = fork();
    if (fd == 0) { // child
       char c;
       close(contact[1]); // close the write end.

       while (read(contact[0], &c, 1) > 0)
           printf("c=%c", c);
       close(contact[0]);
       exit(0);
    } else { // parent
        close(contact[0]); // close the read end.

        write(contact[1], "X", 1);
        close(contact[1]); // child sees EOF
        wait(NULL);
        exit(EXIT_SUCCESS);
    }
    return 1;
}