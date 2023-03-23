chapters := chapter-02 \
			chapter-03 \
			chapter-04 \
			chapter-05 \
			chapter-06

%:
	@for chapter in $(chapters); do (cd $${chapter}; make $@); done
