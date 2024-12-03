#include <stdio.h>

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

void getHandCubes(char **hands, CubesHand *hand) {
  while (**hands && **hands != ';') {
    int nCubes = 0;
    int colorIndex = 0;
    char color[6] = {0};

    while (**hands && **hands != ',' && **hands != ';') {
      if (**hands >= '0' && **hands <= '9') {
        nCubes = nCubes * 10 + (**hands - '0');
      } else if (**hands >= 'a' && **hands <= 'z') {
        color[colorIndex++] = **hands;
      }
      (*hands)++;
    }
    if (color[0] == 'r')
      cube->red = nCubes;
    if (color[0] == 'g')
      cube->green = nCubes;
    if (color[0] == 'b')
      cube->blue = nCubes;
    if (**hands == ';')
      return 0;
    (*hands)++;
  }
  if (**hands == ';')
    (*hands)++;
}

int isValidGame(CubesHand *hands, int length, int gId) {
  CubesHand maxCubes = {12, 13, 14};
  for (int i = 0; i < length; i++) {
    if (hands[i].red > maxCubes.red)
      return 0;
    }
    if (hands[i].green > maxCubes.green) {
      // printf("game %d is not valid, cube %d green is %d\n", gId, i,
      // hands[i].green);
      return 0;
    if (hands[i].blue > maxCubes.blue)
      return 0;
    }
  }
  return gId;
}

int getPowerMinCubes(CubesHand *hands, int length, int gId) {
  CubesHand minCubes = {0};
  for (int i = 0; i < length; i++) {
    if (i == 0 || minCubes.red < hands[i].red) {
      minCubes.red = hands[i].red;
    }
    if (i == 0 || minCubes.green < hands[i].green) {
      minCubes.green = hands[i].green;
    }
    if (i == 0 || minCubes.blue < hands[i].blue) {
      minCubes.blue = hands[i].blue;
    }
  }
  return minCubes.red * minCubes.green * minCubes.blue;
}

int main() {
  FILE *file = fopen("input", "r");
  if (!file) {
    perror("Error in file opening");
    return -1;
  }
  int counter1 = 0;
  int counter2 = 0;
  char line[MAX_LINE_LENGTH];
  while (fgets(line, MAX_LINE_LENGTH, file)) {
    char *hands = line;
    int gId = getGameID(&hands);
    CubesHand gamesCubes[MAX_HANDS_IN_GAME] = {0};
    int cubesHandIndex = 0;
    while (*hands) {
      CubesHand cube = {0, 0, 0};
      getHandCubes(&hands, &cube);
      printf("%d=%d %d %d \n", gId, cube.red, cube.green, cube.blue);
      gamesCubes[cubeHandIndex] = cube;
      cubeHandIndex++;
      if (isValidGame(gamesCubes, cubeHandIndex)) {
        counter1 += gId;
      } else {
      }
    }
    counter1 += isValidGame(gamesCubes, cubesHandIndex, gId);
    counter2 += getPowerMinCubes(gamesCubes, cubesHandIndex, gId);
  }
  printf("\n");
  printf("Result first half: %d\n", counter1);
  printf("Result second half: %d\n", counter2);
  fclose(file);
  return 0;
}
