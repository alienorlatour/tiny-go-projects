# Commands

## Create a habit

```shell
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"name":"read a few pages", "weekly_frequency":3}' \
localhost:28710 \
habits.Habits/CreateHabit
```


```json
{
  "habit": {
    "id": "df990e0b-5825-460d-86d2-18dcec19adeb",
    "name": "read a few pages",
    "weeklyFrequency": 3
  }
}
```

## List habits
```shell
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{}' \
localhost:28710 \
habits.Habits/ListHabits
```

```json
{
  "habits": [
    {
      "id": "df990e0b-5825-460d-86d2-18dcec19adeb",
      "name": "read a few pages",
      "weeklyFrequency": 3
    }
  ]
}
```

## Tick a habit

```shell
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"habit_id":"98ab1bbe-41d5-4ed3-8f33-e4f7bec448c8", "timestamp": "2024-01-24T20:24:06+00:00"}' \
localhost:28710 \
habits.Habits/TickHabit
```

```json
{}
```

## Get the status

```shell
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"habit_id":"98ab1bbe-41d5-4ed3-8f33-e4f7bec448c8", "timestamp": "2024-01-24T20:24:06+00:00"}' \
localhost:28710 \
habits.Habits/GetHabitStatus
```

```json
{
  "habit": {
    "id": "df990e0b-5825-460d-86d2-18dcec19adeb",
    "name": "read a few pages",
    "weeklyFrequency": 3
  },
  "ticksCount": 2
}
```
