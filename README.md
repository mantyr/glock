# glock

[![Build Status](https://travis-ci.org/KyleBanks/glock.svg?branch=master)](https://travis-ci.org/KyleBanks/glock)

`glock` is a REST based locking system for distributed applications.

## Core Concepts

The majority of `glock` actions require, at minimum, a `key` and a `secret`.
- The `key` is any string that you use to identify a lock. For example, if you were building a system that sends emails to users on a timed schedule, and you want to ensure only one server/thread sends the email to each user, you may use the email address or username as the key. This value is shared across all instances of your application, to ensure everyone is working with the same `key`.
- The `secret` is a unique value returned by lock commands that allows you to perform additional actions on that key in the future, such as unlocking or extending a lock on the key.

## Options

The following ptions can be specified when running `glock`:

- `-p` *(Default: 7887)*: The port to run `glock` on.
- `-v` *(Default: false)*: Enables verbose output.

## API

All `glock` methods are exposed via a REST API accessible at:

```
<host>:<port>/api/v1.0/<action>
```

### Success Response
For any action, one of two possible JSON responses can be returned. For successful actions, a `Success Response` is returned that looks like so:

```
{
    "success": true,
    "extras": {}
}
```
The `extras` property may or may not exist, depending on the action. The contents of `extras` is documented for each action below.

### Error Response
For failed actions, a `Error Response` is returned:
```
{
    "success": false,
    "error": {
        "code": int,
        "message": string
    }
}
```

The error code of an `Error Response` contains a particular error code for each unique error that can occur, allowing you to take specific action based on the error received.

### /lock

**Path:** */api/v1.0/lock*

**Parameters:**
- `key`: The `key` to attempt to lock.
- `duration` *Optional*:  If specified, the lock on the `key` will be removed automatically after the specified duration, in milliseconds.

**Description:**

Attempts to place a lock on a particular `key`.

**Response:**

If the lock was successful, a `Success Response` is returned containing a `secret` that can be used in future requests to modify the lock.

If the lock failed for any reason, including the `key` already being locked, an `Error Response` will be returned.

**Example:**
```
# Request:
/api/v1.0/lock?key=sampleKey

# Response:
{
    "success": true,
    "extras": {
        "secret": "1234567890-0987654321"
    }
}
```


## Testing

Tests are run using `./test.sh` in the root directory of `glock`.

## License

```
The MIT License (MIT)

Copyright (c) 2016 Kyle Banks

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
