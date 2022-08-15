projects := \
	1_1_supported_levels \
	1_2_object_oriented \
	1_3_new \
	1_4_testing \
	2_1_first_implementation \
	final \
	exercises/1_error_method \
	exercises/2_zero_value_enum \
	exercises/3_all_implementations

%:
	@for project in $(projects); do (cd $${project}; make $@); done
