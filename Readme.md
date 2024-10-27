# Folder Organizer

A simple Go program that organizes files in a specified folder by moving them into subfolders based on their file types (e.g., Images, Videos, Documents, Archives).

## 1. Features

- Automatically creates folders for each file type if they don’t already exist.
- Moves files based on their extensions into corresponding folders.
- Supports the following file types:
    - **Images**: `.jpeg`, `.jpg`, `.png`, `.gif`
    - **Videos**: `.mp4`, `.avi`, `.mov`
    - **Documents**: `.pdf`, `.docx`, `.txt`
    - **Archives**: `.zip`, `.rar`, `.gz`
    - **Windows**: `.exe`, `.com`
    - **Torrents**: `torrent`
    - **Isos**: `.iso`
    
## Prerequisites

- Go (1.16 or later)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/tloszkov/folderorganizer.git
cd folderorganizer
```

## 2. Run the Code

Replace "pathToDownloads := "/home/tloszkov/Downloads" with the actual path you want to organize.

go run main.go

## 3. Usage

The program will:

    Check if subfolders (Images, Videos, Documents, Archives) exist in the specified folder. If not, it creates them.
    Move each file into the appropriate subfolder based on its extension.

Example

If the folder /path/to/Downloads contains the following files:
``
example.jpg
video.mp4
document.pdf
archive.zip
``
After running the program, it will organize them as follows:
```
/path/to/Downloads/
├── Images/
│   └── example.jpg
├── Videos/
│   └── video.mp4
├── Documents/
│   └── document.pdf
└── Archives/
└── archive.zip
```
Error Handling

If an error occurs (e.g., file permission issues), the program will print an error message and terminate.
Contributing




