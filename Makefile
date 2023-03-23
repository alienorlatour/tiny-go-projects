chapters := 02-hello_world \
			03-bookworms \
			04-logger \
			05-gordle \
			06-money_converter
			chapter-07

%:
	@for chapter in $(chapters); do (cd $${chapter}; make $@); done
