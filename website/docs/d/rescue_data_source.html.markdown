---
layout: "rescue"
page_title: "Rescue: rescue_data_source"
sidebar_current: "docs-rescue-data-source"
description: |-
  Sample data source in the Terraform provider animal-rescue.
---

# rescue_data_source

Sample data source in the Terraform provider animal-rescue.

## Example Usage

```hcl
data "rescue_data_source" "example" {
  sample_attribute = "foo"
}
```

## Attributes Reference

* `sample_attribute` - Sample attribute.
