---
layout: "animalrescue"
page_title: "AnimalRescue: animalrescue_data_source"
sidebar_current: "docs-animalrescue-data-source"
description: |-
  Sample data source in the Terraform provider animalrescue.
---

# animalrescue_data_source

Sample data source in the Terraform provider animalrescue.

## Example Usage

```hcl
data "animalrescue_data_source" "example" {
  sample_attribute = "foo"
}
```

## Attributes Reference

* `sample_attribute` - Sample attribute.
