package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`

	Count int `json:"count"`
}

func main() {
	var data = map[string][]string{}
	cmd := exec.Command("git", "log", "--pretty=`%H|%an|%ae|%at`")

	output, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	outputarr := strings.Split(string(output), "\n")

	for _, val := range outputarr {
		if val == "\n" {
			return
		}

		resArry := strings.Split(val, "|")

		if len(resArry) > 1 {

			v, ok := data[resArry[2]]

			if ok {
				vi := append(v, resArry[1])
				data[resArry[2]] = vi

			} else {
				data[resArry[2]] = []string{
					resArry[1],
				}

			}

		}

	}

	var finalCount = []UserData{}

	for id, name := range data {
		// finalCount[id] = len(name)

		finalCount = append(finalCount, UserData{
			Name:  name[0],
			Email: id,
			Count: len(name),
		})
	}

	// fmt.Println(finalCount)

	var userDataStrings [][]string
	for _, user := range finalCount {
		userDataStrings = append(userDataStrings, []string{user.Name, user.Email, fmt.Sprintf("%d", user.Count)})
	}

	ShowTable(userDataStrings)

}

func ShowTable(data [][]string) {
	re := lipgloss.NewRenderer(os.Stdout)
	barStyle := re.NewStyle().Padding(0, 1)
	headerStyle := barStyle.Copy().Foreground(lipgloss.Color("252")).Bold(true)

	CapitalizeHeaders := func(d []string) []string {
		for i := range d {
			d[i] = strings.ToUpper(d[i])
		}
		return d
	}

	headers := []string{
		"Name", "Email", "Contribution",
	}

	t := table.New().Border(lipgloss.NormalBorder()).BorderStyle(re.NewStyle().Foreground(lipgloss.Color("238"))).Headers(CapitalizeHeaders(headers)...).Width(80).Rows(data...).BorderRow(true).StyleFunc(func(row, col int) lipgloss.Style {
		if row == 0 {
			return headerStyle
		}

		return barStyle.Copy().Foreground(lipgloss.Color("252"))
	})

	fmt.Println(t)
}
