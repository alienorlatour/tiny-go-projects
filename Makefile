chapters := 02-hello_world \
			03-bookworms \
			04-log \
			05-gordle \
			06-money_converter

%:
	@for chapter in $(chapters); do (cd $${chapter}; make $@); done
