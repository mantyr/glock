# glock

[![Build Status](https://travis-ci.org/KyleBanks/glock.svg?branch=master)](https://travis-ci.org/KyleBanks/glock)

`glock` is a REST based locking system for distributed applications.

## Core Concepts

The majority of `glock` actions require, at minimum, a `key` and a `secret`.
- The `key` is any string that you use to identify a lock. For example, if you were building a system that sends emails to users on a timed schedule, and you want to ensure only one server/thread sends the email to each user, you may use the email address or username as the key. This value is shared across all instances of your application, to ensure everyone is working with the same `key`.
- The `secret` is unique to each instance of your application. It could be a thread or server identifier, a UUID, or anything that you can ensure each instance of your application will have a unique value for. When a lock is successfully placed on a `key`, this secret is the only way to make modifications to the lock in the future, such as unlocking or extending the lock.   

## Options

The following ptions can be specified when running `glock`:

- `-p` *(Default: 7887)*: The port to run `glock` on.
- `-v` *(Default: false)*: Enables verbose output.

## API

All `glock` methods are exposed via a REST API accessible at:

```
<host>:<port>/api/v1.0/<action>
```

### /lock

**Path:** */api/v1.0/lock*

**Parameters:**
- `key`: The `key` to attempt to lock.
- `secret`: If the lock is successful, the `secret` can be used by future requests to modify the locked `key`, such as `/unlock`.
- `duration` *Optional*:  If specified, the lock on the `key` will be removed automatically after the specified duration, in milliseconds.


## Testing

Tests are run using `./test.sh` in the root directory.

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
