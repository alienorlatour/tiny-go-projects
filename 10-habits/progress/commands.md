grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"name":"clean the kitchen"}' \
localhost:28710 \
habits.Habits/TickHabit
