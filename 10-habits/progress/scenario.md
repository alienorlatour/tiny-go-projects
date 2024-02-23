## Final Scenario

### 1. Create a habit “Write some Go code”

Request:
```bash
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"name":"Write some Go code", "weekly_frequency":3}' \
localhost:28710 \
habits.Habits/CreateHabit
```

Response:
```json
{
  "habit": {
    "id": "94c573f1-df03-45ec-97fc-8b8fc9943472",
    "name": "Write some Go code",
    "weeklyFrequency": 3
  }
}
```

### 2. Create a habit “Write some Go code”

Request:
```bash
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"name":"Read a few pages", "weekly_frequency":5}' \
localhost:28710 \
habits.Habits/CreateHabit
```

Response:
```json
{
  "habit": {
    "id": "96b72dce-7a2e-43ce-9091-0f9fc447b8a1",
    "name": "Read a few pages",
    "weeklyFrequency": 5
  }
}
```

### 3. List habits

Request:
```shell
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{}' \
localhost:28710 \
habits.Habits/ListHabits
```

Response:
```json
{
  "habits": [
    {
      "id": "94c573f1-df03-45ec-97fc-8b8fc9943472",
      "name": "Write some Go code",
      "weeklyFrequency": 3
    },
    {
      "id": "96b72dce-7a2e-43ce-9091-0f9fc447b8a1",
      "name": "Read a few pages",
      "weeklyFrequency": 5
    }
  ]
}
```

### 4. Tick habit "Write some Go code"

Request:
```shell
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"habit_id":"96b72dce-7a2e-43ce-9091-0f9fc447b8a1"}' \
localhost:28710 \
habits.Habits/TickHabit
```

Response:
```json
{
  
}
```

### 5. Get the status of "Write some Go code" habit

Request:
```shell
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"habit_id":"94c573f1-df03-45ec-97fc-8b8fc9943472"}' \
localhost:28710 \
habits.Habits/GetHabitStatus
```

Response:
```json
{
  "habit": {
    "id": "94c573f1-df03-45ec-97fc-8b8fc9943472",
    "name": "Write some Go code",
    "weeklyFrequency": 3
  },
  "ticksCount": 1
}
```

### 6. Tick habit "Read a few pages"

Request:
```shell
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"habit_id":"96b72dce-7a2e-43ce-9091-0f9fc447b8a1"}' \
localhost:28710 \
habits.Habits/TickHabit
```

Response:
```json
{
  
}
```

### 7. Get the status of "Read a few pages" habit

Request:
```shell
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"habit_id":"96b72dce-7a2e-43ce-9091-0f9fc447b8a1"}' \
localhost:28710 \
habits.Habits/GetHabitStatus
```

Response:
```json
{
  "habit": {
    "id": "96b72dce-7a2e-43ce-9091-0f9fc447b8a1",
    "name": "Read a few pages",
    "weeklyFrequency": 5
  },
  "ticksCount": 1
}
```

### 8. Tick habit "Read a few pages" with a timestamp in the past

Request:
```shell
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"habit_id":"96b72dce-7a2e-43ce-9091-0f9fc447b8a1", "timestamp": "2024-01-24T20:24:06+00:00"}' \
localhost:28710 \
habits.Habits/TickHabit
```

Response:
```json
{
  
}
```

### 9. Get the status of "Read a few pages" habit with a timestamp

Request:
```shell
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"habit_id":"96b72dce-7a2e-43ce-9091-0f9fc447b8a1", "timestamp": "2024-01-24T20:24:06+00:00"}' \
localhost:28710 \
habits.Habits/GetHabitStatus
```

Response:
```json
{
  "habit": {
    "id": "96b72dce-7a2e-43ce-9091-0f9fc447b8a1",
    "name": "Read a few pages",
    "weeklyFrequency": 5
  },
  "ticksCount": 1
}
```