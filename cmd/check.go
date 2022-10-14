package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gordonjb/posterscript/internal/pathutils"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check PATH",
	Short: "Validate Plex local posters",
	Long: `Checks Plex Movie and TV library folders, and output when it believes that 'poster.ext', 'fanart.ext' and 'seasonxx.ext' local image files are missing.
	
Each subdirectory of the starting directory PATH will be treated as a Library folder (e.g. with root PATH 'Images/Plex Posters' which contains
'TV', 'Movie' and 'Other' folders, checks will only be made within the 'Images/Plex Posters/TV/', 'Images/Plex Posters/Movie/' and 'Images/Plex Posters/Other/' folders)`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		rootDir := args[0]
		excludeList, _ := cmd.Flags().GetStringSlice("exclude")
		showAll, _ := cmd.Flags().GetBool("show-all")
		return check(rootDir, excludeList, showAll)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringSliceP("exclude", "e", []string{},
		"Ignore this `DIRECTORY` when scanning the root path. Can be specified multiple times")
	checkCmd.Flags().BoolP("show-all", "a", false, "Show every validated item, not just failures")
}

func check(rootDir string, excludeList []string, showAll bool) error {
	rootDirEntries, err := os.ReadDir(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Poster", "Fanart", "Season", "Name", "Missing Seasons"})
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:             "Name",
			WidthMax:         60,
			WidthMaxEnforcer: text.Trim,
		},
	})

	for _, rootEntry := range rootDirEntries {
		if !rootEntry.IsDir() {
			continue
		}
		if slices.Contains(excludeList, rootEntry.Name()) {
			continue
		}
		libraryDirEntries, err := os.ReadDir(filepath.Join(rootDir, rootEntry.Name()))
		if err != nil {
			log.Fatal(err)
		}
		for _, libraryEntry := range libraryDirEntries {
			if !libraryEntry.IsDir() {
				continue
			}
			libraryItemEntries, err := os.ReadDir(filepath.Join(rootDir, rootEntry.Name(), libraryEntry.Name()))
			if err != nil {
				log.Fatal(err)
			}
			hasPoster := false
			hasFanart := false
			hasSeasons := false
			missingSeasons := []string{}
			for _, libraryItemEntry := range libraryItemEntries {
				if libraryItemEntry.IsDir() {
					isSpecial := libraryItemEntry.Name() == "Specials"
					isSeason := strings.HasPrefix(libraryItemEntry.Name(), "Season ")
					if isSeason || isSpecial {
						hasSeasons = true
						seasonDirEntries, err := os.ReadDir(filepath.Join(rootDir, rootEntry.Name(), libraryEntry.Name(), libraryItemEntry.Name()))
						if err != nil {
							log.Fatal(err)
						}
						hasSeasonPoster := false
						var seasonSuffix string
						if isSeason {
							seasonSuffix = libraryItemEntry.Name()[7:]
						} else {
							seasonSuffix = "-specials-poster"
						}
						for _, file := range seasonDirEntries {
							if file.IsDir() {
								continue
							}
							if pathutils.Stem(file.Name()) == "season"+seasonSuffix {
								hasSeasonPoster = true
								break
							}
						}
						if !hasSeasonPoster {
							if isSeason {
								missingSeasons = append(missingSeasons, seasonSuffix)
							} else {
								missingSeasons = append(missingSeasons, "Sp")
							}
						}
					}
				} else if !hasPoster || !hasFanart {
					filenameStem := pathutils.Stem(libraryItemEntry.Name())
					if filenameStem == "poster" {
						hasPoster = true
					} else if filenameStem == "fanart" {
						hasFanart = true
					}
				}
			}
			if showAll || !hasFanart || !hasPoster || (hasSeasons && len(missingSeasons) > 0) {
				t.AppendRow(table.Row{
					renderBool(hasPoster),
					renderBool(hasFanart),
					renderSeasons(hasSeasons, missingSeasons),
					rootEntry.Name() + "/" + libraryEntry.Name(),
					strings.Join(missingSeasons, ", ")})
			}
		}
		t.AppendSeparator()
	}
	t.Render()

	return nil
}

func renderBool(boolean bool) string {
	if boolean {
		return "✅"
	} else {
		return "❌"
	}
}

func renderSeasons(hasSeasons bool, missingSeasons []string) string {
	if hasSeasons {
		if len(missingSeasons) == 0 {
			return "✅"
		} else {
			return "❌"
		}
	} else {
		return "➖"
	}
}
