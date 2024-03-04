projects := \
	1_api \
	2_setup \
	3_create \
	4_mocks \
	5_integration \
	7_track

%:
	@for project in $(projects); do (cd $${project}; make $@); done
