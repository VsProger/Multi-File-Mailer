# 📧 Multi-File Mailer API

Multi-File Mailer API is a powerful tool designed to send files
to multiple email addresses, as well as for combining files into ZIP archives.
The project allows you to easily work with multipart/form-data files, simplifying
the task of bulk mailing and file management.

## 🌟 Features:

* **Bulk file distribution** 📬: send a document to multiple recipients at once.
*  **Archiving** 📦: process uploaded files to combine them into a single ZIP archive.
*  **Multiple file type support** 📄: process files of different MIME types such as `application/pdf`, `application/vnd.openxmlformats-officedocument.wordprocessingml.document`, and images.
*  **Instant access to archive information** 📊: get data about archive contents and structure.

## 📚 API Routes

### 1. POST /api/archive/information
Provides information about the downloaded ZIP archive, including the size of the archive, the number of files, and the type of each file.

#### Query parameters
* **file** (mandatory): the archive file (ZIP) whose information is to be retrieved.

#### Example of an answer
```
{
    "filename": "my_archive.zip",
    "archive_size": 4102029,
    "total_size": 6836715,
    "total_files": 2,
    "files": [
        {
            "file_path": "photo.jpg",
            "size": 2516582,
            "mimetype": "image/jpeg"
        },
        {
            "file_path": "directory/document.docx",
            "size": 4320133,
            "mimetype": "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
        }
    ]
}
```

### 2. POST /api/archive/files
Combines the uploaded files into a ZIP archive and returns the archive to the client.

#### Query parameters
* **files**[] (mandatory): array of files to archive. Supported MIME types:
    * `application/vnd.openxmlformats-officedocument.wordprocessingml.document`
    * `application/xml`
    * `image/jpeg`
    * `image/png`

#### Example of a request
```
POST /api/archive/files HTTP/1.1
Content-Type: multipart/form-data; boundary=-{boundary}

-{boundary}
Content-Disposition: form-data; name="files[]"; filename="document.docx"
Content-Type: application/vnd.openxmlformats-officedocument.wordprocessingml.document

{Binary data of DOCX file}
-{boundary}
Content-Disposition: form-data; name="files[]"; filename="avatar.png"
Content-Type: image/png

{Binary data of PNG file}
-{boundary}--
```
#### Example of an answer
Downloadable ZIP file with downloaded files.

### 3. POST /api/mail/file
Sends the specified file to multiple recipients via email.

#### Request parameters
* **file** (mandatory): the file you want to send. MIME types supported:
  * `application/vnd.openxmlformats-officedocument.wordprocessingml.document`
  * `application/pdf`
* **emails** (mandatory): a comma-separated list of email addresses to which the file will be sent.

#### Example of a request
```
POST /api/mail/file HTTP/1.1
Content-Type: multipart/form-data; boundary=-{boundary}

-{boundary}
Content-Disposition: form-data; name="file"; filename="document.pdf"
Content-Type: application/pdf

{Binary data of PDF file}
-{boundary}
Content-Disposition: form-data; name="emails"

email1@example.com,email2@example.com
-{boundary}--
```

#### Example of an answer
``` 
{
  "message": "Файл успешно отправлен на указанные почты."
}
```

## 🛠 Installation

1. Clone the repository:
``` 
git clone https://github.com/VsProger/Multi-File-Mailer-API.git
cd Multi-File-Mailer-API
```
2. Create an .env file with variables for SMTP settings:
``` 
PORT=:8080
SMTP_HOST=smtp.example.com
SMTP_PORT=587
SMTP_USERNAME=your_username
SMTP_PASSWORD=your_password
```
3. Install the dependencies and run the application:
``` 
go mod tidy
go run ./cmd/web/main.go
```
Additionally, you can use commands in the `Makefile`
* `make run`: to build and run the application
* `make clean`: to clean up binary
* `make test`: to run tests
* `make docker-run`: to build and run the application  in the docker
  
_(make sure you have make and docker)_
## ⚙️ Configuration
Environment variables required to set up the SMTP server and API port:
- `PORT`: Port on which the API server will run (e.g., `8080`).
- `SMTP_HOST`: Address of the SMTP server (e.g., `smtp.gmail.com` for Gmail).
- `SMTP_PORT`: Port of the SMTP server (e.g., `587` for TLS).
- `SMTP_USERNAME`: Username for authentication (usually your email).
- `SMTP_PASSWORD`: Password or app-specific password for authentication.

## 🧪 API Testing
Using Postman:
1. Open Postman and create a new POST request.
2. Enter your server URL (http://localhost:8080/api/mail/file) and configure Body in form-data format to send files. 
3. Add a file key for the file and emails for a list of email addresses.

## 🚀 How it works
1. **File processing:** Uploaded files are validated against supported file types.
2. **Archive creation:** The API can combine multiple files into a single ZIP archive and send it to the user. 
3. **Mailing:** The API uses the specified SMTP credentials to send files to a list of recipients.

## 🛡️ Security
For Gmail and other services with two-factor authentication, it is recommended to use application passwords. Make sure your .env file is not added to the repository (add it to .gitignore).

## 📄 Licence
This project is distributed under the MIT licence.
