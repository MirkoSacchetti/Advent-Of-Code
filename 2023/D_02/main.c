#include <stdio.h>

#define MAX_LINE_LENGTH 1025

typedef struct {
  int red;
  int green;
  int blue;
} CubesHand;

int getGameID(char *game) {
  int gameID = 0;
  while (*game && *game != ':') {
    if (*game >= '0' && *game <= '9')
      gameID = gameID * 10 + (*game - '0');
    game++;
  }
  game++;
  return gameID;
}

int getCubesHand(char *game, CubesHand *cube) {
  while (*game && *game != ';') {
    int nCubes = 0;
    int colorIndex = 0;
    char color[12];
    while (*game && *game != ',') {
      // check numbers or words or ignore the rest
      if (*game >= '0' && *game <= '9') {
        nCubes = nCubes * 10 + (*game - '0');
      } else if (*game >= 'a' && *game <= 'z') {
        color[colorIndex++] = game[0];
        printf("++++ %s\n", color);
      }

      game++;
    }

    printf("++++ %s\n", color);
    game++;
  }
  return 0;
}

int isValidGame(char *game) { return 1; }

int main() {
  FILE *file = fopen("test", "r");
  if (!file) {
    perror("Err. in file opening");
    return -1;
  }
  int counter1 = 0;
  int counter2 = 0;
  char line[MAX_LINE_LENGTH];
  while (fgets(line, MAX_LINE_LENGTH, file)) {
    char *ptr = line;
    CubesHand cube;
    int gId = getGameID(ptr);
    getCubesHand(ptr, &cube);
  }
  printf("Result first half: %d\n", counter1);
  printf("Result second half: %d\n", counter2);
  fclose(file);
  return 0;
}
