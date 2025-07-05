NAME_PUSH_SWAP = push-swap
NAME_CHECKER = checker

all: $(NAME_PUSH_SWAP) $(NAME_CHECKER)

$(NAME_PUSH_SWAP):
	@go build -o $(NAME_PUSH_SWAP) ./cmd/push-swap

$(NAME_CHECKER):
	@go build -o $(NAME_CHECKER) ./cmd/checker

clean:
	@rm -f $(NAME_PUSH_SWAP) $(NAME_CHECKER)

fclean: clean

re: fclean all

test: all
	@echo "Testing push-swap with simple cases..."
	@echo "Test 1: Simple case"
	@./$(NAME_PUSH_SWAP) "2 1 3" | ./$(NAME_CHECKER) "2 1 3"
	@echo "Test 2: Medium case" 
	@./$(NAME_PUSH_SWAP) "4 67 3 87 23" | ./$(NAME_CHECKER) "4 67 3 87 23"
	@echo "Test 3: Error handling"
	@./$(NAME_PUSH_SWAP) "0 one 2 3" 2>/dev/null || echo "Error handling works"
	@echo "Test 4: Empty input"
	@./$(NAME_PUSH_SWAP) || echo "Empty input handled"
	@echo "Test 5: Duplicate check"
	@./$(NAME_PUSH_SWAP) "1 2 1" 2>/dev/null || echo "Duplicate detection works"
	@echo "Test 6: Unit tests with coverage"
	@go test -cover -count=1 ./...
	@echo "All tests completed successfully!"

.PHONY: all clean fclean re test
