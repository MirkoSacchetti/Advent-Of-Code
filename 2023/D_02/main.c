#include <stdio.h>

typedef struct {
    int red;
    int green;
    int blue;
} Cubes;

int getGameID(char *game, int *index) {
}
bool isValidGame(char *game, int *index) {
    Cubes maxCubes = {red: 12, green: 13. bluee:14}
}

int main() {
    FILE *file = fopen("test", "r");
    if (!file) {
        perror("Err. in file opening");
        return -1;
    }
    int counter1 = 0
    int counter2 = 0
    while (fgets(line, sizeof(line), file)) {
        int index =0;
        gameID =  getGameID(line, &inmdex)
        if ( checkGame(line, &index) ) counter1 =+ gameID;
    }

    fclose(file);
    printf("Result first half: %d\n", counter1);
    printf("Result second half: %d\n", counter2);
    return 0;
}

