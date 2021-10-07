# yikes

Yikes is a cli-application to simplify the process to maintaining a list of tasks to complete.
It also has commands to help store random notes and retrive them for later use

## Run locally

```sh
git clone https://github.com/shreyaparadkar/yikes.git
# or clone your own fork

cd yikes
go install yikes
```

## Commands List

- **yikes [-h help]** - Get information about different commands

  ```sh
  yikes -h
  ```
  
- **yikes add [todo action]** - Add items to the todo list

  ```sh
  yikes add "Task description"
  ```

- **yikes edit [task no] [updated desc]** - Edit description of the tasks from the list

  ```sh
  yikes edit 3 new task info
  ```
  
- **yikes list [-i incomplete] [-c complete]** - Check the items present todo list

  ```sh
  yikes list 
  
  #to view only incomplete tasks
  yikes list -i
  
  #to view only complete tasks
  yikes list -c
  ```
  
- **yikes done [task no] [-a all]** - Check off tasks from the list

  ```sh
  yikes done 1
  
  #to mark all tasks as done
  yikes done -a
  ```

- **yikes undone [task no] [-a all]** - Undo a checked off task

  ```sh
  yikes undone 1
  
  #to mark all tasks as undone
  yikes undone -a
  ```

- **yikes delete [-a all]** - Remove tasks from the list

  ```sh
  yikes delete 1
  
  #to delete all tasks
  yikes delete -a
  ```

- **yikes notes [-t title] [-d desc] [-c create] [-r remove] [-g get] [-o open]** - Create and manage notes

  ```sh
  #get a list of all availabe notes
  yikes notes -g
  
  #create a new note with no content
  yikes notes -c -t="Title of the note"

  #create a new note with content
  yikes notes -c -t="Title of the note" -d="Description of the note"

  #create a new note and open notepad to edit it
  yikes notes -c -t="Title of the note" -o 

  #open and edit a created note
  yikes notes -t="Title of the note" -o 
  
  #delete a note
  yikes notes -r -t="Title of the note"
  ```
