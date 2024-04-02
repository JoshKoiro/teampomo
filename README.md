# teampomo 🍅
A cli pomodoro timer that integrates with MS Teams to adjust status automagically 🎉 when you are in an active pomorodo session.

## Features
- CLI program that counts down a pomodoro session with customized duration.
- Creates an event in your outlook calendar to represent the pomodoro session with the ability to name the event using the `--task` flag.
- Updates your teams status message to state that you are in a pomodoro session and how long you have remaining.
- Ability to show or hide what you are working on in your status message.

## Installation

Download the latest release of the .exe file from the releases page (coming soon), or download and install from source.

### Building from Source

Install the latest version of golang to be able to build the project.

git clone the repository `git clone https://github.com/JoshKoiro/teampomo.git`

navigate into the `/cmd` directory and run the `go build .` command to compile the binary for your system. You can then run the program making sure to include the `start` flag to start a pomodoro session.

Alternatevly, you can use the `go run . start` command to run the application in place.

## Usage

### Microsoft API Key

You must provide a Microsoft Graph API key in the same folder as the executeable. When the key is aquired simply copy it into a text file titled `key.txt` and place it next to the executeable file.

### Specifying time

When running the command `start` the duration of the pomodoro will default to 25 min, however this can be customized by entering a value specifying the number minutes you would like the pomorodo session to last.
```
./teampomo.exe start 35 # run Pomodoro for 35 minutes.
```
Alternatively, you can specify the suffix `sec` after the numerical value to define the number of seconds you would like the timer to last. This is good for testing functionality as well as providing some alternate use cases for the program.
```
./teampomo.exe start 35sec # race the timer by running for only 35 seconds
```
Please note that when specifying a duration in seconds, it will not create a calendar event. A calendar event is only created when specifying the duration in minutes without the `sec` suffix.

### --task flag

You are able to note the task that you are working on for the duration of the pomodoro by using the `--task` flag with the name of the task as a single word with no spaces, or in quotes following the flag.
```
./teampomo.exe start --task "cleaning up email"
```

This command will provide a description of the task in the calendar event like this:

**Pomodoro - cleaning up email"**

### --public flag

If a task is provided in the command, you are able to add a `--public` flag to allow the task description to be published to your status message. Without specifying the `--public` flag, any task name that is entered using the `--task` flag is hidden from the Teams status message and will only display "In a Pomodoro session for the next x minutes".