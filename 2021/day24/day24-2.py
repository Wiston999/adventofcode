from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

class ALU(object):
    def __init__(self, steps, inputs):
        self.steps = steps
        self.registers = {
            'w': 0,
            'x': 0,
            'y': 0,
            'z': 0,
        }
        self.inputs = inputs

    def set_inputs(self, inputs):
        self.inputs = inputs

    def reset(self):
        self.registers = {
            'w': 0,
            'x': 0,
            'y': 0,
            'z': 0,
        }
        self.inputs = []

    def value(self, v):
        return self.registers[v]

    def compute(self):
        for s in map(str.strip, self.steps):
            splitted = s.split(' ')
            logger.debug('Operation %s on state %s', splitted, self.registers)
            if splitted[0] == 'inp':
                v = self.inputs.pop(0)
                reg = splitted[-1]
                self.registers[reg] = int(v)
            if splitted[0] == 'add':
                reg1 = splitted[1]
                reg2 = splitted[2]
                if reg2 in 'wxyz':
                    self.registers[reg1] += self.registers[reg2]
                else:
                    self.registers[reg1] += int(reg2)
            if splitted[0] == 'mul':
                reg1 = splitted[1]
                reg2 = splitted[2]
                if reg2 in 'wxyz':
                    self.registers[reg1] *= self.registers[reg2]
                else:
                    self.registers[reg1] *= int(reg2)
            if splitted[0] == 'div':
                reg1 = splitted[1]
                reg2 = splitted[2]
                if reg2 in 'wxyz':
                    self.registers[reg1] = self.registers[reg1] // self.registers[reg2]
                else:
                    self.registers[reg1] = self.registers[reg1] // int(reg2)
            if splitted[0] == 'mod':
                reg1 = splitted[1]
                reg2 = splitted[2]
                if reg2 in 'wxyz':
                    self.registers[reg1] = self.registers[reg1] % self.registers[reg2]
                else:
                    self.registers[reg1] = self.registers[reg1] % int(reg2)
            if splitted[0] == 'eql':
                reg1 = splitted[1]
                reg2 = splitted[2]
                if reg2 in 'wxyz':
                    self.registers[reg1] = int(self.registers[reg1] == self.registers[reg2])
                else:
                    self.registers[reg1] = int(self.registers[reg1] == int(reg2))
        return self.registers

def z_trim(alu):
    for step in alu.steps:
        if step.startswith('add x -'):
            return abs(int(step.split(' ')[-1]))

    return None

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('--start', type=int, default=99999999999999,
            help='Amount of numbers to check')
    arg_parser.add_argument('--limit', type=int, default=1000000,
            help='Amount of numbers to check')
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

    lines = args.input.readlines()
    inputs = [i for i, l in enumerate(lines) if l.startswith('inp')] + [len(lines)]

    alus = [ALU(lines[inputs[i]:inputs[i+1]], []) for i in range(14)]
    computes = {}
    for w in range(1, 10):
        alus[0].reset()
        alus[0].set_inputs([w])
        alus[0].compute()
        computes[alus[0].value('z')] = [w]

    for i, alu in enumerate(alus[1:]):
        logger.info('Computing value: %s (%03d)', i+1, len(computes))
        current = computes.copy()
        computes = {}
        z_target = z_trim(alu)
        for w in range(1, 10):
            for z in current:
                if z_target is not None and (z % 26) - z_target != w:
                    continue
                alu.reset()
                alu.set_inputs([w])
                alu.registers = {'x': 0, 'y': 0, 'z': z, 'w': w}
                alu.compute()
                if alu.value('z') not in computes:
                    computes[alu.value('z')] = current[z] + [w]

    result = ''.join(map(str, computes[0]))
    print ("Result is", result, file=args.output)

if __name__ == '__main__':
    main()
