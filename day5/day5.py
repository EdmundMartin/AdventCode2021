from typing import List, Tuple, Dict
from collections import defaultdict


class Position:

    def __init__(self, x: int, y: int):
        self.x = x
        self.y = y

    def intersections(self, other: 'Position', calculate_diag=False) -> List[Tuple[int, int]]:
        intersections = []
        if self.x == other.x:
            min_y = min(self.y, other.y)
            max_y = max(self.y, other.y)
            for i in range(min_y, max_y + 1):
                intersections.append((self.x, i))
        elif self.y == other.y:
            min_x = min(self.x, other.x)
            max_x = max(self.x, other.x)
            for i in range(min_x, max_x + 1):
                intersections.append((i, self.y))
        elif calculate_diag:
            x = self.x
            y = self.y
            n = x - other.x
            if n < 0:
                n = -n
            if other.x < x:
                x_step = -1
            else:
                x_step = 1
            if other.y < y:
                y_step = -1
            else:
                y_step = 1
            for i in range(n + 1):
                intersections.append((x, y))
                x += x_step
                y += y_step
        return intersections

    def __str__(self):
        return f"<Position: {self.x}, {self.y}>"


def parse_pairing(pairing: str) -> Position:
    pairing = pairing.strip()
    cords = pairing.split(',')
    return Position(int(cords[0]), int(cords[1]))


def read_inputs(filename: str) -> List[Tuple[Position, Position]]:
    results = []
    with open(filename, 'r') as input_file:
        lines = input_file.readlines()
    for line in lines:
        first, second = line.split('->')
        results.append((parse_pairing(first), parse_pairing(second)))
    return results


def count_danger(counts: Dict[Tuple[int, int], int]) -> int:
    total = 0
    for danger_lvl in counts.values():
        if danger_lvl >= 2:
            total += 1
    return total


def calculate_amount(inputs: List[Tuple[Position, Position]], calculate_diag: bool) -> int:
    results = defaultdict(int)
    for pairings in inputs:
        a, b = pairings
        intersects = a.intersections(b, calculate_diag=calculate_diag)
        for i in intersects:
            results[i] += 1
    return count_danger(results)


if __name__ == '__main__':
    res = read_inputs("input.txt")
    print(f"Part 1: {calculate_amount(res, False)}")
    print(f"Part 2: {calculate_amount(res, True)}")