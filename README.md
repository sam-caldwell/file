File
====

## Description

Simple file utilities.

## Packages

| Package | Description |
|---------|-------------|
| 

## Functions

| Feature                     | Description                                                                                                                                                            | 
|-----------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `CountFilesWithName`        | Walk a given path and count the number of files found.                                                                                                                 |
| `Create`                    | This is a wrapper for `os.WriteFile()` because I'm tired of golang changing things requiring a bunch of effort to fix.                                                 |
| `CreateRandomImage`         | This will create a file with a specified size, filled with random data                                                                                                 |
| `CreateTempFile`            | Creates a temporary file in a temporary directory using a unique identifier (UUID) and returns  the file path. It handles the cleanup of the temporary file after use. |
| `CreateTempFileWithContent` | Create a temporary file with the given content                                                                                                                         | 
| `GetExtension`              | Extracts the extension of a given file name and returns it as a string.                                                                                                |
| `HasExtension`              | Checks if a file has a specific extension and returns a boolean result.                                                                                                |
| `Existsp`                   | Determines whether a given file exists and returns a boolean value.                                                                                                    |
| `GetExtension`              | Extracts the extension of a given file name and returns it as a string.                                                                                                |
| `HasExtensions`             | Checks if a file has any of the specified extensions from a given list.                                                                                                |
| `IsValidYaml`               | Validates if a YAML file can be opened and unmarshalled without error.                                                                                                 |
| `HasJsonExtension`          | Checks if a file has a JSON extension (case-insensitive) and returns a boolean result.                                                                                 |
| `HasYamlExtension`          | Checks if a file has a YAML extension (case-insensitive) and returns a boolean result.                                                                                 |



