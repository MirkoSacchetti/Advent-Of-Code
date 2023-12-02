#include <stdio.h>

int isDigit(char c) {
  if (c >= '0' && c <= '9')
    return c - '0';
  return -1;
}

int isLetters(char *line, int index) {
  char *words[] = {"one", "two",   "three", "four", "five",
                   "six", "seven", "eight", "nine"};
  int length = sizeof(words) / sizeof(words[0]);
  for (int i = 0; i < length; i++) {
    char *l = &line[index];
    while (*words[i]) {
      if (*l != *words[i])
        break;
      words[i]++;
      l++;
    }
    if (*words[i] == '\0')
      return i + 1;
  }
  return -1;
}

int main() {
  FILE *file;
  file = fopen("input", "r");
  if (file == NULL) {
    perror("Err. in file opening");
    return -1;
  }
  char line[1024];
  int counter1 = 0;
  int counter2 = 0;
  while (fgets(line, sizeof(line), file)) {
    // part one
    int a = -1;
    int b = -1;
    for (int i = 0; line[i] != '\0'; i++) {
      int digit = isDigit(line[i]);
      if (digit != -1) {
        if (a == -1) {
          a = digit;
          b = digit;
        } else {
          b = digit;
        }
      }
    }
    counter1 += a * 10 + b;
    // part two
    a = -1;
    b = -1;
    for (int i = 0; line[i] != '\0'; i++) {
      int digit = isDigit(line[i]);
      if (digit == -1)
        digit = isLetters(line, i);
      if (digit != -1) {
        if (a == -1) {
          a = digit;
          b = digit;
        } else {
          b = digit;
        }
      }
    }
    counter2 += a * 10 + b;
  }
  fclose(file);
  printf("Result first half: %d\n", counter1);
  printf("Result second half: %d\n", counter2);
  return 0;
}
