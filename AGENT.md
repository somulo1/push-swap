# Push-Swap Project Notes

## Commands
- Build: `make all`
- Test: `make test`
- Coverage: `go test -cover`
- Run tests: `go test ./...`

## Current Issues
- ❌ Algorithm produces too many instructions for 100 numbers (need < 700, getting ~1000+)
- ❌ Some 6-element cases may still fail the <9 instruction requirement

## Algorithm Performance Requirements
- 2-3 elements: ≤ 3 instructions
- 5 elements: ≤ 12 instructions  
- 100 elements: ≤ 700 instructions
- 500 elements: ≤ 5500 instructions

## Standard Push-Swap Algorithm
The project requires implementing an efficient sorting algorithm using two stacks.
Most successful implementations use:
1. Simple algorithms for ≤5 elements
2. Radix sort for larger arrays
3. Bit manipulation for optimal performance

## Current Status
- ✅ All functional requirements met
- ✅ Error handling complete
- ✅ Test coverage >70%
- ❌ Performance optimization needed
