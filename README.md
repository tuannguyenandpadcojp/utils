# utils
Utilizations packages use for Go projects

## Packages
| package                  | coverage                                                                           |
| ------------------------ | -----------------------------------------------------------------------------------|
| [base62][base62-package] | [![codecov](https://codecov.io/gh/tuannguyenandpadcojp/utils/branch/main/graph/badge.svg?token=APZD94NB9H&flag=base62)](https://codecov.io/gh/tuannguyenandpadcojp/utils) |
| [valctx][valctx-package] | [![codecov](https://codecov.io/gh/tuannguyenandpadcojp/utils/branch/main/graph/badge.svg?token=APZD94NB9H&flag=valctx)](https://codecov.io/gh/tuannguyenandpadcojp/utils) |

### base62
A package use to encode int64 to string and vice versa using base62 encoding.

### valctx
A package use to create a new context that inherits values from a parent context but doesn't propagate its cancellation or deadline to child contexts.

[base62-package]: ./base62
[valctx-package]: ./valctx
