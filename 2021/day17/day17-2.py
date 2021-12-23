from __future__ import print_function
import argparse
import logging
import sys

import re

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

def test_trajectory(traj, target):
    current = [0, 0]
    trajectory = traj.copy()
    path = []
    while True:
        previous = current.copy()
        path.append(current.copy())

        current[0] += trajectory[0]
        current[1] += trajectory[1]

        trajectory[0] += 1 if trajectory[0] < 0 else -1 if trajectory[0] > 0 else 0
        trajectory[1] -= 1

        logger.debug('Trajectory %s to target %s at %s', trajectory, target, current)
        if target['x_start'] <= current[0] <= target['x_end'] and target['y_start'] <= current[1] <= target['y_end']:
            return True, path
        elif current[0] > target['x_end'] or current[1] < target['y_start']: # overshot
            logger.debug('Trajectory %s overshot', traj)
            break
        elif current[0] == previous[0] and current[1] < target['y_start']: # Stuck at X
            logger.debug('Trajectory %s stuck at x', traj)
            break
    return False, path

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('-y', '--max-y', type=int, default=1000,
            help='Max Y coordinate for trajectory')
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

    regexp = re.match(
        r'target area: x=(?P<x_start>-?\d+)\.\.(?P<x_end>-?\d+), y=(?P<y_start>-?\d+)\.\.(?P<y_end>-?\d+)',
        args.input.read()
    )
    target = {k: int(v) for k, v in regexp.groupdict().items()}
    logger.info('Target area is: %s', target)
    success_trajectories = []

    for x in range(1, target['x_end'] + 1): # higher x will overpass target in 1 step
        y = 1
        for y in range(target['y_start'], args.max_y):
            hit, path = test_trajectory([x, y], target)
            logger.info('Tested trajectory (%s, %s) --> %s', x, y, hit)
            if hit:
                logger.info('Found hit trajectory: (%s, %s)', x, y)
                success_trajectories.append([x, y])


    result = len(success_trajectories)
    logger.debug('Trajectories are: %s', success_trajectories)
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
