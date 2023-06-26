SnapSDK Specification Draft 1.0.0
================================

This is the `SnapSDK` spec. The intent is to eventually propose these draft changes to Open API.

This draft proposes a new SnapSDK specification. SnapSDK provides a common structure for defining the interface of SDKs and libraries, focusing on simplicity and ease of use for developers.

Structure
---------

A SnapSDK document consists of several sections, including info, namespace, objects, methods, and definitions.

Info
----

info provides metadata about the SDK or library.

```yaml
    info:
      version: "1.0.0"
      title: SDK or Library Title
``` 

Namespace
---------

namespace is a string that defines the namespace for the SDK or library methods.

```yaml
    namespace: Namespace
```    

Objects and Methods
-------------------

objects is a dictionary that provides a layer of abstraction for methods of the SDK or library. Each object can contain various methods.

methods is a dictionary that describes the functions provided by the SDK or library. Each method is an object that can contain various properties, such as operationId, summary, description, parameters, and returnType.

```yaml
    objects:
      object1:
        methods:
          method1:
            operationId: operationIdentifier
            summary: Summary of the method.
            description: Detailed description of the method.
            parameters:
              - name: parameterName
                type: parameterType
            returnType:
              type: returnType
``` 

Each method's parameters is an array of parameters that the method accepts. Each parameter has a name and a type.

The returnType of a method describes the type of value that the method returns. It has a type property and, for complex types, an items property that refers to a definition.

Definitions
-----------

definitions is a dictionary of data types used in the SDK or library. These can be used to describe complex return types, parameter types, and so on.

 ```yaml   
    definitions:
      DataType:
        type: object
        properties:
          propertyName:
            type: propertyType
```    

Each data type can have properties, which is a dictionary of property names to property types.

Types
-----

Types in SnapSDK should be defined to have a clear mapping to the basic data types in JavaScript, Rust, Go, and Python.

*   `string`: Maps to string in JavaScript and Go, str in Python, and String in Rust.
*   `integer`: Maps to number in JavaScript, int in Python and Go, and i32 or i64 in Rust.
*   `boolean`: Maps to boolean in JavaScript and Go, bool in Python and Rust.
*   `array`: Maps to Array in JavaScript, list in Python, slice in Go, and Vec in Rust.
*   `object`: Maps to Object in JavaScript, dict in Python, map or a struct in Go, and struct in Rust.

Example
-------

Here is an example of a SnapSDK specification:

```yaml
    info:
      version: "1.0.0"
      title: Dog API SDK
    namespace: Dog
    objects:
      dogDatabase:
        methods:
          getDog:
            operationId: getDog
            summary: Returns a dog by ID.
            description: Returns detailed information about a dog with the given ID.
            parameters:
              - name: id
                type: integer
            returnType:
              $ref: "#/definitions/Dog"
          getDogs:
            operationId: getDogs
            summary: Returns a list of dogs.
            description: Returns a list of dogs in the system.
            returnType:
              type: array
              items:
                $ref: "#/definitions/Dog"
    definitions:
      Dog:
        type: object
        properties:
          id:
            type: integer
          name:
            type: string
          breed:
            type: string
          age:
            type: integer
          color:
            type: string
          weight:
            type: number
          owner:
            type: string
```

This specification describes an SDK or library "Dog" with two methods under the "dogDatabase" object: getDog, which returns a Dog object for a given dog ID, and getDogs, which returns a list of Dog objects.