import logging
import sys

import itertools

logger = logging.getLogger('day22')

BEST = sys.maxsize

class Spell(object):
    def __init__(self, name, c=0, d=0, a=0, h=0, r=0, t=0):
        self.name = name
        self.cost = c
        self.damage = d
        self.armor = a
        self.heal = h
        self.recharge = r
        self.turns = t

    def __repr__(self):
        return 'Spell<"{}" c={} d={} a={} h={} r={} t={}>'.format(
            self.name,
            self.cost,
            self.damage,
            self.armor,
            self.heal,
            self.recharge,
            self.turns,
        )

class Player(object):
    def __init__(self, name, hit=100, mana=500):
        self.name = name
        self.armor = 0
        self.damage = 0
        self.effects = []
        self.hit_points = hit
        self.mana = mana

    def __repr__(self):
        return 'Player<"{}" ({}) attack={} defense={} m={}>'.format(
            self.name,
            self.hit_points,
            self.damage,
            self.armor_points,
            self.mana
        )

    def copy(self):
        n = Player(self.name, self.hit_points, self.mana)
        n.armor = self.armor
        n.damage = self.damage
        n.effects = [[e[0], e[1]] for e in self.effects]
        return n

    @property
    def effect_names(self):
        return [e.name for e, _ in self.effects]

    @property
    def armor_points(self):
        return self.armor

    def start_turn(self):
        armor = 0
        for i, effect in enumerate(self.effects):
            logger.debug('Applying spell %s', effect)
            self.hit_points += effect[0].heal
            self.hit_points -= effect[0].damage
            self.mana += effect[0].recharge
            armor += effect[0].armor
            effect[1] -= 1
        self.armor = armor
        self.effects = filter(lambda x: x[1] > 0, self.effects)

    def add_effect(self, e):
        self.effects.append([e, e.turns])

    def defense(self, attack_points):
        self.hit_points -= max(1, attack_points - self.armor_points)

def find_best_game(me, enemy, spells, spell_path, spent=0, depth=0, hard_mode=False):
    global BEST
    logger.debug(
        'Recursive level %03d (%s) spent %03d mana. %s vs %s. Applied: %s',
        depth,
        'hard' if hard_mode else 'easy',
        spent,
        me,
        enemy,
        ', '.join(spell_path)
    )
    # Don't bother continuing computing
    if spent > BEST:
        return
    # Apply effects
    me.start_turn()
    enemy.start_turn()
    if hard_mode and depth % 2 == 0:
        me.hit_points -= 1

    if me.hit_points <= 0:
        logger.debug('I was defeated')
        return

    if enemy.hit_points <= 0:
        logger.debug('Boss defeated using %05d mana', spent)
        if spent < BEST:
            logger.info(
                'Boss defeated using %05d mana at depth %03d, new best: %s',
                spent,
                depth,
                ', '.join(spell_path),
            )
            BEST = spent
        return

    logger.debug('My current effects: %s', me.effect_names)
    logger.debug('Boss current effects: %s', enemy.effect_names)
    if depth % 2 == 0: # My turn
        for s_name, s in spells.items():
            # Cannot afford
            if s.cost > me.mana:
                continue
            # Effect already applied
            if s_name in me.effect_names or s_name in enemy.effect_names:
                continue
            n_me = me.copy()
            n_enemy = enemy.copy()
            n_me.mana -= s.cost
            # Damaging effects to enemy
            if s.turns > 0 and s.damage > 0:
                n_enemy.add_effect(s)
            # Positive effects for me
            elif s.turns > 0:
                n_me.add_effect(s)
            # Instant attack
            else:
                n_me.hit_points += s.heal
                n_enemy.hit_points -= s.damage
            logger.debug('Using spell %s', s_name)
            sp = [sn for sn in spell_path] + [s_name]
            find_best_game(n_me, n_enemy, spells, sp, spent + s.cost, depth + 1, hard_mode)
    else: # Boss turn, plain damage
        me.defense(enemy.damage)
        find_best_game(me, enemy, spells, spell_path, spent, depth + 1, hard_mode)

SPELLS = {
    'Magic Missile': Spell('Magic Missile', 53, d=4),
    'Drain': Spell('Drain', 73, d=2, h=2),
    'Shield': Spell('Shield', 113, a=7, t=6),
    'Poison': Spell('Poison', 173, d=3, t=6),
    'Recharge': Spell('Recharge', 229, r=101, t=5),
}
