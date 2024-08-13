# Installation

To use the Marzban Terraform Provider, add the following to your Terraform configuration:

```hcl
terraform {
  required_providers {
    marzban = {
      source  = "yourusername/marzban"
      version = "0.1.0"
    }
  }
}

provider "marzban" {
  # Optional: provider configuration
}
```

## Provider Configuration

The provider configuration can include the following optional parameters:

- `api_url`: The base URL for the Marzban API.
- `token`: The API token for authentication.

Example:

```hcl
provider "marzban" {
  api_url = "https://api.marzban.example.com"
  token   = "your_api_token"
}
