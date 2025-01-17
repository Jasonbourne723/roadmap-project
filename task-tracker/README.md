# Task Tracker

Task tracker is a project used to track and manage your tasks. In this task, you will build a simple command line interface (CLI) to track what you need to do, what you have done, and what you are currently working on. This project will help you practice your programming skills, including working with the filesystem, handling user inputs, and building a simple CLI application.


Sample solution for the task-tracker challenge from roadmap.sh.

https://roadmap.sh/projects/task-tracker

Clone the repository and run the following command:

```
git clone -b main https://github.com/Jasonbourne723/roadmap-project.git
cd task-tracker
```

Run the following command to build and run the project:
```go
go build -o tt.exe main.go

// add 
.\tt.exe add leran-go

// update
.\tt.exe update 1 learn-c#

// del
.\tt.exe del 1

// mark
.\tt.exe mark 1 inprogress
.\tt.exe mark 1 todo
.\tt.exe mark 1 done

// list
.\tt.exe list
.\tt.exe list todo
.\tt.exe list inprogress
.\tt.exe list done

```