#include <stdio.h>

int main() {
    char userInput[256];
    printf("Enter input: ");
    scanf("%255s", userInput);
    printf(userInput);
    return 0;
}
