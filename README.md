# ASCII Art Generator

A Go program that converts images into ASCII art. This tool takes an image file (JPEG or PNG) and transforms it into a text-based representation using carefully selected ASCII characters.

## Features

- Supports both JPEG and PNG image formats
- Customizable output dimensions (width and height)
- Adjustable ASCII character palette for different visual effects
- Simple grayscale conversion algorithm for accurate brightness representation

## How It Works

1. The program reads an input image file (`image_1.jpg` by default)
2. Resizes the image to the specified dimensions (default: 50x25)
3. Converts each pixel to grayscale using the formula:  
   `gray = (0.299R + 0.587G + 0.114B) / 256`
4. Maps the grayscale value to one of 17 ASCII characters (from darkest to lightest):  
   `" ", ".", ",", ":", ";", "o", "x", "X", "|", "/", "=", "(", ")", "&", "%", "#", "@"`
5. Outputs the resulting ASCII art to the console

## Requirements

- Go (1.16 or higher recommended)
- Image libraries (automatically handled by Go modules)

## Installation

1. Clone this repository:
```bash
git clone https://github.com/actionjacks/convertImageToASCII.git
cd convertImageToASCII
```

2. Build the program:
```bash
go build
```

## Usage

Place your image file (`image_1.jpg` or `image_1.png`) in the project directory.

Run the program:
```bash
./ascii-art-generator
```

Or for a specific image:
```bash
go run main.go your_image.jpg
```

For custom dimensions, modify the `defaultWidth` and `defaultHeight` variables in `main.go`.

## Example Output

```
               .....,..,. .... .......           .
          ,:::,,,,. .. .            ....          
       .;xx;,                         .,,,.       
      ;XX,                             .,::,.     
     :/|,                            ....,::;,    
    ;//.                             .....o;;;.   
   : ,=.             .  .              . :o;x,,   
.. , .x,:           ,,.....:.      .,....oXoX,,.  
..., .|Xo,         ..  . .,:  ...,:;:...;=(;o.o,  
  .: X#=:,              .    ...,X|;....,;|xX;..  
  ...&&(;..         .,X|:     ;)=;,::;,. ,x|XX,   
  ,//o/=x,.,,,:;;,....:,    ..:x::;;:.....,o;:o   
  .%#x:x   .,,;;o:...     ..... .:;o;:,  .,; oX   
   X)|;:    .......... .    .  ....,,.,.  :x:xo...
.. .;((                  . .,;,. ..,,,:.  ,,x||...
... ,o(|,.          ...  .. .,,:.. .       .x/,   
    ,; ;x;:...   .         ...,;, ..    ...o;.    
     x , :,..      .      o;::;x;.......,..:, .   
     |,,, ,,.   .,,.     .:::;;:,:..:,. . .,. ... 
     ,X,   :;,  ....... .o o;.,:.,.,;:,, .,.     .
      .xx;.  ,  .,   .   ,.,,,:,,:.,...;::.       
        .:;:;;,    ..  . ...,:: ..,  ,,,.         
           .,;x;.          ..o,   .:...           
               ,:...      .:o;.   ..              
                  ...                             
```

## Customization

- Adjust the `asciiChars` array in `main.go` to change the character palette.
- Modify the resize algorithm in `convertImageToASCII()` for different scaling methods.
- Change the grayscale conversion coefficients for different brightness effects.

## License

MIT License - Free to use and modify
