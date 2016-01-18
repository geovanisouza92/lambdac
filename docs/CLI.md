
# Command-line reference

**Command**: `$ lambdac runtime`
**REST Method**: `GET /runtimes`
**Status Code**: `HTTP 200 OK`
**Description**: List all registered runtimes.
**Aliases**: `rt`

---

**Command**: `$ lambdac runtime create`
**REST Method**: `POST /runtimes`
**Status Code**: `HTTP 201 Created`
**Description**: Creates a new runtime.
**Flags**:

- `-t`, `--template`: Template file to create multiple runtimes at once (override all other flags)
- `-n`, `--name`: Runtime name
- `-l`, `--label`: Runtime label
- `-i`, `--image`: Runtime image
- `-c`, `--command`: Runtime command
- `-a`, `--agent`: Run functions in this runtime as daemon, processing multiple events with lower latency
- `-d`, `--driver`: Runtime backend driver to create instances
- `-o`, `--driver-opt`: Runtime driver options used to create instances

---

**Command**: `$ lambdac runtime info`
**REST Method**: `GET /runtimes/{id}`
**Status Code**: `HTTP 200 OK`
**Description**: Gets runtime detailed information.
**Flags**:

- `-r`, `--runtime`: Runtime ID or name

---

**Command**: `$ lambdac runtime destroy`
**REST Method**: `DELETE /runtimes/{id}`
**Status Code**: `HTTP 410 Gone`
**Description**: Remove a runtime.
**Aliases**: `rm`
**Flags**:

- `-r`, `--runtime`: Runtime ID or name
- `-f`, `--force`: Remove runtime and functions at once

---

**Command**: `$ lambdac daemon`
**Description**: Start the daemon.
