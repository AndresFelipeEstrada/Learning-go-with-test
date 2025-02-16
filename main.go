package main

import "github.com/fatih/color"

func main() {
	// Create a new color object
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")

	// Or just add them to New()
	d := color.New(color.FgCyan, color.Bold)
	d.Printf("This prints bold cyan %s\n", "too!.")

	// Mix up foreground and background colors, create new mixes!
	red := color.New(color.FgRed)

	boldRed := red.Add(color.Bold)
	boldRed.Println("This will print text in bold red.")

	whiteBackground := red.Add(color.BgYellow)
	whiteBackground.Println("Red text with white background.")

	// Mix with RGB color codes
	color.RGB(255, 128, 0).AddBgRGB(0, 0, 0).Println("orange with black background")

	color.BgRGB(255, 128, 0).AddRGB(255, 255, 255).Println("orange background with white foreground")
}
