package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Jira tracking in git commit
func main() {
	out, err := exec.Command("git", "branch").Output()

	if err != nil {
		log.Fatal(err)
	}

	output := (string(out))

	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "* ") {

			var stripped = strings.Replace(line, "*", "", -1)
			commitName := strings.Trim(stripped, " ")

			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Enter Time: ")
			time, _, _ := reader.ReadLine()
			timeStr := string(time)

			fmt.Print("Enter Comment: ")
			comment, _, _ := reader.ReadLine()
			commentStr := string(comment)

			resultCommand := fmt.Sprintf("%s #time %s #comment %s", commitName, timeStr, commentStr)

			fmt.Println(resultCommand)
			out, err = exec.Command("git", "commit", "-m", resultCommand).Output()
			fmt.Println(string(out))
			break
		}
	}
}
