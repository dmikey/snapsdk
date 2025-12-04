# DogsApp SDK

## Object: dog
### Definitions
- id: integer
- name: string
- breed: string
- age: integer

### Methods
#### listDogs
- Summary: Returns a list of dogs.
- Description: Returns a list of dogs in the system.
- Usage Examples:
  - Go: `Dog.ListDogs()`
  - Python: `dog.listDogs()`
  - JavaScript: `dog.listDogs()`
  - Rust: `dog::listDogs()`
- Returns: array

#### getDog
- Summary: Returns a dog by ID.
- Description: Returns detailed information about a dog with the given ID.
- Parameters: id (integer)
- Usage Examples:
  - Go: `Dog.GetDog(id (integer))`
  - Python: `dog.getDog(id (integer))`
  - JavaScript: `dog.getDog(id (integer))`
  - Rust: `dog::getDog(id (integer))`
- Returns: Dog


##### Generated with Snap version: 1.0
