from __future__ import print_function
import argparse
import logging
import sys

import itertools

from utils import *

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger('day21')

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
    weapons = {
        'dagger': Weapon(8, 4, 0),
        'shortsword': Weapon(10, 5, 0),
        'warhammer': Weapon(25, 6, 0),
        'longsword': Weapon(40, 7, 0),
        'greataxe': Weapon(74, 8, 0),
    }
    armors = {
        'leather': Armor(13, 0, 1),
        'chainmail': Armor(31, 0, 2),
        'splintmail': Armor(53, 0, 3),
        'bandedmail': Armor(75, 0, 4),
        'platemail': Armor(102, 0, 5),
    }
    rings = {
        'damage+1': Ring(25, 1, 0),
        'damage+2': Ring(50, 2, 0),
        'damage+3': Ring(100, 3, 0),
        'defense+1': Ring(20, 0, 1),
        'defense+2': Ring(40, 0, 2),
        'defense+3': Ring(80, 0, 3),
    }

    me = Player('me')
    base_enemy = Player('enemy')

    for l in args.input:
        if 'Hit Points' in l:
            base_enemy.hit_points = int(l.split(' ')[-1])
        if 'Damage' in l:
            base_enemy.set_weapon(Weapon(d=int(l.split(' ')[-1])))
        if 'Armor' in l:
            base_enemy.set_armor(Armor(a=int(l.split(' ')[-1])))

    for weapon, armor, ring in generate_equipment(weapons, armors, rings):
        me = Player('me')
        enemy = base_enemy.copy()
        me.set_weapon(weapon)
        me.set_armor(armor)
        me.set_rings(ring)
        logger.info('Playing %s vs %s', me, enemy)
        if game(me, enemy) is me:
            result = me.equipment_value
            logger.info('Won the game')
            break

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
