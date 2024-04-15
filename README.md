# meta_gin: A Meta Framework for Gin

**WARNING:**
- This project is currently in the alpha stage. **Do not use in production environments.**
- The structure of the project is subject to change until it reaches version 1.0.0.

## Motivation

As a software engineer who values efficiency, I found myself repeatedly writing similar code for different projects. This inspired the creation of `meta_gin`, a framework designed to simplify and streamline web development with Gin and Gorm.

## Design Principles

`meta_gin` is crafted with the following considerations to empower developers:

- **Flexibility in Structure**: Developers have full control over their project structure.
- **Mandatory Tools**: Utilizes [Gin](https://github.com/gin-gonic/gin) and [Gorm](https://gorm.io) for routing and ORM functionality.
- **CRUD Operations Simplified**: Write less CRUD code, focus more on unique business logic.
- **Configuration Simplicity**: Minimize setup and configuration code to reduce bugs and development time.
- **Efficiency**: Less code means fewer bugs.

## Features

Efficient CRUD operations with minimal code:

![CRUD Example](https://github.com/mukezhz/meta_gin/assets/43813670/681dcb65-1dea-47c8-b01f-87e26d67cf7e)

For usage examples, please check the `examples` directory in the repository.

## Installation
- Initiate a go project:
```
go mod init <module>
````
- Add a meta_gin as dependency:
```
go get github.com/mukezhz/meta_gin
```

## Roadmap

- [ ] Support for custom middleware integration.
- [ ] Extension points for custom CRUD logic.
- [ ] Streamline the framework structure for enhanced clarity and reduced bloat.
- [ ] Development of a demo project using `meta_gin`.

## Contributing

As I continue to learn more about Go and the Gin framework, I welcome your feedback and contributions:
- Please start by opening an [issue](https://github.com/mukezhz/meta_gin/issues/new) to discuss potential changes before submitting a pull request.

**Let's code something amazing together! Happy coding! 🙇**

