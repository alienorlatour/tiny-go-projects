projects := \
	progress \
	1_api \
	2_setup \
	3_create \
	progress

%:
	@for project in $(projects); do (cd $${project}; make $@); done
