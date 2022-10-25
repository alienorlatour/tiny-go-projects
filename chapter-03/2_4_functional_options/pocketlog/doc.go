/*
Package pocketlog exposes an API to log your work.

First, instantiate a logger with pocketlog.New, and giving it a threshold level.
Messages of lesser criticality won't be logged.

Sharing the logger is the responsibility of the caller.

The logger can be called to logf messages on three levels:
  - Debug: mostly used to debug code, follow step-by-step processes
  - Info: valuable messages providing  insights to the milestones of a process
  - Error: error messages to understand what went wrong

The New() function accepts a variety of configuration functions.
One of them lets the user define the output to which logs will be written.
It is the responsibility of the caller of this library to ensure thread safety of the writer.
*/
package pocketlog
