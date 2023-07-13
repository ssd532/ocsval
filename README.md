# Omni CSV Validator

Omni CSV validator validates CSV files against a schema.

## Schema

The `Config` JSON file is used to define validation rules for a CSV file and has two sections: `global` and `rules`.

### Global

The `global` object defines rules that apply to the entire CSV file.

```json
  "global": {
    "maxRecords": 100,
    "maxColumns": 5,
    "separator": ","
  },
```

- `maxRecords` (integer): This specifies the maximum number of records (lines) that the CSV file should contain.
- `maxColumns` (integer): This is the maximum number of columns that each record in the CSV file should contain.
- `separator` (string): This is the character that separates columns within each record (line) in the CSV file.

### Rules

The `rules` object is an array where each item represents a validation rule for a particular column of the CSV. The type of rules we currently support is called `column`.

Each `column` rule includes these properties:

```json
    {
      "type": "column",
      "column": "Name",
      "required": true,
      "dataType": "string"
    },
```

- `type` (string): This specifies the type of rule. Currently the package only supports `column` type rules.
- `column` (string): This is the name of the column that the rule applies to.
- `required` (boolean): This specifies if the column must be present in each record. If `required` is set to true and the column is missing in a record, it will be flagged as an error.
- `dataType` (string): This specifies the expected data type for the column. Currently `string` and `integer` are supported. If the data in the column does not match the specified data type, it will be flagged as an error.

### Future

For now, our schema supports the `type` value of "column", which allows us to define rule validations for specific CSV columns. Expanding this, we could consider adding additional types that would allow us to do more complex validations. Some of these could include:

1. `"record"`: This could allow us to define rules that validate an entire record (row). For example, we may want to ensure that at least one of several columns is filled, which is not possible with "column"-level validation.

2. `"file"`: This would allow us to define rules that sit at a higher level than our current "maxRecords" and "maxColumns" global rules. Possible `file` level rules could enforce that the file must have a specific set of columns, or that specific columns must appear in a certain order.

3. `"inter-column"`: This could allow validations that span multiple columns. For example, you might want to validate that the value in "start date" column is earlier than the "end date" column in the same record.

4. `"custom"`: For a user-defined rule. The user would provide a custom function that implements their unique rule.

For these new `type` values to work, each would require a different set of additional properties. For example, the `custom` type might need a `customFunction` property that references a user-supplied validation function.

In designing these rule types, you would consider the most common validation scenarios you need to support, beginning with the most general types (like `column`) and adding more specific ones as needed.