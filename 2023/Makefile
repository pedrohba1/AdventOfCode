.ONESHELL: # Applies to every targets in the file!


run:
	@echo "--- day $(DAY) solution ---"
	@echo "--- part 1 ---"
	@cd $(DAY)/part1/ && time go run main.go
	@cd ../../
	@echo ""
	@echo "--- part 2 ---"
	@cd $(DAY)/part2 && time go run main.go
	@cd ../../ 
	@echo "--- bonus ---"
	@if [ -d "./$(DAY)/bonus" ]; then \
		cd ./$(DAY)/bonus; \
		files=$$(ls **/*.go 2> /dev/null); \
		if [ -n "$$files" ]; then \
			for file in $$files; do \
				time go run $$file; \
				echo ""; \
			done; \
		else \
			echo "No Go files found in $(DAY)/bonus"; \
		fi; \
	else \
		echo "no bonus executions"; \
	fi