### CLI ToDo Application

Simple CLI ToDo Application written in GoLang. Commands:

### Add new item to list
- go run ./ -add "Buy milk" 

| #   | Title         | Completed | Created At                 | Completed At               |
|-----|---------------|-----------|----------------------------|----------------------------|
| 0   | Buy tea       | OK        | Tue, 19 Nov 2024 16:26:33 CET | Tue, 19 Nov 2024 16:52:41 CET |
| 1   | Go for a walk | X         | Mon, 25 Nov 2024 13:46:51 CET |                            |
| 2   | Buy milk      | X         | Mon, 25 Nov 2024 13:47:04 CET |                            |
| 3   | Buy sugar     | X         | Mon, 25 Nov 2024 13:47:41 CET |                            |

### Display all items on list
- go run ./ -list

### Edit item on the list
- go run ./ -edit "2:Do something"

### Mark as done/undone
- go run ./ -toggle 2 