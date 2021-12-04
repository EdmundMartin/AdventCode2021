from typing import List, Dict, Tuple, Optional


def matrix_to_set(matrix: List[List[int]]) -> Dict[int, Tuple[int, int]]:
    result = dict()
    for row, _ in enumerate(matrix):
        for col, _ in enumerate(matrix):
            result[matrix[row][col]] = (row, col)
    return result


class Board:

    def __init__(self, matrix: List[List[int]]):
        self.matrix = matrix
        self.values = matrix_to_set(matrix)
        self.marked = set()

    def __str__(self):
        return f"<Board: {self.matrix}>"

    def __contains__(self, item):
        return item in self.values

    def _mark_position(self, number: int) -> Tuple[int, int]:
        pos = self.values[number]
        self.marked.add(pos)
        return pos

    def _is_winner(self, row_number, col_number):
        target_row = [(row_number, i) for i in range(5)]
        target_col = [(i, col_number) for i in range(5)]
        if all(i in self.marked for i in target_row) or all(i in self.marked for i in target_col):
            return True
        return False

    def check_win(self, number: int) -> bool:
        pos = self._mark_position(number)
        if len(self.marked) < 5:
            return False
        return self._is_winner(*pos)

    def get_unmarked_numbers(self) -> List[int]:
        unmarked = []
        for key, value in self.values.items():
            if value not in self.marked:
                unmarked.append(key)
        return unmarked


def read_input(filename: str):
    with open(filename, 'r') as f:
        contents = f.readlines()
    bingo_numbers = [int(c) for c in contents[0].rstrip().split(',')]
    boards = []
    current_board = []
    for line in contents[2:]:
        if len(line) == 1:
            boards.append(Board(current_board))
            current_board = []
        else:
            cleaned_line = line.strip().split(' ')
            current_board.append([int(c) for c in cleaned_line if c != ''])
    boards.append(Board(current_board))
    return bingo_numbers, boards


def iterate_first_winner(numbers: List[int], board_list: List[Board]) -> Tuple[int, Board]:
    for number in numbers:
        for board in board_list:
            if number in board:
                is_win = board.check_win(number)
                if is_win:
                    return number, board
    raise AttributeError("No winning board")


def iterate_last_winner(numbers: List[int], board_list: List[Board]) -> Tuple[int, Board]:
    last_number, last_winner = None, None
    seen_winners = set()
    for number in numbers:
        for idx, board in enumerate(board_list):
            if idx in seen_winners:
                continue
            if number in board:
                is_win = board.check_win(number)
                if is_win and idx not in seen_winners:
                    last_number, last_winner = number, board
                    seen_winners.add(idx)
    return last_number, last_winner


def part1(numbers, board_list) -> int:
    win_num, win_board = iterate_first_winner(numbers, board_list)
    unmarked = win_board.get_unmarked_numbers()
    return sum(unmarked) * win_num


def part2(numbers, board_list) -> int:
    win_num, win_board = iterate_last_winner(numbers, board_list)
    unmarked = win_board.get_unmarked_numbers()
    return sum(unmarked) * win_num


if __name__ == '__main__':
    import time
    start = time.time()
    numbers, board_list = read_input("input.txt")
    print(part1(numbers, board_list))
    print(part2(numbers, board_list))
    print(time.time() - start)