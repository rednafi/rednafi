<!-- Zero width character is used to put extra blank lines before and after code -->

<h3>

```python
​
from dataclasses import dataclass
from typing import Tuple

@dataclass
class Stack:
    languages   : Tuple[str, ...] = ("Python", "Go", "Bash")
    databases   : Tuple[str, ...] = ("PostgreSQL", "Mongo", "Redis")
    misc        : Tuple[str, ...] = ("Docker", "Celery")
    ongoing     : Tuple[str, ...] = ("Django", "GraphQL")
​
```
</h3>
