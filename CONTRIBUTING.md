# Contributing to SREnity

Thank you for considering contributing to SREnity! This guide outlines how you can contribute to the project and provides information about the project's architecture and structure.

## Hexagonal Architecture

SREnity follows the Hexagonal Architecture, also known as Ports and Adapters. This architecture promotes a clean and modular design, separating the core business logic from external concerns such as input/output mechanisms.

### Project Structure

- `./src/`: Main codebase.
- `./src/cli/`: CLI code for command-line interfaces.
- `./src/domain/`: Logic code representing the core business domain.
- `./src/entities/`: Data-related code, such as structs or data models.
- `./src/repositories/`: Input and output adapters.

Feel free to explore the respective directories for more details on each component.

## How to Contribute

1. Fork the repository.
2. Create a new branch: `git checkout -b feature/your-feature`.
3. Make your changes and commit them: `git commit -m 'Add new feature'`.
4. Push to the branch: `git push origin feature/your-feature`.
5. Create a pull request with a clear description of the changes.

### Development Setup

To set up your development environment:

1. Clone the repository: `git clone https://github.com/yourusername/SREnity.git`.
3. Build and run the project `cd src/; go run -c <config file> help`.

### Code Guidelines

- Follow the Go coding standards.
- Write meaningful commit messages.
- Keep pull requests focused on a single task or feature.

### Testing

Ensure that your changes include appropriate tests. Run tests locally before submitting a pull request.

### Reporting Issues

If you encounter any issues or have suggestions, please open an issue on the GitHub repository.

## Code of Conduct

Please note that this project follows the [Code of Conduct](CODE_OF_CONDUCT.md). Be respectful and considerate when interacting with the community.

Thank you for contributing to SREnity!
