chapters := chapter-02 \
			chapter-03 \
			chapter-04

%:
	@for chapter in $(chapters); do (cd $${chapter}; make $@); done
