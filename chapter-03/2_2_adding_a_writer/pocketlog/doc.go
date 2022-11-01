/*
Package pocketlog exposes an API to log your work.

First, instantiate a logger with pocketlog.New, and giving it a threshold level and an output writer.
Messages of lesser criticality won't be logged.

Sharing the logger is the responsibility of the caller.

The logger can be called to log messages on three levels:
  - Debug: mostly used to debug code, follow step-by-step processes
  - Info: valuable messages providing  insights to the milestones of a process
  - Error: error messages to understand what went wrong

It is the responsibility of the caller of this library to ensure thread safety of the writer.
*/
package pocketlog
