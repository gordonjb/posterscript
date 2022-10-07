package cmd

import (
	"archive/zip"
	"fmt"
	"io"

	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/gordonjb/posterscript/internal/pathutils"
	"github.com/spf13/cobra"
)

// tpdbarchiveCmd represents the tpdbarchive command
var tpdbarchiveCmd = &cobra.Command{
	Use:     "tpdbarchive FILE",
	Aliases: []string{"archive", "tpdb", "a"},
	Short:   "Unpack TPDb archive to Plex format",
	Long: `Unzip a poster set from The Poster Database into the correct folder structure for Plex to pick them up.
	
Supports Movie and TV sets.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		archivePath := args[0]
		return tPDbArchive(archivePath)
	},
}

func init() {
	rootCmd.AddCommand(tpdbarchiveCmd)
}

var specialsFolderName = "Specials"
var specialsPosterName = "season-specials-poster"
var posterName = "poster"

func tPDbArchive(archivePath string) error {
	r := regexp.MustCompile(` - (Specials)| - Season (\d+)`)
	source := filepath.Join(archivePath)
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, f := range reader.File {
		fStem := pathutils.Stem(f.Name)
		isSeason := r.MatchString(fStem)
		var destinationStem string
		var destinationFolder string
		if isSeason {
			matchIndex := r.FindStringIndex(fStem)
			title := fStem[:matchIndex[0]]
			matches := r.FindStringSubmatch(fStem)
			isSpecial := matches[1] == "Specials"
			subfolderName := specialsFolderName
			seasonPosterName := specialsPosterName
			if !isSpecial {
				subFolderInt, err := strconv.Atoi(matches[2])
				if err != nil {
					return err
				}
				subfolderName = fmt.Sprintf("Season %02d", subFolderInt)
				seasonPosterName = fmt.Sprintf("season%02d", subFolderInt)
			}
			destinationFolder = filepath.Join(title, subfolderName)
			destinationStem = seasonPosterName
		} else {
			destinationFolder = fStem
			destinationStem = posterName
		}
		err := os.MkdirAll(destinationFolder, os.ModePerm)
		if err != nil {
			return err
		}
		posterPath := filepath.Join(destinationFolder, destinationStem+filepath.Ext(f.Name))
		fmt.Printf("Unzipping to '%s'\n", posterPath)

		posterFile, err := os.OpenFile(posterPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		originFile, err := f.Open()
		if err != nil {
			return err
		}
		_, err = io.Copy(posterFile, originFile)
		if err != nil {
			return err
		}
		posterFile.Close()
		originFile.Close()
	}

	// Files will be output in current directory
	//destination, err := filepath.Abs(filepath.Join(directoryname))
	if err != nil {
		return err
	}

	return nil
}
