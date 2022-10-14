package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

// libibCmd represents the libib command
var libibCmd = &cobra.Command{
	Use:   "libib FILE",
	Short: "Filter libib library by tags",
	Long:  `Reads in a CSV export of a Libib library, and filters entries based on combinations of tags.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		includeList, _ := cmd.Flags().GetStringSlice("include")
		excludeList, _ := cmd.Flags().GetStringSlice("exclude")
		csvPath := args[0]
		return libib(csvPath, includeList, excludeList)
	},
}

func init() {
	rootCmd.AddCommand(libibCmd)

	libibCmd.Flags().StringSliceP("include", "i", []string{},
		"A libib `TAG`. Can be specified multiple times. Items in the library must have *all* of these TAGs, or they will not be returned")
	libibCmd.Flags().StringSliceP("exclude", "e", []string{},
		"A libib `TAG`. Can be specified multiple times. If an item in the library has *any* of these TAGs, it will not be returned")
}

func libib(csvPath string, includeList []string, excludeList []string) error {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
	}

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	headerRow, err := csvReader.Read()
	if err == io.EOF {
		log.Fatal("No lines in file")
	}
	if err != nil {
		log.Fatal(err)
	}
	headers := make(map[string]int)
	hasTags := false
	for index, key := range headerRow {
		if key == "tags" {
			hasTags = true
		}
		headers[key] = index
	}
	if !hasTags {
		log.Fatal("no tag col")
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle(fmt.Sprintf("Items matching tag sets: include [%s], exclude [%s]", strings.Join(includeList, ", "), strings.Join(excludeList, ", ")))
	t.AppendHeader(table.Row{"Name", "Tags"})
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:             "Name",
			WidthMax:         60,
			WidthMaxEnforcer: text.Trim,
		},
	})

	lineCount := 0
	returnCount := 0
	for {

		row, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		lineCount++
		row_tags := strings.Split(row[headers["tags"]], ",")

		matchesIncludeList := true
		for _, v := range includeList {
			if !slices.Contains(row_tags, v) {
				matchesIncludeList = false
				break
			}
		}

		matchesExcludeList := true
		for _, v := range excludeList {
			if slices.Contains(row_tags, v) {
				matchesExcludeList = false
				break
			}
		}

		if matchesIncludeList && matchesExcludeList {
			t.AppendRow(table.Row{
				row[headers["title"]],
				fmt.Sprintf("[%s]", row[headers["tags"]])})
			returnCount++
		}
	}
	t.AppendFooter(table.Row{
		fmt.Sprintf("Returned %d of %d records.", returnCount, lineCount)})
	t.Render()
	return nil
}
