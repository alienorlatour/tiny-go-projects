chapters := chapter-01 \
			chapter-02 \
			chapter-03

build:
	@for chapter in $(chapters); do (cd $${chapter}; make build); done

run:
	@for chapter in $(chapters); do (cd $${chapter}; make run); done

test:
	@for chapter in $(chapters); do (cd $${chapter}; make run); done

cover:
	@for chapter in $(chapters); do (cd $${chapter}; make cover); done
