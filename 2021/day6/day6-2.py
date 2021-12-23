from __future__ import print_function
import argparse
import logging
import sys
import time

from queue import LifoQueue
import itertools
import multiprocessing
from multiprocessing.managers import SyncManager

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

class PQManager(SyncManager):
    pass
PQManager.register("LifoQueue", LifoQueue)

def Manager():
    m = PQManager()
    m.start()
    return m

class Worker(multiprocessing.Process):
    MAX_LEN = 1024 * 1024
    def __init__(self, wq, rq, days):
        self.work_q = wq
        self.result_q = rq
        self.max_days = days
        super().__init__()

    def run(self):
        while True:
            task = self.work_q.get()
            if task is None:
                logger.info('Exiting process')
                self.result_q.put(None)
                self.work_q.task_done()
                break

            (generation, fishes) = task
            logger.debug('Received task for generation %d with %d fishes', generation, len(fishes))

            for d in range(generation, self.max_days):
                for i in range(len(fishes)):
                    fishes[i] = fishes[i] - 1
                    if fishes[i] == -1:
                        fishes[i] = 6
                        fishes.append(8)

                if len(fishes) > self.MAX_LEN:
                    logger.debug('Splitting fishes population at %s size', len(fishes))
                    self.work_q.put((d + 1, fishes[:self.MAX_LEN//2]))
                    self.work_q.put((d + 1, fishes[self.MAX_LEN//2:]))
                    break
            else: # Finished for loop without break, means we reached max_days
                logger.debug('Finished computing generation split with result %d', len(fishes))
                self.result_q.put(len(fishes))

            self.work_q.task_done()

def get_fish(n, days, paralellism):
    result = 0
    processes = []
    m = Manager()
    result_queue = multiprocessing.Queue()
    work_queue = m.LifoQueue()

    for i in range(paralellism):
        logger.debug('Notifiying process %02d to finish', i)
        work_queue.put(None)

    work_queue.put((0, [n]))

    for i in range(paralellism):
        logger.info('Starting process %02d', i)
        p = Worker(work_queue, result_queue, days)
        p.start()
        processes.append(p)
        if i == 0:
            time.sleep(10) # Give first worker enough time to fill the LIFO so None is actually get at the end

    for i in range(paralellism):
        logger.debug('Notifiying process %02d to finish', i)
        work_queue.put(None)

    finish_signals = 0
    while True:
        r = result_queue.get()
        if r is None:
            finish_signals += 1
            if finish_signals == paralellism: # Finished with everything
                break
        else:
            result += r
            logger.info('Received result %s, total is %s', r, result)

    for p in processes:
        p.terminate()
    logger.info('Total value for fish %d: %d', n, result)

    return result

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('-p', '--paralellism', type=int, default=multiprocessing.cpu_count(),
            help='Number of processes to be used')
    arg_parser.add_argument('-d', '--days', type=int, default=80,
            help='Number of days for the simulation')
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

    fishes = list(map(int, args.input.read().split(',')))

    fishes_value = [get_fish(i, args.days, args.paralellism) for i in range(7)]

    logger.info('Got fish value for %s days: %s', args.days, fishes_value)

    result = sum(v * fishes.count(i) for i, v in enumerate(fishes_value))

    print ("Result is", result, file=args.output)

if __name__ == '__main__':
    main()
