projects := \
	1_1_basic_main \
	1_2_read_input \
	1_3_isolate_validate \
	1_4_check_for_victory \
	2_1_hints \
	2_2_benchmarking_stringbuilider \
	3_corpus \
	4_config_options

%:
	@for project in $(projects); do (cd $${project}; make $@); done
