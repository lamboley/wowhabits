# WoW Habits Architecture

This project uses a hexagonal architecture. The business logic is in the center and everything technical (HTTP, database, email, ...) is pushed to the edges. The center does not know or care which web framework or database is used. It only talks to the outside world through interfaces it defines itself.
