# THUMBIFY

## About The Project

A GoLang CLI app to create thumbnails for video


### Built With
* [Go](https://go.dev/)
* [FFmpeg](https://ffmpeg.org/)

## Getting Started

### Prerequisites
* FFmpeg
* Go
* Git

### Installation

1. Clone the repo
```sh
git clone https://github.com/colt005/thumbify.git
```
2. Build binary
```sh
go build -o thumbify
```
3. Use the binary to generate thumbnails for your videos
```sh
./thumbify gen -f "video.mp4" -o "output.png"
```

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Acknowledgements

This project was created using the [Cobra CLI](https://github.com/spf13/cobra-cli) 
