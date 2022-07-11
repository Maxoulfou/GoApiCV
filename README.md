# Portfolio API
Devansh's portfolio as a REST API, built with `gin`.

[![Heroku](https://github.com/Devansh3712/portfolio/actions/workflows/heroku.yml/badge.svg)](https://github.com/Devansh3712/portfolio/actions/workflows/heroku.yml)

## Available Endpoints

### `/docs/index.html`
Returns the swagger documentation for the API endpoints.

### `/`
Returns basic user information.
- Request Type: `GET`
- Response Type: `JSON`

### `/about`
Returns the about section of the user.
- Request Type: `GET`
- Response Type: `JSON`

### `/projects`
Returns the pinned projects on GitHub of the user using the GitHub GraphQL API.
- Request Type: `GET`
- Response Type: `JSON`

### `/misc`
Returns miscellaneous data of the user. Currently returns the recent tweet and the latest GitHub commit made by the user.
- Request Type: `GET`
- Response Type: `JSON`

### `/contact`
Send an email to the user.
- Request Type: `POST`
- Response Type: `JSON`
- Body Format:
    ```json
    {
        "email": userEmail,
        "subject": emailSubject,
        "body": emailBody
    }
    ```

### `/resume`
Returns the resume of the user in PDF format.
- Request Type: `GET`
- Response Type: `PDF File`
