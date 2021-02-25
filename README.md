# Truefit Manifest

This is intended to be a fun project to try some different technologies, build something interesting(maybe), and learn things.

## API

_Why Go?_

First thing is first...why go? For fun! Wanted to try something new and compare it with our dotnet stack.

With that said, things are going to be done differently here than you're used to. We are trying to keep everything as idiomatic as possible in terms of doing things the "go way".

Some notes in that regard:

- Avoid bloated, "batteries included" frameworks in favor of smaller, flexible tools. For example, we are using the lightweight Chi router instead of a full blown web framework like Gin

- No ORM. We are using SqlBoiler to generate some helpers for DB access and sql-migrate for migrations. The Go mindset dictates that ORMs too often get in our way and cause performance issues. I think the argument against this approach is going to be decreased dev productivity but let's give it a shot for the sake of trying something new.
