projects := \
	2_solver/2_2_open_maze_image \
	2_solver/2_3_add_solver \
	3_exploring/3_1_find_entrance \
	3_exploring/3_2_send_signal \
	3_exploring/3_3_record_explored_path \
	3_exploring/3_4_spin_routine \
	3_exploring/3_5_stop_listening_short \
	3_exploring/3_6_test_explore_one_routine \
	4_show_result \
	5_notify_treasure_found/5_1_track_all_routines \
	5_notify_treasure_found/5_2_send_quit_signal \
	6_visualisation/6_1_color_explored_pixels \
	6_visualisation/6_2_animate_the_exploration \
	builder

%:
	@for project in $(projects); do (cd $${project}; make $@); done
