# StepQA
StepQA is a CLI tool written in Go that allows users to manage Q&A items, including adding, updating, deleting, and searching for them.
It supports the ability to keep track of which steps have been completed for a particular Q&A item.

## Requirements:
* Go version 1.16 or higher
* PostgreSQL
## Installation
* Clone the repository or download the code.
* Run go mod download to download the dependencies.
* Create a PostgreSQL database.
* Setup the environment variables for the db
```
export STEPQA_DB_HOST=your_database_host
export STEPQA_DB_USERNAME=your_database_username
export STEPQA_DB_PASSWORD=your_database_password
export STEPQA_DB_NAME=your_database_name
```
* Run go build to build the executable.
* Run the StepQA executable to start using the tool.
## Usage
```
List QAs by ID - This option lists all Q&A items that have the specified IDs.
List all Q&A - This option lists all Q&A items in the database.
Search for Q&A - This option allows the user to search for Q&A items by a search term.
Add Q&A - This option allows the user to add a new Q&A item.
Update Q&A - This option allows the user to update an existing Q&A item.
Delete Q&A - This option allows the user to delete an existing Q&A item.
Toggle step completion - This option allows the user to toggle the completion of steps for a specific Q&A item.
Reset all step checkmarks - This option resets the completion status for all steps in all Q&A items.
Reset step checkmarks for a specific ID - This option resets the completion status for all steps in the Q&A item with the specified ID.
Exit - This option exits the program.
```

## Examples
Adding a new Q&A item

$ ./qa
### StepQA ###
```
Welcome to StepQA!

Choose an action:
> List QAs by ID
  List all Q&A
  Search for Q&A
  Add Q&A
  Update Q&A
  Delete Q&A
  Toggle step completion
  Reset all step checkmarks
  Reset step checkmarks for a specific ID
  Exit
```
### Add Q&A
```
Topic: Go
Subtopic: Programming
Question: What is Go?
Answer: Go is an open-source programming language that makes it easy to build simple, reliable, and efficient software.

Add a step (leave blank when done): Install Go
Add a step (leave blank when done): Write a Go program
Add a step (leave blank when done): Test the program
Add a step (leave blank when done): 

Q&A added successfully.
```

### Updating an existing Q&A item
$ ./qa
### StepQA ###
```
Welcome to StepQA!

Choose an action:
> List QAs by ID
  List all Q&A
  Search for Q&A
  Add Q&A
  Update Q&A
  Delete Q&A
  Toggle step completion
  Reset all step checkmarks
  Reset step checkmarks for a specific ID
  Exit
Update Q&A

Enter the ID of the Q&A to update: 1

Topic: Go
Subtopic: Programming
Question: What is Go?
Answer: Go is an open-source programming language that makes it easy to build simple, reliable, and efficient software.

1. [ ] Install Go
2. [ ] Write a Go program
3. [ ] Test the program
Add a step (leave blank when done): Run the program

Q&A updated successfully.
```
## SQL Database setup:

Note that the SQL script also adds some "default values" just remove them and leave the CREATE DATABASE entry.
