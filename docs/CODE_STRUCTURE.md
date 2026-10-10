# WoW Habits Code Structure

Description about important directories in our repository.

### bin/

- Build binaries are put here.

### cmd/

- Each directory here builds its own binary.

#### cmd/wowhabits/

- Entry point for the `wowhabits` binary.

### docs/

- Project documentation.
- More information is found in [README.md](./README.md).

### internal/

- Various packages to do all sorts of different things.

#### internal/api/

- The HTTP REST API server code.

### frontend/

- Web frontend, built with [Vue 3](https://vuejs.org/).
