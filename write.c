#define _GNU_SOURCE
#include <unistd.h>
#include <sys/syscall.h>
#include <sys/types.h>
int
main(int argc, char *argv[])
{
    char ogromnoyeNazvanieStroki[19] = "abhaziyabezobraziya";
    write(1,ogromnoyeNazvanieStroki, 19);
}