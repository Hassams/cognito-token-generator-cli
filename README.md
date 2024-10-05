# JWT CLI for AWS Cognito

A cross-platform command-line tool to authenticate with AWS Cognito and generate JSON Web Tokens (JWT). This tool securely encrypts user credentials using AES-256 encryption with a passphrase, supports copying the generated JWT token to the system clipboard, and is written in Go, making it highly portable across Linux, macOS, and Windows.

## Features

- **User Authentication**: Authenticate with AWS Cognito using the `USER_PASSWORD_AUTH` flow.
- **JWT Token Generation**: Retrieves an Access Token (JWT) upon successful authentication.
- **Secure Credential Management**: User credentials (passwords) are encrypted using AES-256 encryption with a passphrase.
- **Copy JWT to Clipboard**: Automatically copies the generated JWT token to the system clipboard for ease of use.
- **Cross-Platform Support**: Works on Linux, macOS, and Windows.
- **Clear Saved Credentials**: Easily clear saved AWS and user credentials via CLI commands.

## Requirements

- **Go 1.17+**: You need Go installed to build this project.
- **AWS Cognito**: You need a configured AWS Cognito User Pool and Client ID.

## Installation

### Clone the Repository

```bash
git clone https://github.com/hassams/cognito-token-generator-cli.git
cd jwt-cli-tool

### Build the Application

To build the CLI application for your platform, run the following command:

# For macOS or Linux
go build -o jwtcli

# For Windows
go build -o jwtcli.exe
```

This will generate the `jwtcli` binary (or `jwtcli.exe` for Windows) that you can run directly.

## Usage

The CLI tool allows you to authenticate with AWS Cognito and generate JWT tokens. The tokens are copied to the system clipboard for easy access.

### Running the CLI

To generate a JWT token, simply run the tool:

```bash
./jwtcli
```

You will be prompted to enter:
- An encryption passphrase to secure your credentials.
- AWS Cognito Client ID, User Pool ID, and Region (only the first time, as they can be saved for future use).
- Your username and password for authentication.

Upon successful authentication, the JWT token will be displayed and copied to your clipboard.

### Commands

1. **Generate JWT Token**: Simply run the tool without arguments:

```bash
./jwtcli
```

2. **Clear AWS Credentials**: Use this command to clear saved AWS Cognito credentials (Client ID, User Pool ID, Region):

```bash
./jwtcli clear-aws-credentials
```

3. **Clear User Credentials**: Use this command to clear saved user credentials (username and encrypted password):

```bash
./jwtcli clear-user-credentials
```

4. **Show Version**: Display the current version of the CLI tool:

```bash
./jwtcli version
```

### Example Workflow

1. **Run the tool**:

```bash
./jwtcli
```

2. **Enter the encryption passphrase**: This passphrase is used to encrypt and decrypt your stored credentials.
3. **Enter AWS Cognito Credentials**: You'll be prompted for the Client ID, User Pool ID, Region, username, and password (these can be saved for future use).
4. **Get the JWT token**: After successful authentication, the JWT token will be generated, displayed in the terminal, and copied to your system clipboard.

### Environment Variables

You can set environment variables to avoid entering AWS Cognito credentials each time:

```bash
export AWS_COGNITO_CLIENT_ID="your-client-id"
export AWS_COGNITO_USER_POOL_ID="your-user-pool-id"
export AWS_REGION="your-region"
```

If these environment variables are set, the tool will automatically use them without prompting for the Client ID, User Pool ID, and Region.

## Security

- **AES-256 Encryption**: Passwords are encrypted using AES-256 encryption with a passphrase provided by the user at runtime.
- **No Hardcoded Secrets**: No secrets are hardcoded in the codebase, and user credentials are stored securely in encrypted form.

## Contributing

Contributions are welcome! If you have any ideas, issues, or improvements, feel free to submit a pull request or create an issue.

To contribute:
1. Fork the project.
2. Create a new branch (`git checkout -b feature/your-feature-name`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push the branch (`git push origin feature/your-feature-name`).
5. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- AWS SDK for Go: [github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2)
- Cobra CLI: [github.com/spf13/cobra](https://github.com/spf13/cobra)
- Clipboard management: [github.com/atotto/clipboard](https://github.com/atotto/clipboard)
