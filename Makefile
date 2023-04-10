chapters := 02-hello_world \
			03-bookworms_digest \
			04-log_story \
			05-gordle \
			06-money_converter

%:
	@for chapter in $(chapters); do (cd $${chapter}; make $@); done
