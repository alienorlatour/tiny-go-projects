projects := \
	progress \
	1_api \
	2_setup \
	3_create

%:
	@for project in $(projects); do (cd $${project}; make $@); done
