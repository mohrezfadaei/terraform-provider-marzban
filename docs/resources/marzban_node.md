# marzban_node Resource

This resource manages a node in Marzban.

## Example Usage

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

## Arguments

- `name` - (Required) The name of the node.
- `address` - (Required) The address of the node.
- `port` - (Optional) The port of the node. Defaults to `62050`.
- `api_port` - (Optional) The API port of the node. Defaults to `62051`.
- `usage_coefficient` - (Optional) The usage coefficient of the node. Defaults to `1.0`.
- `add_as_new_host` - (Optional) Whether to add as a new host. Defaults to `true`.

## Attributes

- `node_id` - The ID of the node.
