
# Command-line reference

**Command**: `$ lambdac list`
**REST Method**: `GET /functions`
**Status Code**: `HTTP 200 OK`
**Description**: List all functions.
**Aliases**: `ls`

---

**Command**: `$ lambdac create`
**REST Method**: `POST /functions`
**Status Code**: `HTTP 201 Created`
**Description**: Creates a function.
**Aliases**: `new`
**Arguments**: `code path` could be a file or folder.
**Flags**:

- `-n`, `--name`: Function name
- `-r`, `--runtime`: Function runtime
- `-h`, `--handler`: Function handler
- `-d`, `--description`: Function description (optional)
- `-t`, `--timeout`: Function timeout (default: 3s ; units: ns, us (or µs), ms, s, m, h)
- `-m`, `--memory`: Function memory limit (default: 32 MB)
- `-i`, `--instances`: Function max instances
- `-e`, `--env`: Function environment variables (format: VARIABLE=VALUE ; can be informed multiple times)

---

**Command**: `$ lambdac info`
**REST Method**: `GET /function/{id}`
**Status Code**: `HTTP 200 OK`
**Description**: Gets detailed information about a function.
**Flags**:

- `-f`, `--function`, `$LAMBDAC_FUNCTION`: Function ID or name

---

**Command**: `$ lambdac config`
**REST Method**: `PUT /functions/{id}`
**Status Code**: `HTTP 202 Accepted`
**Description**: Changes function instances configuration.
**Flags**:

- `-f`, `--function`, `$LAMBDAC_FUNCTION`: Function ID or name
- `-h`, `--handler`: Function handler
- `-d`, `--description`: Function description (optional)
- `-t`, `--timeout`: Function timeout (default: 3s ; units: ns, us (or µs), ms, s, m, h)
- `-m`, `--memory`: Function memory limit (default: 32 MB)
- `-i`, `--instances`: Function max instances

---

**Command**: `$ lambdac destroy`
**REST Method**: `DELETE /function/{id}`
**Status Code**: `HTTP 410 Gone`
**Description**: Destroy a function.
**Aliases**: `rm`
**Flags**:

- `-f`, `--function`, `$LAMBDAC_FUNCTION`: Function ID or name

---

**Command**: `$ lambdac env`
**REST Method**: `GET /functions/{id}/env`
**Status Code**: `HTTP 200 OK`
**Description**: List environment variables.
**Flags**:

- `-f`, `--function`, `$LAMBDAC_FUNCTION`: Function ID or name

---

**Command**: `$ lambdac env set`
**REST Method**: `PUT /functions/{id}/env`
**Status Code**: `HTTP 202 Accepted`
**Description**: Adds/changes environment variables.
**Arguments**: `VARIABLE=VALUE VARIABLE=VALUE ...` key-value pairs representing environment variables.
**Flags**:

- `-f`, `--function`, `$LAMBDAC_FUNCTION`: Function ID or name

---

**Command**: `$ lambdac env unset`
**REST Method**: `DELETE /functions/{id}/env`
**Status Code**: `HTTP 410 Gone`
**Description**: Removes environment variables.
**Arguments**: `VARIABLE VARIABLE ...` keys to remove.
**Flags**:

- `-f`, `--function`, `$LAMBDAC_FUNCTION`: Function ID or name

---

**Command**: `$ lambdac pull`
**REST Method**: `GET /functions/{id}/code`
**Status Code**: `HTTP 200 OK`
**Description**: Gets function code.
**Arguments**: `code path (default: $CWD)` folder to put the code.
**Flags**:

- `-f`, `--function`, `$LAMBDAC_FUNCTION`: Function ID or name

---

**Command**: `$ lambdac push`
**REST Method**: `PUT /functions/{id}/code`
**Status Code**: `HTTP 202 Accepted`
**Description**: Updates function code.
**Arguments**: `code path (default: $CWD)` could be a file or folder.
**Flags**:

- `-f`, `--function`, `$LAMBDAC_FUNCTION`: Function ID or name

---

**Command**: `$ lambdac ps`
**REST Method**: `GET /functions/{id}/ps`
**Status Code**: `HTTP 200 OK`
**Description**: List function instances.
**Flags**:

- `-f`, `--function`, `$LAMBDAC_FUNCTION`: Function ID or name

---

**Command**: `$ lambdac logs`
**REST Method**: `GET /functions/{id}/logs`
**Status Code**: `HTTP 200 OK`
**Description**: Display function instances logs.
**Flags**:

- `-f`, `--function`, `$LAMBDAC_FUNCTION`: Function ID or name
- `-t`, `--tail`: Number of lines to show from the end of logs (< 1 means all lines)

---

**Command**: `$ lambdac stats`
**REST Method**: `GET /functions/{id}/stats`
**Status Code**: `HTTP 200 OK`
**Description**: Shows function statistics.
**Flags**:

- `-f`, `--function`, `$LAMBDAC_FUNCTION`: Function ID or name

---

**Command**: `$ lambdac invoke`
**REST Method**: `POST /function/{id}/invoke`
**Status Code**: `HTTP 202 Accepted`
**Description**: Calls a function.
**Arguments**: `event data` all arguments are sent as event payload.
**Flags**:

- `-f`, `--function`, `$LAMBDAC_FUNCTION`: Function ID or name

---

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
