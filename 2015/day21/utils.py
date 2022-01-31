import logging

import itertools

logger = logging.getLogger('day21')

class Item(object):
    def __init__(self, c=0, d=0, a=0):
        self.cost = c
        self.damage = d
        self.armor = a

    def __repr__(self):
        return 'Item<cost={} damage={} defense={}>'.format(
            self.cost,
            self.damage,
            self.armor
        )

class Weapon(Item): pass
class Armor(Item): pass
class Ring(Item): pass

class Player(object):
    def __init__(self, name, hit=100):
        self.name = name
        self.weapon = None
        self.armor = None
        self.rings = []
        self.hit_points = hit

    def copy(self):
        n = Player(self.name, self.hit_points)
        n.weapon = self.weapon
        n.armor = self.armor
        n.rings = self.rings
        return n

    def set_weapon(self, w):
        self.weapon = w

    def set_armor(self, a):
        self.armor = a

    def set_rings(self, r):
        self.rings = r

    def add_ring(self, r):
        if len(self.rings) < 2:
            self.rings.append(r)
            return True
        return False

    @property
    def equipment_value(self):
        v = 0
        if self.weapon:
            v += self.weapon.cost
        if self.armor:
            v += self.armor.cost
        v += sum(r.cost for r in self.rings)
        return v

    @property
    def armor_points(self):
        a = 0
        if self.armor:
            a += self.armor.armor
        a += sum(r.armor for r in self.rings)
        return a

    @property
    def attack_points(self):
        a = 0
        if self.weapon:
            a += self.weapon.damage
        a += sum(r.damage for r in self.rings)
        return a

    def defense(self, attack_points):
        self.hit_points -= max(1, attack_points - self.armor_points)

    def __repr__(self):
        return 'Player<"{}" ({}) a={} d={}>'.format(
            self.name,
            self.hit_points,
            self.attack_points,
            self.armor_points
        )

def game(p1, p2):
    turn = 0
    while p1.hit_points > 0 and p2.hit_points > 0:
        attacker = p1 if turn % 2 == 0 else p2
        defender = p2 if turn % 2 == 0 else p1
        logger.debug(
            '%s attacks %s (%03d) with %s attack points',
            attacker.name,
            defender.name,
            defender.hit_points,
            attacker.attack_points
        )
        defender.defense(attacker.attack_points)
        logger.debug('%s has %03d hit points', defender.name, defender.hit_points)
        turn += 1
    return p1 if p1.hit_points > p2.hit_points else p2

def generate_equipment(weapons, armors, rings, reverse=False):
    combinations = list(itertools.product(
        weapons.keys(),
        [None] + armors.keys(),
        [None] + rings.keys(),
        [None] + rings.keys(),
    ))
    logger.debug('Generated %s equipment combinations', len(combinations))

    unsorted_combs = []
    for c in combinations:
        logger.debug('Combination generated: %s', c)
        if c[2] == c[3] and c[2] is not None: # Same ring cannot be repeated
            continue
        w = weapons[c[0]]
        a = armors[c[1]] if c[1] else None
        r = filter(lambda x: x is not None, (
            rings[c[2]] if c[2] else None,
            rings[c[3]] if c[3] else None
        ))
        v = w.cost + (a.cost if a else 0) + sum(ri.cost for ri in r)
        unsorted_combs.append((w, a, r, v))

    for c in sorted(unsorted_combs, key=lambda x: x[3], reverse=reverse):
        logger.debug('Yielding %s', c)
        yield (c[0], c[1], c[2])
