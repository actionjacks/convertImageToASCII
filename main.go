package main

import (
	"fmt"
	"image"
	_ "image/jpeg" // Import for JPEG support
	_ "image/png"  // Import for PNG support
	"log"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

// asciiChars is an array of characters used to represent pixels
var asciiChars = []string{" ", ".", ",", ":", ";", "o", "x", "X", "|", "/", "=", "(", ")", "&", "%", "#", "@"}

// convertImageToASCII converts an image to ASCII art
// Now the function accepts width and height as arguments.
func convertImageToASCII(img image.Image, width, height int) string {
	// Resize the image
	resized := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
	bounds := resized.Bounds()
	// height := bounds.Max.Y  // We don't need this, we use the height parameter
	width = bounds.Max.X // Make sure width is correct after resizing.

	var asciiArt strings.Builder
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get pixel color
			r, g, b, _ := resized.At(x, y).RGBA()
			// Convert to grayscale (simpler calculations)
			gray := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 256
			// Select ASCII character based on brightness
			charIndex := int(gray * float64(len(asciiChars)-1) / 255)
			asciiArt.WriteString(asciiChars[charIndex])
		}
		asciiArt.WriteString("\n")
	}
	return asciiArt.String()
}

func main() {
	// Default image file name
	defaultImageFile := "image_1.jpg" // You can change to "obraz.png" or another

	// Default dimensions
	defaultWidth := 50  // You can adjust
	defaultHeight := 25 // You can adjust

	// Check if file exists
	_, err := os.Stat(defaultImageFile)
	if os.IsNotExist(err) {
		log.Fatalf("Image file '%s' doesn't exist in the current folder.", defaultImageFile)
	}

	filePath := defaultImageFile // Use default file name

	// Open image file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Cannot open file: %v", err)
	}
	defer file.Close()

	// Decode image
	img, format, err := image.Decode(file) // image.Decode returns format
	if err != nil {
		log.Fatalf("Cannot decode image: %v", err)
	}
	fmt.Printf("Image format: %s\n", format) // Print the format.

	// Convert image to ASCII art
	// Now we pass width and height.
	asciiArt := convertImageToASCII(img, defaultWidth, defaultHeight)

	// Display ASCII art
	fmt.Println(asciiArt)
}
