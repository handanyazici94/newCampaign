package io

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadCommandsFromFile(inputFileName string) [][]string {
	commands := [][]string{}

	f := OpenFile(inputFileName)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		command := strings.Fields(scanner.Text() + "\t")
		commands = append(commands, command)
	}
	err := f.Close()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return commands
}

func OpenFile(inputFileName string) (f *os.File) {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return f
}

func CreateAndWriteOutputFile(outputFileName string, command string) error {
	file, err := os.OpenFile(outputFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	WriteToFile(command, file)

	return file.Close()
}

func WriteToFile(output string, f *os.File) {
	_, err := f.WriteString(output + "\n")
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
