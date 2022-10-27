# birthday-reminder

## about

REST API to remind one of one's friends' birthdays. Currently only able to store said birthdays.

## usage

Example run:

```bash
go run main.go
curl -LiX localhost:8080/person/ # get all persons
curl -LiX localhost:8080/person/foo # get foo (see model/person.go; birthday is represented here as a string of the form YYYY-MM-DD)
curl -LiX localhost:8080/person/foo/birthday # get foo's birthday
```

## todo

- golanci-lint
- implement scheduler for setting reminders
- deploy to Lambda?
