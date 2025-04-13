def parse_file(filename):
    boards = []
    draws = []
    with open(filename, "r") as file:
        content = file.read().strip()
    blocks = content.split("\n\n")
    # first block is the list of integers draws
    draws = list(map(int, blocks[0].split(",")))
    # after the first line, the rest are boards, rows of integers
    boards = [
        [list(map(int, row.split())) for row in board.splitlines()]
        for board in blocks[1:]
    ]
    return draws, boards

def part_one(data):
    draws = data[0]
    boards = data[1]
    for n in draws:
        for board in boards:
            bingo = mark_number(board, n)
            if bingo > 0:
                return bingo

def part_two(data):
    draws = data[0]
    boards = data[1]
    last_bingo = 0
    for n in draws:
        for board in boards:
            bingo = mark_number(board, n)
            if bingo > 0:
                last_bingo = bingo
                # only ONE bingo for board!
                boards.remove(board)
    return last_bingo

def check_bingo(board, row, col):
    score = 0
    bingo = True
    #check the row of the number
    for i in range(len(board)):
        if board[i][col] != "*":
            bingo = False
            break
    if not bingo:
        #check the column of the number
        bingo = True
        for j in range(len(board[row])):
            if board[row][j] != "*":
                bingo = False
                break
    if bingo:
        for i in range(len(board)):
            for j in range(len(board[i])):
                if not board[i][j] == "*":
                    score +=board[i][j]
    return score

def mark_number( board, number):
    for i in range(len(board)):
        for j in range(len(board[i])):
            if board[i][j] == number:
                board[i][j] = "*"
                bingo = check_bingo(board, i, j)
                if bingo > 0:
                    return  bingo * number
    return 0 

if __name__ == "__main__":
    data = parse_file("input.txt")
    # reusing the same boards in part two is acceptable
    # even though part_one modifies the boards....
    print("Part one:", part_one(data))
    print("Part two:", part_two(data))
