projects := \
	1_1_our_first_program \
	1_3_greet \
	2_1_parlez_vous_francais \
	3_1_phrasebook \
	4_1_flags

build:
	@for project in $(projects); do (cd $${project}; make build); done

run:
	@for project in $(projects); do (cd $${project}; make run); done

test:
	@for project in $(projects); do (cd $${project}; make test); done

cover:
	@for project in $(projects); do (cd $${project}; make cover); done
