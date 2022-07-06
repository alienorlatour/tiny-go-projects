chapters := chapter-02 \
			chapter-03 \
			chapter-04

build:
	@for chapter in $(chapters); do (cd $${chapter}; make build); done

run:
	@for chapter in $(chapters); do (cd $${chapter}; make run); done

test:
	@for chapter in $(chapters); do (cd $${chapter}; make test); done

cover:
	@for chapter in $(chapters); do (cd $${chapter}; make cover); done
