chapters := chapter-02 \
			chapter-03 \
			chapter-04 \
			chapter-07

%:
	@for chapter in $(chapters); do (cd $${chapter}; make $@); done
