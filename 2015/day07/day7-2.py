from __future__ import print_function
import argparse
import logging
import sys

import re
import functools

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

class WireMap(object):
    def __init__(self):
        self.data = {
            '0': Wire('0', [0], 'identity'),
            '1': Wire('1', [1], 'identity'),
        }

    def add_wire(self, data):
        op, output = data.split(' -> ')
        output = output.strip()

        if re.match('^\d+$', op):
            operation = 'identity'
            inputs = [int(op)]
            self.data[op] = Wire(op, [int(op)], 'identity')
        elif re.match('^\w+$', op):
            operation = 'wire-identity'
            inputs = [op]
        elif re.match('^\w+ AND \w+', op):
            operation = 'and'
            inputs = op.split(' AND ')
        elif re.match('^\w+ OR \w+', op):
            operation = 'or'
            inputs = op.split(' OR ')
        elif re.match('^NOT \w+', op):
            operation = 'not'
            inputs = [op[4:]]
        elif re.match('^\w+ RSHIFT \d+', op):
            operation = 'rshift ' + op.split(' RSHIFT ')[-1]
            inputs = [op.split(' RSHIFT')[0]]
        elif re.match('^\w+ LSHIFT \d+', op):
            operation = 'lshift ' + op.split(' LSHIFT ')[-1]
            inputs = [op.split(' LSHIFT')[0]]
        else:
            logger.warning('Unknown operation: %s', data)


        if operation != 'identity':
            for i in inputs:
                if i not in self.data:
                    self.data[i] = Wire(i, [], None)

            inputs = [self.data[i] for i in inputs]


        logger.debug('Adding wire %s with inputs %s and operation %s', output, inputs, operation)
        if output not in self.data:
            self.data[output] = Wire(output, inputs, operation)
        else:
            self.data[output].operation = operation
            self.data[output].add_input(inputs)

    def wire_dependencies(self, name):
        pending = [self.data[name]]
        visited = set()
        while pending:
            p = pending.pop(0)
            yield self.data[p.name].inputs
            if self.data[p.name].operation != 'identity' and p.name not in visited:
                pending.extend(self.data[p.name].inputs)
                visited.add(p.name)

    def wire_output(self, name):
        return self.data[name].output

class Wire(object):
    def __init__(self, name, inputs, operation):
        self.name = name
        self.inputs = inputs
        self.operation = operation
        self.cache = None

    def add_input(self, inputs):
        self.inputs.extend(inputs)

    @property
    def output(self):
        if self.cache:
            return self.cache

        output = None
        logger.debug('Operation on %s', self)
        if self.operation == 'identity':
            output = self.inputs[0]
        elif self.operation == 'wire-identity':
            output = self.inputs[0].output
        elif self.operation == 'and':
            output = self.inputs[0].output & self.inputs[1].output
        elif self.operation == 'or':
            output = self.inputs[0].output | self.inputs[1].output
        elif self.operation == 'not':
            output = ~self.inputs[0].output
        elif self.operation.startswith('rshift'):
            output = self.inputs[0].output >> int(self.operation.split(' ')[-1])
        elif self.operation.startswith('lshift'):
            output = self.inputs[0].output << int(self.operation.split(' ')[-1])
        else:
            logger.error('Unknown operation %s', self.operation)

        self.cache = output
        return output

    def __repr__(self):
        return 'Wire<name: {} operation: {} inputs: [{}]>'.format(
            self.name,
            self.operation,
            ', '.join(i.name if isinstance(i, Wire) else str(i) for i in self.inputs)
        )

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('wire', type=str,
            help='Wire to compute its output')
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

    wm = WireMap()
    for l in args.input:
        wm.add_wire(l.strip())

    for w in sorted(wm.data.keys()):
        print(w, '->', wm.data[w], file=args.output)
    dependencies = list(wm.wire_dependencies(args.wire))
    logger.info('Computing value of %s that depends on %s (%s)', args.wire, dependencies, len(dependencies))
    wire_a_value = wm.wire_output('a')
    for n, w in wm.data.items():
        if w.name != 'b':
            w.cache = None
        else:
            w.cache = wire_a_value

    result = wm.wire_output('a')

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
