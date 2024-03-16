# teampomo 🍅
A cli pomodoro timer that integrates with MS Teams to adjust status automagically 🪄 when you are in an active pomorodo session.

## Features
- CLI program that counts down a pomodoro session with customized duration.
- Creates an event in your outlook calendar to represent the pomodoro session.
- Updates your teams status message to state that you are in a pomodoro session and how long you have remaining.

## Installation

Download the latest release of the .exe file from the releases page here, or downlad and install from source.

### Building from Source
Install the latest version of golang to be able to build the project.

git clone the repository `git clone https://github.com/JoshKoiro/teampomo.git`

navigate into the `/cmd` directory and run the `go build .` command to compile the binary for your system. You can then run the program making sure to include the `start` flag to start a pomodoro session.

Alternatevly, you can use the `go run . start` command to run the application in place.

## Usage

The binary application expects the `start` flag in order to start a pomodoro. You may specify the duration of the pomodoro session by entering a numerical value like `start 35` for a 35 minute pomodoro or even `start 45sec` to you want to try to race the clock!