from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

class Cave(object):
    def __init__(self, name):
        self.name = name
        self.connections = []

    @property
    def small(self):
        return self.name == self.name.lower()

    @property
    def large(self):
        return not self.small

    @property
    def start(self):
        return self.name == 'start'

    @property
    def end(self):
        return self.name == 'end'

    def add_connection(self, connection):
        self.connections.append(connection)

    def __repr__(self):
        return '{} --> [{}]'.format(self.name, ','.join(c.name for c in self.connections))

def pathfinder(caves, current_c, visited):
    paths = []
    if current_c == 'end':
        logger.info('Found path to end: %s', visited)
        return visited + [current_c]

    if current_c in visited and caves[current_c].small:
        impossible = False
        if visited.count(current_c) >= 2:
            impossible = True

        if any(visited.count(c) >= 2 for c in visited if caves[c].small):
            impossible = True

        if impossible:
            logger.debug('Found impossible path: %s -> %s', visited, current_c)
            return [None]

    visited.append(current_c)
    for c in caves[current_c].connections:
        # Can't go back to start
        if c.name == 'start':
            continue
        path = pathfinder(caves, c.name, visited.copy())
        logger.debug('Temporal path found: %s', path)
        if len(path) > 1 and path[-1] == 'end':
            logger.info('Added path: %s', path)
            logger.info('Visited was: %s', visited)
            paths.extend(path)
    return paths


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

    caves = {}
    for l in args.input:
        c1, c2 = l.strip().split('-')
        if c1 not in caves:
            caves[c1] = Cave(c1)
        if c2 not in caves:
            caves[c2] = Cave(c2)

        caves[c1].add_connection(caves[c2])
        caves[c2].add_connection(caves[c1])

    logger.info('Cave system: %s', '\n'.join(str(c) for c in caves.values()))


    paths = pathfinder(caves, 'start', [])
    logger.info('Found %d paths: %s', paths.count('end'), paths)
    result = paths.count('end')
    print ("Result is", result, file=args.output)

if __name__ == '__main__':
    main()
