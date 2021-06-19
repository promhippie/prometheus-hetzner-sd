// +build ignore

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/promhippie/prometheus-hetzner-sd/pkg/action"
)

func main() {
	labels := []string{}

	for _, label := range action.Labels {
		labels = append(labels, label)
	}

	sort.Strings(labels)

	f, err := os.Create("docs/partials/labels.md")

	if err != nil {
		fmt.Printf("failed to create file")
		os.Exit(1)
	}

	defer f.Close()

	f.WriteString(
		"* `__address__`\n",
	)

	for _, row := range labels {
		if strings.HasSuffix(row, "_") {
			row = row + "<name>"
		}

		f.WriteString(fmt.Sprintf(
			"* `%s`\n",
			row,
		))
	}
}
