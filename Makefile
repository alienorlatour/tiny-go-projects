chapters := 02-hello_world \
			03-bookworms \
			04-logger \
			05-gordle \
			06-money_converter \
			07-generic_cache \
			08-http_gordle \
			09-maze_solver \
			10-habits \
			11-templates \
			12-multiplatform


%:
	@for chapter in $(chapters); do (cd $${chapter}; make $@); done
