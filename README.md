# Contributing to FileStruc

Thanks for your interest in contributing to FileStruc! This guide will help you set up, build, and test the project locally, as well as provide guidelines for submitting changes. Currently, the main focus of the project is adding file structures for more programming languages. If you're interested in contributing to this, this guide will help you get started.

## Getting Started

### 1. Fork the Repository  
Fork this repository by clicking the "Fork" button at the top of this page.

### 2. Clone Your Fork  
```bash
git clone https://github.com/matheuslinkdev/filestruc.git
cd filestruc
```

### 3. Install Dependencies  
Make sure you have Golang installed. Run the following command to download the dependencies:  
```bash
go mod download
```

## Building and Running the Project

### Local Installation  
To build and run the CLI locally:  
```bash
go run cmd/main/main.go
```  
Or, to build the binary:  
```bash
go build -o filestruc cmd/main/main.go
./filestruc --help
```

### Using Docker

#### Build the Docker Image  
```bash
docker build -t filestruc .
```

#### Run the Container  
```bash
docker run --rm filestruc
```

## Making Changes

### 1. Create a New Branch for Each Feature or Bug Fix  
If you're contributing a new file structure for a programming language, create a new branch with a descriptive name:  
```bash
git checkout -b feature/add-<language>-structure
```

### 2. Implement the File Structure  
In your branch, navigate to the appropriate directory for language structures and add the necessary file structure for the language you're working on. Make sure the structure is well-organized and matches the format of existing ones.

### 3. Commit Your Changes with Descriptive Messages  
```bash
git commit -m "Add file structure for <language>"
```

### 4. Push Your Branch to Your Forked Repository  
```bash
git push origin feature/add-<language>-structure
```

### 5. Open a Pull Request  
Open a Pull Request from your branch to the main repository.

## Code Style  
Please follow the general Go code style guidelines. Run `go fmt` to format your code before committing.

## Issues and Bugs  
If you find any issues or bugs, feel free to open an issue in the Issues tab. Make sure to provide clear information about the bug and how to reproduce it.

## License  
By contributing to this project, you agree that your contributions will be licensed under the MIT License
