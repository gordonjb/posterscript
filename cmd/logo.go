package cmd

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// logoCmd represents the logo command
var logoCmd = &cobra.Command{
	Use:   "logo FILE",
	Short: "Overlay a logo file over a poster file, for example to add network logos",
	RunE: func(cmd *cobra.Command, args []string) error {
		posterPath := args[0]
		logoPath, _ := cmd.Flags().GetString("logoDir")
		return logo(posterPath, logoPath)
	},
}

func init() {
	rootCmd.AddCommand(logoCmd)
	logoCmd.Flags().StringP("logoDir", "l", "logos", "Directory containing overlay images")
}

func logo(posterPath string, logoPath string) error {
	logoDirEntries, err := os.ReadDir(logoPath)
	if err != nil {
		log.Fatal(err)
	}
	optionList := []string{}
	for _, entry := range logoDirEntries {
		if !entry.IsDir() {
			optionList = append(optionList, entry.Name())
		}
	}
	var logo string
	prompt := &survey.Select{
		Message:  "Select logos:",
		Options:  optionList,
		PageSize: 10,
	}
	survey.AskOne(prompt, &logo)

	posterFile, err := os.Open(posterPath)
	if err != nil {
		log.Fatal(err)
	}

	posterImage, _, err := image.Decode(posterFile)
	if err != nil {
		log.Fatal(err)
	}
	defer posterFile.Close()

	logoFile, err := os.Open(filepath.Join(logoPath, logo))
	if err != nil {
		log.Fatal(err)
	}
	logoImage, err := png.Decode(logoFile)
	if err != nil {
		log.Fatal(err)
	}
	defer logoFile.Close()

	bounds := image.Rect(0, 0, 1000, 1500)
	newPoster := image.NewRGBA(bounds)
	draw.CatmullRom.Scale(newPoster, newPoster.Rect, posterImage, posterImage.Bounds(), draw.Over, nil)
	draw.Draw(newPoster, bounds, logoImage, image.Point{X: 0, Y: 0}, draw.Over)

	output, err := os.Create(posterPath + ".logo.jpg")
	if err != nil {
		log.Fatal(err)
	}
	jpeg.Encode(output, newPoster, &jpeg.Options{Quality: 95})
	defer output.Close()
	return nil
}
