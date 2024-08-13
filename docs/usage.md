# Usage

## Creating a Node

To create a node using the Marzban Terraform Provider, add the following to your Terraform configuration:

```hcl
resource "marzban_node" "example" {
  name              = "example_node"
  address           = "192.168.1.1"
  port              = 62050
  api_port          = 62051
  usage_coefficient = 1.0
  add_as_new_host   = true
}
```

## Applying the Configuration

1. Initialize the Terraform configuration:

   ```sh
   terraform init
   ```

2. Validate the configuration (optional but recommended):

   ```sh
   terraform validate
   ```

3. Apply the configuration to create the node:

   ```sh
   terraform apply
   ```

This will create a node in Marzban with the specified configuration.
