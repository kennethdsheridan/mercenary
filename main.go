package main

// The import block imports multiple packages that are required in this application.
import (
	"fmt"                                                                // Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	"github.com/mitchellh/go-ps"                                         // This package provides functions to find and list processes.
	"gitlab.servicenow.net/cce/hweng/hardware-test/mercenary.git/errs"   // Custom package for managing application-specific errors.
	"gitlab.servicenow.net/cce/hweng/hardware-test/mercenary.git/logger" // Custom package for logging.
	"go.uber.org/zap"                                                    // Package zap provides fast, structured, leveled logging.
	"os"                                                                 // Package os provides a platform-independent interface to operating system functionality.
	"strconv"                                                            // Package strconv implements conversions to and from string representations of basic data types.
	"syscall"                                                            // Package syscall contains an interface to the low-level operating system primitives.
	"time"                                                               // Package time provides functionality for measuring and displaying time.
)

// The entry point of the application.
func main() {
	// Remove or truncate the log file before starting the application
	err := os.Truncate("appLog.log", 0)
	if err != nil {
		fmt.Println("Error removing the log file: " + err.Error())
		return
	}
	// stores all the command-line arguments passed. If the number of arguments is not 2,
	// it logs an error message and exits the program.
	if len(os.Args) != 2 {
		logger.Error("Usage: ./mercenary <pid> or ./mercenary <process name>")
		os.Exit(1)
	}

	// The target process to be terminated, it can be a PID or a process name.
	target := os.Args[1]
	// Try to convert the target into an integer, which should be the PID.
	pid, err := strconv.Atoi(target)

	var process ps.Process // A variable to hold the process to be terminated.
	// If target can be converted into a PID, it means it is a PID.
	if err == nil {
		// Then, try to find the process by its PID.
		process, err = ps.FindProcess(pid)
		if err != nil {
			// If an error occurs while finding the process, it logs the error and exits the program.
			logErrorAndExit(err, errs.MercError{Msg: fmt.Sprintf("Failed to find process with PID '%d'.", pid)})
		}
	} else {
		// If the target cannot be converted into a PID, it means it is a process name.
		// In this case, we get all the running processes.
		processes, err := ps.Processes()
		if err != nil {
			// If an error occurs while getting the process list, it logs the error and exits the program.
			logErrorAndExit(err, errs.MercError{Msg: "Failed to get process list."})
		}

		// Loop through the processes to find the one with the given name.
		for _, proc := range processes {
			if proc.Executable() == target {
				process = proc
				break
			}
		}

		// If no process is found with the given name, it logs a message and exits the program.
		if process == nil {
			logger.Info(fmt.Sprintf("No process found with the name '%s'.", target))
			os.Exit(1)
		}
	}

	// After finding the process, try to terminate it gracefully by sending a SIGTERM signal.
	logger.Info(fmt.Sprintf("Sending SIGTERM to process with PID '%d'...", process.Pid()))
	if err := syscall.Kill(process.Pid(), syscall.SIGTERM); err != nil {
		// If an error occurs while sending the signal, it logs the error and exits the program.
		logErrorAndExit(err, errs.MercError{Msg: "Failed to send SIGTERM."})
	}

	// After sending the SIGTERM signal, it waits for 6 seconds for the process to terminate.
	time.Sleep(6 * time.Second)

	// After waiting, it checks if the process is still running.
	process, err = ps.FindProcess(process.Pid())
	if err == nil && process != nil {
		// If the process is still running, it tries to kill it by sending a SIGKILL signal.
		logger.Info(fmt.Sprintf("Process with PID '%d' still running. Sending SIGKILL...", process.Pid()))
		if err := syscall.Kill(process.Pid(), syscall.SIGKILL); err != nil {
			// If an error occurs while sending the signal, it logs the error and exits the program.
			logErrorAndExit(err, errs.MercError{Msg: "Failed to send SIGKILL."})
		} else {
			logger.Info(fmt.Sprintf("Process with PID '%d' was killed.", process.Pid()))
		}
	} else {
		// If the process is not running, it means it has been terminated gracefully.
		logger.Info(fmt.Sprintf("Process with PID '%d' has been terminated gracefully.", process.Pid()))
	}
}

// logErrorAndExit is a function that takes an error and a custom application error as arguments. It logs the error and exits the program.
// This function is useful to avoid repetitive error handling code.
func logErrorAndExit(err error, appErr errs.MercError) {
	// Log the error in the console and in the log file (appLog.log).
	logger.Error(err.Error(), zap.Error(&appErr))
	// Terminate the program with exit code 1, indicating that an error has occurred.
	os.Exit(1)
}
