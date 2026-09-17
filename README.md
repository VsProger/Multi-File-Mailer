# Multi-File Mailer API

A Go HTTP service for bulk file delivery over SMTP and for building ZIP archives from `multipart/form-data` uploads. Built for the [Doodocs Days 2.0](https://github.com/doodocs/doodocs-days/tree/main/backend) backend challenge.

## Overview

Two operations that web backends routinely need and routinely get wrong: accepting a set of uploaded files and returning them as one archive, and mailing a document to many recipients in a single request. Both hinge on handling `multipart/form-data` correctly — streaming the parts, validating what was actually uploaded rather than trusting the client, and failing clearly when a file is not what it claims to be.

The service exposes three endpoints and holds no state: no database, no queue, no session.

## Technical Approach

**Content-type validation by inspection, not by claim.** A `Content-Type` header and a file extension are both client-controlled. Uploads are identified from their magic bytes via `h2non/filetype`, and requests carrying a disallowed type are rejected before any processing.

**Archive handling.** `archive/zip` from the standard library builds archives in place and reads uploaded ones for inspection, reporting per-entry path, size and MIME type alongside compressed and uncompressed totals.

**Mail delivery.** SMTP credentials come from the environment, never from the request. The recipient list is parsed from a comma-separated field and the attachment is delivered to each address.

**Layering.** Handlers parse and validate HTTP; services hold the archive and mail logic and know nothing about `http.Request`; utilities cover MIME detection, archive assembly and environment loading. That separation is what makes the archive service unit-testable without standing up a server.

## Technologies

Go 1.21 · `net/http` (standard-library routing, no framework) · `archive/zip` · `h2non/filetype` · `jordan-wright/email` · `joho/godotenv` · Docker · Make

## API

### `POST /api/archive/information`

Inspects an uploaded ZIP archive.

**Parameters** — `file` (required): the ZIP archive.

```json
{
    "filename": "my_archive.zip",
    "archive_size": 4102029,
    "total_size": 6836715,
    "total_files": 2,
    "files": [
        { "file_path": "photo.jpg", "size": 2516582, "mimetype": "image/jpeg" },
        { "file_path": "directory/document.docx", "size": 4320133, "mimetype": "application/vnd.openxmlformats-officedocument.wordprocessingml.document" }
    ]
}
```

### `POST /api/archive/files`

Combines uploaded files into a ZIP archive and returns it.

**Parameters** — `files[]` (required). Accepted types: `application/vnd.openxmlformats-officedocument.wordprocessingml.document`, `application/xml`, `image/jpeg`, `image/png`.

```http
POST /api/archive/files HTTP/1.1
Content-Type: multipart/form-data; boundary=-{boundary}

-{boundary}
Content-Disposition: form-data; name="files[]"; filename="document.docx"
Content-Type: application/vnd.openxmlformats-officedocument.wordprocessingml.document

{binary}
-{boundary}--
```

Responds with the ZIP file as a download.

### `POST /api/mail/file`

Sends one file to several recipients.

**Parameters** — `file` (required; `application/pdf` or `.docx`), `emails` (required; comma-separated addresses).

```json
{ "message": "Файл успешно отправлен на указанные почты." }
```

## How to Run

```bash
git clone https://github.com/VsProger/Multi-File-Mailer.git
cd Multi-File-Mailer
```

Create a `.env` file:

```
PORT=:8080
SMTP_HOST=smtp.example.com
SMTP_PORT=587
SMTP_USERNAME=your_username
SMTP_PASSWORD=your_password
```

For Gmail and other providers with two-factor authentication, use an app-specific password. `.env` is gitignored and must stay that way.

```bash
go mod tidy
go run ./cmd/web/main.go
```

Or through the Makefile: `make run`, `make test`, `make clean`, `make docker-run`.

Quick check:

```bash
curl -F "file=@document.pdf" -F "emails=a@example.com,b@example.com" \
     http://localhost:8080/api/mail/file
```

## Project Structure

```
├── cmd/web/main.go           entry point
├── application/              HTTP handlers and routing
│   ├── app.go                    server setup, route table
│   ├── archive_handlers.go
│   └── email_handlers.go
├── services/                 business logic, independent of net/http
│   ├── archive_services.go
│   ├── archive_services_test.go
│   └── email_service.go
├── utils/                    MIME detection, archive assembly, mail, env
├── models/                   request and response types
├── config/                   configuration loading
├── logger/                   levelled logging
├── Dockerfile
└── Makefile
```

## Future Improvements

- **Bound the upload size.** There is currently no explicit limit, so a large multipart body can consume memory; `http.MaxBytesReader` plus streaming straight to a temporary file would fix it.
- **Report per-recipient delivery results.** Mailing is all-or-nothing today; returning per-address status would let clients retry only what failed.
- **Move mail off the request path** into a worker with retry and backoff, so a slow SMTP server stops blocking the HTTP response.
- **Extend test coverage** from the archive service to the handlers, using `httptest` with synthetic multipart bodies.

## License

MIT — see [LICENSE](LICENSE).
