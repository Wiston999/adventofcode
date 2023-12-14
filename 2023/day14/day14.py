import sys
import numpy as np

def read_array_from_file(file_path):
    with open(file_path, 'r') as f:
        lines = f.readlines()
    return np.array([list(line.strip()) for line in lines])

def tilt_north(arr):
    nrows, ncols = arr.shape
    for col in range(ncols):
        segment = arr[:, col]
        separator_indices = np.where(segment == '#')[0]
        segment_starts = np.concatenate(([0], separator_indices + 1))
        segment_ends = np.concatenate((separator_indices, [nrows]))

        for start, end in zip(segment_starts, segment_ends):
            segment_slice = segment[start:end]
            num_rocks = np.count_nonzero(segment_slice == 'O')
            segment_slice[:num_rocks] = 'O'
            segment_slice[num_rocks:] = '.'
    return arr

def tilt_west(arr):
    return np.rot90(tilt_north(np.rot90(arr, 3)), 1)

def tilt_south(arr):
    return np.rot90(tilt_north(np.rot90(arr, 2)), 2)

def tilt_east(arr):
    return np.rot90(tilt_north(np.rot90(arr, 1)), 3)

def run_cycle(arr):
    arr = tilt_north(arr)
    arr = tilt_west(arr)
    arr = tilt_south(arr)
    arr = tilt_east(arr)
    return arr

def calculate_total_load(arr):
    rock_positions = np.where(arr == 'O')
    return np.sum(arr.shape[0] - rock_positions[0])

def run_cycles(arr, cycles): # DP & Mesmorization!
    seen_states = {}
    cycle_count = 0

    while cycle_count < cycles:
        # Convert the array to a string for hashing
        state_key = arr.tobytes()  
        if state_key in seen_states:
            # Detected a repeating pattern
            cycle_period = cycle_count - seen_states[state_key]
            remaining_cycles = (cycles - cycle_count) % cycle_period
            return run_cycles(arr, remaining_cycles)

        seen_states[state_key] = cycle_count
        arr = run_cycle(arr)
        cycle_count += 1

    return arr
    
array = read_array_from_file(sys.argv[1])

#### Part 1
total_load1 = calculate_total_load(tilt_north(array))
print(f"Part 1 = {total_load1}")

#### Part 2
total_load2 = calculate_total_load(run_cycles(array, 1000000000))
print(f"Part 2 = {total_load2}")
