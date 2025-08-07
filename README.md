# Go Templates

It's personal collection of project templates for [go-scaffold](https://github.com/go-scaffold/go-scaffold) to quickly generate initial project structures for different types of Go applications. I'm happy  if it'll be handful for someone else too.

## Templates

### Go Package (`package`)
Template for generating a basic Go package structure.

### HTTP API (`http_api`)
A lightweight HTTP service template that provides a foundation for building REST APIs and microservices with configurable endpoints and port binding.

## Usage

To use these templates with first install [go-scaffold](https://github.com/go-scaffold/go-scaffold). Then by following examples that provided in every template create your values file to generate project tailored to your needs.

```bash
# Generate a new project using a template
go-scaffold generate http_api  my-awesome-service --values ./http_api/values.yaml 
```

## Template Structure

Each template follows a consistent structure:
- `template.yaml` - Template configuration
- `values.yaml` - Default values for the template
- Project files and directories with Go template syntax

## Contributing

1. Fork the repository
2. Create a new template directory
3. Add your template files with appropriate Go template syntax
4. Update this README with your template description
5. Submit a pull request

## Available Templates

- **http_api** - REST API and microservice template

More templates coming soon!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
