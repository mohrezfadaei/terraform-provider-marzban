# Marzban Terraform Provider

The Marzban Terraform Provider allows you to manage Marzban nodes and configurations via Terraform, providing a unified and automated way to interact with the Marzban API.

## Features

- **Create and Manage Nodes**: Easily create, read, update, and delete Marzban nodes.
- **Configuration Management**: Manage your Marzban configurations directly from your Terraform scripts.
- **Authentication**: Securely authenticate with the Marzban API using tokens.

## Documentation

Comprehensive documentation is available to help you get started and effectively use the provider. Please refer to the following sections:

- [Installation](./docs/installation.md): Learn how to install and configure the provider.
- [Usage](./docs/usage.md): See examples and details on how to use the provider.
- [Resource: marzban_node](./docs/resources/marzban_node.md): Detailed documentation for the `marzban_node` resource.

## Getting Started

Here’s a quick example to get you started with the Marzban Terraform Provider:

```hcl
terraform {
  required_providers {
    marzban = {
      source  = "mohrezfadaei/marzban"
      version = "0.1.0"
    }
  }
}

provider "marzban" {
  api_url = "https://api.marzban.example.com"
  token   = "your_api_token"
}

resource "marzban_node" "example" {
  name              = "example_node"
  address           = "192.168.1.1"
  port              = 62050
  api_port          = 62051
  usage_coefficient = 1.0
  add_as_new_host   = true
}
```

## Contributing

We welcome contributions to enhance the Marzban Terraform Provider. If you are interested in contributing, please get in touch with me at [<mohrezfadaei@gmail.com>](mailto:mohrezfadaei@gmail.com).

## License

This project is licensed under the Apache License. See the [LICENSE](./LICENSE) file for details.

## Support

If you encounter any issues or have questions, please open an issue on our [GitHub repository](https://github.com/mohrezfadaei/terraform-provider-marzban).

---

For more information about Terraform, visit the official [Terraform website](https://terraform.io).

<a href="https://terraform.io">
    <img src="https://www.datocms-assets.com/2885/1629941242-logo-terraform-main.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>
