projects := \
	1_1_supported_levels \
	1_2_object_oriented \
	1_3_new \
	1_4_testing \
	1_5_documentation \
	2_1_first_implementation \
	2_2_adding_a_writer \
	2_3_factorisation \
	3_1_functional_options \
	3_2_testing \
	exercises/1_error_method \
	exercises/2_zero_value_enum \
	exercises/3_all_implementations \
	exercises/4_testing_at_random \
	exercises/5_logging_level \
	exercises/6_alternative_exposition \
	exercises/7_avoiding_long_messages \
	exercises/8_json_output

%:
	@for project in $(projects); do (cd $${project}; make $@); done
