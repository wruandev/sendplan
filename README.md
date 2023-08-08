# Sendplan
Command Line Interface to save your plan for sending job application.

## Requirements
- Go version 1.20 (tested).

## Features
- Show all saved plans.
- Add a new plan.
- Delete an existing plan.
- Check and notify all saved plans that almost reach deadline.
- Save all plan data into json file.

## Commands
### Show all saved plans
```sh-session
$ sendplan list
```

### Add new plan
```sh-session
$ sendplan add
```

### Delete a plan
```sh-session
$ sendplan delete <id>
```

### Check and notify all near deadline plans
```sh-session
$ sendplan check
```

### Display help
```sh-session
$ sendplan --help
```

## License

MIT