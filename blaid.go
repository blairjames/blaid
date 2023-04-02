package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// Configure block devices for redundancy
const (
	redundantDriveOne = ""
	redundantDriveTwo = ""
	home              = ""
)

// Directories replicated across three hardware devices
func tierOneDirs() []string {
	return []string{
		home + "/Documents",
		home + "/python",
		home + "/go",
		home + "/docker",
		home + "/scripts",
	}
}

// Directories replicated across two hardware devices
func tierTwoDirs() []string {
	return []string{
		home + "/minecraft",
	}
}

// Write log message to system journal
func writeLog(message string) {
	if message == "" {
		errorMessage := "ERROR! The log message input to writeLog() cannot be empty"
		log.Fatal(errorMessage)
	}
	appName := "blaid"
	cmdString := "echo " + message + " | systemd-cat -p4 -t " + appName
	_, err := exec.Command("bash", "-c", cmdString).Output()
	log.Println(message)
	if err != nil {
		log.Fatal("ERROR! Failed to write log to system journal\n" +
			"Log message: " + message + "\n" + err.Error())
	}
}

// Ensure disk is available for sync
func testDriveIsAvailable(drive string) (bool, error) {
	if drive == "" {
		errorMessage := "ERROR! The 'drive' parameter of testDriveIsAvailable() cannot be empty"
		writeLog(errorMessage)
		log.Fatal(errorMessage)
	}
	_, err := os.Stat(drive)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Ensure disk has sufficient capacity for sync
func driveUsage(drive string) int {
	if drive == "" {
		errorMessage := "ERROR! The 'drive' parameter of driveUsage() cannot be empty"
		writeLog(errorMessage)
		log.Fatal(errorMessage)
	}
	cmd := "df --output=pcent " + drive + " | tr -dc '0-9'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		writeLog("ERROR! failed to get disk usage of " + drive + " running: " +
			cmd + "\n" + err.Error())
	}
	usage, err := strconv.Atoi(string(out))
	if err != nil {
		writeLog("ERROR! failed to convert disk usage to int for " + drive +
			"\n" + err.Error())
	}
	return usage
}

// Sync each directory to the appropriate redundant drives
func runSync(directories []string, destinationDrive string, deleteDestinationFiles bool) {
	if directories == nil || destinationDrive == "" {
		errorMessage := "ERROR! The 'directories' and 'destinationDrive'" +
			" parameters of runSync() cannot be empty"
		writeLog(errorMessage)
		log.Fatal(errorMessage)
	}
	for _, dir := range directories {
		writeLog("Syncing: " + dir + " to " + destinationDrive)
		rsyncCmd := exec.Command("rsync", "-a", dir, destinationDrive)
		if deleteDestinationFiles {
			rsyncCmd = exec.Command("rsync", "-a", "--delete", dir, destinationDrive)
		}
		_, err := rsyncCmd.Output()
		if err != nil {
			writeLog("ERROR! Blaid failed to sync " + dir + " to " + destinationDrive +
				"\n" + err.Error())
		}
		if err == nil {
			writeLog("Directory: " + dir + " was synchronized successfully to " + destinationDrive)
		}
	}
}

func main() {
	var redundantDrives = []string{redundantDriveOne, redundantDriveTwo}
	var deleteDestinationFiles bool = false
	for {
		for _, drive := range redundantDrives {
			var currentUsageInt int = driveUsage(drive)
			var currentUsageStr string = strconv.Itoa(currentUsageInt)
			var criticalDiskLimit int = 92
			var highWaterDiskUsage int = 79
			if available, err := testDriveIsAvailable(drive); !available {
				errorMessage := "ERROR! " + drive + " is not available\n" + err.Error()
				writeLog(errorMessage)
				log.Fatal(errorMessage)
			}
			writeLog("Block Device: " + drive + " is available for sync")
			writeLog("Current usage of " + drive + " is " + currentUsageStr + "%")
			if currentUsageInt > criticalDiskLimit {
				errorMessage := "ERROR! " + drive + " does not have available capacity to allow sync\n"
				writeLog(errorMessage)
				log.Fatal(errorMessage)
			}
			if currentUsageInt > highWaterDiskUsage {
				writeLog("Drive: " + drive + " has high current usage," +
					" enabeling deletion of extraneous files from redundant drives")
				deleteDestinationFiles = true
			}
			writeLog("Drive: " + drive + " has capacity available for sync\n" +
				"Value of deleteDestinationFiles: " + strconv.FormatBool(deleteDestinationFiles) + "\n" +
				"----- Syncing tier one directories to " + drive + " -----")
			runSync(tierOneDirs(), drive, deleteDestinationFiles)
		} // for
		writeLog("----- Syncing tier two directories to " + redundantDriveOne + " -----\n" +
			"Value of deleteDestinationFiles: " + strconv.FormatBool(deleteDestinationFiles))
		runSync(tierTwoDirs(), redundantDriveOne, deleteDestinationFiles)
		// Sleep for one minute
		writeLog("Sleeping for one minute")
		time.Sleep(time.Second * 60)
	} // infinite loop
}
