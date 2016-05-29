# dps_app-cookbook

TODO: Enter the cookbook description here.

## Supported Platforms

TODO: List your supported platforms.

## Attributes

<table>
  <tr>
    <th>Key</th>
    <th>Type</th>
    <th>Description</th>
    <th>Default</th>
  </tr>
  <tr>
    <td><tt>['dps_app']['bacon']</tt></td>
    <td>Boolean</td>
    <td>whether to include bacon</td>
    <td><tt>true</tt></td>
  </tr>
</table>

## Usage

### dps_app::default

Include `dps_app` in your node's `run_list`:

```json
{
  "run_list": [
    "recipe[dps_app::default]"
  ]
}
```

## License and Authors

Author:: Jose M. Robles (jrobles@bellsouth.net)
