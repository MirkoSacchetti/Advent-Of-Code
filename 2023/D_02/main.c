#include <stdio.h>
#include <stdlib.h>

#define MAX_LINE_LENGTH 1025
#define MAX_HANDS_IN_GAME 25

typedef struct {
  int red;
  int green;
  int blue;
} CubesHand;

int getGameID(char **hands) {
  int gameID = 0;
  while (**hands && **hands != ':') {
    if (**hands >= '0' && **hands <= '9')
      gameID = gameID * 10 + (**hands - '0');
    (*hands)++;
  }
  return gameID;
}

int getHandCubes(char **hands, CubesHand *cube) {
  while (**hands && **hands != ';') {
    int nCubes = 0;
    int colorIndex = 0;
    char color[12];
    while (**hands && **hands != ',' && **hands != ';') {
      // check numbers or words or ignore the rest
      if (**hands >= '0' && **hands <= '9') {
        nCubes = nCubes * 10 + (**hands - '0');
      } else if (**hands >= 'a' && **hands <= 'z') {
        color[colorIndex++] = **hands;
      }
      (*hands)++;
    }
    color[colorIndex] = '\0';
    if (color[0] == 'r')
      cube->red = nCubes;
    if (color[0] == 'g')
      cube->green = nCubes;
    if (color[0] == 'b')
      cube->blue = nCubes;
    (*hands)++;
  }
  return 0;
}

int isValidGame(CubesHand *hands, int length) {
  CubesHand maxCubes = {12, 13, 14};
  for (int i = 0; i < length; i++) {
    printf("%d - ", hands[i].red);
    printf("%d - ", hands[i].green);
    printf("%d\n", hands[i].blue);
    if (hands[i].red > maxCubes.red)
      return 0;
    if (hands[i].green > maxCubes.green)
      return 0;
    if (hands[i].green > maxCubes.blue)
      return 0;
  }
  return 1;
}

int main() {
  FILE *file = fopen("test", "r");
  if (!file) {
    perror("Err. in file opening");
    return -1;
  }
  int counter1 = 0;
  int counter2 = 0;
  char *line = malloc(MAX_LINE_LENGTH);
  while (fgets(line, MAX_LINE_LENGTH, file)) {
    char *hands = line;
    int gId = getGameID(&hands);
    CubesHand gamesCubes[MAX_HANDS_IN_GAME];
    int cubeHandIndex = 0;
    while (*hands) {
      CubesHand cube = {0, 0, 0};
      getHandCubes(&hands, &cube);
      gamesCubes[cubeHandIndex] = cube;
      cubeHandIndex++;
      if (isValidGame(gamesCubes, cubeHandIndex)) {
      }
      counter1 += gId;
    }
  }

  printf("Result first half: %d\n", counter1);
  printf("Result second half: %d\n", counter2);
  fclose(file);
  return 0;
}
