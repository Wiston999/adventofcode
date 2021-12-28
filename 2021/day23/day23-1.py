from __future__ import print_function
import argparse
import logging
import sys

import collections
from queue import PriorityQueue as pq
import re

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

CURRENT_BEST = sys.maxsize
EXAMINED = {}

class SolutionList(object):
    def __init__(self):
        self.list = []
        self.best = Solution([([], 30000)])

    def add(self, s):
        self.list.append(s)
        if self.best is None or s.score < self.best.score:
            self.best = s

class Solution(object):
    def __init__(self, steps):
        self.steps = steps

    @property
    def score(self):
        return sum(s[1] for s in self.steps)

class State(object):
    ROOMS = {
            'A': [(2, 3), (3, 3)],
            'B': [(2, 5), (3, 5)],
            'C': [(2, 7), (3, 7)],
            'D': [(2, 9), (3, 9)],
            }

    HALLWAY = [(1, i) for i in range(1, 12)]

    HALLWAY_MOVES = {
            (1, 1): [(1, 2), (1, 4), (1, 6), (1, 8), (1, 10), (1, 11)],
            (1, 2): [(1, 1), (1, 4), (1, 6), (1, 8), (1, 10), (1, 11)],
            (1, 4): [(1, 1), (1, 2), (1, 6), (1, 8), (1, 10), (1, 11)],
            (1, 6): [(1, 1), (1, 2), (1, 4), (1, 8), (1, 10), (1, 11)],
            (1, 8): [(1, 1), (1, 2), (1, 4), (1, 6), (1, 10), (1, 11)],
            (1, 10): [(1, 1), (1, 2), (1, 4), (1, 6), (1, 8), (1, 11)],
            (1, 11): [(1, 1), (1, 2), (1, 4), (1, 6), (1, 8), (1, 10)],
    }

    COST = {
            'A': 1,
            'B': 10,
            'C': 100,
            'D': 1000,
            }

    def __init__(self, d):
        self.state = [list(l) for l in d.splitlines()]

    def __repr__(self):
        return '\n'.join(''.join(l) for l in self.state)

    def __lt__(self, other):
        return self.score() < other.score()

    def hash(self):
        return ''.join(''.join(l) for l in self.state[1:-1]).encode()

    def is_final(self):
        return all(self.state[v[0][0]][v[0][1]] == k and self.state[v[1][0]][v[1][1]] == k for k, v in State.ROOMS.items())

    def copy(self):
        other = State('')
        other.state = [list(l) for l in self.state]
        return other

    def clear(self, start, end):
        clear = True
        max_i = max(start, end)
        min_i = min(start, end)
        return all(self.state[1][i] in '.#' for i in range(min_i + 1, max_i + 1) if i != start)

    def generate_new(self):
        states = []
        for p in State.HALLWAY:
            element = self.state[p[0]][p[1]]
            if element != '.':
                p0 = State.ROOMS[element][0]
                p1 = State.ROOMS[element][1]
                if self.clear(p[1], p0[1]):
                    if self.state[p1[0]][p1[1]] == '.' and self.state[p0[0]][p0[1]] == '.':
                        ns = self.copy()
                        ns.state[p[0]][p[1]] = '.'
                        ns.state[p1[0]][p1[1]] = element
                        cost = State.COST[element] * (abs(p[1] - p0[1]) + 2)
                        states.append((ns, cost))
                    if self.state[p0[0]][p0[1]] == '.' and self.state[p1[0]][p1[1]] == element:
                        ns = self.copy()
                        ns.state[p[0]][p[1]] = '.'
                        ns.state[p0[0]][p0[1]] = element
                        cost = State.COST[element] * (abs(p[1] - p0[1]) + 1)
                        states.append((ns, cost))

        for element, positions in State.ROOMS.items():
            if self.state[positions[1][0]][positions[1][1]] not in [element, '.']:
                wrong_element = self.state[positions[1][0]][positions[1][1]]
                if self.state[positions[0][0]][positions[0][1]] == '.':
                    ns = self.copy()
                    ns.state[positions[0][0]][positions[0][1]] = wrong_element
                    ns.state[positions[1][0]][positions[1][1]] = '.'
                    states.append((ns, State.COST[wrong_element]))
            if self.state[positions[0][0]][positions[0][1]] not in [element, '.'] or \
                (self.state[positions[0][0]][positions[0][1]] == element and self.state[positions[1][0]][positions[1][1]] not in [element, '.']):
                wrong_element = self.state[positions[0][0]][positions[0][1]]
                for _, y in State.HALLWAY_MOVES:
                    if self.clear(positions[0][1], y):
                        ns = self.copy()
                        ns.state[1][y] = wrong_element
                        ns.state[positions[0][0]][positions[0][1]] = '.'
                        states.append((ns, State.COST[wrong_element] * (1 + abs(y - positions[0][1]))))

        return sorted(states, key=lambda x: x[1])

    def locate(self, e):
        location = []
        for i, l in enumerate(self.state):
            for j, c in enumerate(l):
                if c == e:
                    location.append((i, j))
        return location

    def score(self):
        score = 0
        for element, positions in State.ROOMS.items():
            locations = self.locate(element)
            for l in locations:
                if l not in positions:
                    score += abs(l[0] - positions[1][0]) + abs(l[1] - positions[1][1])
        return score

def search(steps, sl, depth=0):
    state = steps[-1][0]
    EXAMINED[state.hash()] = sum(s[1] for s in steps)
    for s, cost in state.generate_new():
        solution = Solution(steps + [(s, cost)])
        h = s.hash()
        if s.is_final():
            sl.add(solution)
            logger.info('(%03d) Current best: %s - %s', depth, solution.score, '\n'.join(str(s) for s in sl.best.steps))
        elif (h not in EXAMINED or EXAMINED[h] > solution.score) and sl.best.score > solution.score:
        # elif sl.best.score > solution.score and depth < 15:
            if logger.isEnabledFor(logging.DEBUG):
                logger.debug('(%02d) Exploring state (%04d) (%03d) %s', depth, cost, s.score(), s)
            result = search(solution.steps, sl, depth+1)

def dijkstra(state):
    visited = set()
    pending = pq()
    pending.put((0, state))
    costs = collections.defaultdict(lambda: sys.maxsize)
    best = sys.maxsize
    while not pending.empty():
        current = pending.get()
        visited.add(str(current[1]))
        for s, cost in current[1].generate_new():
            if str(s) not in visited:
                new_cost = current[0] + cost
                if new_cost < costs[str(s)]:
                    costs[str(s)] = new_cost
                    pending.put((new_cost, s))
                    if s.is_final() and new_cost < best:
                        best = new_cost
                        logger.info('New best: %s', best)

    return best

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('-i', '--input', type=argparse.FileType('r'), default=sys.stdin,
            help='Intput file, use - for stdin')
    arg_parser.add_argument('-o', '--output', type=argparse.FileType('w'), default=sys.stdout,
            help='Output file, use - for stdout')
    arg_parser.add_argument('-l', '--loglevel', type=str.upper, default='info',
            choices=['DEBUG', 'INFO', 'WARNING', 'ERROR'], help='Output file (when new set is created)')
    args = arg_parser.parse_args(argv)

    logger.setLevel(args.loglevel)

    logger.debug('Log level: %s', args.loglevel)
    logger.debug('Input file: %s', args.input.name)
    logger.debug('Output file: %s', args.output.name)
    result = 0

    state = State(args.input.read())

    logger.info('Read state: %s', state)
    sl = SolutionList()
    search([(state, 0)], sl)
    result = sl.best.score
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
