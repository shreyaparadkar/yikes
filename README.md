# yikes

Yikes is a cli-application to simplify the process to maintaining a list of tasks to complete.
It also has commands to help store random notes and retrive them for later use

### Commands List

- **yikes [-h help]** - Get information about different commands
- **yikes add [todo action]** - Add items to the todo list
- **yikes list [-i incomplete] [-c complete]** - Check the items present todo list
- **yikes done [task no] [-a all]** - Check off tasks from the list
- **yikes undone [task no] [-a all]** - Undo a checked off task
- **yikes delete [-a all]** - Remove tasks from the list
- **yikes notes [-t title] [-d desc] [-c create] [-r remove] [-g get] [-o open]** - Create and manage notes

### Run locally
```
git clone https://github.com/shreyaparadkar/yikes.git
# or clone your own fork

cd yikes
go install yikes
```
