from __future__ import print_function
import argparse
import logging
import sys

import math
import re
import itertools
from collections import Counter

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

def matrix_mul(A, B):
    result = [[0] * len(B[0]) for _ in range(len(B))]
    for i in range(len(A)):
        for j in range(len(B[0])):
            for k in range(len(B)):
                result[i][j] += A[i][k] * B[k][j]

    return (round(result[0][0]), round(result[1][0]), round(result[2][0]))

def rotation_matrix(axis, rotation):
    matrix = [[0] * 3 for _ in range(3)]
    if axis == 'x':
        matrix = [
            [1, 0, 0],
            [0, math.cos(rotation * math.pi / 180), -math.sin(rotation * math.pi / 180)],
            [0, math.sin(rotation * math.pi / 180), math.cos(rotation * math.pi / 180)],
        ]
    if axis == 'y':
        matrix = [
            [math.cos(rotation * math.pi / 180), 0, math.sin(rotation * math.pi / 180)],
            [0, 1, 0],
            [-math.sin(rotation * math.pi / 180), 0, math.cos(rotation * math.pi / 180)],
        ]
    if axis == 'z':
        matrix = [
            [math.cos(rotation * math.pi / 180), -math.sin(rotation * math.pi / 180), 0],
            [math.sin(rotation * math.pi / 180), math.cos(rotation * math.pi / 180), 0],
            [0, 0, 1],
        ]
    return matrix

class Scanner(object):
    def __init__(self, id_):
        self.id = id_
        self.probes = []
        self.position = None

    def add_probe(self, x, y, z):
        self.probes.append((x, y, z))

    def rotate(self, rotation):
        probes = self.probes.copy()

        for axis, rot in rotation:
            for i, p in enumerate(probes):
                probes[i] = matrix_mul(rotation_matrix(axis, rot), [[p[0]], [p[1]], [p[2]]])

        return probes

    def __repr__(self):
        return '<Scanner {}>'.format(self.id)


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

    curr_scanner = None
    scanners = []
    for l in args.input:
        l = l.strip()
        if l.startswith('--- scanner'):
            if curr_scanner is not None:
                scanners.append(curr_scanner)
            scanner_id = re.sub('[^\d]+(\d+)[^\d]+', '\\1', l)
            curr_scanner = Scanner(scanner_id)
        elif len(l) > 0:
            curr_scanner.add_probe(*list(map(int, l.split(','))))
    scanners.append(curr_scanner)

    logger.info('Read %d scanners', len(scanners))

    positioned = [scanners[0]]
    positioned[0].position = (0, 0, 0)
    unpositioned = scanners[1:]
    result = len(scanners[0].probes)
    logger.info('Unpositioned: %s', unpositioned)
    while len(unpositioned) > 0:
        unpos = unpositioned.pop(0)
        logger.debug('Trying to locate %s', unpos.id)
        positioned_len = len(positioned)
        for pos in positioned:
            for rotation in itertools.product([('x', 0), ('x', 90), ('x', 180), ('x', 270)], [('y', 0), ('y', 90), ('y', 180), ('y', 270)], [('z', 0), ('z', 90), ('z', 180), ('z', 270)]):
                rotated_j = unpos.rotate(rotation)
                diffs = [(b[0] - a[0], b[1] - a[1], b[2] - a[2]) for a, b in itertools.product(pos.probes, rotated_j)]
                common = Counter(d for d in diffs).most_common(1)
                if common[0][1] >= 12:
                    logger.debug(
                        'Found offset %s at rotation %s, scanner%s-scanner%s',
                        common[0][0],
                        rotation,
                        pos.id,
                        unpos.id
                    )
                    s = Scanner(unpos.id)
                    s.probes = [(p[0] - common[0][0][0], p[1] - common[0][0][1], p[2] - common[0][0][2]) for p in rotated_j]
                    s.position = common[0][0]
                    positioned.append(s)
                    logger.info('New unique probes: %s + %s', result, len(rotated_j) - common[0][1])
                    result += len(rotated_j) - common[0][1]
                    break
            if positioned_len != len(positioned):
                logger.info('New positioned found: %d', len(positioned))
                break
        else:
            unpositioned.append(unpos)


    result = max(sum(abs(b - a) for a, b in zip(l.position, r.position)) for l, r in itertools.product(positioned, positioned))
    print ("Result is", result, file=args.output)

if __name__ == '__main__':
    main()
