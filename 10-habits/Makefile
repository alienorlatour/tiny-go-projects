projects := \
	progress \
	1_api \
	2_setup \
	3_create \
	4_mocks \
	5_integration \
	progress

%:
	@for project in $(projects); do (cd $${project}; make $@); done
