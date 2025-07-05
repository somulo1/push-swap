# Push-Swap

A sorting algorithm project that implements a non-comparative sorting algorithm using two stacks and a limited set of operations.

## Description

Push-Swap consists of two programs:

1. **push-swap**: Calculates and displays the smallest program using push-swap instruction language that sorts integer arguments received.

2. **checker**: Takes integer arguments and reads instructions on standard input. Executes them and displays "OK" if integers are sorted, otherwise "KO".

## Available Operations

- `pa`: push the top first element of stack b to stack a
- `pb`: push the top first element of stack a to stack b
- `sa`: swap first 2 elements of stack a
- `sb`: swap first 2 elements of stack b
- `ss`: execute sa and sb
- `ra`: rotate stack a (shift up all elements by 1)
- `rb`: rotate stack b
- `rr`: execute ra and rb
- `rra`: reverse rotate a (shift down all elements by 1)
- `rrb`: reverse rotate b
- `rrr`: execute rra and rrb

## Building

```bash
make
# or
make all
```

This will create two executables: `push-swap` and `checker`.

## Usage

### Push-Swap

```bash
./push-swap "4 67 3 87 23"
./push-swap "2 1 3 6 5 8"
```

### Checker

```bash
echo -e "pb\npb\nra\nsa\nrrr\npa\npa" | ./checker "2 1 3 6 5 8"
./push-swap "4 67 3 87 23" | ./checker "4 67 3 87 23"
```

## Testing

```bash
make test
```

## Algorithm

The implementation uses different sorting strategies based on stack size:

- **2-3 elements**: Direct comparison and swapping
- **4-5 elements**: Move smallest elements to stack b, sort remaining, then merge back
- **6+ elements**: Radix sort using binary representation

## Error Handling

The programs handle the following errors:
- Non-integer arguments
- Duplicate numbers
- Invalid operations (checker only)

Errors are displayed as "Error" followed by a newline on stderr.
