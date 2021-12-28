from __future__ import print_function
import argparse
import logging
import sys

import collections
from queue import PriorityQueue as pq

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

CURRENT_BEST = sys.maxsize
EXAMINED = {}

class SolutionList(object):
    def __init__(self):
        self.list = []
        self.best = Solution([([], 20000)])

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
    HALLWAY = list(range(11))

    POSISION_ROOM = {
        11: 'A', 15: 'A',
        12: 'B', 16: 'B',
        13: 'C', 17: 'C',
        14: 'D', 18: 'D',
    }

    ROOMS = {
        'A': [11, 15],
        'B': [12, 16],
        'C': [13, 17],
        'D': [14, 18],
    }

    ROOM_HALLWAY = {
        11: 2,
        12: 4,
        13: 6,
        14: 8,
        15: 2,
        16: 4,
        17: 6,
        18: 8,
    }

    #
    # 0 1 2 3 4 5 6 7 8 9 10
    #     11  12  13  14
    #     15  16  17  18
    # . . . . . . . . . . .
    #   # B # C # B # D # #
    #   # A # D # C # A # 
    #

    MOVES = {
        0: { 11: 3, 12: 5, 13: 7, 14: 9, 15: 4, 16: 6, 17: 8, 18: 10 },
        1: { 11: 2, 12: 4, 13: 6, 14: 8, 15: 3, 16: 5, 17: 7, 18: 9 },
        3: { 11: 2, 12: 2, 13: 4, 14: 6, 15: 3, 16: 3, 17: 5, 18: 7 },
        5: { 11: 4, 12: 2, 13: 2, 14: 4, 15: 5, 16: 3, 17: 3, 18: 5 },
        7: { 11: 6, 12: 4, 13: 2, 14: 2, 15: 7, 16: 5, 17: 3, 18: 3 },
        9: { 11: 8, 12: 6, 13: 4, 14: 2, 15: 9, 16: 7, 17: 5, 18: 3 },
        10: { 11: 9, 12: 7, 13: 5, 14: 3, 15: 10, 16: 8, 17: 6, 18: 4 },
        11: { 0: 3, 1: 2, 3: 2, 5: 4, 7: 6, 9: 8, 10: 9 },
        12: { 0: 5, 1: 4, 3: 2, 5: 2, 7: 4, 9: 6, 10: 7 },
        13: { 0: 7, 1: 6, 3: 4, 5: 2, 7: 2, 9: 4, 10: 5 },
        14: { 0: 9, 1: 8, 3: 6, 5: 4, 7: 2, 9: 2, 10: 3 },
        15: { 0: 4, 1: 3, 3: 3, 5: 5, 7: 7, 9: 9, 10: 10 },
        16: { 0: 6, 1: 5, 3: 3, 5: 3, 7: 5, 9: 7, 10: 8 },
        17: { 0: 8, 1: 7, 3: 5, 5: 3, 7: 3, 9: 5, 10: 6 },
        18: { 0: 10, 1: 9, 3: 7, 5: 5, 7: 3, 9: 3, 10: 4 },
    }

    COST = {
        'A': 1,
        'B': 10,
        'C': 100,
        'D': 1000,
    }

    REVERSE_COST = {
        0: '.',
        1: 'A',
        10: 'B',
        100: 'C',
        1000: 'D',
    }

    def __init__(self, d):
        self.state = [0 for _ in range(max(State.MOVES) + 1)]
        if d is not None:
            # Hallway parse
            for i, c in enumerate(d[1][1:-1]):
                if c in State.COST:
                    self.state[i] = State.COST[c]
            # Rooms parse
            for i, l in enumerate(d[2:-1]):
                for j, c in enumerate(l.replace('#', ' ').split(' ')[3:7], start=i*4 + 11):
                    if c in State.COST:
                        self.state[j] = State.COST[c]

    def __repr__(self):
        st =  '#############\n#{}#\n'.format(''.join(State.REVERSE_COST[self.state[v]] for v in State.HALLWAY))
        for i in range(11, max(State.MOVES.keys())+1, 4):
            st += '###{}#{}#{}#{}###\n'.format(
                State.REVERSE_COST[self.state[i]],
                State.REVERSE_COST[self.state[i+1]],
                State.REVERSE_COST[self.state[i+2]],
                State.REVERSE_COST[self.state[i+3]],
            )
        st += '#############\n'
        return st

    def __lt__(self, other):
        return self.score() < other.score()

    def hash(self):
        return ','.join(map(str, self.state))

    def is_final(self):
        return all(self.state[x] == State.COST[k] for k, v in State.ROOMS.items() for x in v)

    def copy(self):
        other = State(None)
        other.state = self.state.copy()
        return other

    def clear(self, start, end):
        max_i = max(start, end)
        min_i = min(start, end)
        return 0 == sum(self.state[i] for i in range(min_i, max_i + 1) if i != start)

    def move(self, start, end):
        value = self.state[start]
        ns = self.copy()
        ns.state[start] = 0
        ns.state[end] = value
        cost = value * State.MOVES[start][end]
        return ns, cost

    def can_leave(self, start):
        room = State.ROOMS[State.POSISION_ROOM[start]]
        for r in room:
            if r < start and self.state[r] > 0:
                return False
        return True

    def generate_new(self):
        input_states, output_states = [], []
        for p, v in enumerate(self.state):
            if v > 0:
                rooms = State.ROOMS[State.REVERSE_COST[v]]
                good_room = all(self.state[r] == v or self.state[r] == 0 for r in rooms)
                full_room = all(self.state[r] == v for r in rooms)
                if full_room: continue
                if p < 11:
                    for r in rooms[::-1]:
                        if not self.clear(p, State.ROOM_HALLWAY[r]): break
                        if self.state[r] == 0 and r in State.MOVES[p]:
                            input_states.append(self.move(p, r))
                        elif self.state[r] == v and r in State.MOVES[p]:
                            continue
                        else:
                            break
                for m in State.MOVES[p]:
                    if self.state[m] == 0:
                        if m < 11 and p < 11: continue
                            # hallway_states.append(self.move(p, m))
                        elif p >= 11 and m < 11:
                            if p in rooms and good_room: continue
                            if self.can_leave(p) and self.clear(State.ROOM_HALLWAY[p], m):
                                output_states.append(self.move(p, m))

        input_states.sort(key=lambda x: x[1] / 1000.0 + x[0].score())
        if any(s.is_final() for s, c in input_states):
            output_states = []
        else:
            output_states.sort(key=lambda x: x[1])

        return input_states + output_states

    def score(self):
        score = 0
        for p, v in enumerate(self.state):
            if v > 0 and p not in State.ROOMS[State.REVERSE_COST[v]]:
                score += v * abs(p - State.ROOMS[State.REVERSE_COST[v]][-1])
        return score

def search(steps, sl, depth=0):
    state = steps[-1][0]
    EXAMINED[state.hash()] = sum(s[1] for s in steps)
    for s, cost in state.generate_new():
        solution = Solution(steps + [(s, cost)])
        h = s.hash()
        if sl.best.score > solution.score:
            if s.is_final():
                sl.add(solution)
                logger.info('(%03d) Current best: %s - %s', depth, solution.score, '\n'.join(str(s) for s in sl.best.steps))
            elif (h not in EXAMINED or EXAMINED[h] > solution.score):
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
                if new_cost < costs[str(s)] and new_cost < best:
                    costs[str(s)] = new_cost
                    pending.put((new_cost, s))
                    if s.is_final():
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

    state = State(args.input.readlines())
    logger.info('Read state: %s', state)
    sl = SolutionList()
    search([(state, 0)], sl)
    logger.info ('B&B best: %s - %s', sl.best.score, '\n'.join(str(s[0]) + '\n' + str(s[1]) for s in sl.best.steps))
    logger.info('Explored %s states', len(EXAMINED))
    result = dijkstra(state)
    logger.info ('Dijkstra best: %s', sl.best.score)
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
