from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

def increase_energy(octopuses, position, bumped):
    x, y = position
    result = 1

    bumped.add(position)

    for n_x in range(x-1, x+2):
        for n_y in range(y-1, y+2):
            if (n_x, n_y) in octopuses and (n_x, n_y) not in bumped:
                octopuses[(n_x, n_y)] += 1
                if octopuses[(n_x, n_y)] > 9:
                    logger.debug('Octopus %s flashed by %s reflect', (n_x, n_y), (x, y))
                    bumped.add((n_x, n_y))
                    flashed, bumped = increase_energy(octopuses, (n_x, n_y), bumped)
                    result += flashed

    return result, bumped

def print_octopuses(octopuses, m_x=10, m_y=10):
    result = ''
    for x in range(m_x):
        for y in range(m_y):
            result += '{} '.format(octopuses[(x, y)])
        result += '\n'
    return result

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
    octopuses = {(x,y): int(e) for x, l in enumerate(args.input) for y, e in enumerate(l.strip())}

    logger.debug('Read octopuses energy: %s', octopuses)

    while True:
        octopuses = {(x, y): e + 1 for (x, y), e in octopuses.items()}
        # These cannot flash again
        flashed = set((x, y) for (x, y), e in octopuses.items() if e > 9)
        logger.debug('Octopuses at step %s before recursivity', result + 1)
        logger.debug('\n'+print_octopuses(octopuses))
        for x, y in flashed.copy():
            logger.info('Octopus %s flashed at step %03d' , (x, y), result + 1)
            new_flashes, flashed = increase_energy(octopuses, (x, y), flashed)
            logger.debug('%03d octopuses flashed too: %s', new_flashes, flashed)

        octopuses = {(x, y): e if e <= 9 else 0 for (x, y), e in octopuses.items()}

        logger.debug('Octopuses at step %s after recursivity', result + 1)
        logger.debug('\n'+print_octopuses(octopuses))

        result += 1
        if len(flashed) == len(octopuses):
            break

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
