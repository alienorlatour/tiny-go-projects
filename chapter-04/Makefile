projects := \
	pre-refactoring \
	final

build:
	@for project in $(projects); do (cd $${project}; make build); done

run:
	@for project in $(projects); do (cd $${project}; make run); done

test:
	@for project in $(projects); do (cd $${project}; make test); done

cover:
	@for project in $(projects); do (cd $${project}; make cover); done
